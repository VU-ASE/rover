import { writable } from 'svelte/store';

/**
 * This store describes *all* outgoing connections: both the roverd REST API as well as the passthrough webRTC server
 */

type ConnectionStatus =
	| {
			type: 'connecting';
	  }
	| {
			type: 'connected';
	  }
	| {
			type: 'disconnected';
			error: string;
	  };

type ConnectionStore = {
	roverd: {
		status: ConnectionStatus;
	};
	passthrough: {
		status: ConnectionStatus;
	};
};

const connectionStore = writable<ConnectionStore>({
	roverd: {
		status: { type: 'connecting' }
	},
	passthrough: {
		status: { type: 'connecting' }
	}
});

/**
 * Utilities
 */

const allLoaded = (store: ConnectionStore): boolean =>
	store.roverd.status.type !== 'connecting' && store.passthrough.status.type !== 'connecting';

const allConnected = (store: ConnectionStore): boolean =>
	store.roverd.status.type === 'connected' && store.passthrough.status.type === 'connected';

const allDisconnected = (store: ConnectionStore): boolean =>
	store.roverd.status.type === 'disconnected' && store.passthrough.status.type === 'disconnected';

const someDisconnected = (store: ConnectionStore): boolean =>
	store.roverd.status.type === 'disconnected' || store.passthrough.status.type === 'disconnected';

export { connectionStore, allLoaded, allConnected, allDisconnected, someDisconnected };
