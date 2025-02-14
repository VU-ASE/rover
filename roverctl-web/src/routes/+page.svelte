<script lang="ts">
	import Authenticate from '../views/authenticate.svelte';
	import Connecting from '../views/connecting.svelte';
	import Connected from '../views/connected/index.svelte';
	import { connectionStore } from '$lib/store/connection';
	import { initServerConnection } from '$lib/events/connection';
	import { authStore } from '$lib/store/auth';
	import { get } from 'svelte/store';

	// We want this to be explicit state (instead of derived state from the peerconnection) so that you can still view old data if the connection suddenly fails
	let step: 'authenticate' | 'connecting' | 'connected' = 'authenticate';

	const onConnect = (blockScreen: boolean) => {
		// Generate random client ID
		const clientId = 'client-' + Math.random().toString(36).substr(2, 9);

		const as = get(authStore);
		if (as.passthroughUrl) {
			initServerConnection(as.passthroughUrl, clientId);
			if (blockScreen) {
				step = 'connecting';
			}
		}
	};
</script>

<div class=" h-full w-full flex justify-center items-center animate-fade-out-container">
	{#if step === 'authenticate'}
		<Authenticate on:connect={() => onConnect(true)} />
	{:else if step === 'connecting'}
		<Connecting
			on:connect={() => {
				onConnect(true);
			}}
			on:reauthenticate={() => {
				step = 'authenticate';
			}}
			on:connected={() => {
				step = 'connected';
			}}
		/>
	{:else if step === 'connected'}
		<Connected
			on:connect={() => {
				onConnect(false);
			}}
		/>
	{/if}
</div>
