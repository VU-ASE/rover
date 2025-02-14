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
	import { CIRCULAR_BUFFER_SIZE } from '$lib/store/circularbuffer';
</script>

<div class="card variant-soft p-2 px-4">
	<div class="flex flex-row items-center justify-between">
		<div>
			<h2>Buffer size</h2>

			<p class="text-gray-400 text-sm">
				{CIRCULAR_BUFFER_SIZE} frames
			</p>
		</div>
	</div>
</div>
