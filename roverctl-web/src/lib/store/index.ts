/**
 * This is the global state that keeps track of all services, their state, their configuration and their sensors.
 * It also keeps the "stream" that allows for playback and scrubbing
 */

import { get, writable } from 'svelte/store';
import { createServiceStore, type ServiceStore } from './service';
import { DebugOutput } from 'ase-rovercom/gen/debug/debug';
import { SensorOutput } from 'ase-rovercom/gen/outputs/wrapper';
import { TuningState } from 'ase-rovercom/gen/tuning/tuning';
import { notify } from '$lib/events/notifications';

type GlobalState = {
	// Used for scrubbing and playback
	paused: Date | null; // if null, we are not paused, if a date, we are paused at that date (and time)
	millisecondsPreserved: number; // how many seconds of data do we keep?
	scrubOffsetMilliseconds: number; // how many seconds back in time are we currently?

	// Used to show status when the car is (dis)connected
	carConnected: boolean;
	carOffset: number; // The clock offset between the car and the forwarding server, will be <0 if the car is behind

	// A map of all services (as nested stores), indexed by their name
	services: Map<string, ServiceStore>;

	// Used to show the camera feed in the UI
	cameraFeeds: {
		service: string;
		endpoint: string;
		sensorId: number;
	}[];

	// Used to show a different ordering of the services
	fullscreenMode: boolean;

	// The tuning parameters are received periodically from the car, the overrides are custom set by the user
	tuning: TuningState | undefined;
	tuningOverrides: ({
		key: string;
	} & (
		| {
				type: 'integer';
				value: number;
		  }
		| {
				type: 'float';
				value: number;
		  }
		| {
				type: 'string';
				value: string;
		  }
	))[];
};

const createGlobalStore = () => {
	const store = writable<GlobalState>({
		paused: null,
		millisecondsPreserved: 10 * 1000,
		scrubOffsetMilliseconds: 0,
		services: new Map(),
		carConnected: false,
		carOffset: 0,
		cameraFeeds: [],
		fullscreenMode: false,
		tuning: undefined,
		tuningOverrides: []
	});
	const { subscribe, update, set } = store;

	return {
		// Required functions
		subscribe,
		update,
		set: (newState: GlobalState) => {
			// Must be minimally 1 second preserved
			const millisecondsPreserved = Math.max(newState.millisecondsPreserved, 1 * 1000);
			// We can not scrub further back than the preserved data
			const scrubOffsetMilliseconds = Math.min(
				newState.scrubOffsetMilliseconds,
				millisecondsPreserved
			);

			return set({
				...newState,
				millisecondsPreserved: millisecondsPreserved,
				scrubOffsetMilliseconds: scrubOffsetMilliseconds
			});
		},
		// Add an incoming sensor frame to the global state, pass it down
		addFrame: (frame: DebugOutput) => {
			// We only need to update when the stream does not exist yet (and not paused)
			const val = get(store);

			if (val.paused) {
				return;
			}

			if (!frame.service || !frame.endpoint) {
				console.error('Frame without service or endpoint', frame);
				return;
			}

			// try to parse the sensor data
			let parsedData: SensorOutput;
			try {
				parsedData = SensorOutput.decode(frame.message);
			} catch (e) {
				console.error('Could not parse sensor data', e);
				return;
			}

			if (parsedData.cameraOutput) {
				parsedData.timestamp = frame.sentAt; // todo: check
			}

			if (!parsedData.sensorId) {
				console.error('Sensor data without sensorId', parsedData);
				return;
			}

			const service = frame.service;
			const endpoint = frame.endpoint;
			const sentAt = new Date(frame.sentAt); // convert ms timestamp to Date

			// Is this service already in the global state?
			const serviceStore = val.services.get(service.name);
			if (!serviceStore) {
				update((oldState) => {
					const newState = { ...oldState };
					const newService = createServiceStore({
						name: service.name,
						pid: -1,
						endpoints: new Map()
					});
					newService.addFrame(endpoint, sentAt, parsedData);
					newState.services.set(service.name, newService);
					return newState;
				});
			} else {
				serviceStore.addFrame(endpoint, sentAt, parsedData);
			}
		},
		toggleStream: () => {
			update((oldState) => {
				let newPaused: Date | null = null;
				if (oldState.paused && oldState.carConnected) {
					newPaused = null;
				} else {
					newPaused = new Date();
				}

				return {
					...oldState,
					paused: newPaused
				};
			});
		},
		pauseStream: () => {
			update((oldState) => {
				return {
					...oldState,
					paused: oldState.paused ? oldState.paused : new Date()
				};
			});
		},
		setCarConnected: (connected: boolean) => {
			update((oldState) => {
				return {
					...oldState,
					carConnected: connected
				};
			});
		},
		setCarOffset: (offset: number) => {
			update((oldState) => {
				return {
					...oldState,
					carOffset: offset
				};
			});
		},
		// updateServicelist: (list: ServiceList) => {
		// 	update((oldState) => {
		// 		const newState = { ...oldState };

		// 		for (const service of list.services) {
		// 			if (!service.identifier?.name) {
		// 				continue;
		// 			}

		// 			// Do we already have this service?
		// 			const serviceStore = newState.services.get(service.identifier?.name);
		// 			if (!serviceStore) {
		// 				// Insert
		// 				const newService = createServiceStore({
		// 					...service,
		// 					endpoints: new Map()
		// 				});
		// 				newState.services.set(service.identifier?.name, newService);
		// 			} else {
		// 				// Update
		// 				// - replace the identifier
		// 				// - update the status
		// 				// - update the options
		// 				// - update the dependencies
		// 				// - update the endpoints, but do not touch the streams

		// 				const newService = get(serviceStore);
		// 				newService.identifier = service.identifier;
		// 				newService.status = service.status;
		// 				newService.options = service.options;
		// 				newService.dependencies = service.dependencies;
		// 				newService.registeredAt = service.registeredAt;

		// 				// Filter all endpoints that are not in the new service
		// 				newService.endpoints = new Map(
		// 					[...newService.endpoints].filter(([key]) =>
		// 						service.endpoints.find((e) => e.name === key)
		// 					)
		// 				);

		// 				// Add all endpoints that are not in the old service
		// 				for (const endpoint of service.endpoints) {
		// 					const endpointName = endpoint.name;
		// 					const endpointStore = newService.endpoints.get(endpointName);
		// 					if (!endpointStore) {
		// 						newService.endpoints.set(endpointName, {
		// 							...endpoint,
		// 							streams: new Map()
		// 						});
		// 					}
		// 				}

		// 				serviceStore.set(newService);
		// 			}
		// 		}

		// 		// Update the status of all services that are no longer in the service list to STOPPED
		// 		for (const [name, serviceStore] of newState.services) {
		// 			if (!list.services.find((s) => s.identifier?.name === name)) {
		// 				const service = get(serviceStore);
		// 				service.status = ServiceStatus.STOPPED;
		// 				serviceStore.set(service);
		// 			}
		// 		}

		// 		return newState;
		// 	});
		// },
		addCameraFeed: (service: string, endpoint: string, sensorId: number) => {
			update((oldState) => {
				if (
					oldState.cameraFeeds.find(
						(f) => f.service === service && f.endpoint === endpoint && f.sensorId === sensorId
					)
				) {
					return oldState;
				}

				const newState = { ...oldState };
				newState.cameraFeeds.push({ service, endpoint, sensorId });
				return newState;
			});
		},
		setTuning: (tuning: TuningState) => {
			update((oldState) => {
				// If there is a tuning override in the current state that has the same name but a mismatching type, remove it
				const tuningOverrides = oldState.tuningOverrides.filter((override) => {
					const tuningParam = tuning.dynamicParameters.find((param) => {
						return param.number?.key === override.key || param.string?.key === override.key;
					});

					// If there is no match, the tuning override should be removed anyway
					if (!tuningParam) {
						return false;
					}

					// If there is a match, check if the types match
					if (tuningParam.number && override.type === 'float') {
						return true;
					} else if (tuningParam.number && override.type === 'integer') {
						return true;
					} else if (tuningParam.string && override.type === 'string') {
						return true;
					}

					return false;
				});

				return {
					...oldState,
					tuningOverrides: tuningOverrides,
					tuning: tuning
				};
			});
		},
		setTuningOverride: (key: string, value: string) => {
			update((oldState) => {
				let tuningOverrides = oldState.tuningOverrides;

				// Short path, if the value is empty, remove it from the overrides
				if (value === '') {
					tuningOverrides = tuningOverrides.filter((override) => {
						return override.key !== key;
					});

					return {
						...oldState,
						tuningOverrides: tuningOverrides
					};
				}

				// Check if there is a tuning parameter with the same key
				const tuningParam = oldState.tuning?.dynamicParameters.find((param) => {
					return param.number?.key === key || param.string?.key === key;
				});

				// We can't override a value that does not exist
				if (!tuningParam) {
					return oldState;
				}

				// Parse based on the original tuning parameter type
				if (tuningParam.string) {
					tuningOverrides = tuningOverrides.filter((override) => {
						return override.key !== key;
					});
					tuningOverrides.push({
						key: key,
						type: 'string',
						value: value
					});
				} else if (tuningParam.number) {
					const floatValue = parseFloat(value);
					if (isNaN(floatValue)) {
						notify('error', `Tuning value '${key}' must be a float`);
						return oldState;
					}

					tuningOverrides = tuningOverrides.filter((override) => {
						return override.key !== key;
					});
					tuningOverrides.push({
						key: key,
						type: 'float',
						value: floatValue
					});
				}

				return {
					...oldState,
					tuningOverrides: tuningOverrides
				};
			});
		}
	};
};

type GlobalStore = ReturnType<typeof createGlobalStore>;

const globalStore = createGlobalStore();

export { globalStore };
export type { GlobalStore, GlobalState };
