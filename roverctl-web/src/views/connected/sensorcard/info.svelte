<script lang="ts">
	import type { SensorStreamStore } from '$lib/store/sensorstream';
	import { AccordionItem } from '@skeletonlabs/skeleton';
	import Chart from './chart.svelte';
	import ChartIcon from '~icons/uil/chart-line';
	import WarningIcon from '~icons/si/warning-fill';
</script>

<div class="flex space-x-2 items-center">
	<WarningIcon class="text-gray-400 text-xl" />
	<p class="text-gray-400 text-sm">
		Advanced service information depends on roverd connectivity, which is not supported at this time
	</p>
</div>
