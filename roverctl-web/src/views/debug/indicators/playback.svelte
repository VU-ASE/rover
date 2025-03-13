<script lang="ts">
	import { Circle } from 'svelte-loading-spinners';
	import { styles } from '$lib/utils/styles';
	import CheckIcon from '~icons/tdesign/check-circle-filled';
	import CloseCircle from '~icons/tdesign/close-circle-filled';
	import WarningIcon from '~icons/si/warning-fill';
	import { connectionStore } from '$lib/store/connection';
	import { createEventDispatcher } from 'svelte';
	import { globalStore } from '$lib/store';
	import PauseIcon from '~icons/fluent/pause-12-filled';
	import PlayIcon from '~icons/fluent/play-12-filled';
</script>

<div class="card variant-soft p-2 px-4">
	<div class="flex flex-row items-center justify-between">
		<div class="flex flex-col items-start">
			<h2>Stream</h2>
			{#if $globalStore.paused && $globalStore.carConnected}
				<p class="text-gray-400 text-sm">Paused</p>
			{:else if !$globalStore.carConnected}
				<p class="text-gray-400 text-sm">Paused (Rover disconnected)</p>
			{:else}
				<p class="text-green-400 text-sm">Active</p>
			{/if}
		</div>

		<button
			on:click={globalStore.toggleStream}
			disabled={!$globalStore.carConnected}
			aria-describedby="Play/pause the stream"
			class={'btn-icon variant-filled ' +
				(!$globalStore.carConnected
					? 'variant-soft-secondary'
					: $globalStore.paused
						? 'variant-soft-success'
						: 'variant-soft-secondary')}
		>
			{#if $globalStore.paused}
				<PlayIcon class="text-white" />
			{:else}
				<PauseIcon class="text-white" />
			{/if}
		</button>
	</div>
</div>
