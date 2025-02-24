<script lang="ts">
	import type { ServiceStore } from '$lib/store/service';
	import { Accordion, AccordionItem } from '@skeletonlabs/skeleton';
	import ChartCollection from './chartCollection.svelte';
	import GearIcon from '~icons/ph/gear-fill';
	import ConfigIcon from '~icons/icon-park-outline/setting-config';
	import TextIcon from '~icons/oui/logstash-filter';

	import InfoIcon from '~icons/lucide/info';
	import Configuration from './configuration.svelte';

	export let serviceStore: ServiceStore;
</script>

<div class="flex flex-col w-full gap-y-1">
	<div class="flex flex-row w-full gap-x-2 items-center">
		<h2 class="font-mono text-secondary-700">
			{$serviceStore.name}
		</h2>
	</div>
	<div class="card">
		<Accordion>
			{#each $serviceStore.endpoints as [key, endpoint]}
				{#each endpoint.streams as [key2, stream]}
					<ChartCollection streamStore={stream} endpoint={key} />
				{/each}
			{/each}
			<AccordionItem>
				<svelte:fragment slot="lead">
					<ConfigIcon />
				</svelte:fragment>
				<svelte:fragment slot="summary">Tuning</svelte:fragment>
				<svelte:fragment slot="content">
					<Configuration serviceName={$serviceStore.name} />
				</svelte:fragment>
			</AccordionItem>
		</Accordion>
	</div>
</div>
