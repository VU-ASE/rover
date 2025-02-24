<script lang="ts">
	import PlaybackFrameButton from './playbackframebutton.svelte';
	import { globalStore } from '$lib/store';

	const sliderMin = 0;
	const sliderMax = 1000;
	const sliderRange = sliderMax - sliderMin;
	$: sliderValue = sliderMax;

	$: updateScrub(sliderValue);
	const updateScrub = (sliderVal: number) => {
		if (sliderVal !== sliderMax && !$globalStore.paused) {
			globalStore.pauseStream();
		}

		const scrubValue = sliderRange - sliderVal;
		const proportion = scrubValue / sliderRange;
		const offset = Math.ceil(proportion * $globalStore.millisecondsPreserved);
		$globalStore.scrubOffsetMilliseconds = offset;
	};

	$: updatePlayhead($globalStore.scrubOffsetMilliseconds);
	const updatePlayhead = (offset: number) => {
		const proportion = offset / $globalStore.millisecondsPreserved;
		const scrubValue = sliderRange - Math.ceil(proportion * sliderRange);
		sliderValue = scrubValue;
	};
</script>

<div
	class="flex items-center w-full px-4 py-2 gap-x-2 transition-opacity duration-200 hover:opacity-100"
>
	<input
		type="range"
		min={sliderMin}
		max={sliderMax}
		bind:value={sliderValue}
		class="w-full appearance-none focus:outline-none"
		style="--webkit-appearance: none; appearance: none;"
	/>
	<style>
		input[type='range']::-webkit-slider-thumb {
			cursor: grab;
			margin-top: -6px;
			background: white;
			width: 15px;
			height: 15px;
			border-radius: 50%;
			box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
		}
		input[type='range']::-moz-range-thumb {
			cursor: grab;
			background: white;
			width: 15px;
			height: 15px;
			border-radius: 50%;
			box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
		}
		input[type='range']::-webkit-slider-runnable-track {
			background: white;
			height: 4px;
			border-radius: 9999px;
		}
		input[type='range']::-moz-range-track {
			background: white;
			height: 4px;
			border-radius: 9999px;
		}
	</style>
	<PlaybackFrameButton direction="backward" />
	<PlaybackFrameButton direction="forward" />
</div>
