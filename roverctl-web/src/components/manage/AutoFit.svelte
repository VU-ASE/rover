<script lang="ts">
	import { onMount } from 'svelte';
	import { useSvelteFlow, useNodes } from '@xyflow/svelte';
	import type { FitViewOptions } from '@xyflow/svelte';

	const { fitView } = useSvelteFlow();
	const nodes = useNodes();

	const fitOptions: FitViewOptions = {
		maxZoom: 1.2,
		minZoom: 0.6
	};

	// Auto fit view when component mounts
	onMount(() => {
		fitView(fitOptions);
	});

	// Auto fit view when nodes change
	$: {
		if ($nodes.length > 0) {
			fitView(fitOptions);
		}
	}
</script>

<!-- No UI needed, just runs fitView() when nodes change -->
