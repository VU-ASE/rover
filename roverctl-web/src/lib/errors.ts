import { AxiosError } from 'axios';
import type { PipelineSetError, GenericError, BuildError } from './openapi';

const errorToText = (error: unknown): string => {
	if (error instanceof AxiosError) {
		const data = error.response?.data;
		if (data && typeof data === 'object') {
			// Check if the data conforms to the RoverdError schema
			if ('errorType' in data && 'errorValue' in data) {
				const { errorType, errorValue } = data as {
					errorType: string;
					errorValue: unknown;
				};

				switch (errorType) {
					case 'generic': {
						// Cast errorValue to GenericError if possible
						const generic = errorValue as GenericError;
						return generic.message;
					}
					case 'build': {
						// Cast errorValue to BuildError
						const build = errorValue as BuildError;
						const buildLog =
							Array.isArray(build.build_log) && build.build_log.length > 0
								? '\n' + build.build_log.join('\n')
								: '';
						return buildLog;
					}
					case 'pipeline_set': {
						// Cast errorValue to PipelineSetError
						const pipelineSet = errorValue as PipelineSetError;
						const ve = pipelineSet.validation_errors;
						const errors: string[] = [];

						if (ve.unmet_streams && ve.unmet_streams.length > 0) {
							for (const us of ve.unmet_streams) {
								errors.push(`Enabled service '${us.source}' depends on stream '${us.stream}' from service '${us.target}', however this service was not enabled or does not output this stream.
                                    -> Enable service '${us.target}' OR disable service '${us.source}'
                                    `);
							}
						}
						if (ve.unmet_services && ve.unmet_services.length > 0) {
							for (const us of ve.unmet_services) {
								errors.push(`Enabled service '${us.source}' depends on service '${us.target}', however this service was not enabled.
                                    -> Enable service '${us.target}' OR disable service '${us.source}'
                                    `);
							}
						}
						if (ve.duplicate_services && ve.duplicate_services.length > 0) {
							for (const ds of ve.duplicate_services) {
								errors.push(`A service with name '${ds}' was enabled more than once.
                                    -> Enable only one '${ds}' service`);
							}
						}
						if (ve.duplicate_aliases && ve.duplicate_aliases.length > 0) {
							for (const da of ve.duplicate_aliases) {
								errors.push(`An alias with name '${da}' was enabled more than once.
                                    -> Enable only one service with '${da}' alias`);
							}
						}
						if (ve.aliases_in_use && ve.aliases_in_use.length > 0) {
							for (const aiu of ve.aliases_in_use) {
								errors.push(`The alias '${aiu}' is already in use by a named service and cannot be used.
                                    -> Disable service '${aiu}' OR disable the service that uses this alias`);
							}
						}
						return errors.join('\n\n') || 'Unknown error occurred setting pipeline';
					}
					default:
						// Unknown errorType; fallback to the Axios error message.
						return error.message;
				}
			}
		}
		// If no structured error data is available, fallback to the Axios error message.
		return error.message;
	}
	// Fallback for non-Axios errors.
	return 'An unexpected error occurred.';
};

export { errorToText };
