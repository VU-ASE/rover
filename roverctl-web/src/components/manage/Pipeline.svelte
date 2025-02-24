<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WarningIcon from '~icons/ix/warning-filled';

	import WifiIcon from '~icons/material-symbols/wifi';
	import WifiOffIcon from '~icons/material-symbols/wifi-off';
	import StartIcon from '~icons/ic/round-play-circle';
	import StopIcon from '~icons/ic/round-stop-circle';
	import CheckIcon from '~icons/ic/sharp-check';
	import RemoveIcon from '~icons/ic/sharp-remove';
	import PlusIcon from '~icons/subway/add-1';

	import { useStore } from '@xyflow/svelte';
	import type { Edge, Node } from '@xyflow/svelte';
	import { Circle, DoubleBounce } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { isError, useMutation, useQuery, useQueryClient } from '@sveltestack/svelte-query';
	import {
		HealthApi,
		PipelineApi,
		ServicesApi,
		type FullyQualifiedService,
		type PipelineGet200ResponseEnabledInnerService,
		type ServicesAuthorServiceVersionGet200Response
	} from '$lib/openapi';
	import colors from 'tailwindcss/colors';
	import Navbar from '../../components/Navbar.svelte';
	import { Modal, Tab, TabGroup } from '@skeletonlabs/skeleton';
	import ErrorOverlay from '../../components/ErrorOverlay.svelte';
	import { writable } from 'svelte/store';
	import { SvelteFlow, Background, Controls, MarkerType } from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import ServiceNode from './ServiceNode.svelte';
	import dagre from 'dagre';
	import type { Writable } from 'svelte/store';
	import type { PipelineNode, PipelineNodeData } from './type';
	import { onMount } from 'svelte';
	import { color } from 'chart.js/helpers';
	import AutoFit from './AutoFit.svelte';
	import { SlideToggle } from '@skeletonlabs/skeleton';
	import ServiceItem from './ServiceItem.svelte';
	import { AxiosError } from 'axios';
	import { errorToText } from '$lib/errors';
	import InstallServiceModal from './InstallServiceModal.svelte';

	const queryClient = useQueryClient();

	// Flowchart data
	const nodeTypes = {
		service: ServiceNode
	};
	const nodes: Writable<PipelineNode[]> = writable([]);
	const edges: Writable<Edge[]> = writable([]);

	const edgesFromEnabledServices = (enabled: PipelineNodeData[]) => {
		// For each dependency, create an edge
		const newEdges: Edge[] = [];
		for (const e of enabled) {
			for (const input of e.service.inputs) {
				const source = input.service;
				const target = e.fq.name;
				console.log('trying to create edge between', source, target);

				for (const stream of input.streams) {
					const id = `${source}-${target}-${stream}`;

					// Check if the target service exists, and if the stream is present
					if (
						enabled.find((s) => s.fq.name === target) &&
						enabled.find((s) => s.fq.name === source)
					) {
						// Add if not already present
						if (!newEdges.find((edge) => edge.id === id)) {
							newEdges.push({
								id,
								source,
								target,
								animated: false
							});
						}
					}
				}
			}
		}

		return newEdges;
	};

	// Query enabled pipeline services once, and resolve their inputs/outputs to create edges
	const nodesQuery = useQuery(
		'pipelineNodes',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);
			const sapi = new ServicesApi(config.roverd.api);

			// Fetch enabled services in the pipeline
			const pipeline = await papi.pipelineGet();

			// For each enabled service, query its specific information (inputs, outputs, build time)
			const services: {
				service: ServicesAuthorServiceVersionGet200Response;
				fq: FullyQualifiedService;
			}[] = [];
			for (const e of pipeline.data.enabled) {
				const fq = e.service.fq;
				const service = await sapi.servicesAuthorServiceVersionGet(fq.author, fq.name, fq.version);
				if (service.data) {
					services.push({
						service: service.data,
						fq: fq
					});
				}
			}

			const newNodes: PipelineNode[] = pipeline.data.enabled.map((e) => {
				// Try to find the service information
				const service = services.find(
					(s) =>
						s.fq.name === e.service.fq.name &&
						s.fq.author === e.service.fq.author &&
						s.fq.version === e.service.fq.version
				);
				if (!service) {
					throw new Error(
						'Service ' +
							e.service.fq.name +
							' was enabled but not installed on the Rover (author: ' +
							e.service.fq.author +
							', version: ' +
							e.service.fq.version +
							')'
					);
				}

				return {
					// Services take the role of "as", when it is set
					id: e.service.fq.as || e.service.fq.name,
					position: { x: 0, y: 0 }, // required but later replaced using dagre
					type: 'service',
					// These width and height values are hardcoded in the custom node component as well, so they must match
					width: 200,
					height: 80,
					deletable: false,
					draggable: false,
					data: {
						fq: e.service.fq,
						service: service.service,
						process: e.process,
						onDelete: () => removeService(service.fq)
					}
				};
			});

			const newEdges = edgesFromEnabledServices(newNodes.map((n) => n.data));
			createAndSetGraph(newNodes, newEdges);
			return pipeline.data;
		},
		{
			enabled: true, // run once on mount
			refetchOnMount: false,
			staleTime: Infinity // no background refetch
		}
	);

	const createAndSetGraph = (newNodes: PipelineNode[], newEdges: Edge[]) => {
		// Create a daggerable graph to automatically layout the nodes
		const graph = new dagre.graphlib.Graph();
		graph.setGraph({ rankdir: 'LR' }); // Top-to-Bottom layout
		graph.setDefaultEdgeLabel(() => ({}));

		// Add nodes to graph
		newNodes.forEach((node) => graph.setNode(node.id, { width: node.width, height: node.height }));

		// Add edges to graph
		newEdges.forEach((edge) => graph.setEdge(edge.source, edge.target));

		// Compute layout
		dagre.layout(graph);

		// Apply computed positions
		const positionedNodes = newNodes.map((node) => ({
			...node,
			position: { x: graph.node(node.id).x, y: graph.node(node.id).y }
		}));

		nodes.set(positionedNodes);
		edges.set(newEdges);
	};

	const addService = (service: PipelineNodeData) => {
		$stopPipeline.mutate();

		// Check if service is already present
		if ($nodes.some((n) => n.data.fq.name === service.fq.name)) {
			return;
		}

		// Add the service to the pipeline
		const newNodes = [
			...$nodes,
			{
				id: service.fq.as || service.fq.name,
				position: { x: 0, y: 0 },
				type: 'service',
				width: 200,
				height: 80,
				deletable: false,
				draggable: false,
				data: {
					...service,
					onDelete: () => removeService(service.fq)
				}
			}
		];

		// Add edges from the new service
		const newEdges = edgesFromEnabledServices(newNodes.map((n) => n.data));

		createAndSetGraph(newNodes, newEdges);
	};

	const removeService = (fq: FullyQualifiedService) => {
		$stopPipeline.mutate();

		// Check if service is already present
		if (!$nodes.some((n) => n.data.fq.name === fq.name)) {
			return;
		}

		// Remove the service from the pipeline
		const newNodes = $nodes.filter((n) => n.data.fq.name !== fq.name);

		// Add edges from the new service
		const newEdges = edgesFromEnabledServices(newNodes.map((n) => n.data));

		createAndSetGraph(newNodes, newEdges);
	};

	onMount(() => {
		// Fetch the pipeline on mount
		$nodesQuery.refetch();
	});

	const pipelineQuery = useQuery(
		'pipeline',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);
			const sapi = new ServicesApi(config.roverd.api);

			// Fetch enabled services in the pipeline
			const pipeline = await papi.pipelineGet();
			return pipeline.data;
		},
		{
			keepPreviousData: false
		}
	);

	// Whether the "install service" modal is open
	$: serviceModalOpen = false;

	let selectedService: FullyQualifiedService | null = null;
	$: selectedService;

	const servicesQuery = useQuery(
		'availableServices',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const sapi = new ServicesApi(config.roverd.api);

			// Fetch all available services
			// todo: this can be optimized in the future, requesting only services for opened authors/services/versions
			const services = await sapi.fqnsGet();

			// For each fqn, get detailed information as necessary to add to the pipeline
			const detailedServices: PipelineNodeData[] = await Promise.all(
				services.data.map(async (fqn) => {
					const service = await sapi.servicesAuthorServiceVersionGet(
						fqn.author,
						fqn.name,
						fqn.version
					);
					return {
						fq: fqn,
						service: service.data
					};
				})
			);

			return detailedServices;
		},
		{
			enabled: serviceModalOpen, // run when service added
			refetchOnMount: true,
			staleTime: 100
		}
	);

	/**
	 * Sequential actions for starting the actual pipeline
	 */

	const stopPipeline = useMutation(
		'stopPipeline',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);
			// Best-effort based, ignore errors (i.e. if already stopped)
			try {
				const response = await papi.pipelineStopPost();
				return response.data;
			} catch {}
		},
		{
			// Invalidate the pipeline query regardless of mutation success or failure
			onSettled: () => {
				queryClient.invalidateQueries('pipeline');
				$startPipeline.reset();
				$buildService.reset();
				$savePipeline.reset();
			}
		}
	);

	const buildService = useMutation('buildService', async (fq: FullyQualifiedService) => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const sapi = new ServicesApi(config.roverd.api);
		const response = await sapi.servicesAuthorServiceVersionPost(fq.author, fq.name, fq.version);
		return response.data;
	});

	const savePipeline = useMutation('savePipeline', async (services: PipelineNodeData[]) => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const papi = new PipelineApi(config.roverd.api);
		const response = await papi.pipelinePost(
			services.map((s) => ({
				fq: s.fq
			}))
		);
		return response.data;
	});

	const startPipeline = useMutation(
		'startPipeline',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);
			const response = await papi.pipelineStartPost();
			return response.data;
		},
		{
			// Invalidate the pipeline query regardless of mutation success or failure
			onSettled: () => {
				queryClient.invalidateQueries('pipeline');
			}
		}
	);

	const startConfiguredPipeline = async () => {
		const services = $nodes;
		$stopPipeline.reset();
		$buildService.reset();
		$savePipeline.reset();
		$startPipeline.reset();
		selectedService = null;

		await $stopPipeline.mutateAsync();
		await Promise.all(services.map((n) => $buildService.mutateAsync(n.data.fq)));
		await $savePipeline.mutateAsync(services.map((n) => n.data));
		await $startPipeline.mutateAsync();
	};

	// Filter services from the service query if available

	let tabSet: number = 0;

	let pipelineStarting = false;
	$: pipelineStarting =
		$stopPipeline.isLoading ||
		$startPipeline.isLoading ||
		$buildService.isLoading ||
		$savePipeline.isLoading;
</script>

<div class="h-[90vh] sm:h-[30vh] overflow-hidden relative">
	<!-- Pipeline flowchart column -->
	<div class="md:col-span-2 lg:col-span-3 h-full overflow-hidden">
		<SvelteFlow
			{nodes}
			{edges}
			{nodeTypes}
			style="background-color: transparent;"
			proOptions={{ hideAttribution: true }}
			on:nodeDelete={({ detail }) => {
				console.log('delete', detail);
			}}
		>
			<!-- <Background bgColor={colors.slate[900]} patternColor={colors.slate[500]} /> -->
			<Background
				bgColor={'transparent'}
				class="bg-transparent"
				patternColor={colors.slate[500]}
				gap={20}
			/>
			<AutoFit />
		</SvelteFlow>

		<!-- Loading Overlay -->
		{#if $nodesQuery.isLoading}
			<div class="absolute inset-0 flex justify-center items-center">
				<div class="absolute inset-0 bg-secondary-700 bg-opacity-20"></div>
				<div class="flex flex-row items-center gap-2 text-secondary-200">
					<Circle size="20" color={colors.zinc[400]} />
					<p>fetching pipeline</p>
				</div>
			</div>
		{:else if $nodesQuery.isError}
			<div class="absolute inset-0 flex justify-center items-center">
				<div class="absolute inset-0 bg-error-300 bg-opacity-20"></div>
				<div class="flex flex-col text-center text-white items-center">
					<WarningIcon class="text-2xl relative mb-2" />
					<h1>Could not fetch pipeline</h1>
					<p>{$nodesQuery.error || 'An unknown error occurred'}</p>
				</div>
			</div>
		{/if}
	</div>
</div>

<div class="w-full px-4">
	{#if $pipelineQuery.data}
		{#if $pipelineQuery.data.status === 'started'}
			<div class="w-full card p-2 px-4">
				<div class="flex flex-row justify-between items-center w-full">
					<div class="flex flex-col">
						<div class="flex flex-row items-center gap-2">
							<DoubleBounce size="15" color={colors.green[400]} />
							<p class="text-white text-xl">
								Pipeline is <span class="text-green-500">running</span>
								{#if $stopPipeline.isLoading}
									, stopping...
								{/if}
							</p>
						</div>
					</div>

					<div class="flex flex-col">
						<button
							on:click={() => $stopPipeline.mutate()}
							type="button"
							class="btn text-orange-500"
							disabled={$stopPipeline.isLoading}
						>
							<StopIcon />
							<span>Stop execution</span>
						</button>
					</div>
				</div>
			</div>
		{:else if $nodes.length === 0}
			<div class="w-full card p-2 px-4">
				<div class="flex flex-row justify-between items-center w-full">
					<div class="flex flex-col">
						<div class="flex flex-row items-center gap-2">
							<p class="text-white text-xl">
								Pipeline is <span class="text-secondary-700">empty</span>
							</p>
						</div>
					</div>

					<div class="flex flex-col">
						<button
							on:click={startConfiguredPipeline}
							type="button"
							class="btn text-primary-500"
							disabled
						>
							<StartIcon />
							<span>Start execution</span>
						</button>
					</div>
				</div>
			</div>
		{:else if $pipelineQuery.data.status === 'startable' || ($pipelineQuery.data.status === 'empty' && $nodes.length > 0)}
			<div class="w-full card p-2 px-4">
				<div class="flex flex-row justify-between items-center w-full">
					<div class="flex flex-col">
						<div class="flex flex-row items-center gap-2">
							<p class="text-white text-xl">
								Pipeline is <span class="text-blue-400">startable</span>
							</p>
						</div>
					</div>

					<div class="flex flex-col">
						<button
							on:click={startConfiguredPipeline}
							type="button"
							class="btn text-primary-500"
							disabled={pipelineStarting}
						>
							<StartIcon />
							<span>Start execution</span>
						</button>
					</div>
				</div>
			</div>
		{:else}
			<div class="w-full card p-2 px-4">
				<div class="flex flex-row justify-between items-center w-full">
					<div class="flex flex-col">
						<div class="flex flex-row items-center gap-2">
							<DoubleBounce size="15" color={colors.green[400]} />
							<p class="text-white text-xl">
								Pipeline is <span class="text-orange-500">running</span>
								{#if $stopPipeline.isLoading}
									, stopping...
								{/if}
							</p>
						</div>
					</div>

					<div class="flex flex-col">
						<button
							on:click={() => $stopPipeline.mutate()}
							type="button"
							class="btn text-orange-500"
							disabled={$stopPipeline.isLoading}
						>
							<StopIcon />
							<span>Stop execution</span>
						</button>
					</div>
				</div>
			</div>
		{/if}
	{:else if $pipelineQuery.isError}
		<div class="w-full card p-2 px-4">
			<div class="flex flex-row justify-between items-center w-full">
				<div class="flex flex-col">
					<div class="flex flex-row items-center gap-2">
						<p class="text-white text-xl">
							<span class="text-error-500">Error</span> fetching pipeline
						</p>
					</div>
				</div>

				<div class="flex flex-col">
					<button
						on:click={startConfiguredPipeline}
						type="button"
						class="btn text-primary-500"
						disabled
					>
						<StartIcon />
						<span>Start execution</span>
					</button>
				</div>
			</div>
		</div>
	{:else}
		<div class="w-full card p-2 px-4">
			<div class="flex flex-row justify-between items-center w-full">
				<div class="flex flex-col">
					<div class="flex flex-row items-center gap-2">
						<Circle size="20" color={colors.gray[200]} />
						<p class="text-white text-xl">
							<span class="text-secondary-700">fetching</span> pipeline status
						</p>
					</div>
				</div>

				<div class="flex flex-col">
					<button
						on:click={startConfiguredPipeline}
						type="button"
						class="btn text-primary-500"
						disabled
					>
						<StartIcon />
						<span>Start execution</span>
					</button>
				</div>
			</div>
		</div>
	{/if}

	<div class="grid grid-cols-1 lg:grid-cols-5 h-[calc(70vh-8.5rem)] overflow-hidden gap-4 mt-4">
		<!-- Sidebar (1/5 width on large screens) -->
		<div class="flex flex-col h-full gap-4">
			<div class="card lg:col-span-1 overflow-y-auto flex flex-col h-full">
				{#if $servicesQuery.data && $servicesQuery.data.length > 0}
					{#each $servicesQuery.data as service}
						<button
							class={`p-2 px-4 
								${
									selectedService &&
									selectedService.name === service.fq.name &&
									selectedService.author === service.fq.author &&
									selectedService.version === service.fq.version
										? 'card variant-soft-primary'
										: ''
								}

						flex flex-row w-full gap-2 items-center text-left btn`}
							on:click={() => (selectedService = service.fq)}
						>
							{#if $nodes.some((node) => node.data.fq.name === service.fq.name && node.data.fq.author === service.fq.author && node.data.fq.version === service.fq.version)}
								<button
									class="w-6 h-6 card variant-outline-primary text-primary-400"
									on:click={() => removeService(service.fq)}
								>
									<CheckIcon />
								</button>
							{:else}
								<button
									on:click={() => addService(service)}
									class="w-6 h-6 card variant-outline-tertiary text-primary-400"
								>
								</button>
							{/if}

							<div class="flex flex-col w-full">
								<h1 class="text-sm text-secondary-700">{service.fq.author}</h1>
								<p class="text-md">
									{service.fq.name}
									<span class="text-secondary-800">
										{service.fq.version}
									</span>
								</p>
							</div>
						</button>
					{/each}
				{:else if $servicesQuery.data}
					<div
						class="flex w-full h-full items-center justify-center text-center text-secondary-600 p-4"
					>
						<div class="flex flex-row items-center gap-2">
							<p>There are no services installed yet!</p>
						</div>
					</div>
				{:else if $servicesQuery.isError}
					<div class="p-4 text-red-400">
						<p>Could not fetch installed services</p>
						<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
							{errorToText($servicesQuery.error)}
						</div>
					</div>
				{:else}
					<div
						class="flex w-full h-full items-center justify-center text-center text-secondary-600"
					>
						<div class="flex flex-row items-center gap-2">
							<Circle size="10" color={colors.gray[200]} />
							<p>Fetching installed services</p>
						</div>
					</div>
				{/if}
			</div>

			<button
				class="w-full btn variant-ghost-primary text-primary-400 flex flex-row gap-0 items-center"
				on:click={() => (serviceModalOpen = true)}
			>
				<p>Install a service</p>
				<span class="text-xs"><PlusIcon /></span>
			</button>
		</div>

		<!-- Main Content (4/5 width on large screens) -->
		<div class=" card variant-ghost lg:col-span-4 overflow-y-auto">
			{#if selectedService}
				<TabGroup>
					<Tab bind:group={tabSet} name="tab1" value={0}>
						<span>Details</span>
					</Tab>
					<Tab bind:group={tabSet} name="tab2" value={1}>Logs</Tab>
					<!-- <Tab bind:group={tabSet} name="tab3" value={2}>Performance</Tab> -->
					<!-- Tab Panels --->
					<svelte:fragment slot="panel">
						{#key selectedService.name + selectedService.author + selectedService.version}
							<ServiceItem
								fq={selectedService}
								process={$pipelineQuery.data?.enabled.find(
									(e) =>
										selectedService &&
										e.service.fq.name === selectedService.name &&
										e.service.fq.author === selectedService.author &&
										e.service.fq.version === selectedService.version
								)?.process}
								{tabSet}
							/>
						{/key}
					</svelte:fragment>
				</TabGroup>
			{:else}
				<div class="flex flex-col p-4">
					{#if $stopPipeline.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">
							Stopped running pipeline
						</div>
					{:else if $stopPipeline.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not stop previous pipeline:
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($stopPipeline.error)}
							</div>
						</div>
					{:else if $stopPipeline.isLoading}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400">
							Stopping pipeline...
						</div>
					{/if}

					{#if $buildService.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">
							Built all services
						</div>
					{:else if $buildService.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not build service {$buildService.variables?.name}:
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($buildService.error)}
							</div>
						</div>
					{:else if $buildService.isLoading}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400">
							Building service {$buildService.variables?.name}...
						</div>
					{/if}

					{#if $savePipeline.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">Saved pipeline</div>
					{:else if $savePipeline.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not save pipeline
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($savePipeline.error)}
							</div>
						</div>
					{:else if $savePipeline.isLoading}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400">
							Saving pipeline...
						</div>
					{/if}

					{#if $startPipeline.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">
							Started pipeline
						</div>
					{:else if $startPipeline.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not start pipeline:
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($startPipeline.error)}
							</div>
						</div>
					{:else if $startPipeline.isLoading}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400">
							Starting pipeline...
						</div>
					{/if}
				</div>
			{/if}
		</div>
	</div>
</div>

<InstallServiceModal bind:modalOpen={serviceModalOpen} />
