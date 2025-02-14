<script lang="ts">
	import { getClosestTimestamp, getLatestFrame, getLatestNFrames } from '$lib/store/circularbuffer';
	import type { CircularBuffer } from '$lib/store/circularbuffer';
	import type { TimestampedSensorOutput } from '$lib/store/sensorstream';
	import { createId } from '@paralleldrive/cuid2';
	import { onMount } from 'svelte';
	import { globalStore } from '$lib/store';
	import { get } from 'svelte/store';
	import PictureIcon from '~icons/line-md/image-filled';
	import DrawingIcon from '~icons/fluent/whiteboard-16-filled';
	import { Circle } from 'svelte-loading-spinners';
	import { styles } from '$lib/utils/styles';

	let syncWorker: Worker | undefined = undefined;

	// Viewers can toggle the camera (jpegs) and canvas objects
	$: renderCamera = true;
	$: renderCanvas = true;

	onMount(() => {
		// Every 5 seconds, check if there are new camera sources available
		const interval = setInterval(() => {
			for (const [serviceName, service] of $globalStore.services) {
				for (const [endpointName, endpoint] of get(service).endpoints) {
					for (const [sensorId, stream] of endpoint.streams) {
						const streamData = get(stream);
						const lastFrame = getLatestFrame(streamData.received);
						if (
							lastFrame?.sensorData.cameraOutput?.debugFrame?.jpeg ||
							lastFrame?.sensorData.cameraOutput?.debugFrame?.canvas
						) {
							globalStore.addCameraFeed(serviceName, endpointName, sensorId);
						}
					}
				}
			}
		}, 5000);

		// Stop the interval when the component is destroyed
		return () => clearInterval(interval);
	});

	const loadWorker = async () => {
		const SyncWorker = await import('$lib/worker/livestream.worker?worker');
		syncWorker = new SyncWorker.default();

		const canvas = document.getElementById(canvasId) as HTMLCanvasElement;
		if (!canvas) return;
		const workerCanvas = canvas.transferControlToOffscreen();

		// Give ownership of the HTML Canvas to the worker
		syncWorker.postMessage(
			{
				domCanvas: workerCanvas
			},
			[workerCanvas]
		);
	};
	onMount(loadWorker);

	const canvasId = createId();

	$: cameraIndex = 0;
	$: cameraFeed = $globalStore.cameraFeeds[cameraIndex];
	$: cameraService = $globalStore.services.get(cameraFeed?.service || '');
	$: cameraStream = cameraService
		? get(cameraService)
				.endpoints.get(cameraFeed?.endpoint || '')
				?.streams.get(cameraFeed?.sensorId || -1)
		: undefined;

	function handleCameraUpdate(
		currentFrame: TimestampedSensorOutput,
		renderCam: boolean,
		renderCanv: boolean
	) {
		// Try to get the jpeg bytes and canvas objects from the current frame, by parsing the debugFrame as CameraOutput
		const jpeg = renderCam ? currentFrame?.sensorData.cameraOutput?.debugFrame?.jpeg : null;
		const canvasData = renderCanv
			? currentFrame?.sensorData.cameraOutput?.debugFrame?.canvas
			: null;

		// Send everything to the worker to render
		syncWorker?.postMessage({
			jpeg: jpeg,
			canvasData: canvasData
		});
	}

	const getCurrentFrame = (circbuf: CircularBuffer<TimestampedSensorOutput>) => {
		let timestamp = 0;
		if ($globalStore.paused) {
			timestamp =
				$globalStore.paused.valueOf() -
				$globalStore.scrubOffsetMilliseconds +
				$globalStore.carOffset;
		} else {
			timestamp = Date.now() - $globalStore.scrubOffsetMilliseconds + $globalStore.carOffset;
		}

		const foundFrame = getClosestTimestamp(circbuf, timestamp);
		return foundFrame;
	};

	onMount(() => {
		// Start a function to keep polling for chart data at the refresh rate
		let updater = requestAnimationFrame(function chartUpdater() {
			// Don't add data when the stream is paused (or when the chart is not initialized)
			if ($cameraStream && syncWorker) {
				const foundFrame = getCurrentFrame($cameraStream?.received);
				if (foundFrame) {
					handleCameraUpdate(foundFrame, renderCamera, renderCanvas);
				}
			}
			updater = requestAnimationFrame(chartUpdater);
		});
		return () => cancelAnimationFrame(updater);
	});
</script>

<div class={'w-full'}>
	<div class={'card variant-soft w-full h-full relative'}>
		{#if $globalStore.cameraFeeds.length <= 0}
			<div class={'absolute top-0 left-0 w-full h-full flex items-center justify-center'}>
				<div class={'flex flex-col gap-y-4 items-center justify-center'}>
					<Circle
						size="30"
						color={styles.colors.primary}
						unit="px"
						duration="1s"
						pause={!!$globalStore.paused}
					/>
					<div class="text-center px-4">
						<h2 class="text-gray-200">Waiting for camera output</h2>
						<p class="text-gray-300">Camera-compatible outputs will be shown here automatically.</p>
					</div>
				</div>
			</div>
		{/if}

		<div class={'relative'}>
			<canvas id={canvasId} class={'w-full h-full'} />
			{#if cameraIndex < 0}
				<div
					class={'absolute top-0 left-0 w-full h-full flex items-center justify-center bg-gray-700'}
				>
					<div class={'flex flex-col gap-y-20'}>
						<div class="text-center px-4">
							<h2 class="text-gray-200">
								Stream
								<span class="text-orange-300">disabled</span>
							</h2>
							<p class="text-gray-300">Select a feed to resume streaming</p>
						</div>
					</div>
				</div>
			{/if}
		</div>
		<div class={'flex flex-row justify-between px-4 py-2'}>
			{#if $globalStore.cameraFeeds.length > 0}
				<div class={'flex flex-row gap-x-2 items-center'}>
					<p class={'text-gray-400 whitespace-nowrap'}>Camera feed:</p>
					<select
						bind:value={cameraIndex}
						class="px-4 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					>
						<option value={-1}>disabled</option>
						{#each $globalStore.cameraFeeds as feed, index (feed.service)}
							<option value={index}>
								{feed.service}.{feed.endpoint}.{feed.sensorId}
							</option>
						{/each}
					</select>
				</div>

				<div class={'flex flex-row gap-x-4 items-center'}>
					<button
						on:click={() => {
							renderCamera = !renderCamera;
							if (!renderCamera && !renderCanvas) {
								renderCanvas = true;
							}
						}}
						aria-describedby="Toggle camera rendering"
						class={'btn-icon variant-filled ' +
							(renderCamera ? 'variant-soft-success' : 'variant-soft-secondary')}
					>
						<PictureIcon />
					</button>

					<button
						on:click={() => {
							renderCanvas = !renderCanvas;
							if (!renderCamera && !renderCanvas) {
								renderCamera = true;
							}
						}}
						aria-describedby="Toggle canvas rendering"
						class={'btn-icon variant-filled ' +
							(renderCanvas ? 'variant-soft-success' : 'variant-soft-secondary')}
					>
						<DrawingIcon />
					</button>
				</div>
			{/if}
		</div>
	</div>
</div>
