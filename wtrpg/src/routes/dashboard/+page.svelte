<script>
	// $props() is for taking from parent
	let show_game_creation = $state(false);
	const todos = $state([
		{ id: 1, done: false, description: 'Name of the game', value: '' },
		{ id: 2, done: false, description: 'RPG system played', value: '' },
		{ id: 4, done: false, description: 'Description', value: '' }
	]);
	const game_ready = $derived(todos.every((tbd) => tbd.done));
</script>

<div class="container">
	<div>
		{#if show_game_creation}
			<div role="presentation" class="background_when_creating_game">
				<ul>
					{#each todos as tbd (tbd.id)}
						<li class={{ done: tbd.done }}>
							<input type="checkbox" bind:checked={tbd.done} disabled />
							<span>{tbd.description}</span>
							<input
								type="text"
								value={tbd.value}
								oninput={(target) => {
									tbd.done = target.currentTarget.value !== '';
									tbd.value = target.currentTarget.value;
								}}
							/>
						</li>
					{/each}
					<button
						disabled={!game_ready}
						onclick={() => {
							show_game_creation = !show_game_creation;
							todos.forEach((tbd) => {
								tbd.value = '';
								tbd.done = false;
							});
						}}
					>
						Create Game!
					</button>
				</ul>
			</div>
		{/if}
		<div class="controls">
			<button onclick={() => (show_game_creation = !show_game_creation)}>
				{show_game_creation ? 'close' : 'create_game'}
			</button>
		</div>
	</div>
	<button>
		<a href="/game">Get yo game on twin</a>
	</button>
	<button>
		<a href="/login">Logging out? Fineee</a>
	</button>
</div>

<style>
	.container {
		display: flex;
		gap: 1em;
	}
	.controls {
		position: absolute;
		padding: 1em;
		z-index: 10;
		left: 25%;
		top: 25%;
	}
	.background_when_creating_game {
		position: fixed;
		display: flex;
		justify-content: center;
		align-items: center;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		backdrop-filter: blur(20px);
		z-index: 5;
	}
</style>
