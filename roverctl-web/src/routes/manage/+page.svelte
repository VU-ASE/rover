<script lang="ts">
	import SteeringIcon from '~icons/ix/steering';
	import DebugIcon from '~icons/ix/chart-curve-spline';
	import WarningIcon from '~icons/ix/warning-filled';
	import WifiIcon from '~icons/material-symbols/wifi';
	import WifiOffIcon from '~icons/material-symbols/wifi-off';

	import { Circle } from 'svelte-loading-spinners';
	import { config } from '$lib/config';
	import { useQuery } from '@sveltestack/svelte-query';
	import { HealthApi } from '$lib/openapi';
	import colors from 'tailwindcss/colors';
	import Navbar from '../../components/Navbar.svelte';

	import { Modal } from '@skeletonlabs/skeleton';
	import ErrorOverlay from '../../components/ErrorOverlay.svelte';
	import Pipeline from '../../components/manage/Pipeline.svelte';

	const statusQuery = useQuery(
		'status',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			// Fetch status
			const hapi = new HealthApi(config.roverd.api);
			const status = await hapi.statusGet();
			return status.data;
		},
		{
			staleTime: 10, // 10 seconds
			keepPreviousData: true
		}
	);
</script>

<Navbar />
<div class="h-[50vh]">
	<Pipeline />
</div>

<ErrorOverlay />
