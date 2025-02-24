<script lang="ts">
	import type { SensorStreamStore } from '$lib/store/sensorstream';
	import { AccordionItem } from '@skeletonlabs/skeleton';
	import Chart from './chart.svelte';
	import ChartIcon from '~icons/uil/chart-line';

	/**
	 * We need to create a separate component to iterate over the nested stores, since stores can only be dereferenced at the top level.
	 * This also allows us to pass colors without having two identical colors next to each other.
	 */

	export let streamStore: SensorStreamStore;
	export let endpoint: string;

	const availableChartColors = [
		'white'
		//   token("colors.red.600"),
		//   token("colors.blue.600"),
		//   token("colors.green.600"),
		//   token("colors.purple.600"),
		//   token("colors.orange.600"),
		//   token("colors.sky.600"),
		//   token("colors.teal.600"),
	];
</script>

{#each Object.entries($streamStore.chartItems) as [key], index}
	<AccordionItem>
		<svelte:fragment slot="lead">
			<ChartIcon />
		</svelte:fragment>
		<svelte:fragment slot="summary">
			{endpoint}.<strong>{key}</strong>
		</svelte:fragment>
		<svelte:fragment slot="content">
			<Chart
				{key}
				{streamStore}
				color={availableChartColors[index % availableChartColors.length] || 'text-blue-600'}
			/>
		</svelte:fragment>
	</AccordionItem>
{/each}
