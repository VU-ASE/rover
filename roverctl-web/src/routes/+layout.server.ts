import { redirect } from '@sveltejs/kit';

export function load({ cookies, url }) {
	const hasVisited = cookies.get('visited');

	if (!hasVisited && url.pathname !== '/dangers') {
		throw redirect(302, '/dangers'); // Redirect to the specific page
	}

	return {}; // Continue loading the layout normally
}
