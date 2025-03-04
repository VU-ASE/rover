<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WifiIcon from '~icons/material-symbols/wifi';
	import WifiOffIcon from '~icons/material-symbols/wifi-off';

	import { Circle } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { isError, useMutation, useQuery } from '@sveltestack/svelte-query';
	import { HealthApi } from '$lib/openapi';
	import colors from 'tailwindcss/colors';
	import Navbar from '../../components/Navbar.svelte';

	import { Modal } from '@skeletonlabs/skeleton';
	import ErrorOverlay from '../../components/ErrorOverlay.svelte';
	import Pipeline from '../../components/manage/Pipeline.svelte';
	import { onMount } from 'svelte';
	import { errorToText } from '$lib/errors';

	import FireIcon from '~icons/game-icons/celebration-fire';
	import DisconnectIcon from '~icons/wpf/disconnected';
	import SpeedIcon from '~icons/ic/baseline-speed';

	import { browser } from '$app/environment';

	const accept = () => {
		document.cookie = 'visited=true; path=/; max-age=' + 60 * 60 * 24 * 365;
		window.location.href = '/';
	};

	import WarningIcon from '~icons/ix/warning-filled';
</script>

<div class="flex h-full w-full items-center justify-center animate-fade-out-container">
	<div class="flex flex-col gap-2 items-center text-center animate-fade-in">
		<h1 class="text-4xl text-secondary-700 mb-4">
			Understand the <span class="text-error-500">dangers</span>
		</h1>

		<div
			class="card p-4 w-full px-6 *: variant-soft
		flex flex-row items-center justify-start space-x-4"
		>
			<FireIcon class="text-error-500 text-4xl" />
			<p class="text-secondary-700">
				Never, <strong>ever</strong> leave charging batteries unattended
			</p>
		</div>
		<div
			class="card p-4 w-full px-6 *: variant-soft
	flex flex-row items-center justify-start space-x-4"
		>
			<DisconnectIcon class="text-error-500 text-4xl" />
			<p class="text-secondary-700">
				Unplug power <strong>immediately</strong> when you see the disconnect symbol
			</p>
		</div>
		<div
			class="card p-4 w-full px-6 *: variant-soft
flex flex-row items-center justify-start space-x-4"
		>
			<SpeedIcon class="text-error-500 text-4xl" />
			<p class="text-secondary-700">
				Do <strong>not</strong> drive at full speed without notifying peers
			</p>
		</div>

		<p class="text-secondary-700 mt-2">
			You can (and should) read the full document
			<a
				href="https://ase.vu.nl/docs/tutorials/First%20Steps/dangers"
				class="underline text-primary-400">here</a
			>
		</p>

		<button class="btn variant-soft-error mt-2" on:click={accept}> I understand </button>
	</div>
</div>
