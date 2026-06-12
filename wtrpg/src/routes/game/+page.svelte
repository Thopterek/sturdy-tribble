<script lang="ts">
	import Dsix from '$lib/assets/dsix.svelte';
	import Dtwenty from '$lib/assets/dtwenty.svelte';
	let chat_selection = $state('');
	const onselectionchange = () => {
		chat_selection = document.getSelection()?.toString() ?? '';
	};
</script>

<svelte:document {onselectionchange} />

<h1>Game Lobby</h1>

<button>
	<a href="/dashboard">Check that profile</a>
</button>

<p>
	Selected: {chat_selection};
</p>
<Dsix />
<Dtwenty />
