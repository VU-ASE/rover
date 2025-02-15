import {
	PUBLIC_ROVERD_HOST,
	PUBLIC_ROVERD_PORT,
	PUBLIC_ROVERD_USERNAME,
	PUBLIC_ROVERD_PASSWORD,
	PUBLIC_PASSTHROUGH_HOST,
	PUBLIC_PASSTHROUGH_PORT
} from '$env/static/public';
import { Configuration } from './openapi';

/**
 * This file is used to parse the environment variables
 * and exposes a configuration object that can be used throughout
 * the application.
 */

type AppConfig = {
	roverd: RoverdConfig;
	passthrough?: PassthroughConfig;
};

// How to reach the REST API exposed by roverd
type RoverdConfig = {
	host: string;
	port: number;
	// Authorization
	username: string;
	password: string;
	// openapi configuration object
	api: Configuration;
};

// Used for debugging using passthrough and the transceiver service on the Rover
type PassthroughConfig = {
	host: string;
	port: number;
};

type ParsedConfig =
	| ({
			success: true;
	  } & AppConfig)
	| {
			success: false;
			error: string;
	  };

const parseConfigFromEnv = (): ParsedConfig => {
	try {
		let roverdHost = PUBLIC_ROVERD_HOST;
		if (!roverdHost) {
			throw new Error('ROVERD_HOST environment variable was not set but is required');
		}
		// Strip schema (http:// or https://) from the host
		roverdHost = roverdHost.replace(/(^\w+:|^)\/\//, '');

		const roverdPortStr = PUBLIC_ROVERD_PORT;
		if (!roverdPortStr) {
			throw new Error('ROVERD_PORT environment variable was not set but is required');
		}
		// Try to parse the port as an integer
		const roverdPort = parseInt(roverdPortStr, 10);
		if (isNaN(roverdPort)) {
			throw new Error('ROVERD_PORT environment variable was not a valid integer');
		}

		const roverdUsername = PUBLIC_ROVERD_USERNAME;
		if (!roverdUsername) {
			throw new Error('ROVERD_USERNAME environment variable was not set but is required');
		}

		const roverdPassword = PUBLIC_ROVERD_PASSWORD;
		if (!roverdPassword) {
			throw new Error('ROVERD_PASSWORD environment variable was not set but is required');
		}

		// Try to parse the passthrough host and port
		let passthroughHost = PUBLIC_PASSTHROUGH_HOST;
		const passthroughPortStr = PUBLIC_PASSTHROUGH_PORT;
		let passthroughPort: number | undefined;
		if (passthroughHost && passthroughPortStr) {
			passthroughHost = passthroughHost.replace(/(^\w+:|^)\/\//, '');
			passthroughPort = parseInt(passthroughPortStr, 10);
			if (isNaN(passthroughPort)) {
				throw new Error('PASSTHROUGH_PORT environment variable was not a valid integer');
			}
		}

		return {
			success: true,
			roverd: {
				host: roverdHost,
				port: roverdPort,
				username: roverdUsername,
				password: roverdPassword,
				api: new Configuration({
					basePath: `http://${roverdHost}:${roverdPort}`,
					username: roverdUsername,
					password: roverdPassword
				})
			},
			passthrough:
				passthroughHost && passthroughPort
					? { host: passthroughHost, port: passthroughPort }
					: undefined
		} as const;
	} catch (err) {
		console.error('Could not parse configuration from environment variables:', err);
		let msg = 'Unknown error, see console.';
		if (err instanceof Error) {
			msg = err.message;
		}
		return { success: false, error: msg };
	}
};

const config = parseConfigFromEnv();

export { config };
