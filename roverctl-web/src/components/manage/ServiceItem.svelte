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

	export let tabSet: number;
	export let fq: FullyQualifiedService;
	export let process: PipelineGet200ResponseEnabledInnerProcess | undefined;

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
	{#if tabSet === 0}
		{#if $serviceQuery.data}
			<div class="flex flex-col gap-2">
				{#if fq.author.toLowerCase() === 'vu-ase'}
					<div class=" text-blue-500 border-l-4 border-blue-500 pl-2">
						<p>
							This service is developed by the VU-ASE team. <br />You can find its source code
							<a
								class="text-blue-300"
								href={`
                                    https://github.com/vu-ase/${fq.name}
                                `}>here</a
							>.
						</p>
					</div>
				{/if}
				<div class="grid grid-cols-2 gap-4">
					<div>
						<h1>Inputs</h1>
						{#if $serviceQuery.data.inputs.length === 0}
							<p class="text-gray-500">This service does not depend on any inputs</p>
						{:else}
							{#each $serviceQuery.data.inputs as input}
								<p class="pl-2 font-mono">
									- {input.service}
								</p>

								{#each input.streams as stream}
									<p class="pl-8 font-mono">- {stream}</p>
								{/each}
							{/each}
						{/if}
					</div>
					<div>
						<h1>Outputs</h1>
						{#if $serviceQuery.data.outputs.length === 0}
							<p class="text-gray-500">This service does not produce any outputs</p>
						{:else}
							{#each $serviceQuery.data.outputs as output}
								<p class="pl-2 font-mono">
									- {output}
								</p>
							{/each}
						{/if}
					</div>
				</div>
				{#if fq.as}
					<div>
						<h1>Impersonation</h1>
						<p>
							This service impersonates the <span class="font-mono">{fq.as || 'controller'}</span> service
						</p>
					</div>
				{/if}

				{#if $serviceQuery.data.configuration.length > 0}
					<div>
						<h1>Configurable options</h1>
						<p>
							The following options can be modified through <span class="font-mono"
								>service.yaml</span
							>
						</p>
						{#each $serviceQuery.data.configuration as config}
							<p class="pl-2 font-mono">
								- {config.name}
							</p>
						{/each}
					</div>
				{/if}
			</div>
		{:else if $serviceQuery.isError}
			<p class="text-red-500">{$serviceQuery.error}</p>
		{:else if $serviceQuery.isLoading}
			<div class="flex flex-row items-center gap-2">
				<Circle color={colors.gray[500]} size="20" />
				<p class="text-gray-500">Loading...</p>
			</div>
		{/if}
	{:else if tabSet === 1}
		{#if $logsQuery.data}
			<div class=" text-gray-100 font-mono h-full text-sm overflow-y-auto mb-2">
				{#each $logsQuery.data as log}
					<p>{@html highlight(log)}</p>
				{/each}
			</div>
		{:else if $logsQuery.isError}
			<p class="text-red-500">{$logsQuery.error}</p>
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
