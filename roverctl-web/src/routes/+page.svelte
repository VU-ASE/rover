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
			{#if $statusQuery.isError}
				<div class="flex flex-col text-center">
					<p class="text-error-500">
						Could not fetch Rover status: ({$statusQuery.error})
					</p>
				</div>
			{:else if $statusQuery.data}
				<div class="flex flex-col items-center text-center gap-1">
					<h1 class="text-4xl text-secondary-700">
						Rover <span class="text-primary-500">{$statusQuery.data.rover_id}</span>
					</h1>
					<h2 class="text-xl text-secondary-700 font-mono">{$statusQuery.data.rover_name}</h2>
				</div>
			{:else}
				<div class="flex flex-col items-center text-center gap-1">
					<div class="flex flex-row items-center gap-4 text-zinc-400">
						<h1 class="text-4xl text-secondary-700">Rover</h1>
						<Circle size="30" color={'#0089d9'} />
					</div>
					<Circle size="20" color={colors.zinc[500]} />
				</div>
			{/if}

			<div class="flex flex-col gap-4">
				<a class="block card card-hover p-4 px-6 w-full" href="/manage">
					<div class="flex flex-row gap-4 w-full pr-4 items-center">
						<SteeringIcon class="text-4xl text-primary-500" />
						<div class="flex flex-col">
							<h1 class="">Manage</h1>
							<p class="text-secondary-800">Configure pipelines and services</p>
						</div>
					</div>
				</a>
				{#if config.passthrough}
					<a class="block card card-hover p-4 px-6 w-full" href="/debug">
						<div class="flex flex-row gap-4 w-full pr-4 items-center">
							<DebugIcon class="text-4xl text-primary-500" />
							<div class="flex flex-col">
								<h1>Tune and Debug</h1>
								<p class="text-secondary-800">Modify service behavior on the fly</p>
							</div>
						</div>
					</a>
				{:else}
					<div class="block card variant-soft p-4 px-6 w-full">
						<p class="text-sm text-secondary-600 mb-2">Currently unavailable</p>
						<div class="flex flex-row gap-4 w-full pr-4 items-center text-secondary-900">
							<DebugIcon class="text-4xl" />
							<div class="flex flex-col">
								<h1>Tune and Debug</h1>
								<p>Modify service behavior on the fly</p>
							</div>
						</div>
					</div>
				{/if}
			</div>
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
	</div>
</div>
