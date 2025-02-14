import { persisted } from 'svelte-persisted-store';

/**
 * Can identify roverd by either its IP address or the rover ID (which will be used to compute the IP address)
 */
type RoverdAddress = {
	type: 'ip';
	address: string;
};
type RoverdId = {
	type: 'id';
	id: number;
};
type RoverdLocation = RoverdAddress | RoverdId;

/**
 * Authentication store
 */

type AuthStore = {
	roverdLocation: RoverdLocation;
	username: string;
	password: string;
	enableRoverd: boolean;
	passthroughUrl: string | null;
};

const authStore = persisted('authStore', <AuthStore>{
	roverdLocation: { type: 'id', id: 0 },
	username: '',
	password: '',
	enableRoverd: false,
	passthroughUrl: ''
});

/**
 * Utilities
 */

const validAuthStore = (store: AuthStore): boolean => {
	let valid = !!store.passthroughUrl && store.passthroughUrl.length > 0;

	if (store.enableRoverd) {
		valid =
			valid &&
			store.roverdLocation !== undefined &&
			store.username.length > 0 &&
			store.password.length > 0;

		if (store.roverdLocation.type === 'id') {
			return store.roverdLocation.id > 0 && valid;
		} else {
			return store.roverdLocation.address.length > 0 && valid;
		}
	}

	return valid;
};

const getRoverdBaseUrl = (store: AuthStore): string => {
	if (store.enableRoverd) {
		if (store.roverdLocation.type === 'id') {
			return `http://192.168.0.${100 + store.roverdLocation.id}`;
		} else {
			return `http://${store.roverdLocation.address}`;
		}
	} else {
		return '';
	}
};

export { authStore, validAuthStore, getRoverdBaseUrl };
