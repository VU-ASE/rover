<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WarningIcon from '~icons/ix/warning-filled';
	import { Circle } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { useQuery } from '@sveltestack/svelte-query';
	import { HealthApi } from '$lib/openapi';
	import colors from 'tailwindcss/colors';

	const statusQuery = useQuery('status', async () => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		// Fetch status
		const hapi = new HealthApi(config.roverd.api);
		const status = await hapi.statusGet();
		return status.data;
	});
</script>

<div class="h-full w-full flex justify-center items-center animate-fade-out-container">
	<div class="flex flex-col gap-4 items-center align-center">
		{#if config.success}
			{#if $statusQuery.isLoading}
				<div class="flex flex-col text-center">
					<h1>Rover</h1>
					<div class="flex flex-row items-center gap-2 text-zinc-400">
						<Circle size="10" color={colors.zinc[400]} />
						<p class="text-zinc-400">loading status</p>
					</div>
				</div>
			{:else if $statusQuery.isError}
				<div class="flex flex-col text-center">
					<h1>Rover</h1>
					<p class="text-error-400">unavailable ({$statusQuery.error})</p>
				</div>
			{:else if $statusQuery.data}
				<div class="flex flex-col text-center">
					<h1>
						Rover {$statusQuery.data.rover_id} ({$statusQuery.data.rover_name})
					</h1>
					<p class="text-success-400">available (roverd {$statusQuery.data.version})</p>
				</div>
			{/if}

			<a class="block card card-hover p-4 px-6 w-full" href="/manage">
				<div class="flex flex-row gap-4 w-full pr-4 items-center">
					<SteeringIcon class="text-2xl" />
					<div class="flex flex-col">
						<h1>Manage</h1>
						<p>Configure pipelines and services</p>
					</div>
				</div>
			</a>
			{#if config.passthrough}
				<a class="block card card-hover p-4 px-6 w-full" href="/manage">
					<div class="flex flex-row gap-4 w-full pr-4 items-center">
						<DebugIcon class="text-2xl" />
						<div class="flex flex-col">
							<h1>Tune and Debug</h1>
							<p>Modify service behavior on the fly</p>
						</div>
					</div>
				</a>
			{:else}
				<div class="block card variant-soft p-4 px-6 w-full">
					<div class="flex flex-row gap-4 w-full pr-4 items-center text-zinc-500">
						<DebugIcon class="text-2xl" />
						<div class="flex flex-col">
							<h1>Tune and Debug</h1>
							<p>Modify service behavior on the fly</p>
							<p class="text-sm text-error-300">Not available in current configuration</p>
						</div>
					</div>
				</div>
			{/if}
		{:else}
			<div class="block text-error-300">
				<div class="flex flex-row gap-4 w-full pr-4 items-start">
					<WarningIcon class="text-2xl" />
					<div class="flex flex-col">
						<h1>Not started properly</h1>
						<p class="text-sm">
							The configuration could not be loaded.<br />
							Are you sure that you instantiated roverctl with the correct parameters?
						</p>
						<div class="card variant-ghost-error mt-2 p-2 px-4 code">
							{config.error}
						</div>
					</div>
				</div>
			</div>
		{/if}
		<p class="text-gray-400 text-sm">
			<strong>roverctl-web</strong> v0.0.1
		</p>
	</div>
</div>
