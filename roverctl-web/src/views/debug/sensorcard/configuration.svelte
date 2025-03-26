<script lang="ts">
	import type { SensorStreamStore } from '$lib/store/sensorstream';
	import { AccordionItem } from '@skeletonlabs/skeleton';
	import Chart from './chart.svelte';
	import ChartIcon from '~icons/uil/chart-line';
	import WarningIcon from '~icons/si/warning-fill';
	import {
		Configuration,
		PipelineApi,
		ServicesApi,
		type ServicesAuthorServiceVersionGet200Response
	} from '$lib/openapi';
	import stripAnsi from 'strip-ansi';
	import { onDestroy, onMount } from 'svelte';
	import { createRemoteResult } from '$lib/utils/remoteResult';
	import { Circle } from 'svelte-loading-spinners';
	import { styles } from '$lib/utils/styles';
	import { createMapStore } from '$lib/utils/map';
	import { sendTuningState } from '$lib/events/tuning';
	import { globalStore } from '$lib/store';
	import { connectionStore } from '$lib/store/connection';
	import { config } from '$lib/config';
	import { useQuery } from '@sveltestack/svelte-query';
	import colors from 'tailwindcss/colors';
	import { errorToText } from '$lib/errors';
	import SendIcon from '~icons/ic/baseline-send';

	import type { ServiceStore } from '$lib/store/service';

	export let serviceStore: ServiceStore;

	// Data structure to keep track of the tuning data already sent out
	const sentData = createMapStore<string, string | number>();
	const tuningData = createMapStore<string, string | number>();

	const serviceQuery = useQuery(['configuration', $serviceStore.realName], async () => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const papi = new PipelineApi(config.roverd.api);
		const sapi = new ServicesApi(config.roverd.api);

		const s = await papi.pipelineGet();
		const e = s.data.enabled.find((e) => e.service.fq.name === $serviceStore.realName);
		if (!e) {
			throw new Error('Service not found');
		}

		const sd = await sapi.servicesAuthorServiceVersionGet(
			e?.service.fq.author,
			e?.service.fq.name,
			e?.service.fq.version
		);
		return sd.data;
	});

	const sendTuningData = () => {
		const service = $serviceQuery.data;
		if (!service) {
			return;
		}

		// Get all non-empty tuning data
		const data = Array.from($tuningData.entries()).filter(([, value]) => value !== undefined);

		// Send the data to passthrough
		if ($connectionStore.dataChannel) {
			sendTuningState($connectionStore.dataChannel, {
				timestamp: Date.now(),
				dynamicParameters: data.map(([key, value]) => {
					// Try to find this in the fetched configuration
					const option = service.configuration.find((option) => option.name === key);
					if (!option) {
						return {};
					}

					if (option.type === 'number') {
						return {
							number: {
								key: key,
								value: typeof value === 'string' ? parseFloat(value) : value
							}
						};
					} else {
						return {
							string: {
								key: key,
								value: value.toString()
							}
						};
					}
				})
			});

			// If no error, add all the sent values to the sentData store
			for (const [key, value] of data) {
				sentData.add(key, value);
			}
		}
	};
</script>

<div>
	{#if $serviceQuery.data}
		<div class="grid grid-cols-1 lg:grid-cols-3 gap-2 mb-4">
			{#each $serviceQuery.data.configuration as option}
				<div>
					<label class="label flex flex-col items-start w-full">
						<span class="text-secondary-700 font-mono">{option.name} </span>
						<div class="flex flex-row items-stretch w-full">
							{#if !option.tunable}
								<div class="bg-secondary-400 w-1"></div>
							{:else if $sentData.get(option.name) !== $tuningData.get(option.name)}
								<div class="bg-orange-400 w-1"></div>
							{:else}
								<div class="bg-green-400 w-1"></div>
							{/if}

							<input
								class="input w-full"
								type="text"
								placeholder={option.value.toString()}
								disabled={!option.tunable}
								on:keyup={(e) => {
									if (e.target) {
										// @ts-ignore we know there is a value
										tuningData.add(option.name, e.target.value);
									}
								}}
							/>
						</div>
					</label>
				</div>
			{/each}
		</div>
		<button
			on:click={sendTuningData}
			class="flex flex-row gap-4 btn variant-ghost-primary text-primary-500 w-full"
		>
			Send tuning data

			<SendIcon />
		</button>
	{:else if $serviceQuery.isError}
		<div class="flex flex-col gap-2 mt-0 px-2 text-start">
			<p class="text-error-400">Could not fetch configuration</p>
			<div class="card p-2 px-4 text-red-500 font-mono whitespace-pre-line">
				{errorToText($serviceQuery.error)}
			</div>
		</div>
	{:else}
		<div class="w-full flex flex-row gap-2 items-center text-secondary-700">
			<Circle size="10" color={colors.zinc[400]} />
			<p>fetching configuration</p>
		</div>
	{/if}
</div>
