import { writable } from 'svelte/store';

export function createMapStore<K, V>(initialValue: Map<K, V> = new Map()) {
	const { subscribe, set, update } = writable(initialValue);

	return {
		subscribe,
		set,
		update,
		// Add an entry to the map
		add(key: K, value: V) {
			update((map) => {
				map.set(key, value);
				return new Map(map); // Ensure reactivity by creating a new Map instance
			});
		},
		// Remove an entry from the map
		remove(key: K) {
			update((map) => {
				map.delete(key);
				return new Map(map); // Ensure reactivity by creating a new Map instance
			});
		},
		// Clear the entire map
		clear() {
			set(new Map());
		}
	};
}
