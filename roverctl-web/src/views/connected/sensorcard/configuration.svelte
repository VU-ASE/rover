<script lang="ts">
	import type { SensorStreamStore } from '$lib/store/sensorstream';
	import { AccordionItem } from '@skeletonlabs/skeleton';
	import Chart from './chart.svelte';
	import ChartIcon from '~icons/uil/chart-line';
	import WarningIcon from '~icons/si/warning-fill';
	import { authStore, getRoverdBaseUrl } from '$lib/store/auth';
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

	export let serviceName: string;

	// Data structure to keep track of the tuning data already sent out
	const sentData = createMapStore<string, string | number>();
	const tuningData = createMapStore<string, string | number>();

	// Create an API from the auth store
	const config = new Configuration({
		basePath: getRoverdBaseUrl($authStore),
		username: $authStore.username,
		password: $authStore.password
	});
	const papi = new PipelineApi(config);
	const sapi = new ServicesApi(config);

	const serviceQuery = createRemoteResult<ServicesAuthorServiceVersionGet200Response>();
	const fetchConfiguration = async () => {
		if (!$authStore.enableRoverd) return;

		serviceQuery.start();
		// First try to find the service FQN
		try {
			const s = await papi.pipelineGet();
			const e = s.data.enabled.find((e) => e.service.name === serviceName);
			if (!e) {
				throw new Error('Service not found');
			}

			const sd = await sapi.servicesAuthorServiceVersionGet(
				e?.service.author,
				e?.service.name,
				e?.service.version
			);
			serviceQuery.success(sd.data);
		} catch (err) {
			if (err instanceof Error) {
				serviceQuery.errorOccurred(err.message);
			} else {
				serviceQuery.errorOccurred('An unknown error occurred');
			}
		}
	};

	let intervalId: NodeJS.Timeout;
	let refetchPeriodically = true;
	$: refetchPeriodically;

	function refetch() {
		if (!refetchPeriodically) {
			return;
		}

		fetchConfiguration();
	}

	onMount(() => {
		fetchConfiguration();
		intervalId = setInterval(refetch, 1000);
	});

	onDestroy(() => {
		clearInterval(intervalId);
	});

	const sendTuningData = () => {
		const config = $serviceQuery.data;
		if (!config) {
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
					const option = config.configuration.find((option) => option.name === key);
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

{#if !$authStore.enableRoverd}
	<div class="flex space-x-2 items-center">
		<WarningIcon class="text-gray-400 text-xl" />
		<p class="text-gray-400 text-sm">
			Dynamic configuration depends on roverd metadata, which was disabled by you.
		</p>
	</div>
{:else}
	<div>
		{#if $serviceQuery.data}
			<div class="grid grid-cols-1 lg:grid-cols-3 gap-2 mb-4">
				{#each $serviceQuery.data.configuration as option}
					<div>
						<label class="label flex flex-col items-start w-full">
							<span>{option.name} </span>
							<div class="flex flex-row items-stretch w-full">
								{#if !option.tunable}
									<div class="bg-gray-400 w-1"></div>
								{:else if $sentData.get(option.name) !== $tuningData.get(option.name)}
									<div class="bg-orange-200 w-1"></div>
								{:else}
									<div class="bg-green-200 w-1"></div>
								{/if}

								<input
									class="input w-full"
									type="text"
									placeholder={option.value.toString()}
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
			<button on:click={sendTuningData} class="btn btn-sm variant-filled-secondary">Send</button>
		{:else}
			<div>loading</div>
		{/if}
	</div>
{/if}
