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
</script>

<div class="card variant-soft p-2 px-4">
	<div class="flex flex-row items-center justify-between">
		<div>
			<h2>Time offset</h2>

			{#if $connectionStore.isConnecting || $connectionStore.server?.connectionState === 'connecting'}
				<p class="text-gray-400 text-sm">Not connected</p>
			{:else if $connectionStore.server?.connectionState === 'connected'}
				<p class="text-gray-400 text-sm">
					{$globalStore.carOffset} ms
				</p>
			{:else}
				<p class="text-gray-400 text-sm">Not connected</p>
			{/if}
		</div>
	</div>
</div>
