/**
 * Utilities and types for all incoming communication over the control channel
 */

import { globalStore } from '$lib/store';
import { ConnectionState, ControlError } from 'rovercom/gen/control/control';

const handleIncomingControlMessage = async (e: MessageEvent<ArrayBuffer>) => {
	globalStore.setCarConnected(true);

	// Parse the control message as a ConnectionState first
	const controlMessage = ConnectionState.decode(new Uint8Array(e.data));
	if (controlMessage.client) {
		if (controlMessage.client === 'car') {
			console.log('Car connection state changed connected = ', controlMessage.connected);
			globalStore.setCarConnected(controlMessage.connected);
			// Pause the stream if the car disconnected
			if (!controlMessage.connected) {
				globalStore.pauseStream();
			}
		} else {
			console.log(
				'Unknown client',
				controlMessage.client,
				'connected = ',
				controlMessage.connected
			);
		}
		return;
	}

	// Try to parse the control message as a ControlError
	const error = ControlError.decode(new Uint8Array(e.data));
	if (error.message) {
		console.error('Control error', error.message);
		return;
	}

	console.error('Unknown control message', controlMessage);
};

export { handleIncomingControlMessage };
