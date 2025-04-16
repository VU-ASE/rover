/**
 * This utility helps to determine the type of sensor data that was sent, by trying to decode a TimestampedSensorOutput
 */

import type { TimestampedSensorOutput } from '$lib/store/sensorstream';

// import type { TimestampedSensorOutput } from '$lib/new/state/sensorstream';
// import type { SensorType } from '$lib/ui/sensorfeed/types';

const trySensorType = (sensorData: TimestampedSensorOutput): string | undefined => {
	// Is this a video?
	try {
		const cameraData = sensorData.sensorData.cameraOutput;
		if (cameraData && (cameraData.debugFrame || cameraData.trajectory)) {
			return 'video';
		}
	} catch (e) {
		// Not a video
	}

	// Is this controller output?
	try {
		if (
			sensorData.sensorData.controllerOutput?.leftThrottle ||
			sensorData.sensorData.controllerOutput?.rightThrottle
		) {
			return 'controllerOutput';
		}
	} catch (e) {
		// Not a controller output
	}

	// Is this distance output?
	try {
		if (sensorData.sensorData.distanceOutput?.distance) {
			return 'distanceSensor';
		}
	} catch (e) {
		// Not a distance sensor
	}

	// Is this speed output?
	try {
		if (sensorData.sensorData.speedOutput?.rpm) {
			return 'speedSensor';
		}
	} catch (e) {
		// Not a speed sensor
	}

	return undefined;
};

// Because the type frame can (but shouldn't) change during a sensor stream, we can get the populated parameters from the frame dynamically
const getParametersFromFrame = (frame: TimestampedSensorOutput) => {
	// Get the entries (key,value) from all possible types
	// and filter out the ones that are not defined numbers
	// note: this is not recursive (todo: use lodash deep?)
	let entries: [string, number][] = [];

	if (frame.sensorData.controllerOutput) {
		entries = [
			...(Object.entries(frame.sensorData.controllerOutput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];
	}
	if (frame.sensorData.distanceOutput) {
		entries = [
			...(Object.entries(frame.sensorData.distanceOutput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];
	}
	if (frame.sensorData.rpmOuput) {
		entries = [
			...(Object.entries(frame.sensorData.rpmOuput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];
	}
	if (frame.sensorData.batteryOutput) {
		entries = [
			...(Object.entries(frame.sensorData.batteryOutput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];
	}
	if (frame.sensorData.energyOutput) {
		entries = [
			...(Object.entries(frame.sensorData.energyOutput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];
	}
	if (frame.sensorData.genericIntScalar) {
		entries = [[frame.sensorData.genericIntScalar.key, frame.sensorData.genericIntScalar.value]];
	}
	if (frame.sensorData.genericFloatScalar) {
		entries = [
			[frame.sensorData.genericFloatScalar.key, frame.sensorData.genericFloatScalar.value]
		];
	}
	if (frame.sensorData.imuOutput) {
		entries = [
			...(Object.entries(frame.sensorData.imuOutput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];

		// The IMU outputs are Vectors (see communication definitions, so we need to flatten them)
		// Add the magnemeter values
		entries.push(['magnetometerX', frame?.sensorData?.imuOutput?.magnetometer?.x ?? 0]);
		entries.push(['magnetometerY', frame?.sensorData?.imuOutput?.magnetometer?.y ?? 0]);
		entries.push(['magnetometerZ', frame?.sensorData?.imuOutput?.magnetometer?.z ?? 0]);
		// Add the gyroscope values
		entries.push(['gyroscopeX', frame?.sensorData?.imuOutput?.gyroscope?.x ?? 0]);
		entries.push(['gyroscopeY', frame?.sensorData?.imuOutput?.gyroscope?.y ?? 0]);
		entries.push(['gyroscopeZ', frame?.sensorData?.imuOutput?.gyroscope?.z ?? 0]);
		// Add the accelerometer values
		entries.push(['accelerometerX', frame?.sensorData?.imuOutput?.accelerometer?.x ?? 0]);
		entries.push(['accelerometerY', frame?.sensorData?.imuOutput?.accelerometer?.y ?? 0]);
		entries.push(['accelerometerZ', frame?.sensorData?.imuOutput?.accelerometer?.z ?? 0]);
		// Add the linear acceleration values
		entries.push([
			'linearAccelerationX',
			frame?.sensorData?.imuOutput?.linearAccelerometer?.x ?? 0
		]);
		entries.push([
			'linearAccelerationY',
			frame?.sensorData?.imuOutput?.linearAccelerometer?.y ?? 0
		]);
		entries.push([
			'linearAccelerationZ',
			frame?.sensorData?.imuOutput?.linearAccelerometer?.z ?? 0
		]);
		// velocity
		entries.push(['velocityX', frame?.sensorData?.imuOutput?.velocity?.x ?? 0]);
		entries.push(['velocityY', frame?.sensorData?.imuOutput?.velocity?.y ?? 0]);
		entries.push(['velocityZ', frame?.sensorData?.imuOutput?.velocity?.z ?? 0]);
		// Euler
		entries.push(['eulerX', frame?.sensorData?.imuOutput?.euler?.x ?? 0]);
		entries.push(['eulerY', frame?.sensorData?.imuOutput?.euler?.y ?? 0]);
		entries.push(['eulerZ', frame?.sensorData?.imuOutput?.euler?.z ?? 0]);
	}
	if (frame.sensorData.luxOutput) {
		entries = [
			...(Object.entries(frame.sensorData.luxOutput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];
	}
	if (frame.sensorData.speedOutput) {
		entries = [
			...(Object.entries(frame.sensorData.speedOutput).filter(
				([, value]) => typeof value === 'number'
			) as [string, number][])
		];
	}

	return entries;
};

// This function takes in all frames received so far (for a single sensor) and returns the datasets that can be plotted
// in an x,y format (x being the timestamp, y being the value)
const getDatasetsFromFrames = (
	received: TimestampedSensorOutput[]
): {
	datasets: {
		key: string;
		data: { x: number; y: number }[];
	}[];
	sensorId: number;
} | null => {
	if (received.length <= 0) {
		return null;
	}

	// Get the parameters from the first frame
	const firstFrame = received[0];
	if (!firstFrame) {
		return null;
	}
	const parameters = getParametersFromFrame(firstFrame);

	const datasets = parameters.map(([k]) => {
		const datasetData = received.map((frame) => {
			// merge to one object to take keys from
			const merged = {
				...frame.sensorData.controllerOutput,
				...frame.sensorData.distanceOutput,
				...frame.sensorData.speedOutput
			};

			const key = k as keyof typeof merged;

			return {
				x: frame.sensorData.timestamp,
				y: merged[key]
			};
		});

		return {
			key: k,
			data: datasetData as { x: number; y: number }[]
		};
	});

	return {
		datasets,
		sensorId: firstFrame.sensorData.sensorId
	};
};

export { trySensorType, getParametersFromFrame, getDatasetsFromFrames };
