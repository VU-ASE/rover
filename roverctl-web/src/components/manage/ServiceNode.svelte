<script lang="ts">
	import type { PipelineGet200ResponseEnabledInnerService } from '$lib/openapi';
	import { Handle, MarkerType, Position, type Edge, type NodeProps } from '@xyflow/svelte';
	import type { Writable } from 'svelte/store';
	import DeleteIcon from '~icons/ic/baseline-delete';

	import CheckmarkIcon from '~icons/heroicons/check-badge-20-solid';
	import type { PipelineNode, PipelineNodeData } from './type';
	import { createEventDispatcher } from 'svelte';

	type $$Props = NodeProps & {
		data: PipelineNodeData;
	};
	export let data: PipelineNodeData;
	$$restProps;
</script>

<div class="card variant-outline-primary w-[200px] h-[80px] relative">
	<!-- NB: The width and height values are hardcoded and correspond to the node w/h in pipeline.svelte, required for dagre to work -->

	<!-- Overlapping Source and Target Handles -->
	<Handle
		type="source"
		position={Position.Right}
		id="center-source"
		class="!absolute !transform !translate-x-[-50%] !translate-y-[-50%] !left-[50%] !top-[50%] w-4 h-4 opacity-0"
		style="pointer-events: none;"
	/>
	<Handle
		type="target"
		position={Position.Right}
		id="center-target"
		class="!absolute !transform !translate-x-[-50%] !translate-y-[-50%] !left-[50%] !top-[50%] w-4 h-4 opacity-0"
		style="pointer-events: none;"
	/>

	<!-- Todo: show errors when we already see validation/missing dependency errors -->
	<!-- Show a red small banner 50% over the bottom edge of the node -->
	<!-- <div
			class="absolute bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2 bg-red-500 text-white text-xs flex items-center justify-center px-2 py-1 rounded-md"
		>
			<span class="whitespace-nowrap">Unmet dependencies</span>
		</div> -->

	<!-- {#if data.onDelete}
		<button
			class="absolute top-0 right-0 translate-x-1/2 -translate-y-1/2 w-6 h-6 rounded-full bg-red-500 hover:bg-red-600 text-white flex items-center justify-center shadow-md"
			on:click={() => data.onDelete && data.onDelete()}
			title="Delete Node"
		>
			<DeleteIcon />
		</button>
	{/if} -->

	<div class="flex flex-col p-2 px-4">
		<div class="flex flex-row items-center gap-1 text-sm">
			{#if data.fq.author.toLowerCase() === 'vu-ase'}
				<span class="text-primary-500">
					<CheckmarkIcon />
				</span>
			{/if}
			<p class="text-secondary-700">{data.fq.author}</p>
		</div>
		<h1 class="truncate w-full overflow-hidden whitespace-nowrap text-primary-500 font-mono">
			{data.fq.name}
		</h1>
		<p class="text-xs truncate w-full overflow-hidden whitespace-nowra text-secondary-800">
			{data.fq.version}
		</p>
	</div>
</div>
