<script lang="ts">
	import { config } from '$lib/config';
	import { PipelineApi } from '$lib/openapi';
	import { useEmergencyBrake } from '$lib/queries/pipeline';
	import { globalStore } from '$lib/store';
	import { useMutation, useQueryClient } from '@sveltestack/svelte-query';
	import { onDestroy, onMount } from 'svelte';
	import { toasts } from 'svelte-toasts';

	const queryClient = useQueryClient();
	const emergencyBrake = useEmergencyBrake();
	const stopPipeline = useMutation(
		'stopPipeline',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);
			const response = await papi.pipelineStopPost();
			return response.data;
		},
		{
			// Invalidate the pipeline query regardless of mutation success or failure
			onSettled: () => {
				queryClient.invalidateQueries('pipeline');
			}
		}
	);

	function handleKeydown(event: KeyboardEvent) {
		// Check for Ctrl+S (Windows/Linux) or Cmd+S (Mac)
		if ((event.ctrlKey || event.metaKey) && event.key === 'e') {
			event.preventDefault();
			$emergencyBrake
				.mutateAsync()
				.then(() => {
					toasts.add({
						title: 'Emergency Reset Applied',
						description: 'Rover was reset',
						duration: 10000,
						placement: 'bottom-right',
						type: 'info',
						theme: 'dark',
						onClick: () => {},
						onRemove: () => {}
					});
				})
				.catch((error) => {
					toasts.add({
						title: 'Emergency Reset Failed',
						description: error.message,
						duration: 10000,
						placement: 'bottom-right',
						type: 'error',
						theme: 'dark',
						onClick: () => {},
						onRemove: () => {}
					});
				});
		} else if ((event.ctrlKey || event.metaKey) && event.key === 's') {
			event.preventDefault();
			$stopPipeline.mutateAsync().then(() => {
				toasts.add({
					title: 'Pipeline Stopped',
					description: 'The pipeline was stopped',
					duration: 10000,
					placement: 'bottom-right',
					type: 'info',
					theme: 'dark',
					onClick: () => {},
					onRemove: () => {}
				});
			});
		} else if ((event.ctrlKey || event.metaKey) && event.key === 'p') {
			event.preventDefault();
			globalStore.pauseStream();
			toasts.add({
				title: 'Debug collection Paused',
				description: 'New frames will be discarded',
				duration: 10000,
				placement: 'bottom-right',
				type: 'info',
				theme: 'dark',
				onClick: () => {},
				onRemove: () => {}
			});
		}
	}

	// Keyboard handler
	onMount(() => {
		window.addEventListener('keydown', handleKeydown);

		onDestroy(() => {
			window.removeEventListener('keydown', handleKeydown);
		});
	});
</script>
