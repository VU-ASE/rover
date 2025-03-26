import { TuningState } from 'ase-rovercom/gen/tuning/tuning';

const sendTuningState = async (channel: RTCDataChannel, state: TuningState) => {
	const serialized = TuningState.encode(state).finish();
	channel.send(serialized);
};

export { sendTuningState };
