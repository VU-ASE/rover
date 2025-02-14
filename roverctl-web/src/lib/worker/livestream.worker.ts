import type { CameraSensorOutput_DebugFrame } from 'ase-rovercom/gen/outputs/camera';

let latestJpeg: Uint8Array | null = null;

const populateCanvas = async (
	jpeg: Uint8Array | null,
	canvasData: CameraSensorOutput_DebugFrame['canvas'],
	canvas: OffscreenCanvas,
	ctx: CanvasRenderingContext2D
) => {
	// clear the canvas
	ctx.clearRect(0, 0, canvas.width, canvas.height);

	// draw a white background
	ctx.fillStyle = 'black';
	ctx.fillRect(0, 0, canvas.width, canvas.height);

	if (jpeg && jpeg !== latestJpeg) {
		latestJpeg = jpeg;
		// create a blob from the jpeg
		const blob = new Blob([jpeg], { type: 'image/jpeg' });
		// convert to bitmap
		const bmp = await createImageBitmap(blob);

		canvas.width = bmp.width;
		canvas.height = bmp.height;

		// Show on the canvas
		ctx.drawImage(
			bmp,
			0,
			0,
			bmp.width,
			bmp.height, // source rectangle
			0,
			0,
			canvas.width,
			canvas.height
		);
	}

	if (canvasData) {
		if (!canvasData?.width || !canvasData?.height) {
			console.error('Canvas descriptor has no width or height', canvasData);
			return;
		}

		// Scale factor to make sure that the image fits the canvas
		// const scaleFactor = Math.min(
		//   canvas.width / canvasData.width,
		//   canvas.height / canvasData.height
		// );

		// draw the objects
		if (canvasData?.objects) {
			for (const object of canvasData.objects) {
				if (object.circle) {
					// We assume 1 point radius circles for now, so we can draw them as rectangles
					const { center } = object.circle;
					const x = center?.x ?? 0;
					const y = center?.y ?? 0;

					if (object.circle.color) {
						ctx.fillStyle = `rgb(${object.circle.color.r}, ${object.circle.color.g}, ${object.circle.color.b})`;
					} else {
						ctx.fillStyle = 'green';
					}

					// todo: add support for radii and different objects
					ctx.fillRect(x, y, 4, 4);
					// ctx.fillRect(x, y, x + 1, y + 1);
				} else {
					console.error('Unknown object type', object);
				}
			}
		}
	}
};

// Global context
let globalCanvas: OffscreenCanvas | null = null;
let globalCanvasCtx: CanvasRenderingContext2D | null = null;

onmessage = (e) => {
	// dom canvas is a reference to the HTML canvas element
	// jpeg is the jpeg bytes we have received
	// canvasData is the list of objects that we need to draw on the canvas
	const { domCanvas, jpeg, canvasData } = e.data;
	if (!domCanvas && !jpeg && !canvasData) {
		return;
	}

	if (domCanvas) {
		if (!globalCanvas) {
			globalCanvas = domCanvas;
		}

		if (!globalCanvasCtx) {
			globalCanvasCtx = domCanvas.getContext('2d');
		}

		if (!globalCanvasCtx) {
			console.error('Could not get canvas context', e.data);
			return;
		}
	} else {
		if (!globalCanvasCtx || !globalCanvas) {
			console.error('Canvas or context not initialized');
			return;
		}

		populateCanvas(jpeg, canvasData, globalCanvas, globalCanvasCtx);
	}
};

export {};
