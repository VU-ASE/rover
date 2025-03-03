// import { getToastStore } from '@skeletonlabs/skeleton';

/**
 * Utility functions to report errors or successes to the user and UI, without bothering with styling in the compute layer
 */

const notify = (type: 'success' | 'error' | 'info', message: string) => {
	// const toastStore = getToastStore();

	if (type === 'error') {
		console.error(message);
	} else {
		console.log(type, message);
	}

	// toastStore.trigger({
	// 	// title: title ?? type.charAt(0).toUpperCase() + type.slice(1),
	// 	message: message
	// 	// duration: 10000,
	// 	// placement: "bottom-right",
	// 	// theme: "dark",
	// 	// type: type,
	// });
};

export { notify };
