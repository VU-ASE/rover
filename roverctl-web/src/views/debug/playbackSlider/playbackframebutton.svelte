<script lang="ts">
	import { globalStore } from '$lib/store';
	import BackwardIcon from '~icons/fluent/previous-20-filled';
	import ForwardIcon from '~icons/fluent/next-20-filled';

	export let direction: 'backward' | 'forward' = 'backward';

	$: btnEnabled =
		(direction === 'backward' &&
			$globalStore.scrubOffsetMilliseconds < $globalStore.millisecondsPreserved - 10) ||
		(direction === 'forward' && $globalStore.scrubOffsetMilliseconds > 10);

	const onClick = () => {
		if (
			direction === 'backward' &&
			$globalStore.scrubOffsetMilliseconds < $globalStore.millisecondsPreserved - 10
		) {
			$globalStore.scrubOffsetMilliseconds += 10;
		} else if (direction === 'forward' && $globalStore.scrubOffsetMilliseconds > 10) {
			$globalStore.scrubOffsetMilliseconds -= 10;
		}

		if (!$globalStore.paused) {
			globalStore.pauseStream();
		}
	};
</script>

<button
	title={direction === 'backward' ? 'Previous frame' : 'Next frame'}
	on:click={onClick}
	class={`cursor-pointer h-8 w-8 rounded-full ${
		btnEnabled ? 'bg-white text-gray-400' : 'bg-gray-300 text-gray-500 cursor-not-allowed'
	}`}
	disabled={!btnEnabled}
>
	<div class="flex items-center justify-center h-full w-full">
		<span class="mb-0 pb-0 text-md">
			{#if direction === 'backward'}
				<BackwardIcon />
			{:else}
				<ForwardIcon />
			{/if}
		</span>
	</div>
</button>
