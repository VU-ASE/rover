import { controlChannelLabel, dataChannelLabel } from '$lib/constants/channels';
import { iceServers } from '$lib/constants/connection';
import { connectionStore } from '$lib/store/connection';
import { notify } from '$lib/events/notifications';
import { handleIncomingDataMessage } from '$lib/events/data';
import { handleIncomingControlMessage } from './control';
import { globalStore } from '$lib/store';

const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

const initServerConnection = async (serverAddress: string, clientId: string) => {
	connectionStore.update(() => ({
		isConnecting: true,
		error: null
	}));

	try {
		// Create a new peer connection
		console.log(`Attempting to connect to server at ${serverAddress}`);
		const pc = new RTCPeerConnection({
			iceServers,
			iceCandidatePoolSize: 2
		});

		// Strip the trailing / from the server address
		if (serverAddress.endsWith('/')) {
			serverAddress = serverAddress.slice(0, -1);
		}
		if (!serverAddress.startsWith('http://') && !serverAddress.startsWith('https://')) {
			serverAddress = `http://${serverAddress}`;
		}

		// Handle ICE candidates received from the server
		pc.onicecandidate = async (e) => {
			// Sleep to make sure we don't send ICE candidates before our SDP offer was sent
			await sleep(5000);

			if (e.candidate) {
				try {
					console.log('Received ICE candidate', e.candidate, 'sending to server');
					const response = await fetch(`${serverAddress}/client/ice`, {
						method: 'POST',
						body: JSON.stringify({
							...e.candidate,
							id: clientId
						})
					});
					if (!response.ok) {
						const { message } = (await response.json()) as { message: string };

						if (message) {
							throw new Error(`ICE candidate sending failed: ${message}`);
						} else {
							throw new Error(`ICE candidate sending failed: ${response.statusText}`);
						}
					}
					const remoteCandidates = await response.json();
					console.log(
						'Successfully sent ICE candidates to server. Received remote candidates',
						remoteCandidates
					);

					// Add all ICE candidates, so that we know how to reach the server
					for (const candidate of remoteCandidates) {
						const c = new RTCIceCandidate(candidate);
						await pc.addIceCandidate(c);
					}
				} catch (e) {
					console.error(e);
					if (e instanceof Error) {
						throw new Error(
							`Failed to send ICE candidate: ${e.message}. Are you sure the passthrough server is running and accessible?`
						);
					}
					throw e;
				}
			}
		};

		// Register a connection state change handler, update the store if a connection is lost
		pc.onconnectionstatechange = () => {
			console.log(
				'Server connection state has changed. New connection state is',
				pc.connectionState
			);
			if (
				pc.connectionState === 'disconnected' ||
				pc.connectionState === 'failed' ||
				pc.connectionState === 'closed'
			) {
				connectionStore.reset();
				globalStore.setCarConnected(false);
			} else {
				connectionStore.update(() => ({
					server: pc
				}));
			}

			switch (pc.connectionState) {
				case 'connected':
					notify('success', 'Successfully connected to server');
					break;
				case 'disconnected':
					notify('error', 'Please reconnect to continue');
					break;
			}
		};

		// Register a data channel creation handler (the server will create the data channel for us)
		pc.ondatachannel = (e) => {
			notify(
				'error',
				`Server created new data channel ${e.channel.label}. This is not supposed to happen.`
			);
		};

		// Create the data and control channels
		const controlChan = pc.createDataChannel(controlChannelLabel);
		controlChan.addEventListener('message', handleIncomingControlMessage);
		const dataChan = pc.createDataChannel(dataChannelLabel);
		dataChan.addEventListener('message', handleIncomingDataMessage);

		// Create an SDP offer for the server to answer
		const offer = await pc.createOffer();
		await pc.setLocalDescription(offer); // this will trigger ICE candidate gathering

		// Send the offer to the HTTP server
		console.log('Remote connection configured, will send final offer to server');
		try {
			const response = await fetch(`${serverAddress}/client/sdp`, {
				method: 'POST',
				body: JSON.stringify({
					offer: offer,
					id: clientId
				})
			});
			if (!response.ok) {
				const { message } = (await response.json()) as { message: string };

				if (message) {
					throw new Error(`SDP sending failed: ${message}`);
				} else {
					throw new Error(`SDP sending failed: ${response.statusText}`);
				}
			}

			// The SDP answer created by the server
			console.log('Received SDP answer from server');
			const answer = await response.json();
			const remoteDesc = new RTCSessionDescription(answer);
			pc.setRemoteDescription(remoteDesc);
			console.log('Remote description was set, connection should be established');
		} catch (e) {
			console.error(e);
			if (e instanceof Error) {
				throw new Error(
					`Failed to send SDP offer: ${e.message}. Are you sure the passthrough server is running and accessible?`
				);
			}
			throw e;
		}

		connectionStore.update(() => ({
			clientId: clientId,
			server: pc,
			controlChannel: controlChan,
			dataChannel: dataChan,
			isConnecting: false
		}));
	} catch (e) {
		globalStore.setCarConnected(false);
		if (e instanceof Error) {
			notify('error', e.message);
			connectionStore.update(() => ({
				error: e.message,
				isConnecting: false
			}));
		} else {
			connectionStore.update(() => ({
				error: 'An unknown error occurred',
				isConnecting: false
			}));
			notify('error', 'An unknown error occurred');
		}
	}
};

export { initServerConnection };
