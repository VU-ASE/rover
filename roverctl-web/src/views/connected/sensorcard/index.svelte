<script lang="ts">
	import type { ServiceStore } from '$lib/store/service';
	import { Accordion, AccordionItem } from '@skeletonlabs/skeleton';
	import ChartCollection from './chartCollection.svelte';
	import GearIcon from '~icons/ph/gear-fill';
	import ConfigIcon from '~icons/icon-park-outline/setting-config';
	import TextIcon from '~icons/oui/logstash-filter';
	import Logs from './logs.svelte';
	import Info from './info.svelte';
	import InfoIcon from '~icons/lucide/info';
	import Configuration from './configuration.svelte';

	export let serviceStore: ServiceStore;
</script>

<div class="flex flex-col w-full gap-y-1">
	<div class="flex flex-row w-full gap-x-2 items-center">
		<h2>
			{$serviceStore.name}
		</h2>
	</div>
	<div class="card variant-soft">
		<Accordion>
			<AccordionItem>
				<svelte:fragment slot="lead">
					<InfoIcon />
				</svelte:fragment>
				<svelte:fragment slot="summary">About</svelte:fragment>
				<svelte:fragment slot="content">
					<Info />
				</svelte:fragment>
			</AccordionItem>
			{#each $serviceStore.endpoints as [key, endpoint]}
				{#each endpoint.streams as [key2, stream]}
					<ChartCollection streamStore={stream} endpoint={key} />
				{/each}
			{/each}
			<AccordionItem>
				<svelte:fragment slot="lead">
					<ConfigIcon />
				</svelte:fragment>
				<svelte:fragment slot="summary">Configuration</svelte:fragment>
				<svelte:fragment slot="content">
					<Configuration serviceName={$serviceStore.name} />
				</svelte:fragment>
			</AccordionItem>
			<AccordionItem>
				<svelte:fragment slot="lead">
					<TextIcon />
				</svelte:fragment>
				<svelte:fragment slot="summary">Logs</svelte:fragment>
				<svelte:fragment slot="content">
					<Logs serviceName={$serviceStore.name} />
				</svelte:fragment>
			</AccordionItem>
		</Accordion>
	</div>
</div>
