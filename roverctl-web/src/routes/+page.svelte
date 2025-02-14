<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WarningIcon from '~icons/ix/warning-filled';

	import { config } from '$lib/config';
</script>

<div class="h-full w-full flex justify-center items-center animate-fade-out-container">
	<div class="flex flex-col gap-4 items-center align-center">
		{#if config.success}
			<a class="block card card-hover p-4 px-6 w-full" href="/manage">
				<div class="flex flex-row gap-4 w-full pr-4 items-center">
					<SteeringIcon class="text-2xl" />
					<div class="flex flex-col">
						<h1>Manage</h1>
						<p>Configure pipelines and services</p>
					</div>
				</div>
			</a>
			{#if config.config.passthrough}
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
				<div class="block card variant-soft p-4 px-6">
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
