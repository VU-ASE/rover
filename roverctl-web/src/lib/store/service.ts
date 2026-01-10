/**
 * This is all state associated with a service store.
 */

import { SensorOutput } from 'rovercom/gen/outputs/wrapper';
import type { ServiceEndpoint, ServiceIdentifier } from 'rovercom/gen/debug/debug';
import { get, writable } from 'svelte/store';
import { createSensorStreamStore, type SensorStreamStore } from './sensorstream';

type ServiceState = Omit<ServiceIdentifier, 'endpoints'> & {
	endpoints: Map<
		string, // endpoint name
		ServiceEndpoint & {
			// Add a map of sensor stream stores to the service state,
			// each endpoint can have multiple sensor streams, as one endpoint can expose multiple sensors (using different sensor ids)
			streams: Map<number, SensorStreamStore>;
		}
	>;
	realName: string; // name as unmodified by the "as" field
};

const createServiceStore = (service: ServiceState) => {
	const store = writable(service);
	const { subscribe, update, set } = store;

	return {
		subscribe,
		update,
		set,
		addFrame: (endpoint: ServiceEndpoint, sentAt: Date, sensorData: SensorOutput) => {
			// We only need to update when the stream does not exist yet
			const val = get(store);

			// Each endpoint has a map of sensor streams, try to fetch this map or create a new one
			const streamMap =
				val.endpoints.get(endpoint.name)?.streams || new Map<number, SensorStreamStore>();

			// Does there already exist a stream for this sensor id?
			const sensorStream = streamMap.get(sensorData.sensorId);
			if (!sensorStream) {
				const newStream = createSensorStreamStore(sensorData.sensorId);
				newStream.addFrame(sensorData, sentAt);
				streamMap.set(sensorData.sensorId, newStream);
			} else {
				sensorStream.addFrame(sensorData, sentAt);
			}

			// Each service has a map of endpoints, try to fetch the endpoint for which we just added a stream
			const key = endpoint.name;
			const endpointItem = val.endpoints.get(key);
			if (!endpointItem) {
				update((oldState) => {
					const newState = { ...oldState };
					const newEndpoint = { ...endpoint, streams: streamMap };
					newState.endpoints.set(key, newEndpoint);
					return newState;
				});
			} else {
				update((oldState) => {
					const newState = { ...oldState };
					const newEndpoint = { ...endpointItem, streams: streamMap };
					newState.endpoints.set(key, newEndpoint);
					return newState;
				});
			}
		}
	};
};

type ServiceStore = ReturnType<typeof createServiceStore>;

export { createServiceStore };
export type { ServiceStore, ServiceState };
