<script lang="ts">
	import { createId } from '@paralleldrive/cuid2';
	import Chart, { type ChartDataset, type ChartItem } from 'chart.js/auto';
	import 'chartjs-adapter-luxon';
	import { onMount } from 'svelte';
	import type { SensorStreamStore } from '$lib/store/sensorstream';
	import { globalStore } from '$lib/store';
	import { getClosestTimestamp } from '$lib/store/circularbuffer';

	export let streamStore: SensorStreamStore;
	export let key: string;
	export let color: string;

	let value: number | undefined = undefined;
	$: value;

	const chartId = createId();

	// Chart config
	const lineStyle = {
		borderWidth: 0.5,
		pointBorderWidth: 0.5,
		pointRadius: 0.8
	};

	let chart: Chart | null = null;

	const dataset: ChartDataset = {
		label: key,
		data: [],
		borderColor: color,
		...lineStyle
	};

	onMount(() => {
		// Create a new chart with combined data
		const chartCtx = document.getElementById(chartId);
		chart = new Chart(chartCtx as ChartItem, {
			type: 'scatter',
			data: {
				datasets: [dataset]
			},

			plugins: [
				{
					id: 'verticalLign',
					beforeDatasetsDraw: (chart, options, el) => {
						const ctx = chart.ctx;
						const xAxis = chart.scales.x!;
						const yAxis = chart.scales.y!;

						// The timestamp to use
						let timestamp = Date.now();
						if ($globalStore.paused) {
							timestamp =
								$globalStore.paused.valueOf() -
								$globalStore.scrubOffsetMilliseconds +
								$globalStore.carOffset;
						} else {
							timestamp = timestamp - $globalStore.scrubOffsetMilliseconds + $globalStore.carOffset;
						}

						const circbuf = $streamStore.chartItems[key];
						if (!circbuf) {
							return;
						}
						const entry = getClosestTimestamp(circbuf, timestamp);
						if (entry) {
							value = entry.y;
						}

						// Add a playhead line
						const xPosition = xAxis.getPixelForValue(timestamp);
						ctx.save();
						ctx.strokeStyle = 'rgba(194, 194, 194, 0.7)'; // Set the line color
						ctx.lineWidth = 2; // Set the line width
						ctx.beginPath();
						ctx.moveTo(xPosition, yAxis.top);
						ctx.lineTo(xPosition, yAxis.bottom);
						ctx.stroke();
						ctx.restore();
					}
				}
			],
			options: {
				// parsing: false,
				// normalized: true,
				maintainAspectRatio: false,
				// @ts-ignore
				// showLine: true,
				animation: false,
				scales: {
					x: {
						// min-max
						min:
							new Date(
								// Subtract the milliseconds from the current time
								new Date().getTime() - $globalStore.millisecondsPreserved
							).getTime() + $globalStore.carOffset,
						max: Date.now() + $globalStore.carOffset,
						display: false,
						ticks: {
							sampleSize: 3
						}
					},
					y: {
						ticks: {
							color: 'white'
						},
						grid: {
							display: false
						}
					}
				},
				plugins: {
					legend: {
						display: false
					},
					tooltip: {
						enabled: false
					}
				}
			}
		});

		// Start a function to keep polling for chart data at the refresh rate
		let updater = requestAnimationFrame(function chartUpdater() {
			// Don't add data when the stream is paused (or when the chart is not initialized)
			if (chart) {
				const curr = Date.now() + $globalStore.carOffset;

				// The first dataset should always be defined (this is the actual plotted data)
				const chartData = $streamStore.chartItems[key]?.buffer;
				if (chartData) {
					chart.data.datasets[0]!.data = chartData.filter((item) => item);
				}

				// Update the min and max of the x-axis if the stream is not paused
				if (!$globalStore.paused && chart.options.scales?.x) {
					chart.options.scales.x.min = curr - $globalStore.millisecondsPreserved;
					chart.options.scales.x.max = curr;
				} else {
					// If paused, set the x-axis to the moment of pausing and the moment of pausing - the milliseconds preserved
					if (chart.options.scales?.x) {
						chart.options.scales.x.min =
							$globalStore.paused!.valueOf() - $globalStore.millisecondsPreserved;
						chart.options.scales.x.max = $globalStore.paused!.valueOf();
					}
				}

				// update the chart
				chart.update('none');
			}
			updater = requestAnimationFrame(chartUpdater);
		});
		return () => cancelAnimationFrame(updater);
	});
</script>

<div>
	<div class={'w-full h-96 relative'}>
		<canvas id={chartId} />
	</div>
</div>
