<script lang="ts">
	import { authStore, validAuthStore } from '$lib/store/auth';
	import { createEventDispatcher } from 'svelte';

	$: authValid = validAuthStore($authStore);

	const dispatch = createEventDispatcher();
	function handleConnect() {
		dispatch('connect');
	}
</script>

<div
	class="space-y-6 text-center flex flex-col items-center w-6/12 animate-fade-in animate-fade-out"
>
	<div class="space-y-4 text-center flex flex-col items-center w-full">
		<label class="label flex flex-col items-start w-full">
			<span>Passthrough URL</span>
			<input
				class="input w-full"
				type="url"
				placeholder="E.g. localhost:7500"
				bind:value={$authStore.passthroughUrl}
			/>
		</label>
		<!-- <label class="label flex flex-col items-start w-full">
			<span>Rover ID</span>
			<input
				class="input w-full"
				type="number"
				placeholder="E.g. 12"
				bind:value={$authStore.roverdLocation.id}
			/>
			<p class="text-sm">
				Or
				<button
					on:click={() => ($authStore.roverdLocation = { type: 'ip', address: '' })}
					class="text-primary-500 hover:text-primary-600">use an IP instead</button
				>
			</p>
		</label> -->

		<label class="flex items-center justify-items-start space-x-2 w-full">
			<input class="checkbox" type="checkbox" bind:checked={$authStore.enableRoverd} />
			<p>Use metadata and configuration options from roverd</p>
		</label>

		{#if $authStore.enableRoverd}
			{#if $authStore.roverdLocation.type === 'id'}
				<label class="label flex flex-col items-start w-full">
					<span>Rover ID</span>
					<input
						class="input w-full"
						type="number"
						placeholder="E.g. 12"
						bind:value={$authStore.roverdLocation.id}
					/>
					<p class="text-sm">
						Or
						<button
							on:click={() => ($authStore.roverdLocation = { type: 'ip', address: '' })}
							class="text-primary-500 hover:text-primary-600">use an IP instead</button
						>
					</p>
				</label>
			{:else}
				<label class="label flex flex-col items-start w-full">
					<span>Rover IP address</span>
					<input
						class="input w-full"
						type="text"
						placeholder="E.g. 192.168.1.112"
						bind:value={$authStore.roverdLocation.address}
					/>
					<p class="text-sm">
						Or
						<button
							on:click={() => ($authStore.roverdLocation = { type: 'id', id: 12 })}
							class="text-primary-500 hover:text-primary-600">use the Rover ID instead</button
						>
					</p>
				</label>
			{/if}

			<div
				class="flex flex-col md:flex-row
			space-y-4
			md:space-y-0 md:space-x-8 w-full"
			>
				<label class="label flex flex-col items-start w-full">
					<span>Username</span>
					<input
						class="input w-full"
						type="text"
						placeholder="E.g. debix"
						bind:value={$authStore.username}
					/>
				</label>
				<label class="label flex flex-col items-start w-full">
					<span>Password</span>
					<input
						class="input w-full"
						type="password"
						placeholder="E.g. debix"
						bind:value={$authStore.password}
					/>
				</label>
			</div>
		{/if}
	</div>
	<div class="flex justify-center space-x-2">
		<button class="btn variant-filled" disabled={!authValid} on:click={handleConnect}>
			Connect
		</button>
	</div>
</div>
