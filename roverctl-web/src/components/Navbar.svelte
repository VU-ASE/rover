<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WifiIcon from '~icons/material-symbols/wifi';
	import WifiOffIcon from '~icons/material-symbols/wifi-off';

	import { Circle } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { useQuery } from '@sveltestack/svelte-query';
	import { HealthApi } from '$lib/openapi';
	import colors from 'tailwindcss/colors';

	export let page: 'manage' | 'debug' = 'manage';

	const statusQuery = useQuery(
		'status',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			// Fetch status
			const hapi = new HealthApi(config.roverd.api);
			const status = await hapi.statusGet();
			return status.data;
		},
		{
			staleTime: 10, // 10 seconds
			keepPreviousData: true
		}
	);

	const activeClass = (p: 'manage' | 'debug') => {
		return p === page ? 'card variant-soft-primary' : 'card variant-soft-surface';
	};
</script>

<div class="w-full flex flex-row p-2 justify-between card items-center">
	<div class="flex flex-row items-center gap-2">
		<!-- Want to show when (re)fetching as well -->
		{#if $statusQuery.isSuccess && $statusQuery.data && !$statusQuery.isFetching}
			<div
				class="card variant-soft-success p-1 px-2 h-full flex items-center justify-center text-lg"
			>
				<WifiIcon />
			</div>
		{:else if $statusQuery.isError}
			<div class="card variant-soft-error p-1 px-2 h-full flex items-center justify-center text-lg">
				<WifiOffIcon />
			</div>
		{:else}
			<div
				class="card variant-soft-primary p-1 px-2 h-full flex items-center justify-center text-lg"
			>
				<Circle size="20" color={colors.white} />
			</div>
		{/if}

		<!-- Want to show the (stale) name, even if the Rover is going offline -->
		{#if $statusQuery.data}
			<div class="flex flex-col text-secondary-600">
				<h1 class="text-sm font-mono">{$statusQuery.data.rover_name}</h1>
				<p class="text-xs">
					Rover <span class="text-primary-500">{$statusQuery.data.rover_id}</span>
				</p>
			</div>
		{:else if $statusQuery.isError}
			<div class="flex flex-col">
				<h1 class="text-sm">Unreachable</h1>
				<p class="text-xs">Could not probe</p>
			</div>
		{:else}
			<div class="flex flex-col">
				<h1 class="text-sm">Loading...</h1>
				<p class="text-xs">Loading...</p>
			</div>
		{/if}
	</div>

	<div class="flex flex-row items-center gap-2">
		<a href="/manage" class={`${activeClass('manage')} p-1 px-2 flex flex-row items-center gap-1`}>
			<SteeringIcon class="text-sm" />
			<p class="">manage</p>
		</a>
		<a href="/debug" class={`${activeClass('debug')} p-1 px-2 flex flex-row items-center gap-1`}>
			<DebugIcon class="text-sm" />
			<p class="">debug</p>
		</a>
	</div>
</div>
