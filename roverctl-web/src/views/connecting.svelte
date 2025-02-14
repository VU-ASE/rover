<script lang="ts">
	import { Circle } from 'svelte-loading-spinners';
	import { styles } from '$lib/utils/styles';
	import CheckIcon from '~icons/tdesign/check-circle-filled';
	import CloseCircle from '~icons/tdesign/close-circle-filled';
	import WarningIcon from '~icons/si/warning-fill';
	import { connectionStore } from '$lib/store/connection';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
	// When the peer connection changes to connected, we can emit a signal to the parent component
	$: {
		if ($connectionStore.server?.connectionState === 'connected') {
			dispatch('connected');
		}
	}

	function reauthenticate() {
		dispatch('reauthenticate');
	}

	function retry() {
		dispatch('connect');
	}
</script>

<div
	class="space-y-6 text-center flex flex-col items-center w-6/12 animate-fade-in animate-fade-out"
>
	<!-- Connection Indicator -->
	<div class="space-y-2 text-center flex flex-col items-center w-full">
		{#if $connectionStore.isConnecting || $connectionStore.server?.connectionState === 'connecting'}
			<!-- Spinner -->
			<div class="flex space-x-4 items-center">
				<Circle size="20" color={styles.colors.primary} unit="px" duration="1s" />
				<p class="text-gray-200">Connecting to passthrough server</p>
			</div>
		{:else if $connectionStore.server?.connectionState === 'connected'}
			<!-- Success Checkmark -->
			<div class="flex space-x-4 items-center animate-fade-in">
				<CheckIcon class="text-green-500 text-2xl" />
				<p class="text-green-500 font-bold">Connected successfully!</p>
			</div>
		{:else}
			<!-- Error Warning -->
			<div class="flex space-x-4 items-center animate-fade-in">
				<CloseCircle class="text-red-500 text-2xl" />
				<p class="text-red-500 font-bold">Failed to connect</p>
			</div>
			{#if $connectionStore.error}
				<p class="text-gray-200 text-sm animate-fade-in">
					{$connectionStore.error}
				</p>
			{/if}
		{/if}

		<!-- Unsupported Message -->
		<div class="flex space-x-2 items-center">
			<WarningIcon class="text-gray-400 text-xl" />
			<p class="text-gray-400 text-sm">
				Connecting to roverd instances is not supported at this time
			</p>
		</div>

		{#if !$connectionStore.isConnecting && $connectionStore.server?.connectionState !== 'connected' && $connectionStore.server?.connectionState !== 'connecting'}
			<div class="flex space-x-4 items-center animate-fade-in">
				<button class="btn variant-filled-surface" on:click={reauthenticate}>
					Re-authenticate
				</button>
				<button class="btn variant-filled" on:click={retry}> Retry </button>
			</div>
		{/if}
	</div>
</div>
