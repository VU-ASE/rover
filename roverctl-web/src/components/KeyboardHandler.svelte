<script lang="ts">
	import { config } from '$lib/config';
	import { PipelineApi } from '$lib/openapi';
	import { useEmergencyBrake } from '$lib/queries/pipeline';
	import { globalStore } from '$lib/store';
	import { useMutation, useQueryClient } from '@sveltestack/svelte-query';
	import { onDestroy, onMount } from 'svelte';
	import toast from 'svelte-french-toast';

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
					toast('Emergency reset applied ', {
						icon: '🚨',
						position: 'bottom-right',
						duration: 4000
					});
				})
				.catch((error) => {
					console.error('Error applying emergency reset:', error);
					toast('Could not apply emergency reset', {
						icon: '⚠️',
						position: 'bottom-right',
						duration: 10000
					});
				});
		} else if ((event.ctrlKey || event.metaKey) && event.key === 's') {
			event.preventDefault();
			$stopPipeline.mutateAsync();
		} else if ((event.ctrlKey || event.metaKey) && event.key === 'p') {
			event.preventDefault();
			globalStore.pauseStream();
			toast('Debug collection paused', {
				icon: '⏸️',
				position: 'bottom-right',
				duration: 4000
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
