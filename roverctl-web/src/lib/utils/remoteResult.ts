import { get, writable, type Writable } from 'svelte/store';

type ResultStatus = 'loading' | 'success' | 'error' | 'reloading' | 'idle';

interface RemoteResultStore<T> extends Writable<RemoteResultState<T>> {
	start: () => void;
	success: (data: T) => void;
	errorOccurred: (error: string) => void;
	isLoading: () => boolean;
	isReloading: () => boolean;
	isError: () => boolean;
	isSuccess: () => boolean;
	hasData: () => boolean;
}

interface RemoteResultState<T> {
	status: ResultStatus;
	data: T | undefined;
	error: string | undefined;
}

export function createRemoteResult<T>(): RemoteResultStore<T> {
	// Initial state
	const initialState: RemoteResultState<T> = {
		status: 'idle',
		data: undefined,
		error: undefined
	};

	// Internal writable store
	const { subscribe, set, update } = writable<RemoteResultState<T>>(initialState);

	// Custom methods
	return {
		subscribe,
		set,
		update,
		start() {
			update((state) => ({
				...state,
				status: state.data ? 'reloading' : 'loading'
			}));
		},
		success(data: T) {
			update(() => ({
				status: 'success',
				data,
				error: undefined
			}));
		},
		errorOccurred(error: string) {
			update(() => ({
				status: 'error',
				data: undefined,
				error
			}));
		},
		isLoading() {
			const state = get(this);
			return state.status === 'loading' || state.status === 'reloading';
		},
		isReloading() {
			const state = get(this);
			return state.status === 'reloading';
		},
		isError() {
			const state = get(this);
			return state.status === 'error';
		},
		isSuccess() {
			const state = get(this);
			return state.status === 'success';
		},
		hasData() {
			const state = get(this);
			return state.data !== undefined;
		}
	};
}
