<script lang="ts">
	import { globalStore } from '$lib/store';
	import BatteryIcon from '~icons/mdi/battery';
	import BatteryAlertIcon from '~icons/mdi/battery-alert';
	import BatteryChargingIcon from '~icons/mdi/battery-charging';

	// From: https://github.com/VU-ASE/display/blob/aeec785d7b5887e3f0bb298c31d9cdc8e5f02f72/src/main.go#L30
	const EMPTY_BATTERY_VOLTAGE = 14.9;
	const FULL_BATTERY_VOLTAGE = 16.8;
	const PLUGGED_IN_BATTERY_VOLTAGE = 16.9;

	function voltageToPercent(voltage: number): { display: string; percent: number | null } {
		if (voltage < EMPTY_BATTERY_VOLTAGE) {
			return { display: 'turning off...', percent: 0 };
		}
		if (voltage > PLUGGED_IN_BATTERY_VOLTAGE) {
			return { display: 'plugged in', percent: null };
		}
		if (voltage >= FULL_BATTERY_VOLTAGE) {
			return { display: '100%', percent: 100 };
		}
		// percentage must be between 0-100 here
		const percentage = Math.floor(
			((voltage - EMPTY_BATTERY_VOLTAGE) / (FULL_BATTERY_VOLTAGE - EMPTY_BATTERY_VOLTAGE)) * 100
		);
		if (percentage < 1) {
			return { display: '< 1%', percent: 0 };
		}
		return { display: `${percentage}%`, percent: percentage };
	}

	function isBatteryDataStale(sentAt: Date): boolean {
		const now = new Date();
		const ageMs = now.getTime() - sentAt.getTime();
		return ageMs > 60000; // 1 minute in milliseconds
	}

	$: batteryInfo = $globalStore.battery
		? {
				...voltageToPercent($globalStore.battery.voltage),
				voltage: $globalStore.battery.voltage,
				isStale: isBatteryDataStale($globalStore.battery.sentAt)
			}
		: null;

	$: textColor =
		batteryInfo && !batteryInfo.isStale && batteryInfo.percent !== null
			? batteryInfo.percent < 20
				? 'text-red-500'
				: batteryInfo.percent < 40
					? 'text-orange-500'
					: 'text-gray-400'
			: 'text-gray-400';
</script>

<div class="card variant-soft p-2 px-4 text-start">
	<h2>Battery</h2>

	{#if !batteryInfo || batteryInfo.isStale}
		<p class="text-gray-400 text-sm">Not available</p>
	{:else}
		<div class="flex items-center gap-2">
			{#if batteryInfo.display === 'plugged in'}
				<BatteryChargingIcon class="w-4 h-4 text-green-500" />
			{:else if batteryInfo.percent !== null && batteryInfo.percent < 20}
				<BatteryAlertIcon class="w-4 h-4 text-red-500" />
			{:else}
				<BatteryIcon class="w-4 h-4 {textColor}" />
			{/if}
			<p class="text-sm {textColor}">
				{batteryInfo.display}
				{#if batteryInfo.display !== 'plugged in'}
					<span class="opacity-75">({batteryInfo.voltage.toFixed(2)}V)</span>
				{/if}
			</p>
		</div>
	{/if}
</div>
