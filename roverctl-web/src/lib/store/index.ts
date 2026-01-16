/**
 * This is the global state that keeps track of all services, their state, their configuration and their sensors.
 * It also keeps the "stream" that allows for playback and scrubbing
 */

import { get, writable } from 'svelte/store';
import { type ServiceStore } from './service';
import { DebugOutput } from 'rovercom/gen/debug/debug';
import { SensorOutput } from 'rovercom/gen/outputs/wrapper';
import { TuningState } from 'rovercom/gen/tuning/tuning';
import { notify } from '$lib/events/notifications';

type GlobalState = {
	// Used for scrubbing and playback
	paused: Date | null; // if null, we are not paused, if a date, we are paused at that date (and time)
	millisecondsPreserved: number; // how many seconds of data do we keep?
	scrubOffsetMilliseconds: number; // how many seconds back in time are we currently?

	// Used to show status when the car is (dis)connected
	carConnected: boolean;
	carOffset: number; // The clock offset between the car and the forwarding server, will be <0 if the car is behind. Currently not used

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

	// If set, this holds a timestamp of a received frame from the car, used to display a warning about clock drift
	clockDrift: Date | null;

	// If set, holds the battery voltage of the car
	battery?: {
		voltage: number;
		sentAt: Date;
	};
};

const defaultState: GlobalState = {
	paused: null,
	millisecondsPreserved: 10 * 1000,
	scrubOffsetMilliseconds: 0,
	services: new Map(),
	carConnected: false,
	carOffset: 0,
	cameraFeeds: [],
	fullscreenMode: false,
	tuning: undefined,
	tuningOverrides: [],
	clockDrift: null,
	battery: undefined
};

const createGlobalStore = () => {
	const store = writable<GlobalState>(defaultState);
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
				parsedData.timestamp = frame.sentAt;
			}

			if (!parsedData.sensorId) {
				console.error('Sensor data without sensorId', parsedData);
				return;
			}

			const service = frame.service;
			const endpoint = frame.endpoint;
			const sentAt = new Date(frame.sentAt); // convert ms timestamp to Date

			// Compare the offset between the car and the local clock
			const carOffset = sentAt.getTime() - Date.now();
			if (Math.abs(carOffset - val.carOffset) > 500 && !val.clockDrift) {
				// More than 500ms difference, and we are not already showing a clock drift warning
				update((oldState) => {
					return {
						...oldState,
						clockDrift: sentAt
					};
				});
			} else if (Math.abs(carOffset - val.carOffset) <= 500 && val.clockDrift) {
				// Less than 500ms difference, and we are showing a clock drift warning, get rid of it
				update((oldState) => {
					return {
						...oldState,
						clockDrift: null
					};
				});
			}

			// Is this service already in the global state?
			const serviceStore = val.services.get(service.name);
			if (!serviceStore && service.name === 'battery') {
				console.log('Received battery info:', parsedData);
				if (parsedData.batteryOutput) {
					update((oldState) => {
						return {
							...oldState,
							battery: {
								voltage: parsedData.batteryOutput?.currentOutputVoltage || 0,
								sentAt: sentAt
							}
						};
					});
				}
			} else if (!serviceStore) {
				console.warn(
					'Received debug data from a service  not found in global state',
					service,
					parsedData
				);
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
		},
		reset: () => {
			set(defaultState);
		}
	};
};

type GlobalStore = ReturnType<typeof createGlobalStore>;

const globalStore = createGlobalStore();

export { globalStore };
export type { GlobalStore, GlobalState };
