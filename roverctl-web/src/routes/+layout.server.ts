import { redirect } from '@sveltejs/kit';

export function load({ cookies, url }) {
	const hasVisited = cookies.get('visited');
	const hasSeenShortcuts = cookies.get('shortcuts');

	if (!hasVisited && url.pathname !== '/dangers') {
		throw redirect(302, '/dangers');
	} else if (!hasSeenShortcuts && url.pathname !== '/shortcuts') {
		throw redirect(302, '/shortcuts');
	}

	return {}; // Continue loading the layout normally
}
