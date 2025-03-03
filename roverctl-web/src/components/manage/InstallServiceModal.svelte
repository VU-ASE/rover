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
	import WebIcon from '~icons/mdi/web';
	import UploadIcon from '~icons/ic/baseline-upload';

	import { Accordion, AccordionItem } from '@skeletonlabs/skeleton';

	import { useStore } from '@xyflow/svelte';
	import type { Edge, Node } from '@xyflow/svelte';
	import { Circle, DoubleBounce } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { isError, useMutation, useQuery, useQueryClient } from '@sveltestack/svelte-query';
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
	import CheckmarkIcon from '~icons/heroicons/check-badge-20-solid';

	const queryClient = useQueryClient();

	export let modalOpen = false;

	let fetchUrl = '';
	$: fetchUrl;

	const fetchService = useMutation(
		'fetchService',
		async (url: string) => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const sapi = new ServicesApi(config.roverd.api);
			const request: FetchPostRequest = {
				url: url
			};
			const response = await sapi.fetchPost(request);
			return response.data;
		},
		{
			// Invalidate the pipeline query regardless of mutation success or failure
			onSettled: () => {
				queryClient.invalidateQueries('pipeline');
				queryClient.invalidateQueries('availableServices');
				queryClient.invalidateQueries('pipelineNodes');
			}
		}
	);

	const fetchLatestRelease = async (repo: string, author: string): Promise<string> => {
		// Construct the GitHub releases URL
		const url = `https://api.github.com/repos/${author}/${repo}/releases`;

		const response = await fetch(url);
		if (!response.ok) {
			throw new Error(`Failed to fetch releases: ${response.status} ${response.statusText}`);
		}

		const releases = await response.json();

		if (!Array.isArray(releases) || releases.length === 0) {
			throw new Error('No releases found');
		}

		// Iterate over releases from latest to older
		for (const release of releases) {
			// Check that release.assets exists and is an array
			if (!release.assets || !Array.isArray(release.assets)) continue;

			// Filter for assets whose names end with .zip
			const zipAssets = release.assets.filter(
				(asset: any) => asset.name && asset.name.endsWith('.zip')
			);

			// The release must have exactly one ZIP asset
			if (zipAssets.length === 1) {
				return zipAssets[0].browser_download_url;
			}
		}

		throw new Error('No release with exactly one ZIP asset found');
	};

	const fetchOfficialService = useMutation(
		'fetchOfficialService',
		async (name: string) => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			// Fetch the latest release ZIP from github under VU-ASE author
			const url = await fetchLatestRelease(name, 'vu-ase');
			const sapi = new ServicesApi(config.roverd.api);
			const request: FetchPostRequest = {
				url: url
			};
			const response = await sapi.fetchPost(request);
			return response.data;
		},
		{
			// Invalidate the pipeline query regardless of mutation success or failure
			onSettled: () => {
				queryClient.invalidateQueries('pipeline');
				queryClient.invalidateQueries('availableServices');
				queryClient.invalidateQueries('pipelineNodes');
			}
		}
	);

	const installBasicPipeline = async () => {
		await $fetchOfficialService.mutateAsync('imaging');
		await $fetchOfficialService.mutateAsync('controller');
		await $fetchOfficialService.mutateAsync('actuator');
	};
</script>

{#if modalOpen}
	<!-- Modal overlay -->
	<div class="fixed inset-0 flex items-center justify-center z-50" aria-modal="true" role="dialog">
		<!-- Dark background overlay -->
		<div class="fixed inset-0 bg-black opacity-50"></div>
		<!-- Modal content container -->
		<div class="bg-surface-600 shadow-lg relative z-50 text-secondary-700 min-w-[40vw]">
			<div class="p-6 pb-4">
				<h2 class="text-xl mb-4 text-secondary-200">Install a Service</h2>
				<p class="">
					There are various options to install services on your Rover. Select one that fits your
					workflow best.
				</p>
			</div>
			<Accordion autocollapse>
				<AccordionItem padding="px-6 py-2">
					<svelte:fragment slot="lead">
						<span class="text-primary-400">
							<UploadIcon />
						</span>
					</svelte:fragment>
					<svelte:fragment slot="summary">
						<span class="text-secondary-200">Upload from disk</span>
					</svelte:fragment>
					<svelte:fragment slot="content">
						To upload services from disk to this Rover, run
						<span class="code">roverctl upload</span> in your terminal.
					</svelte:fragment>
				</AccordionItem>
				<AccordionItem padding="px-6 py-2" on:click={$fetchService.reset}>
					<svelte:fragment slot="lead">
						<span class="text-primary-400">
							<CheckmarkIcon />
						</span>
					</svelte:fragment>
					<svelte:fragment slot="summary">
						<span class="text-secondary-200">Install ASE services</span>
					</svelte:fragment>
					<svelte:fragment slot="content">
						<div class="flex flex-col gap-2">
							<p>Install a basic autonomous driving pipeline on your Rover.</p>
							{#if !$fetchOfficialService.isLoading && !$fetchOfficialService.isSuccess}
								<button class="btn variant-filled-primary" on:click={installBasicPipeline}
									>Install latest version</button
								>
							{:else}
								<button class="btn variant-soft-primary" disabled>Install latest version</button>
							{/if}
							{#if $fetchOfficialService.isLoading}
								<div class="flex flex-row items-center gap-2 mt-0 px-2 text-xs">
									<Circle size="10" color={colors.gray[200]} />
									<p>Installing '{$fetchOfficialService.variables}'</p>
								</div>
							{:else if $fetchOfficialService.isError}
								<div class="flex flex-col gap-2 mt-0 px-2 text-xs">
									<p class="text-error-400">
										Could not install '{$fetchOfficialService.variables}':
									</p>
								</div>
							{:else if $fetchOfficialService.isSuccess}
								<div class="flex flex-row items-center gap-2 mt-0 px-2 text-xs text-success-400">
									<p>Services were installed successfully</p>
								</div>
							{/if}
						</div>
					</svelte:fragment>
				</AccordionItem>

				<AccordionItem padding="px-6 py-2" on:click={$fetchOfficialService.reset}>
					<svelte:fragment slot="lead">
						<span class="text-primary-400">
							<WebIcon />
						</span>
					</svelte:fragment>
					<svelte:fragment slot="summary">
						<span class="text-secondary-200">Install from URL</span>
					</svelte:fragment>
					<svelte:fragment slot="content">
						<div class="w-full flex flex-col gap-2">
							<div class="flex flex-row gap-2 items-end w-full">
								<label class="label w-full">
									<span>URL to install from</span>
									<input
										class="input w-full text-secondary-200"
										bind:value={fetchUrl}
										type="text"
										placeholder="https://example.com/download.zip"
									/>
								</label>

								{#if !$fetchService.isLoading && fetchUrl.length > 4}
									<button
										class="btn variant-filled-primary"
										on:click={() => $fetchService.mutate(fetchUrl)}>install</button
									>
								{:else}
									<button class="btn variant-soft-primary" disabled>install</button>
								{/if}
							</div>

							{#if $fetchService.isLoading}
								<div class="flex flex-row items-center gap-2 mt-0 px-2 text-xs">
									<Circle size="10" color={colors.gray[200]} />
									<p>Installing service</p>
								</div>
							{:else if $fetchService.isError}
								<div class="flex flex-col gap-2 mt-0 px-2 text-xs">
									<p class="text-error-400">Could not install service:</p>
									<div class="card p-2 px-4 text-red-500 font-mono whitespace-pre-line">
										{errorToText($fetchService.error)}
									</div>
								</div>
							{:else if $fetchService.data}
								<div class="flex flex-row items-center gap-2 mt-0 px-2 text-xs text-success-400">
									<p>
										Service '{$fetchService.data.fq.name}' by '{$fetchService.data.fq.author}' was
										installed successfully
									</p>
								</div>
							{/if}
						</div>
					</svelte:fragment>
				</AccordionItem>
			</Accordion>
			<div class="flex justify-end mt-4 p-6 pt-2">
				<button on:click={() => (modalOpen = false)} class="btn variant-soft-secondary">
					Close
				</button>
			</div>
		</div>
	</div>
{/if}
