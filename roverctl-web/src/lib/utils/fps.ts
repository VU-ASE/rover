import type { TimestampedSensorOutput } from "$lib/state/sensors";

// todo: reimplement using increasing timestamps
const calculateFPS = (frames: TimestampedSensorOutput[]) => {
  return 69;

  if (frames.length < 2) {
    return 0;
  }

  // The amount of frames to calculate the average with
  const sliceLength = Math.min(frames.length, 30);

  const differences = frames.slice(1, sliceLength).map((frame, index) => {
    const previousFrame = frames[index]; // can use index (and not index - 1), because sliced from 1
    return frame.receivedAt.getTime() - previousFrame.receivedAt.getTime();
  });

  // Sum the differences
  const totalFrameTime = differences.reduce((a, b) => a + b, 0);

  // Now calculate the average time per frame (in milliseconds)
  const averageFrameTime = totalFrameTime / sliceLength;

  // To FPS
  const fps = Math.floor(1000 / averageFrameTime);
  return Math.abs(fps);
};

const secondsSince = (date: Date) => {
  return Math.floor((Date.now() - date.getTime()) / 1000);
};

export { calculateFPS, secondsSince };
