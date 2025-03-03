<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WifiIcon from '~icons/material-symbols/wifi';
	import WifiOffIcon from '~icons/material-symbols/wifi-off';

	import { Circle } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { isError, useMutation, useQuery } from '@sveltestack/svelte-query';
	import { HealthApi } from '$lib/openapi';
	import colors from 'tailwindcss/colors';
	import Navbar from '../../components/Navbar.svelte';

	import { Modal } from '@skeletonlabs/skeleton';
	import ErrorOverlay from '../../components/ErrorOverlay.svelte';
	import Pipeline from '../../components/manage/Pipeline.svelte';
	import { onMount } from 'svelte';
	import { errorToText } from '$lib/errors';

	import WarningIcon from '~icons/ix/warning-filled';

	const shutdownQuery = useMutation('shutdown', async () => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		// Fetch status
		const hapi = new HealthApi(config.roverd.api);
		const status = await hapi.shutdownPost();
		return status.data;
	});

	onMount(() => {
		$shutdownQuery.mutate();
	});
</script>

<div class="flex h-full w-full items-center justify-center">
	{#if $shutdownQuery.isError}
		<div class="flex flex-col gap-2 mt-0 px-2 items-center">
			<WarningIcon class="text-error-500 text-4xl" />
			<p class="font-bold">Could not shut down Rover</p>
			<p class="text-error-500">Unplug power if you see the disconnect symbol.</p>
			<div class="card p-2 px-4 text-red-500 font-mono whitespace-pre-line">
				{errorToText($shutdownQuery.error)}
			</div>
		</div>
	{:else if $shutdownQuery.isSuccess}
		<div class="flex flex-col gap-2 mt-0 px-2 items-center">
			<WarningIcon class="text-warning-500 text-4xl" />
			<p class="font-bold">Rover shutdown successful</p>
			<p class="text-warning-500">Unplug power when you see the disconnect symbol</p>
		</div>
	{:else}
		<div class="flex flex-col gap-2 mt-0 px-2 items-center">
			<Circle size="20" color={colors.white} />
			<p>Shutting down Rover</p>
		</div>
	{/if}
</div>
