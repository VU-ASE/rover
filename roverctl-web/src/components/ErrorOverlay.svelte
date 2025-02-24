<script lang="ts">
	/**
	 * This component renders an error overlay when roverd reports itself as not operational.
	 * It fetches the status from the roverd API and displays a warning message with instructions.
	 */

	import WarningIcon from '~icons/ix/warning-filled';

	import { config } from '$lib/config';
	import { useQuery } from '@sveltestack/svelte-query';
	import { HealthApi } from '$lib/openapi';

	const statusQuery = useQuery(
		'status',
		async () => {
			if (!config.success) {
				throw new Error('Configuration could not be loaded');
			}

			// Fetch status
			const hapi = new HealthApi(config.roverd.api);
			const status = await hapi.statusGet();
			return status.data;
		},
		{
			staleTime: 10
		}
	);
</script>

{#if $statusQuery.isSuccess && $statusQuery.data && $statusQuery.data.status !== 'operational'}
	<div
		class="w-full h-full absolute top-0 left-0 flex justify-center items-center animate-fade-out-container text-white"
	>
		<!-- Background with 20% opacity -->
		<div class="absolute inset-0 bg-error-500 bg-opacity-90"></div>

		<!-- Content with full opacity -->
		<div class="relative flex flex-col gap-2 items-center text-center">
			<WarningIcon class="text-4xl mb-2" />
			<h1 class="text-xl">roverd is not operational</h1>
			<p>
				The roverd process reported itself as <span class="code variant-filled-error"
					>{$statusQuery.data?.status || 'unknown'}</span
				>.<br />
				Resolve the issue by checking the logs over SSH.
			</p>
			{#if config.success}
				<p class="code text-left w-full variant-glass-secondary p-1 px-2">
					<span class="text-secondary-400">
						# SSH into the Rover <br />
						ssh {config.roverd.username}@{config.roverd.host}<br /><br />

						# View status<br />
						sudo systemctl status roverd<br /><br />

						# View logs<br />
						sudo journalctl -u roverd -n 100 -f<br />
					</span>
				</p>
			{/if}
		</div>
	</div>
{/if}
