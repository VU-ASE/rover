/**
 * All state for the server connection is defined here
 * this is a thin wrapper over the base svelte store
 */

import { writable } from 'svelte/store';

type ConnectionState = {
	clientId: string;
	server: RTCPeerConnection | null;
	controlChannel: RTCDataChannel | null;
	dataChannel: RTCDataChannel | null;
	// Hold connection errors here
	error: string | null;
	isConnecting: boolean;
};

// See: https://svelte.dev/examples/custom-stores
const createConnectionState = () => {
	const initialState = {
		clientId: '',
		server: null,
		metaChannel: null,
		controlChannel: null,
		dataChannel: null,
		error: null,
		isConnecting: false
	};
	const { subscribe, update } = writable<ConnectionState>(initialState);

	return {
		// Required functions
		subscribe,
		update: (fn: (state: ConnectionState) => Partial<ConnectionState>) => {
			update((oldState) => {
				return {
					...oldState,
					...fn(oldState)
				};
			});
		},
		// Custom functions
		reset: () => {
			update(() => initialState);
		}
	};
};

export const connectionStore = createConnectionState();
export type { ConnectionState };
