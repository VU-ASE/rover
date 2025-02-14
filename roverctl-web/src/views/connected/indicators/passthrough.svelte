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
	import RetryIcon from '~icons/akar-icons/arrow-cycle';

	// Allow users to reauthenticate or retry connecting
	const dispatch = createEventDispatcher();
	function retry() {
		dispatch('connect');
	}
</script>

<div class="card variant-soft p-2 px-4">
	<div class="flex flex-row items-center justify-between">
		<div>
			<h2>Passthrough server</h2>

			{#if $connectionStore.isConnecting || $connectionStore.server?.connectionState === 'connecting'}
				<p class="text-gray-400 text-sm">Connecting</p>
			{:else if $connectionStore.server?.connectionState === 'connected'}
				<p class="text-green-400 text-sm">Connected</p>
			{:else}
				<!-- Error Warning -->
				<!-- <p class="text-red-500 font-bold">disconnected</p> -->
				<!-- <button on:click={retry}>(reconnect)</button> -->
				<!-- {#if $connectionStore.error}
					<p class="text-gray-200 text-sm animate-fade-in">
						{$connectionStore.error}
					</p>
				{/if} -->
				<p class="text-orange-400 text-sm">Disconnected</p>
			{/if}
		</div>

		{#if $connectionStore.isConnecting || $connectionStore.server?.connectionState === 'connecting'}
			<Circle size="20" color={styles.colors.primary} unit="px" duration="1s" />
		{:else if $connectionStore.server?.connectionState === 'connected'}
			<!-- <p class="text-green-400 text-sm">Connected</p> -->
		{:else}
			<button
				on:click={retry}
				aria-describedby="Play/pause the stream"
				class={'btn-icon variant-filled variant-soft-secondary'}
			>
				<RetryIcon class="text-white" />
			</button>
		{/if}
	</div>
</div>
