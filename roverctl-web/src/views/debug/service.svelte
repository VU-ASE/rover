<script lang="ts">
	import type { ServiceState, ServiceStore } from '$lib/store/service';

	export let service: ServiceStore | null = null;
</script>

<p>
	{#if $service}
		{$service.name}
	{:else}
		No services available
	{/if}
</p>
