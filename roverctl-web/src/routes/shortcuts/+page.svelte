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
	import { onDestroy, onMount } from 'svelte';
	import { errorToText } from '$lib/errors';

	import FireIcon from '~icons/game-icons/celebration-fire';
	import DisconnectIcon from '~icons/wpf/disconnected';
	import SpeedIcon from '~icons/ic/baseline-speed';
	import StopIcon from '~icons/healthicons/stop';
	import EmergencyIcon from '~icons/mdi/car-emergency';
	import PauseIcon from '~icons/zondicons/pause-outline';

	import { browser } from '$app/environment';

	const accept = () => {
		document.cookie = 'shortcuts=true; path=/; max-age=' + 60 * 60 * 24 * 365;
		window.location.href = '/';
	};

	import WarningIcon from '~icons/ix/warning-filled';
	import Kbd from '../../components/Kbd.svelte';
</script>

<div class="flex h-full w-full items-center justify-center animate-fade-out-container">
	<div class="flex flex-col gap-2 items-center text-center animate-fade-in">
		<h1 class="text-4xl text-secondary-700 mb-4">
			Use these <span class="text-primary-500">shortcuts</span>
		</h1>

		<div
			class="card p-4 w-full px-6 *: variant-soft
		flex flex-row items-center justify-start space-x-4"
		>
			<StopIcon class="text-primary-500 text-4xl" />
			<p class="text-secondary-700">
				Stop the pipeline quickly with <Kbd keys={['ctrl', 's']} />
			</p>
		</div>
		<div
			class="card p-4 w-full px-6 *: variant-soft
	flex flex-row items-center justify-start space-x-4"
		>
			<EmergencyIcon class="text-primary-500 text-4xl" />
			<p class="text-secondary-700">
				Emergency reset the Rover with <Kbd keys={['ctrl', 'e']} />
			</p>
		</div>
		<div
			class="card p-4 w-full px-6 *: variant-soft
flex flex-row items-center justify-start space-x-4"
		>
			<PauseIcon class="text-primary-500 text-4xl" />
			<p class="text-secondary-700">
				Pause debugging capture with <Kbd keys={['ctrl', 'p']} />
			</p>
		</div>

		<button class="btn variant-soft-primary mt-2" on:click={accept}> I understand </button>
	</div>
</div>
