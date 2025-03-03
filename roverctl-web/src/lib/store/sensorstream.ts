/**
 * All state and store associated to a single sensor data stream.
 */

import { writable } from 'svelte/store';
import { SensorOutput } from 'ase-rovercom/gen/outputs/wrapper';
import { getParametersFromFrame } from '$lib/utils/sensorType';
import {
	addFrameToCircularBuffer,
	type CircularBuffer,
	createCircularBuffer
} from './circularbuffer';

// Used to calculate FPS, latency and other statistics
type TimestampedSensorOutput = {
	sentAt: Date;
	receivedAt: Date;
	sensorData: SensorOutput;
};

// As used by the charting engine
type Coordinate = {
	x: number; // this is the timestamp
	y: number; // this is the value
};

// This is one of the individual sensor streams that a user can view in the UI
type SensorStreamState = {
	sensorId: number; // since some outputs may have multiple sensors

	// This is a fixed-size, circular buffer of the last frames with all information added
	received: CircularBuffer<TimestampedSensorOutput>;

	// Timeline entries allow for faster rendering of charts. This is a map of key-value pairs, where the key is the parameter name and the value is an array of timeline entries
	// the goal is to boost performance by keeping the array fixed-size and using modulo operations on the timestamp to quickly find the right index
	chartItems: Record<
		string,
		CircularBuffer<Coordinate> // the key here represents a parameter name on a single sensor, like key "rpm" on sensor "speedsensor"
	>;
};

const createSensorStreamStore = (sensorId: number) => {
	const { subscribe, update, set } = writable<SensorStreamState>({
		sensorId: sensorId,
		received: createCircularBuffer(),
		chartItems: {}
	});

	return {
		// Required functions
		subscribe,
		update,
		set,
		// Custom functions
		addFrame: (frame: SensorOutput, sentAt: Date) => {
			update((oldState) => {
				// Add a timestamp to the frame
				const timestampedFrame: TimestampedSensorOutput = {
					sentAt: sentAt,
					receivedAt: new Date(),
					sensorData: frame
				};

				const newState = oldState;

				// Update the circular buffer
				newState.received = addFrameToCircularBuffer(
					newState.received,
					timestampedFrame,
					timestampedFrame.sensorData.timestamp,
					(frame) => frame.sensorData.timestamp
				);

				// Create or update parallel chart "streams" for each parameter that can be graphed, without needing to normalize and parse the data again for the ChartJS engine
				// (this is a memory-compute time tradeoff, where we prefer short compute time and more memory usage)
				const entries = getParametersFromFrame(timestampedFrame);
				// the key here represents a parameter name on a single sensor, like key "rpm" on sensor "speedsensor"
				for (const [key, value] of entries) {
					// Does there already exist a circular buffer for this key?
					// if not, we create it
					if (!newState.chartItems[key]) {
						newState.chartItems[key] = createCircularBuffer();
					}
					// This is the entry we are going to insert
					const newTimelineEntry = {
						x: timestampedFrame.sentAt.getTime(),
						y: value
					};

					// Update the circular buffer
					const timelineStream = newState.chartItems[key];
					if (!timelineStream) {
						console.error(
							'Timeline stream not found but was added. This error should never occur.'
						);
						continue;
					}

					newState.chartItems[key] = addFrameToCircularBuffer(
						timelineStream,
						newTimelineEntry,
						newTimelineEntry.x,
						(frame) => frame.x
					);
				}

				return newState;
			});
		}
	};
};

type SensorStreamStore = ReturnType<typeof createSensorStreamStore>;

export { createSensorStreamStore };
export type { SensorStreamStore, SensorStreamState, TimestampedSensorOutput };
