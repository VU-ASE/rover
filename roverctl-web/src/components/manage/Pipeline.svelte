<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WarningIcon from '~icons/ix/warning-filled';

	import StartIcon from '~icons/ic/round-play-circle';
	import StopIcon from '~icons/ic/round-stop-circle';
	import CheckIcon from '~icons/ic/sharp-check';
	import PlusIcon from '~icons/subway/add-1';
	import CheckmarkIcon from '~icons/heroicons/check-badge-20-solid';
	import UserIcon from '~icons/heroicons/user-20-solid';

	import { toasts } from 'svelte-toasts';

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
	import { Accordion, AccordionItem, Modal, Tab, TabGroup } from '@skeletonlabs/skeleton';
	import ErrorOverlay from '../../components/ErrorOverlay.svelte';
	import { derived, writable } from 'svelte/store';
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

	import InputIcon from '~icons/ic/baseline-input';
	import OutputIcon from '~icons/ic/baseline-output';
	import ConfigurationIcon from '~icons/ic/baseline-settings';
	import LogsIcon from '~icons/ic/baseline-notes';
	import { ASE_AUTHOR_IDENTIFIER, TRANSCEIVER_IDENTIFIER } from '$lib/constants';
	import { compareVersions } from '$lib/utils/versions';

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
						onDelete: () => removeService(service.fq),
						onSetActive: () => {
							// Set the selected service
							selectedService = service.fq;
						}
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
			staleTime: 1
		}
	);

	const createAndSetGraph = (newNodes: PipelineNode[], newEdges: Edge[]) => {
		// Create a daggerable graph to automatically layout the nodes
		const graph = new dagre.graphlib.Graph();
		graph.setGraph({ rankdir: 'LR' }); // Top-to-Bottom layout
		graph.setDefaultEdgeLabel(() => ({}));

		// Add nodes to graph, don't include the transceiver node in the layout, but do add it to the existing nodes
		newNodes.forEach((node) => {
			if (node.data.fq.name !== TRANSCEIVER_IDENTIFIER) {
				graph.setNode(node.id, { width: node.width, height: node.height });
			}
		});

		// Add edges to graph
		newEdges.forEach((edge) => graph.setEdge(edge.source, edge.target));

		// Compute layout
		dagre.layout(graph);

		// Apply computed positions
		const positionedNodes = newNodes.map((node) => ({
			...node,
			position:
				node.data.fq.name !== TRANSCEIVER_IDENTIFIER
					? { x: graph.node(node.id).x, y: graph.node(node.id).y }
					: { x: 0, y: 0 }
		}));

		nodes.set(positionedNodes);
		edges.set(newEdges);
	};

	const addService = (service: PipelineNodeData) => {
		$stopPipeline.mutate();
		selectedService = service.fq;

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

	const addServiceByName = (author: string, name: string, version?: string) => {
		// Filter services that match the author and name, sort by version descending
		const filtered = $servicesQuery.data?.filter(
			(s) => s.fq.author === author && s.fq.name === name
		);
		if (!filtered) {
			return;
		}

		const sorted = filtered.sort((a, b) => compareVersions(b.fq.version, a.fq.version));
		if (sorted.length < 1) {
			return;
		}

		// If a version is specified, use it, otherwise use the latest version
		const service = sorted.find((s) => s.fq.version === version) || sorted[0];

		// If there is already a service with this name enabled, remove this
		if ($nodes.some((n) => n.data.fq.name === service.fq.name)) {
			removeService(service.fq);
		}

		// Add the service to the pipeline
		addService(service);
	};

	const removeServiceByName = (name: string) => {
		$stopPipeline.mutate();

		// Remove the service from the pipeline
		const newNodes = $nodes.filter((n) => n.data.fq.name !== name);

		// Add edges from the new service
		const newEdges = edgesFromEnabledServices(newNodes.map((n) => n.data));

		createAndSetGraph(newNodes, newEdges);
	};

	const removeService = (fq: FullyQualifiedService) => {
		$stopPipeline.mutate();
		selectedService = fq;

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
			keepPreviousData: false,
			staleTime: 1,
			refetchInterval: 1000,
			onSuccess: (data) => {
				const previousData = $pipelineQuery.data;
				if (previousData && previousData.enabled.length > 0 && data.enabled.length <= 0) {
					toasts.add({
						title: 'Pipeline reset',
						description: 'The existing pipeline was emptied by roverd',
						duration: 10000,
						placement: 'bottom-right',
						type: 'info',
						theme: 'dark',
						onClick: () => {},
						onRemove: () => {}
					});
				}
			},
			onError: () => {
				const previousSuccess = $pipelineQuery.isSuccess;
				if (previousSuccess) {
					toasts.add({
						title: 'Pipeline error',
						description: 'An error occurred while fetching the pipeline status',
						duration: 10000,
						placement: 'bottom-right',
						type: 'error',
						theme: 'dark',
						onClick: () => {},
						onRemove: () => {}
					});
				}
			}
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

	// The active tab selected for service information
	let tabSet: number = 0;
	let pipelineStarting = false;
	$: pipelineStarting =
		$stopPipeline.isLoading ||
		$startPipeline.isLoading ||
		$buildService.isLoading ||
		$savePipeline.isLoading;

	let searchInput = writable('');
	// Organize available services per author, name and version
	const groupedServices = derived(
		[servicesQuery, searchInput], // Reactively depend on servicesQuery and searchInput
		([$servicesQuery, $searchInput]) => {
			if (!$servicesQuery?.data) return [];

			const searchTerm = $searchInput.toLowerCase().trim();

			const serviceMap = new Map<string, Map<string, Set<string>>>();

			for (const service of $servicesQuery.data) {
				const { name, author, version } = service.fq;

				if (!serviceMap.has(author)) {
					serviceMap.set(author, new Map());
				}

				const authorMap = serviceMap.get(author)!;

				if (!authorMap.has(name)) {
					authorMap.set(name, new Set());
				}

				authorMap.get(name)!.add(version);
			}

			// Convert map to array, filter results based on search input
			return Array.from(serviceMap.entries())
				.map(([author, namesMap]) => {
					// Filter names based on search input
					const filteredNames = Array.from(namesMap.entries())
						.map(([name, versionsSet]) => {
							// Filter versions based on search input
							const filteredVersions = Array.from(versionsSet)
								.filter(
									(version) =>
										searchTerm === '' ||
										author.toLowerCase().includes(searchTerm) ||
										name.toLowerCase().includes(searchTerm) ||
										version.toLowerCase().includes(searchTerm)
								)
								.sort((a, b) => compareVersions(b, a)); // Sort versions descending

							return filteredVersions.length > 0 ? { name, versions: filteredVersions } : null;
						})
						.filter(Boolean) as { name: string; versions: string[] }[];

					return filteredNames.length > 0 ? { author, names: filteredNames } : null;
				})
				.filter(Boolean); // Ensure only non-null values are returned
		}
	);
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
					<input
						class="input"
						type="text"
						bind:value={$searchInput}
						placeholder="Search services or authors..."
					/>

					<Accordion>
						{#each $groupedServices as group}
							<AccordionItem regionPanel="px-0 py-0" regionControl="px-3 py-1" open>
								<svelte:fragment slot="lead">
									{#if group?.author.toLowerCase() === ASE_AUTHOR_IDENTIFIER.toLowerCase()}
										<span class="text-primary-400">
											<CheckmarkIcon />
										</span>
									{:else}
										<span class="text-secondary-700">
											<UserIcon />
										</span>
									{/if}
								</svelte:fragment>
								<svelte:fragment slot="summary">
									<span class="text-secondary-200">
										{group?.author}
									</span>
								</svelte:fragment>
								<svelte:fragment slot="content">
									<Accordion>
										{#if group}
											{#each group?.names as service}
												<AccordionItem
													regionPanel="px-0 py-0"
													regionControl="px-3 py-1"
													on:click={() => {}}
												>
													<svelte:fragment slot="lead">
														<div class="flex h-full flex-row items-center">
															{#if $nodes.some((node) => node.data.fq.name === service.name && node.data.fq.author === group.author)}
																<button
																	on:click={(e) => {
																		e.preventDefault();
																		e.stopPropagation();
																		removeServiceByName(service.name);
																	}}
																	class="w-5 h-5 card variant-outline-success text-success-400"
																>
																	<CheckIcon />
																</button>
															{:else}
																<button
																	on:click={(e) => {
																		e.preventDefault();
																		e.stopPropagation();
																		addServiceByName(group.author, service.name);
																	}}
																	class="w-5 h-5 card variant-outline-tertiary text-primary-400"
																>
																</button>
															{/if}
														</div>
													</svelte:fragment>
													<svelte:fragment slot="summary">
														<span class="font-mono text-secondary-800">
															{service.name}
														</span>
													</svelte:fragment>
													<svelte:fragment slot="content">
														{#each service.versions as version}
															<button
																class={`p-2 px-4 
																${
																	selectedService &&
																	selectedService.name === service.name &&
																	selectedService.author === group.author &&
																	selectedService.version === version
																		? 'card variant-soft-primary'
																		: ''
																}
								
														flex flex-row w-full gap-2 items-center text-left btn pl-8

														`}
																on:click={() =>
																	(selectedService = {
																		name: service.name,
																		author: group.author,
																		version: version
																	})}
															>
																{#if $nodes.some((node) => node.data.fq.name === service.name && node.data.fq.author === group.author && node.data.fq.version === version)}
																	<button
																		class="w-5 h-5 rounded-full card variant-outline-success text-success-400"
																		on:click={() =>
																			removeService({
																				name: service.name,
																				author: group.author,
																				version: version
																			})}
																	>
																		<CheckIcon />
																	</button>
																{:else}
																	<button
																		on:click={() =>
																			addServiceByName(group.author, service.name, version)}
																		class="w-5 h-5 rounded-full card variant-outline-tertiary text-primary-400"
																	>
																	</button>
																{/if}

																<div class="flex flex-col w-full">
																	<p class="text-md">
																		<span class="text-secondary-800">
																			{version}
																		</span>
																	</p>
																</div>
															</button>
														{/each}
													</svelte:fragment>
												</AccordionItem>
											{/each}
										{/if}
									</Accordion>
								</svelte:fragment>
							</AccordionItem>
						{/each}
					</Accordion>
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
						<div class="flex flex-row items-center text-md gap-2">
							<InputIcon />
							<span>Inputs</span>
						</div>
					</Tab>
					<Tab bind:group={tabSet} name="tab1" value={1}>
						<div class="flex flex-row items-center text-md gap-2">
							<OutputIcon />
							<span>Outputs</span>
						</div>
					</Tab>
					<Tab bind:group={tabSet} name="tab1" value={2}>
						<div class="flex flex-row items-center text-md gap-2">
							<ConfigurationIcon />
							<span>Configuration</span>
						</div>
					</Tab>
					<Tab bind:group={tabSet} name="tab2" value={3}>
						<div class="flex flex-row items-center text-md gap-2">
							<LogsIcon />
							<span>Logs</span>
						</div>
					</Tab>
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
