<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WarningIcon from '~icons/ix/warning-filled';

	import StartIcon from '~icons/ic/round-play-circle';
	import StopIcon from '~icons/ic/round-stop-circle';
	import CheckIcon from '~icons/ic/sharp-check';
	import PlusIcon from '~icons/heroicons-solid/plus';
	import CheckmarkIcon from '~icons/heroicons/check-badge-20-solid';
	import UserIcon from '~icons/heroicons/user-20-solid';
	import PowerOffIcon from '~icons/ic/round-power';

	import { toasts } from 'svelte-toasts';

	import { SvelteFlowProvider } from '@xyflow/svelte';
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
	import { errorToText, RoverError } from '$lib/errors';
	import InstallServiceModal from './InstallServiceModal.svelte';

	import InputIcon from '~icons/ic/baseline-input';
	import OutputIcon from '~icons/ic/baseline-output';
	import ConfigurationIcon from '~icons/ic/baseline-settings';
	import LogsIcon from '~icons/ic/baseline-notes';
	import { ASE_AUTHOR_IDENTIFIER, TRANSCEIVER_IDENTIFIER } from '$lib/constants';
	import { compareVersions } from '$lib/utils/versions';
	import InstallTransceiverModal from './InstallTransceiverModal.svelte';
	import { serviceConflicts, serviceEqual, serviceIdentifier } from '$lib/utils/service';

	const queryClient = useQueryClient();

	// Flowchart data
	const nodeTypes = {
		service: ServiceNode
	};
	const nodes: Writable<PipelineNode[]> = writable([]);
	const edges: Writable<Edge[]> = writable([]);

	const edgesFromEnabledServices = (enabled: Node<PipelineNodeData>[]) => {
		// For each dependency, create an edge
		const newEdges: Edge[] = [];
		for (const e of enabled) {
			const service = e.data.service;
			const fq = e.data.fq;

			for (const input of service.inputs) {
				const source = input.service;
				const target = serviceIdentifier(fq);
				for (const stream of input.streams) {
					const id = `${source}-${target}-${stream}`;

					// Find the node that matches the source
					const sourceNode = enabled.find(
						(n) =>
							serviceIdentifier(n.data.fq) === source && n.data.service.outputs.includes(stream)
					);

					if (sourceNode) {
						newEdges.push({
							id: id,
							source: sourceNode.id,
							target: e.id, // that is this node
							data: {
								label: stream,
								type: 'input'
							}
						});
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
				const service = services.find((s) => serviceEqual(s.fq, e.service.fq));
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

				return generateNode({
					fq: service.fq,
					service: service.service
				});
			});

			// Create edges from the enabled services
			const newEdges = edgesFromEnabledServices(newNodes);
			createAndSetGraph(newNodes, newEdges);

			return pipeline.data;
		},
		{
			enabled: true, // run once on mount
			refetchOnMount: false,
			staleTime: 1
		}
	);

	const generateNode = (service: PipelineNodeData) => {
		// Create a random id
		const randomId = Math.random().toString(36).substring(7);

		return {
			id: randomId,
			position: { x: 0, y: 0 },
			type: 'service',
			width: 200,
			height: 80,
			deletable: false,
			draggable: false,
			data: {
				...service,
				onSetActive: () => {
					// Set the selected service
					selectedService = service.fq;
				}
			}
		};
	};

	const createAndSetGraph = (newNodes: PipelineNode[], newEdges: Edge[]) => {
		console.log('Was asked to set edges for graph', newNodes, newEdges);

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
				node.data.fq.name !== TRANSCEIVER_IDENTIFIER ? { ...graph.node(node.id) } : { x: 0, y: 0 }
		}));

		nodes.set(positionedNodes);
		edges.set(newEdges);
	};

	const addService = (service: PipelineNodeData) => {
		// Can't modify during execution
		if ($pipelineQuery.data && $pipelineQuery.data.status === 'started') {
			return;
		}

		// If there is a conflicting service already enabled, remove it
		const filteredNodes = $nodes.filter((n) => !serviceConflicts(n.data.fq, service.fq));

		// Add the service to the pipeline
		const newNodes = [...filteredNodes, generateNode(service)];

		// Add edges from the new service
		const newEdges = edgesFromEnabledServices(newNodes);

		createAndSetGraph(newNodes, newEdges);
	};

	const addServiceByName = (author: string, name: string, version?: string) => {
		// Can't modify during execution
		if ($pipelineQuery.data && $pipelineQuery.data.status === 'started') {
			return;
		}

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

		// Add the service to the pipeline
		addService(service);
	};

	const selectServiceByName = (author: string, name: string, version?: string) => {
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

		// Set the selected service
		selectedService = service.fq;
	};

	const removeService = (fq: FullyQualifiedService) => {
		// Can't modify during execution
		if ($pipelineQuery.data && $pipelineQuery.data.status === 'started') {
			return;
		}

		selectedService = fq;

		// Check if service is already present
		if (!$nodes.some((n) => n.data.fq.name === fq.name)) {
			return;
		}

		// Remove the service from the pipeline
		const newNodes = $nodes.filter((n) => n.data.fq.name !== fq.name);

		// Add edges from the new service
		const newEdges = edgesFromEnabledServices(newNodes);

		createAndSetGraph(newNodes, newEdges);
	};

	const removeServiceByName = (name: string) => {
		// Can't modify during execution
		if ($pipelineQuery.data && $pipelineQuery.data.status === 'started') {
			return;
		}

		// Find the node to remove
		const node = $nodes.find((n) => n.data.fq.name === name);
		if (node) {
			removeService(node.data.fq);
		}
	};

	onMount(() => {
		// Fetch the pipeline on mount
		$nodesQuery.refetch();
	});

	// Periodically fetch pipeline status
	const pipelineQuery = useQuery(
		'pipeline',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);

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

				// Show a warning about the service that caused the pipeline to crash (if any)
				for (const enabled of data.enabled) {
					const previousEnabled = previousData?.enabled.find(
						(e) => e.service.fq.name === enabled.service.fq.name
					);

					if (
						// Service must be crashed
						enabled.service.exit !== 0 &&
						// But only send this notification on the first time we know it crashed
						(!previousEnabled ||
							previousEnabled.service.exit === 0 || // Service was not crashed before
							// Or the service was restarted
							(previousData?.last_start &&
								data.last_start &&
								previousData?.last_start < data.last_start))
					) {
						toasts.add({
							title: 'Service error',
							description:
								"Service '" +
								enabled.service.fq.name +
								"' exited with code " +
								enabled.service.exit +
								' and crashed the pipeline. Check its logs for more information.',
							duration: 10000,
							placement: 'bottom-right',
							type: 'warning',
							theme: 'dark',
							onClick: () => {},
							onRemove: () => {}
						});
					}
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
		let services = $nodes;

		// If debug mode is not enabled make sure to not include any transceiver service
		if (!$debugActive.data) {
			services = services.filter((n) => n.data.fq.name !== TRANSCEIVER_IDENTIFIER);
		}

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

			// Exclude transceiver services always
			for (const service of $servicesQuery.data.filter(
				(s) => s.fq.name !== TRANSCEIVER_IDENTIFIER
			)) {
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
			const groups = Array.from(serviceMap.entries())
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

			// Sort by author alphabetically, make sure that vu-ase is always first
			return groups.sort((a, b) => {
				if (!a || !b) {
					return 0;
				}
				if (a.author === ASE_AUTHOR_IDENTIFIER) {
					return -1;
				}
				if (b.author === ASE_AUTHOR_IDENTIFIER) {
					return 1;
				}
				return a.author.localeCompare(b.author);
			});
		}
	);

	// Debug mode is active when:
	// - a transceiver service is enabled
	// - this transceiver service has the same passthrough server specified as the roverctl configuration
	// - roverctl-web was started with debug info environment variables
	const debugActive = useQuery(['debugActive', $nodes], async () => {
		if (!config.success || !config.passthrough) {
			return false;
		}

		const transceiver = $nodes.find((n) => n.data.fq.name === TRANSCEIVER_IDENTIFIER);
		if (!transceiver) {
			return false;
		}

		const fq = transceiver.data.fq;

		// Query the service API to get the configuration for this transceiver
		const sapi = new ServicesApi(config.roverd.api);
		const service = await sapi.servicesAuthorServiceVersionGet(fq.author, fq.name, fq.version);
		if (!service) {
			return false;
		}

		// Find the "passthrough-address" configuration key
		const passthrough = service.data.configuration.find(
			(c) => c.name === 'passthrough-address' && c.type === 'string'
		);
		if (!passthrough) {
			return false;
		}

		// strip the protocol from the roverctl configuration
		const address = passthrough.value.toString().replace(/^https?:\/\//, '');
		return address === config.passthrough.host + ':' + config.passthrough.port;
	});

	const disableDebugMode = () => {
		// Remove the transceiver service
		removeServiceByName(TRANSCEIVER_IDENTIFIER);
		queryClient.invalidateQueries('debugActive');
	};

	const enableDebugMode = useMutation(
		'enableDebugMode',
		async () => {
			if (!config.success) {
				throw new RoverError('Config could not be loaded', 'ERR_CONFIG_INVALID');
			}

			if (!config.passthrough) {
				throw new RoverError('Passthrough was not enabled', 'ERR_PASSTHROUGH_DISABLED');
			}

			// Get all services, check the ones that are transceivers
			const sapi = new ServicesApi(config.roverd.api);
			const services = await sapi.fqnsGet();
			const transceivers = services.data
				.filter((s) => s.name === TRANSCEIVER_IDENTIFIER)
				.sort((a, b) => compareVersions(b.version, a.version));

			for (const transceiver of transceivers) {
				// Does this transceiver expose the same passthrough server as the roverctl configuration?
				const service = await sapi.servicesAuthorServiceVersionGet(
					transceiver.author,
					transceiver.name,
					transceiver.version
				);
				if (!service.data) {
					continue;
				}

				// Find the "passthrough-address" configuration key
				const passthrough = service.data.configuration.find(
					(c) => c.name === 'passthrough-address' && c.type === 'string'
				);
				if (!passthrough) {
					continue;
				}

				// strip the protocol from the roverctl configuration
				const address = passthrough.value.toString().replace(/^https?:\/\//, '');
				if (address === config.passthrough.host + ':' + config.passthrough.port) {
					addService({
						fq: transceiver,
						service: service.data
					});
					return;
				}
			}

			throw new RoverError(
				'No transceiver service with matching passthrough address found',
				'ERR_NO_TRANSCEIVER_INSTALLED'
			);
		},
		{
			onSettled: () => {
				queryClient.invalidateQueries('debugActive');
				queryClient.invalidateQueries('getRelease');
				queryClient.invalidateQueries('downloadFile');
				queryClient.invalidateQueries('modifyZip');
				queryClient.invalidateQueries('uploadZip');
			}
		}
	);
</script>

<div class="h-[90vh] sm:h-[30vh] overflow-hidden relative shrink-0">
	<!-- Pipeline flowchart column -->
	<div
		class="md:col-span-2 lg:col-span-3 h-full overflow-hidden"
		style="touch-action: manipulation;"
	>
		<SvelteFlowProvider>
			<SvelteFlow
				{nodes}
				{edges}
				{nodeTypes}
				style="background-color: transparent;"
				proOptions={{ hideAttribution: true }}
			>
				<Background
					bgColor={'transparent'}
					class="bg-transparent"
					patternColor={colors.slate[500]}
					gap={20}
				/>
				<AutoFit />
			</SvelteFlow>
		</SvelteFlowProvider>

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
					<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
						{errorToText($nodesQuery.error)}
					</div>
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

					<div class="flex flex-row gap-4 items-center">
						<div class="flex flex-row gap-4 items-center">
							<SlideToggle
								name="slider-small"
								checked={!!$debugActive.data}
								background="bg-surface-400"
								active="bg-primary-600"
								size="sm"
								disabled
							/>

							<p class="text-secondary-700">debug mode</p>
						</div>

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
		{:else if $nodes.filter((n) => n.data.fq.name !== TRANSCEIVER_IDENTIFIER).length === 0}
			<div class="w-full card p-2 px-4">
				<div class="flex flex-row justify-between items-center w-full">
					<div class="flex flex-col">
						<div class="flex flex-row items-center gap-2">
							<p class="text-white text-xl">
								Pipeline is <span class="text-secondary-700">empty</span>
							</p>
						</div>
					</div>

					<div class="flex flex-row gap-4 items-center">
						<div class="flex flex-row gap-4 items-center">
							<SlideToggle
								name="slider-small"
								checked={!!$debugActive.data}
								background="bg-surface-400"
								active="bg-primary-600"
								size="sm"
								disabled
							/>

							<p class="text-secondary-700">debug mode</p>
						</div>

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

					<div class="flex flex-row gap-4 items-center">
						<div class="flex flex-row gap-4 items-center">
							<SlideToggle
								name="slider-small"
								checked={!!$debugActive.data}
								background="bg-surface-400"
								active="bg-primary-600"
								size="sm"
								disabled={$debugActive.isLoading || $enableDebugMode.isLoading}
								on:click={(e) => {
									e.preventDefault();
									e.stopPropagation();

									if (!!$debugActive.data) {
										disableDebugMode();
									} else {
										$enableDebugMode.mutate();
									}
								}}
							/>

							<p class="text-secondary-700">debug mode</p>
						</div>

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
								Pipeline is in <span class="text-orange-500">unknown status</span>
							</p>
						</div>
					</div>

					<div class="flex flex-row gap-4 items-center">
						<div class="flex flex-row gap-4 items-center">
							<SlideToggle
								name="slider-small"
								checked={!!$debugActive.data}
								background="bg-surface-400"
								active="bg-primary-600"
								size="sm"
								disabled
							/>

							<p class="text-secondary-700">debug mode</p>
						</div>

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
</div>

<div class="w-full px-4 flex-1 min-h-0 overflow-auto py-4">
	<div class="grid grid-cols-1 lg:grid-cols-5 h-full gap-4">
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
													on:click={() => {
														selectServiceByName(group.author, service.name);
													}}
												>
													<svelte:fragment slot="lead">
														<div class="flex h-full flex-row items-center ml-4">
															{#if $nodes.some((node) => node.data.fq.name === service.name && node.data.fq.author === group.author)}
																{#if $pipelineQuery.data && $pipelineQuery.data.status === 'started'}
																	<button
																		on:click={(e) => {
																			e.preventDefault();
																			e.stopPropagation();
																		}}
																		class="w-5 h-5 card variant-outline-surface text-surface-400"
																	>
																		<CheckIcon />
																	</button>
																{:else}
																	<button
																		on:click={(e) => {
																			e.preventDefault();
																			e.stopPropagation();
																			removeServiceByName(service.name);
																		}}
																		class="w-5 h-5 card variant-outline-secondary text-secondary-400"
																	>
																		<CheckIcon />
																	</button>
																{/if}
															{:else}
																<button
																	on:click={(e) => {
																		e.preventDefault();
																		e.stopPropagation();
																		addServiceByName(group.author, service.name);
																		selectServiceByName(group.author, service.name);
																	}}
																	class="w-5 h-5 card variant-outline-tertiary text-primary-400"
																>
																</button>
															{/if}
														</div>
													</svelte:fragment>
													<svelte:fragment slot="summary">
														<div class="flex flex-row justify-between">
															{#if $pipelineQuery.data && $pipelineQuery.data.enabled.some((s) => s.service.fq.name === service.name && s.service.exit !== 0)}
																<span class="font-mono text-warning-400 whitespace-nowrap truncate">
																	{service.name}
																</span>
															{:else}
																<span
																	class="font-mono text-secondary-800 whitespace-nowrap truncate"
																>
																	{service.name}
																</span>
															{/if}
														</div>
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
																flex flex-row w-full gap-2 items-center text-left btn pl-12`}
																on:click={() =>
																	(selectedService = {
																		name: service.name,
																		author: group.author,
																		version: version
																	})}
															>
																{#if $nodes.some((node) => node.data.fq.name === service.name && node.data.fq.author === group.author && node.data.fq.version === version)}
																	{#if $pipelineQuery.data && $pipelineQuery.data.status === 'started'}
																		<button
																			class="w-5 h-5 rounded-full card variant-outline-surface text-surface-400"
																		>
																			<CheckIcon />
																		</button>
																	{:else}
																		<button
																			class="w-5 h-5 rounded-full card variant-outline-secondary text-secondary-400"
																			on:click={() =>
																				removeService({
																					name: service.name,
																					author: group.author,
																					version: version
																				})}
																		>
																			<CheckIcon />
																		</button>
																	{/if}
																{:else}
																	<button
																		on:click={() =>
																			addServiceByName(group.author, service.name, version)}
																		class="w-5 h-5 rounded-full card variant-outline-tertiary text-primary-400"
																	>
																	</button>
																{/if}

																<div class="flex flex-col w-full">
																	<p class="text-sm text-secondary-800">
																		v{version.replace('v', '')}
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

			<div class="flex flex-row gap-4">
				<button
					class="w-full btn variant-ghost-primary text-primary-400 flex flex-row gap-0 items-center"
					on:click={() => (serviceModalOpen = true)}
				>
					<p>Install a service</p>
					<span class="text-xs"><PlusIcon /></span>
				</button>
				<a
					href="/shutdown"
					class="w-full btn variant-ghost-error text-error-400 flex flex-row gap-0 items-center"
					on:click={() => (serviceModalOpen = true)}
				>
					<p>Shut down</p>
					<span class="text-xs"><PowerOffIcon /></span>
				</a>
			</div>
		</div>

		<!-- Main Content (4/5 width on large screens) -->
		{#if selectedService}
			<div class=" card variant-ghost lg:col-span-4 overflow-y-auto">
				<div class="flex flex-col gap-1 p-2 px-4">
					<p class="text-secondary-700">
						{selectedService.author}
					</p>
					<p class="text-2xl font-mono">
						<span class="text-primary-400">
							{selectedService.name}
						</span>

						{#if selectedService.as}
							as
							<span class="text-secondary-700">
								{selectedService.as}
							</span>
						{/if}
					</p>
					<p class="text-secondary-700">
						v{selectedService.version.replace('v', '')}
					</p>

					<!-- todo: currently, this re-iterates many times over the pipeline query data. This could be prevented by extracting this logic in a separate component
					 and re-using a variable. It is an optimization, but it would also make the code nicer -->
					{#if $pipelineQuery.data && $pipelineQuery.data.enabled.some((s) => selectedService && s.service.exit !== 0 && serviceEqual(s.service.fq, selectedService))}
						<div class="card variant-soft-warning px-4 py-2 mt-1">
							<p class="text-warning-400">
								This service crashed with exit code
								{$pipelineQuery.data.enabled.find(
									(s) => selectedService && serviceEqual(s.service.fq, selectedService)
								)?.service.exit}
								and caused the pipeline to crash.
							</p>
						</div>
					{/if}
				</div>

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
									(e) => selectedService && serviceEqual(e.service.fq, selectedService)
								)?.process}
								{tabSet}
							/>
						{/key}
					</svelte:fragment>
				</TabGroup>
			</div>
		{:else}
			<div class="  lg:col-span-4 overflow-y-auto">
				<div class="flex flex-col">
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
			</div>
		{/if}
	</div>
</div>

<InstallServiceModal bind:modalOpen={serviceModalOpen} />
<InstallTransceiverModal enableMutation={enableDebugMode} />
