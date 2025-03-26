<script lang="ts">
	export let keys: string[] = []; // e.g. ['Ctrl', 'S']
</script>

<span
	class="inline-flex items-center space-x-1 text-sm font-mono text-gray-700 dark:text-gray-300 ml-1"
>
	{#each keys as key}
		<kbd
			class="rounded border border-gray-300 bg-gray-100 px-1.5 py-0.5 text-xs font-semibold shadow-sm dark:border-gray-600 dark:bg-gray-800"
		>
			{key}
		</kbd>
	{/each}
</span>
