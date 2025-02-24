<script lang="ts">
	import { globalStore } from '$lib/store';
</script>

<button
	title="Toggle stream playback"
	on:click={globalStore.toggleStream}
	class="cursor-pointer h-10 w-10 rounded-full
      {$globalStore.paused ? 'bg-green-600 text-green-800' : 'bg-red-600 text-red-800'}"
>
	<div class="flex items-center justify-center h-full w-full">
		<!-- <iconify-icon icon={$globalStore.paused ? 'fa:play' : 'fa:stop'} class="mb-0 pb-0 text-md" /> -->
	</div>
</button>
