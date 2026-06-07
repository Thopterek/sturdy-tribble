<script>
	let chat_selection = $state('');
	const onselectionchange = () => {
		chat_selection = document.getSelection()?.toString() ?? '';
	};
</script>

<svelte:document {onselectionchange} />

<h1>Game Lobby</h1>

<button class="button">
	<a href="/dashboard">Check that profile</a>
</button>

<p>
	Selected: {chat_selection};
</p>

<style>
	.button {
		box-shadow: 4px 4px 0px black; /* hard pixel shadow */
		border: 3px solid black;
		border-radius: 0; /* pixel art has no rounded corners */
	}
</style>
