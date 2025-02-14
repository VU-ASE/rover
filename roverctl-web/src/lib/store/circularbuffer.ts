const CIRCULAR_BUFFER_SIZE = 3000;

// Types and utilities for circular buffers that use timestamps to find the right item
type CircularBuffer<T> = {
	buffer: T[];
	index: number;
	map: Map<number, T>; // timestamp -> T, for quick lookups based on timestamp
};

const createCircularBuffer = <T>(): CircularBuffer<T> => {
	return {
		buffer: new Array(CIRCULAR_BUFFER_SIZE),
		index: -1,
		map: new Map<number, T>()
	};
};

const addFrameToCircularBuffer = <T>(
	buffer: CircularBuffer<T>,
	frame: T,
	timestamp: number,
	// This function is used to get the timestamp from the frame dynamically (so that it works with any frame type T)
	getTimestamp: (frame: T) => number
) => {
	buffer.index = (buffer.index + 1) % buffer.buffer.length;
	buffer.buffer[buffer.index] = frame;
	buffer.map.set(timestamp, frame);
	// did the map grow too large?
	if (buffer.map.size > buffer.buffer.length) {
		// remove the oldest entry
		const oldestIndex = buffer.index + 1;
		const oldest = buffer.buffer[oldestIndex % buffer.buffer.length];
		if (oldest) {
			buffer.map.delete(getTimestamp(oldest));
		}
	}
	return buffer;
};

const getLatestNFrames = <T>(buffer: CircularBuffer<T>, n: number) => {
	const result = [];
	for (let i = 0; i < n; i++) {
		let index = buffer.index - i;
		if (index < 0) {
			index += buffer.buffer.length;
		}
		const entry = buffer.buffer[index % buffer.buffer.length];
		if (entry) {
			result.push(entry);
		}
	}
	return result;
};

const getLatestFrame = <T>(buffer: CircularBuffer<T>) => {
	return buffer.buffer[buffer.index];
};

// If you want the value at a specific timestamp, you can use this function,
// it will return the exact frame at that timestamp, or the closest before it
const getClosestTimestamp = <T>(buffer: CircularBuffer<T>, desiredTimestamp: number) => {
	if (buffer.buffer.length <= 0) {
		return null;
	}

	// Find the exact entry
	const entry = buffer.map.get(desiredTimestamp);
	if (entry) {
		return entry;
	}

	// Find the closest entry before the desired timestamp, max 150 milliseconds back
	for (let i = 0; i <= 5000; i++) {
		const entry = buffer.map.get(desiredTimestamp - i);
		if (entry) {
			return entry;
		}
	}

	return null;
};

export {
	createCircularBuffer,
	addFrameToCircularBuffer,
	getClosestTimestamp,
	getLatestFrame,
	getLatestNFrames,
	CIRCULAR_BUFFER_SIZE
};
export type { CircularBuffer };
