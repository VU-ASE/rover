import type {
	FullyQualifiedService,
	PipelineGet200ResponseEnabledInnerProcess,
	ServicesAuthorServiceVersionGet200Response
} from '$lib/openapi';
import type { Node } from '@xyflow/svelte';

type PipelineNodeData = {
	fq: FullyQualifiedService;
	// Configuration and inputs/outputs
	service: ServicesAuthorServiceVersionGet200Response;
	// If this service is running/was run before
	process?: PipelineGet200ResponseEnabledInnerProcess;
	// Functions to manage state in pipeline
	onSetActive?: () => void;
};

type PipelineNode = Node<PipelineNodeData>;

export type { PipelineNode, PipelineNodeData };
