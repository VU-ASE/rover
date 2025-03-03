<script lang="ts">
	import {
		PipelineApi,
		ServicesApi,
		type FullyQualifiedService,
		type PipelineGet200ResponseEnabledInnerProcess,
		type ServicesAuthorServiceVersionGet200Response
	} from '$lib/openapi';
	import type { PipelineNodeData } from './type';
	import { config } from '$lib/config';
	import { useQuery } from '@sveltestack/svelte-query';
	import stripAnsi from 'strip-ansi';
	import { Circle } from 'svelte-loading-spinners';
	import colors from 'tailwindcss/colors';
	import { errorToText } from '$lib/errors';

	export let tabSet: number;
	export let fq: FullyQualifiedService;
	export let process: PipelineGet200ResponseEnabledInnerProcess | undefined;

	let showOnlyLastRun = true;
	$: showOnlyLastRun;

	const logsQuery = useQuery(
		['logs', fq],
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);

			const logs = await papi.logsAuthorNameVersionGet(fq.author, fq.name, fq.version, 200);
			return logs.data.map((line) => stripAnsi(line));
		},
		{
			staleTime: 3
		}
	);

	const serviceQuery = useQuery(
		['service', fq],
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const sapi = new ServicesApi(config.roverd.api);

			const service = await sapi.servicesAuthorServiceVersionGet(fq.author, fq.name, fq.version);
			return service.data;
		},
		{
			staleTime: 3
		}
	);

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

<div class="px-4">
	{#if tabSet < 3}
		{#if $serviceQuery.data}
			{#if tabSet === 0}
				{#if $serviceQuery.data.inputs.length === 0}
					<p class="text-secondary-700">This service does not depend on any other services.</p>
				{:else}
					<div class="flex flex-col gap-2">
						{#each $serviceQuery.data.inputs as input}
							<div class="flex flex-col gap-1">
								<p class=" text-secondary-400">
									Depends on data from service <span class="font-mono text-success-400">
										{input.service}
									</span>
								</p>

								{#each input.streams as stream}
									<p class="pl-8 text-secondary-400">
										-> This service should expose the <span class="font-mono text-success-400">
											{stream}</span
										> output
									</p>
								{/each}
							</div>
						{/each}
					</div>
				{/if}
			{:else if tabSet === 1}
				{#if $serviceQuery.data.outputs.length === 0}
					<p class="text-secondary-700">This service does not produce data for other services.</p>
				{:else}
					<div class="flex flex-col gap-2">
						{#each $serviceQuery.data.outputs as output}
							<div class="flex flex-col gap-1">
								<p class=" text-secondary-400">
									Exposes the <span class="font-mono text-pink-400">
										{output}
									</span> stream
								</p>
								<p class="text-sm text-secondary-800">
									Add the following code to your service.yaml to read from this stream:
								</p>
								<div class="code variant-ghost-secondary p-2 px-4 mt-1">
									<p class="text-secondary-200">
										inputs:<br />
										&nbsp;&nbsp;- service: {fq.name}<br />
										&nbsp;&nbsp;&nbsp;&nbsp;streams:<br />
										&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;- {output}<br />
									</p>
								</div>
							</div>
						{/each}
					</div>
				{/if}
			{:else if tabSet === 2}
				{#if $serviceQuery.data.configuration.length === 0}
					<p class="text-secondary-700">This service does not expose any configurable options.</p>
				{:else}
					<div class="flex flex-col gap-2">
						<!-- Responsive Container (recommended) -->
						<div class="table-container">
							<!-- Native Table Element -->
							<table class="table table-hover">
								<tbody>
									{#each $serviceQuery.data.configuration as config}
										<tr>
											<td class="font-mono text-secondary-600 whitespace-nowrap">
												{config.name}
											</td>
											<td>
												{#if config.type === 'number'}
													<span class="badge variant-ghost-tertiary">number</span>
												{:else if config.type === 'string'}
													<span class="badge variant-ghost-warning">string</span>
												{:else}
													<span class="badge variant-ghost-secondary">unknown</span>
												{/if}
											</td>
											<td>
												{#if config.tunable}
													<span class="badge variant-glass-success">tunable</span>
												{:else}
													<span class="badge variant-glass-secondary">not tunable</span>
												{/if}
											</td>
											<td>
												{config.value}
											</td>
										</tr>
									{/each}
								</tbody>
							</table>
						</div>
					</div>
				{/if}
			{/if}
		{:else if $serviceQuery.isError}
			<div class=" text-red-500 font-mono whitespace-pre-line">
				{errorToText($serviceQuery.error)}
			</div>
		{:else if $serviceQuery.isLoading}
			<div class="flex flex-row items-center gap-2">
				<Circle color={colors.gray[500]} size="20" />
				<p class="text-gray-500">Loading...</p>
			</div>
		{/if}
	{:else if tabSet >= 3}
		{#if $logsQuery.data}
			{#if !showOnlyLastRun}
				<div class="flex flex-row gap-2 items-center mb-2 text-secondary-700">
					<p>These are the logs from all previous runs.</p>
					<button
						class="text-primary-500"
						on:click={() => {
							showOnlyLastRun = true;
						}}>Show only last run</button
					>
				</div>
			{:else}
				<div class="flex flex-row gap-2 items-center mb-2 text-secondary-700">
					<p>These are the logs from last run.</p>
					<button
						class="text-primary-500"
						on:click={() => {
							showOnlyLastRun = false;
						}}>Show all logs</button
					>
				</div>
			{/if}
			<div class=" text-gray-100 font-mono h-full text-sm overflow-y-auto mb-2">
				{#each $logsQuery.data.filter((line, index) => {
					if (showOnlyLastRun) {
						// Find the last line that includes "roverd spawned"
						const findIndex = $logsQuery.data.findLastIndex( (line) => line.includes('roverd spawned') );
						return index >= findIndex;
					}
					return true;
				}) as log}
					<p>{@html highlight(log)}</p>
				{/each}
			</div>
		{:else if $logsQuery.isError}
			<div class=" text-red-500 font-mono whitespace-pre-line">
				{errorToText($logsQuery.error)}
			</div>
		{:else if $logsQuery.isLoading}
			<div class="flex flex-row items-center gap-2">
				<Circle color={colors.gray[500]} size="20" />
				<p class="text-gray-500">Loading...</p>
			</div>
		{/if}
	{:else if tabSet === 2}
		{#if process}
			{process.cpu}
		{:else}
			<p class="text-gray-500">No performance metrics available</p>
		{/if}
	{/if}
</div>
