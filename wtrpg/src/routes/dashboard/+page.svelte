<script>
	let show_game_creation = $state(false);
	const todos = $state([
		{ id: 1, done: false, description: 'Name of the game' },
		{ id: 2, done: false, description: 'RPG system played' },
		{ id: 4, done: false, description: 'Description' }
	]);
	let { empty, done } = $props();
</script>

<h1>Dashboard</h1>

<div>
	{#if show_game_creation}
		<div role="presentation" class="background_when_creating_game"></div>
		<ul>
			{#each empty as tbd (empty.id)}
				<li></li>
			{/each}
		</ul>
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

<style>
	.controls {
		position: absolute;
		left: 0;
		top: 0;
		padding: 1em;
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
	}
</style>
