<script lang="ts">
	import '../app.postcss';
	import { AppShell } from '@skeletonlabs/skeleton';
	import { initializeStores, Toast } from '@skeletonlabs/skeleton';
	import { browser } from '$app/environment';
	import { ToastContainer, FlatToast } from 'svelte-toasts';

	initializeStores();
	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';
	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser
			}
		}
	});
</script>

<Toast />
<AppShell>
	<title>roverctl-web</title>
	<svelte:fragment slot="header"></svelte:fragment>
	<QueryClientProvider client={queryClient}>
		<slot />
	</QueryClientProvider>
	<ToastContainer placement="bottom-right" let:data>
		<FlatToast {data} />
	</ToastContainer>

	<div class="absolute bottom-0 right-0 w-[25vw] h-[30vh] pr-10 z-[-1] opacity-10 flex items-end">
		<img src="/rover-top.svg" alt="ASE/Rover top view as background pattern" class="w-full" />
	</div>
</AppShell>
