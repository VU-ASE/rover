<script lang="ts">
	import '../app.postcss';
	import { AppShell } from '@skeletonlabs/skeleton';
	import { initializeStores, Toast } from '@skeletonlabs/skeleton';
	import { browser } from '$app/environment';
	import { ToastContainer, FlatToast } from 'svelte-toasts';
	import WarningIcon from '~icons/ix/warning-filled';

	initializeStores();
	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';
	import { config } from '$lib/config';
	import { onDestroy, onMount } from 'svelte';
	import { useEmergencyBrake } from '$lib/queries/pipeline';
	import KeyboardHandler from '../components/KeyboardHandler.svelte';
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
		<KeyboardHandler />
		{#if config.success}
			<slot />
		{:else}
			<div class="h-full w-full flex justify-center items-center animate-fade-out-container">
				<div class="flex flex-col gap-4 items-center align-center">
					<div class="block text-error-300">
						<div class="flex flex-row gap-4 w-full pr-4 items-start">
							<WarningIcon class="text-2xl" />
							<div class="flex flex-col">
								<h1>Not started properly</h1>
								<p class="text-sm">
									The configuration could not be loaded.<br />
									Are you sure that you instantiated roverctl with the correct parameters?
								</p>
								<div class="card variant-ghost-error mt-2 p-2 px-4 text-error-500 font-mono">
									{config.error}
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		{/if}
	</QueryClientProvider>
	<ToastContainer placement="bottom-right" let:data>
		<FlatToast {data} />
	</ToastContainer>

	<div class="absolute bottom-0 right-0 w-[25vw] h-[30vh] pr-10 z-[-1] opacity-10 flex items-end">
		<img src="/rover-top.svg" alt="ASE/Rover top view as background pattern" class="w-full" />
	</div>
</AppShell>
