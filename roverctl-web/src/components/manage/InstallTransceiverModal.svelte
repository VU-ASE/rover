<script lang="ts">
	/**
	 * This modal is shown when enabling debug mode errors, because the necessary service (transceiver)
	 * is not installed on the Rover
	 */

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
	import WebIcon from '~icons/mdi/web';
	import UploadIcon from '~icons/ic/baseline-upload';
	import { compareVersions } from '$lib/utils/versions';
	import { Accordion, AccordionItem } from '@skeletonlabs/skeleton';

	import { useStore } from '@xyflow/svelte';
	import type { Edge, Node } from '@xyflow/svelte';
	import { Circle, DoubleBounce } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import {
		isError,
		Mutation,
		useMutation,
		useQuery,
		useQueryClient,
		type MutationStoreResult
	} from '@sveltestack/svelte-query';
	import {
		HealthApi,
		PipelineApi,
		ServicesApi,
		type FetchPostRequest,
		type FullyQualifiedService,
		type PipelineGet200ResponseEnabledInnerService,
		type ServicesAuthorServiceVersionGet200Response
	} from '$lib/openapi';
	import colors from 'tailwindcss/colors';
	import Navbar from '../Navbar.svelte';
	import { Modal, Tab, TabGroup } from '@skeletonlabs/skeleton';
	import ErrorOverlay from '../ErrorOverlay.svelte';
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
	import { errorToText, RoverError } from '$lib/errors';
	import CheckmarkIcon from '~icons/heroicons/check-badge-20-solid';
	import JSZip from 'jszip';

	import yaml from 'js-yaml';
	import { ASE_AUTHOR_IDENTIFIER, TRANSCEIVER_IDENTIFIER } from '$lib/constants';

	const queryClient = useQueryClient();
	export let enableMutation: MutationStoreResult<void, unknown, void, unknown>;

	/**
	 * Functions to install debug tooling
	 */

	type GitHubRelease = {
		tag_name: string;
		assets: { name: string; browser_download_url: string }[];
	};

	//Get the latest release with exactly one zip file; otherwise, check older releases.
	const getLatestValidRelease = useMutation(
		'getRelease',
		async ({ author, repo }: { author: string; repo: string }) => {
			const releasesUrl = `https://api.github.com/repos/${author}/${repo}/releases`;

			const response = await fetch(releasesUrl);
			if (!response.ok) throw new Error('Failed to fetch releases');

			const releases: GitHubRelease[] = await response.json();

			for (const release of releases) {
				const zipAssets = release.assets.filter((asset) => asset.name.endsWith('.zip'));

				if (zipAssets.length === 1) {
					return {
						version: release.tag_name,
						url: zipAssets[0].browser_download_url
					};
				}
			}

			throw new Error('No valid release found');
		}
	);

	const installTransceiver = async () => {
		// Reset all mutations
		$getLatestValidRelease.reset();

		if (!config.success) {
			throw new RoverError('Config could not be loaded', 'ERR_CONFIG_INVALID');
		}

		if (!config.passthrough) {
			throw new RoverError('Passthrough was not enabled', 'ERR_PASSTHROUGH_DISABLED');
		}

		// Get the latest valid release
		const release = await $getLatestValidRelease.mutateAsync({
			author: ASE_AUTHOR_IDENTIFIER,
			repo: TRANSCEIVER_IDENTIFIER
		});

		// Install from URL via roverd as done in the enableDebugMode mutation
		const sapi = new ServicesApi(config.roverd.api);
		const request: FetchPostRequest = { url: release.url };
		await sapi.fetchPost(request);

		const services = await sapi.fqnsGet();
		const transceivers = services.data
			.filter((s) => s.name === TRANSCEIVER_IDENTIFIER)
			.sort((a, b) => compareVersions(b.version, a.version));

		for (const transceiver of transceivers) {
			// Does this transceiver expose the same passthrough server as the roverctl configuration?
			let service = await sapi.servicesAuthorServiceVersionGet(
				transceiver.author,
				transceiver.name,
				transceiver.version
			);
			if (!service.data) {
				continue;
			}

			// Find the "passthrough-address" configuration key
			let passthrough = service.data.configuration.find(
				(c) => c.name === 'passthrough-address' && c.type === 'string'
			);
			if (!passthrough) {
				continue;
			}

			// Enhancement: try to set the transceiver service.yaml configuration for the passthrough address
			// to the one specified for roverctl.
			const newConfig = service.data.configuration.map((c) => {
				if (
					c.name === 'passthrough-address' &&
					c.type === 'string' &&
					config.success &&
					config.passthrough
				) {
					return {
						...c,
						key: c.name,
						value: 'http://' + config.passthrough.host + ':' + config.passthrough.port
					};
				} else {
					return {
						...c,
						key: c.name
					};
				}
			});

			await sapi.servicesAuthorServiceVersionConfigurationPost(
				transceiver.author,
				transceiver.name,
				transceiver.version,
				newConfig
			);

			// Then try to refetch again
			service = await sapi.servicesAuthorServiceVersionGet(
				transceiver.author,
				transceiver.name,
				transceiver.version
			);
			if (!service.data) {
				continue;
			}

			// Find the "passthrough-address" configuration key
			passthrough = service.data.configuration.find(
				(c) => c.name === 'passthrough-address' && c.type === 'string'
			);
			if (!passthrough) {
				continue;
			}

			const address = passthrough.value.toString().replace(/^https?:\/\//, '');
			if (address === config.passthrough.host + ':' + config.passthrough.port) {
				// 4) Enable debug mode (your existing flow)
				$enableMutation.reset();
				$enableMutation.mutate();

				// Reset mutation we used here
				$getLatestValidRelease.reset();
				return;
			}
		}
		throw new RoverError(
			'Could not configure transceiver to use the configured passthrough server',
			'ERR_PASSTHROUGH_NOT_CONFIGURED'
		);

	};
</script>

{#if $enableMutation.isError}
	<!-- Modal overlay -->
	<div class="fixed inset-0 flex items-center justify-center z-50" aria-modal="true" role="dialog">
		<!-- Dark background overlay -->
		<div class="fixed inset-0 bg-black opacity-50"></div>
		<!-- Modal content container -->
		<div class="bg-surface-600 shadow-lg relative z-50 text-secondary-700 min-w-[40vw]">
			{#if $enableMutation.error instanceof RoverError && $enableMutation.error.code === 'ERR_NO_TRANSCEIVER_INSTALLED'}
				<div class="p-6 pb-4">
					<h2 class="text-xl mb-4 text-secondary-200">Failed to enable Debug Mode</h2>
					<p class="mb-2">
						This Rover is missing the tooling that is required to enable debug mode.
						<br /> Do you want to install this additional tooling now?
					</p>
					<p class="text-sm mb-2">
						<span class="text-primary-400">Note</span>: an internet connection is required.
					</p>

					<!-- Fetching release -->
					{#if $getLatestValidRelease.isLoading}
						<div
							class="flex flex-row gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400"
						>
							<Circle size="20" color={colors.white} />
							<p>Fetching latest release from github</p>
						</div>
					{:else if $getLatestValidRelease.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">
							Found {$getLatestValidRelease.variables?.author}/{$getLatestValidRelease.variables
								?.repo} release {$getLatestValidRelease.data?.version}
						</div>
					{:else if $getLatestValidRelease.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not fetch for {$getLatestValidRelease.variables?.author}/{$getLatestValidRelease
								.variables?.repo}:
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($getLatestValidRelease.error)}
							</div>
						</div>
					{/if}

					<!-- Downloading zip -->
					{#if $downloadFile.isLoading}
						<div
							class="flex flex-row gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400"
						>
							<Circle size="20" color={colors.white} />
							<p>Downloading release</p>
						</div>
					{:else if $downloadFile.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">
							Download successful
						</div>
					{:else if $downloadFile.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not download from {$downloadFile.variables}:
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($downloadFile.error)}
							</div>
						</div>
					{/if}

					<!-- Modifying service YAML to match configured passthrough -->
					{#if $adjustServiceYamlInZip.isLoading}
						<div
							class="flex flex-row gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400"
						>
							<Circle size="20" color={colors.white} />
							<p>Modifying release properties</p>
						</div>
					{:else if $adjustServiceYamlInZip.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">
							Modified service.yaml
						</div>
					{:else if $adjustServiceYamlInZip.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not update service.yaml:
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($adjustServiceYamlInZip.error)}
							</div>
						</div>
					{/if}

					<!-- Upload zip to Rover -->
					{#if $uploadZipToRover.isLoading}
						<div
							class="flex flex-row gap-2 px-4 py-2 border-l-2 border-l-secondary-400 text-secondary-400"
						>
							<Circle size="20" color={colors.white} />
							<p>Uploading tooling to Rover</p>
						</div>
					{:else if $uploadZipToRover.isSuccess}
						<div class="px-4 py-2 border-l-2 border-l-green-500 text-green-600">
							Uploaded transceiver to Rover
						</div>
					{:else if $uploadZipToRover.isError}
						<div class="gap-2 px-4 py-2 border-l-2 border-l-error-400 text-error-400">
							Could not upload transceiver:
							<div class="card mt-2 p-2 px-4 text-red-500 font-mono whitespace-pre-line">
								{errorToText($uploadZipToRover.error)}
							</div>
						</div>
					{/if}
				</div>
				<div class="flex flex-row justify-end mt-4 p-6 pt-2 gap-4">
					<button on:click={$enableMutation.reset} class="btn variant-soft-secondary">
						Close
					</button>
					<button on:click={installTransceiver} class="btn variant-soft-primary"> Install </button>
				</div>
			{:else if $enableMutation.error instanceof RoverError && $enableMutation.error.code === 'ERR_PASSTHROUGH_DISABLED'}
				<div class="p-6 pb-4">
					<h2 class="text-xl mb-4 text-secondary-200">Debug Mode cannot be enabled</h2>
					<p class="mb-2">
						To enable debug mode, you need to restart <span class="code">roverctl</span>
						with the <span class="code">--debug</span> flag.
					</p>
				</div>
				<div class="flex justify-end mt-4 p-6 pt-2">
					<button on:click={$enableMutation.reset} class="btn variant-soft-secondary">
						Close
					</button>
				</div>
			{:else}
				<div class="p-6 pb-4">
					<h2 class="text-xl mb-4 text-secondary-200">Failed to enable Debug Mode</h2>
					<p class="mb-2">Debug mode could not be enabled at this time.</p>
					<div class="card p-2 px-4 text-red-500 font-mono whitespace-pre-line">
						{errorToText($enableMutation.error)}
					</div>
				</div>
				<div class="flex justify-end mt-4 p-6 pt-2">
					<button on:click={$enableMutation.reset} class="btn variant-soft-secondary">
						Close
					</button>
				</div>
			{/if}
		</div>
	</div>
{/if}
