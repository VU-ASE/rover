import { config } from '$lib/config';
import { HealthApi, PipelineApi, ServicesApi, type FullyQualifiedService } from '$lib/openapi';
import { useMutation, useQueryClient } from '@sveltestack/svelte-query';
import type { PipelineNodeData } from '../../components/manage/type';
import { globalStore } from '$lib/store';

/**
 * The query "hooks" are placed in this file since they are reused among multiple views and shortcuts,
 * and we want only a single source of truth for them.
 */

const useStartPipeline = () => {
	const queryClient = useQueryClient();

	return useMutation(
		'startPipeline',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			const papi = new PipelineApi(config.roverd.api);
			const response = await papi.pipelineStartPost();
			return response.data;
		},
		{
			// Invalidate the pipeline query regardless of mutation success or failure
			onSettled: () => {
				queryClient.invalidateQueries('pipeline');
			},
			onSuccess: () => {
				// Reset the debugging data, we now have a new pipeline
				globalStore.reset();
			}
		}
	);
};

const useBuildService = () => {
	return useMutation('buildService', async (fq: FullyQualifiedService) => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const sapi = new ServicesApi(config.roverd.api);
		const response = await sapi.servicesAuthorServiceVersionPost(fq.author, fq.name, fq.version);
		return response.data;
	});
};

const useSavePipeline = () => {
	return useMutation('savePipeline', async (services: PipelineNodeData[]) => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const papi = new PipelineApi(config.roverd.api);
		const response = await papi.pipelinePost(
			services.map((s) => ({
				fq: s.fq
			}))
		);
		return response.data;
	});
};

const useEmergencyBrake = () => {
	return useMutation('emergencyBrake', async () => {
		if (!config.success) {
			throw new Error('Configuration could not be loaded');
		}

		const hapi = new HealthApi(config.roverd.api);
		const response = await hapi.emergencyPost();
		return response.data;
	});
};

export { useStartPipeline, useBuildService, useSavePipeline, useEmergencyBrake };
