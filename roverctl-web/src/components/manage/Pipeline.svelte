<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WarningIcon from '~icons/ix/warning-filled';
	import PlusIcon from '~icons/icons8/plus';

	import WifiIcon from '~icons/material-symbols/wifi';
	import WifiOffIcon from '~icons/material-symbols/wifi-off';
	import type { Edge, Node } from '@xyflow/svelte';
	import { Circle } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { useMutation, useQuery } from '@sveltestack/svelte-query';
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
	import { Modal } from '@skeletonlabs/skeleton';
	import ErrorOverlay from '../../components/ErrorOverlay.svelte';
	import { writable } from 'svelte/store';
	import { SvelteFlow, Background, Controls, MarkerType } from '@xyflow/svelte';
	import '@xyflow/svelte/dist/style.css';
	import ServiceNode from './ServiceNode.svelte';
	import dagre from 'dagre';
	import type { Writable } from 'svelte/store';
	import type { PipelineNode, PipelineNodeData } from './type';
	import { onMount } from 'svelte';

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
			staleTime: 1 // needs to be fresh
		}
	);

	// Whether the "add service" modal is open
	$: serviceModalOpen = false;

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

	const stopPipeline = useMutation('stopPipeline', async () => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const papi = new PipelineApi(config.roverd.api);
		// Best-effort based, ignore errors (i.e. if already stopped)
		try {
			const response = await papi.pipelineStopPost();
			return response.data;
		} catch {}
	});

	const buildService = useMutation('buildService', async (fq: FullyQualifiedService) => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const sapi = new ServicesApi(config.roverd.api);
		const response = await sapi.servicesAuthorServiceVersionPost(fq.author, fq.name, fq.version);
		return response.data;
	});

	const startPipeline = useMutation('startPipeline', async () => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const papi = new PipelineApi(config.roverd.api);
		const response = await papi.pipelineStartPost();
		return response.data;
	});

	const startConfiguredPipeline = async () => {
		await $stopPipeline.mutateAsync();
		await Promise.all($nodes.map((n) => $buildService.mutateAsync(n.data.fq)));
		await $startPipeline.mutateAsync();
	};
</script>

<div class="mx-4 mt-4 border border-solid border-gray-700 h-full relative border-b-0">
	<SvelteFlow
		{nodes}
		{edges}
		{nodeTypes}
		proOptions={{ hideAttribution: true }}
		on:nodeDelete={({ detail }) => {
			console.log('delete', detail);
		}}
	>
		<Background bgColor={colors.slate[900]} patternColor={colors.slate[500]} />
		<Controls class="text-slate-800" />
	</SvelteFlow>

	<!-- Loading Overlay -->
	{#if $nodesQuery.isLoading}
		<div class="absolute inset-0 flex justify-center items-center">
			<!-- Semi-transparent background -->
			<div class="absolute inset-0 bg-gray-200 bg-opacity-20"></div>

			<div class="flex flex-row items-center gap-2 text-zinc-400">
				<Circle size="20" color={colors.zinc[400]} />
				<p>fetching pipeline</p>
			</div>
		</div>

		<!-- Error Overlay -->
	{:else if $nodesQuery.isError}
		<div class="absolute inset-0 flex justify-center items-center">
			<!-- Semi-transparent background -->
			<div class="absolute inset-0 bg-error-300 bg-opacity-20"></div>

			<!-- Fully opaque text/icon -->
			<div class="flex flex-col text-center text-white items-center">
				<WarningIcon class="text-2xl relative mb-2" />
				<h1>Could not fetch pipeline</h1>
				<p>
					{$nodesQuery.error || 'An unknown error occurred'}
				</p>
			</div>
		</div>
	{/if}
</div>

<div
	class="mx-4 mb-4 border border-solid border-gray-700 border-t-0 relative p-4 bg-slate-800 flex flex-row items-center justify-between"
>
	<div class="flex flex-row">
		{#if $stopPipeline.isLoading}
			<p>Stopping pipeline</p>
		{:else if $buildService.isLoading}
			<p>Building service: {$buildService.variables?.name}</p>
		{:else if $buildService.isError}
			<p>Failed to build service ({$buildService.variables?.name}): {$buildService.error}</p>
		{:else if $startPipeline.isLoading}
			<p>Starting pipeline</p>
		{:else if $startPipeline.isSuccess}
			<p>Pipeline started</p>
		{:else if $startPipeline.isError}
			<p>Failed to start pipeline: {$startPipeline.error}</p>
		{:else if $pipelineQuery.data}
			<p>Pipeline is {$pipelineQuery.data.status}</p>
		{:else if $pipelineQuery.isLoading}
			<div class="flex flex-row items-center gap-2 text-zinc-400">
				<Circle size="20" color={colors.zinc[400]} />
				<p>fetching pipeline</p>
			</div>
		{:else if $pipelineQuery.isError}
			<div class="flex flex-col text-center text-white items-center">
				<WarningIcon class="text-2xl relative mb-2" />
				<h1>Could not fetch pipeline</h1>
				<p>
					{$pipelineQuery.error || 'An unknown error occurred'}
				</p>
			</div>
		{/if}
	</div>
	<div class="flex flex-row gap-4">
		<button
			type="button"
			class="btn btn-sm variant-filled"
			on:click={() => (serviceModalOpen = true)}
		>
			<PlusIcon />
			<span class="ml-2">Add Service</span>
		</button>
		<button
			type="button"
			class="btn btn-sm variant-filled"
			on:click={() => startConfiguredPipeline()}
		>
			<!-- <PlusIcon /> -->
			<span class="ml-2">Start Pipeline</span>
		</button>
	</div>
</div>

<!-- Modal Background -->
<div
	class={serviceModalOpen
		? 'fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center'
		: 'hidden'}
>
	<!-- Modal Content -->
	<div class="bg-slate-300 text-slate-700 p-6 rounded-lg shadow-lg">
		{#if $servicesQuery.data}
			{#if $servicesQuery.data.length === 0}
				<p>No services</p>
			{:else}
				<ul>
					{#each $servicesQuery.data as service}
						<button
							type="button"
							class="btn btn-sm variant-filled"
							on:click={() => addService(service)}
						>
							<PlusIcon />
							<span class="ml-2">Add {service.fq.name}</span>
						</button>
					{/each}
				</ul>
			{/if}
		{:else if $servicesQuery.isLoading}
			<div class="flex flex-row items-center gap-2 text-zinc-400">
				<Circle size="20" color={colors.zinc[400]} />
				<p>fetching services</p>
			</div>
		{:else if $servicesQuery.isError}
			<div class="flex flex-col text-center text-white items-center">
				<WarningIcon class="text-2xl relative mb-2" />
				<h1>Could not fetch services</h1>
				<p>
					{$servicesQuery.error || 'An unknown error occurred'}
				</p>
			</div>
		{/if}

		<!-- <h2 class="text-lg font-semibold">Modal Title</h2>
		<p class="mt-2">This is a declarative Tailwind modal.</p> -->

		<!-- Close Button -->
		<button
			class="mt-4 bg-red-500 text-white px-4 py-2 rounded"
			on:click={() => (serviceModalOpen = false)}
		>
			Close
		</button>
	</div>
</div>
