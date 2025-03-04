<script lang="ts">
	import Connected from '../../views/debug/index.svelte';
	import { connectionStore } from '$lib/store/connection';
	import { initServerConnection } from '$lib/events/connection';
	import { get } from 'svelte/store';
	import { config } from '$lib/config';
	import { onMount } from 'svelte';
	import { styles } from '$lib/utils/styles';
	import { Circle } from 'svelte-loading-spinners';

	import CheckIcon from '~icons/tdesign/check-circle-filled';
	import CloseCircle from '~icons/tdesign/close-circle-filled';
	import WarningIcon from '~icons/si/warning-fill';
	import Navbar from '../../components/Navbar.svelte';
	import ErrorOverlay from '../../components/ErrorOverlay.svelte';

	const onConnect = () => {
		if (!config.success || !config.passthrough) {
			console.log('Will not establish connection, no passthrough specified');
			return;
		}

		// Generate random client ID
		const clientId = 'client-' + Math.random().toString(36).slice(2, 9);
		initServerConnection(`http://${config.passthrough.host}:${config.passthrough.port}`, clientId);
	};

	onMount(() => {
		onConnect();
	});
</script>

<Navbar page="debug" />
<div class=" h-full w-full flex justify-center items-center animate-fade-out-container">
	<div class="space-y-2 text-center flex flex-col items-center w-full animate-fade-in">
		{#if !config.success}
			<div class="flex flex-row space-x-4 items-center">
				<WarningIcon class="text-warning-500 text-2xl" />
				<p class="text-warning-500 font-bold">Configuration is invalid</p>
			</div>
			<p class="text-secondary-700">
				Please restart <span class="code">roverctl</span>
			</p>
		{:else if !config.passthrough}
			<div class="flex flex-row space-x-4 items-center">
				<WarningIcon class="text-warning-500 text-2xl" />
				<p class="text-warning-500 font-bold">Debugging is not configured</p>
			</div>
			<p class="text-secondary-700">
				You need to run <span class="code">roverctl</span> with the correct options to enable debugging.
			</p>
		{:else if $connectionStore.isConnecting || $connectionStore.server?.connectionState === 'connecting'}
			<!-- Spinner -->
			<div class="flex space-x-4 items-center">
				<Circle size="20" color={styles.colors.primary} unit="px" duration="1s" />
				<p class="text-gray-200">Connecting to passthrough server</p>
			</div>
		{:else if $connectionStore.server?.connectionState === 'connected'}
			<Connected />
		{:else if $connectionStore.server?.connectionState === 'failed'}
			<!-- Error Warning -->
			<div class="flex space-x-4 items-center">
				<CloseCircle class="text-red-500 text-2xl" />
				<p class="text-red-500 font-bold">Failed to connect</p>
			</div>
			{#if $connectionStore.error}
				<p class="text-secondary-700">
					{$connectionStore.error}
				</p>
			{/if}
		{:else if !$connectionStore.server}
			<!-- Spinner -->
			<div class="flex space-x-4 items-center">
				<Circle size="20" color={styles.colors.primary} unit="px" duration="1s" />
				<p class="text-gray-200">Connecting to passthrough server</p>
			</div>
		{:else}
			<!-- Error Warning -->
			<div
				class="flex space
			-x-4 items-center"
			>
				<CloseCircle class="text-red-500 text-2xl" />
				<p class="text-red-500 font-bold">Failed to connect</p>
			</div>
		{/if}
	</div>
</div>
<ErrorOverlay />
