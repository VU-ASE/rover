<script lang="ts">
	import { globalStore } from '$lib/store';
	import WarningIcon from '~icons/si/warning-fill';

	function getTimeDifference(clockDrift: Date) {
		const now = new Date();
		const drift = clockDrift.getTime() - now.getTime();
		const absDrift = Math.abs(drift);
		const isAhead = drift > 0;

		const seconds = Math.floor(absDrift / 1000);
		const minutes = Math.floor(seconds / 60);
		const hours = Math.floor(minutes / 60);
		const days = Math.floor(hours / 24);

		let timeString;
		if (days > 0) {
			timeString = `${days} day${days !== 1 ? 's' : ''}`;
		} else if (hours > 0) {
			timeString = `${hours} hour${hours !== 1 ? 's' : ''}`;
		} else if (minutes > 0) {
			timeString = `${minutes} minute${minutes !== 1 ? 's' : ''}`;
		} else {
			timeString = `${seconds} second${seconds !== 1 ? 's' : ''}`;
		}

		return {
			timeString,
			direction: isAhead ? 'ahead' : 'behind',
			seconds: drift / 1000
		};
	}

	function formatTime(date: Date) {
		return date.toLocaleString('en-US', {
			year: 'numeric',
			month: 'short',
			day: 'numeric',
			hour: '2-digit',
			minute: '2-digit',
			second: '2-digit',
			hour12: false
		});
	}
</script>

{#if $globalStore.clockDrift}
	{@const now = new Date()}
	{@const diff = getTimeDifference($globalStore.clockDrift)}
	<div
		class="w-full bg-yellow-500 text-yellow-900 px-4 py-3 flex items-start space-x-3 border-b-2 border-yellow-600"
	>
		<div class="flex-1 min-w-0">
			<p class="font-bold text-base mb-1">Clock Drift Detected</p>
			<p class="text-sm mb-2">
				The Rover's clock is <strong>{diff.timeString} {diff.direction}</strong> of your computer's clock.
				This will cause debug information to be lost or misaligned.
			</p>
			<div class="text-xs opacity-90 space-y-0.5 font-mono">
				<div>Rover's clock: <strong>{formatTime($globalStore.clockDrift)}</strong></div>
				<div>Your clock: <strong>{formatTime(now)}</strong></div>
			</div>
		</div>
	</div>
{/if}
