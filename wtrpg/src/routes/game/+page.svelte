<script>
	let chat_selection = $state('');
	const onselectionchange = () => {
		chat_selection = document.getSelection()?.toString() ?? '';
	};
</script>

<svelte:document {onselectionchange} />

<h1>Game Lobby</h1>
<svg width="100" height="100" viewBox="0 0 100 100">
	<rect x="10" y="10" width="50" height="30" fill="red" />
	<circle cx="50" cy="50" r="20" fill="blue" />
	<line x1="0" y1="0" x2="100" y2="100" stroke="black" stroke-width="2" />
</svg>
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
