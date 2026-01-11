/**
 * Utilities and types for all incoming communication over the data channel
 */

import { globalStore } from '$lib/store';
import { Segment } from 'rovercom/gen/segmentation/segmentation';
import { DebugOutput } from 'rovercom/gen/debug/debug';

// This keeps track of the last frame that was sent to the server (and the bytes associated to it)
// we don't need cool svelte stores for it, because we don't need to react to changes
// but we do need to reset it when the connection is lost, so we use some dirty global variables and export a function for it
let highestpacketId = 0;
let frameSegments: Segment[] = [];

const resetMarkerState = () => {
	highestpacketId = 0;
	frameSegments = [];
};

const handleIncomingDataMessage = async (e: MessageEvent<ArrayBuffer>) => {
	globalStore.setCarConnected(true);

	// Parse the segment information and payload data
	const segment = Segment.decode(new Uint8Array(e.data));

	// Is this for a new frame (higher marker) or a frame we were currently parsing already
	if (segment.packetId > highestpacketId) {
		// ignore the previous (partial) frame
		highestpacketId = segment.packetId;
		frameSegments = [segment];
	} else if (segment.packetId === highestpacketId) {
		// check if the segment for this frame already exists
		const existingSegment = frameSegments.find((s) => s.segmentId === segment.segmentId);
		if (!existingSegment) {
			frameSegments.push(segment);
		}
	}

	// check if this finalizes the marker (i.e. we have all segments)
	if (frameSegments.length === segment.totalSegments) {
		// order all the segments by their segment number
		const sortedSegments = frameSegments.sort((a, b) => a.segmentId - b.segmentId);

		// we have all segments, so we can reconstruct the data
		const data = new Uint8Array(sortedSegments.reduce((acc, s) => acc + s.data.byteLength, 0));
		let offset = 0;
		for (const s of sortedSegments) {
			data.set(new Uint8Array(s.data), offset);
			offset += s.data.byteLength;
		}

		// try to decode as DebugOutput
		try {
			const debugMessage = DebugOutput.decode(data);
			if (!debugMessage || !debugMessage.service || debugMessage.message.length === 0) {
				throw new Error('Invalid debug message');
			}

			globalStore.addFrame(debugMessage);
			return;
		} catch (e) {
			console.log('Could not decode as DebugServiceMessage (but this is not an error)', e);
		}
	}

	// just discard the segment, it is probably old
};

export { handleIncomingDataMessage, resetMarkerState };
