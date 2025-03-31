import { config } from '$lib/config';
import { HealthApi, PipelineApi, ServicesApi, type FullyQualifiedService } from '$lib/openapi';
import { useMutation } from '@sveltestack/svelte-query';
import type { PipelineNodeData } from '../../components/manage/type';

/**
 * The query "hooks" are placed in this file since they are reused among multiple views and shortcuts,
 * and we want only a single source of truth for them.
 */

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

export { useBuildService, useSavePipeline, useEmergencyBrake };
