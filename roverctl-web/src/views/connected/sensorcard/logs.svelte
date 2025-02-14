<script lang="ts">
	import WarningIcon from '~icons/si/warning-fill';
	import { authStore, getRoverdBaseUrl } from '$lib/store/auth';
	import { Configuration, PipelineApi, ServicesApi } from '$lib/openapi';
	import stripAnsi from 'strip-ansi';
	import { onDestroy, onMount } from 'svelte';
	import { createRemoteResult } from '$lib/utils/remoteResult';
	import { Circle } from 'svelte-loading-spinners';
	import { styles } from '$lib/utils/styles';

	export let serviceName: string;

	// Create an API from the auth store
	const config = new Configuration({
		basePath: getRoverdBaseUrl($authStore),
		username: $authStore.username,
		password: $authStore.password
	});
	const papi = new PipelineApi(config);

	const logsQuery = createRemoteResult<string[]>();
	const fetchLogs = async () => {
		if (!$authStore.enableRoverd) return;

		logsQuery.start();
		// First try to find the service FQN
		try {
			const s = await papi.pipelineGet();
			const e = s.data.enabled.find((e) => e.service.name === serviceName);
			if (!e) {
				throw new Error('Service not found');
			}

			const l = await papi.logsAuthorNameVersionGet(
				e?.service.author,
				e?.service.name,
				e?.service.version
			);
			logsQuery.success(l.data.map((line) => stripAnsi(line)));
		} catch (err) {
			if (err instanceof Error) {
				logsQuery.errorOccurred(err.message);
			} else {
				logsQuery.errorOccurred('An unknown error occurred');
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

		fetchLogs();
	}

	onMount(() => {
		fetchLogs();
		intervalId = setInterval(refetch, 1000);
	});

	onDestroy(() => {
		clearInterval(intervalId);
	});

	const highlight = (line: string): string => {
		const keywords: Record<string, string> = {
			error: 'text-red-500',
			err: 'text-red-500',
			warning: 'text-yellow-400',
			warn: 'text-yellow-400',
			success: 'text-green-400',
			suc: 'text-green-400',
			info: 'text-blue-400',
			inf: 'text-blue-400',
			debug: 'text-purple-400',
			db: 'text-purple-400',
			wrn: 'text-yellow-400'
		};

		const regex = new RegExp(`\\b(${Object.keys(keywords).join('|')})\\b`, 'gi');

		return line.replace(regex, (match: string) => {
			const colorClass = keywords[match.toLowerCase()]; // Convert match to lowercase
			if (colorClass) {
				return `<span class="${colorClass}">${match}</span>`;
			}
			return match; // In case no match is found (unlikely, but TypeScript-safe)
		});
	};
</script>

{#if !$authStore.enableRoverd}
	<div class="flex space-x-2 items-center">
		<WarningIcon class="text-gray-400 text-xl" />
		<p class="text-gray-400 text-sm">
			Log streaming depends on roverd metadata, which was disabled by you.
		</p>
	</div>
{:else}
	<div>
		<div class="flex flex-row align-center justify-between gap-x-2 mb-2">
			<label class="flex items-center justify-items-start space-x-2 w-full">
				<input class="checkbox" type="checkbox" bind:checked={refetchPeriodically} />
				<p>Refetch logs automatically</p>
			</label>
			{#if $logsQuery.status === 'reloading' || $logsQuery.status === 'loading'}
				<button type="button" disabled class="btn btn-sm variant-filled-secondary"
					>Loading...</button
				>
			{:else}
				<button on:click={fetchLogs} type="button" class="btn btn-sm variant-filled-secondary"
					>Refetch</button
				>
			{/if}
		</div>
		{#if $logsQuery.data}
			<div class="bg-gray-700 text-gray-100 font-mono text-sm p-4 h-96 overflow-y-auto mb-2">
				{#each $logsQuery.data as log}
					<p>{@html highlight(log)}</p>
				{/each}
			</div>
		{:else if $logsQuery.status === 'loading'}
			<div
				class="bg-gray-700 text-gray-100 text-sm p-4 h-96 overflow-y-auto mb-2 flex flex-col gap-y-4 items-center justify-center"
			>
				<Circle size="20" color={styles.colors.primary} unit="px" duration="1s" />
				<div class="text-center px-4">
					<h2 class="text-gray-200">Loading logs</h2>
				</div>
			</div>
		{:else if $logsQuery.status === 'error'}
			<div
				class="bg-gray-700 text-gray-100 text-sm p-4 h-96 overflow-y-auto mb-2 flex flex-col gap-y-4 items-center justify-center"
			>
				<WarningIcon class="text-red-500 text-xl" />
				<div class="text-center px-4">
					<h2 class="text-gray-200">Could not fetch logs</h2>
					<p class="text-gray-400 text-sm">{$logsQuery.error}</p>
				</div>
			</div>
		{/if}
	</div>
{/if}
