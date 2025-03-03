import { TuningState } from 'ase-rovercom/gen/tuning/tuning';

const sendTuningState = async (channel: RTCDataChannel, state: TuningState) => {
	console.log('Sending tuning state', state);
	const serialized = TuningState.encode(state).finish();
	channel.send(serialized);
};

export { sendTuningState };
