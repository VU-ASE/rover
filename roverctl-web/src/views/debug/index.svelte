<script lang="ts">
	import { Circle } from 'svelte-loading-spinners';
	import { styles } from '$lib/utils/styles';
	import CheckIcon from '~icons/tdesign/check-circle-filled';
	import CloseCircle from '~icons/tdesign/close-circle-filled';
	import WarningIcon from '~icons/si/warning-fill';
	import { connectionStore } from '$lib/store/connection';
	import { createEventDispatcher } from 'svelte';
	import { globalStore } from '$lib/store';
	import Service from './service.svelte';
	import ChartCollection from './sensorcard/chartCollection.svelte';
	import SensorCard from './sensorcard/index.svelte';
	import Videostream from './videostream.svelte';
	import PlaybackSlider from './playbackSlider/index.svelte';
	import PlaybackIndicator from './indicators/playback.svelte';
	import PassthroughIndicator from './indicators/passthrough.svelte';
	import OffsetIndicator from './indicators/offset.svelte';
	import BufferSizeIndicator from './indicators/buffer.svelte';
	import CacheSizeIndicator from './indicators/cache.svelte';
	import DelayIndicator from './indicators/delay.svelte';
</script>

<div class="flex flex-col min-h-screen w-full relative">
	<!-- Main Content -->
	{#if $globalStore.services.size <= 0}
		<div class="flex-1 flex items-center justify-center">
			<div class="space-y-2 text-center animate-fade-in animate-fade-out w-full">
				<h1 class="text-green-500 text-lg font-bold">Connection established</h1>
				<p class="text-secondary-700">Waiting for incoming debugging data.</p>
			</div>
		</div>
	{:else}
		<div
			class=" animate-fade-out w-full grid grid-cols-1 md:grid-cols-2 p-2 gap-x-4 gap-y-2 items-start"
		>
			<div class="flex flex-col w-full gap-y-2">
				<Videostream />
				<div class="w-full grid grid-cols-3 gap-x-2 gap-y-2">
					<PlaybackIndicator />
					<PassthroughIndicator />
					<OffsetIndicator />
					<BufferSizeIndicator />
					<CacheSizeIndicator />
					<DelayIndicator />
				</div>
			</div>

			<div class="flex flex-col gap-y-2 w-full">
				{#each Array.from($globalStore.services.values()) as service}
					<SensorCard serviceStore={service} />
				{/each}
			</div>
		</div>
	{/if}

	{#if $globalStore.services.size <= 0}
		<footer class="py-2 flex flex-col items-center text-center">
			<div class="flex space-x-2">
				<p class="text-gray-400">Server status:</p>

				{#if $connectionStore.isConnecting || $connectionStore.server?.connectionState === 'connecting'}
					<!-- Spinner -->
					<div class="flex space-x-4 items-center">
						<Circle size="20" color={styles.colors.primary} unit="px" duration="1s" />
						<p class="text-gray-200">connecting</p>
					</div>
				{:else if $connectionStore.server?.connectionState === 'connected'}
					<!-- Success Checkmark -->
					<p class="text-green-500 font-bold">connected</p>
				{:else}
					<!-- Error Warning -->
					<p class="text-red-500 font-bold">disconnected</p>
					{#if $connectionStore.error}
						<p class="text-gray-200 text-sm animate-fade-in">
							{$connectionStore.error}
						</p>
					{/if}
				{/if}
			</div>
		</footer>
	{/if}
</div>

{#if $globalStore.services.size > 0}
	<div class="absolute bottom-0 left-0 w-full card bg-gray-700">
		<PlaybackSlider />
	</div>
{/if}
