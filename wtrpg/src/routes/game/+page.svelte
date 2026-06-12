<script lang="ts">
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

<div>
	<svg width="64" height="64" viewBox="0 0 16 16">
		<rect class="outline" x="2" y="0" width="12" height="1" />
		<rect class="outline" x="0" y="2" width="1" height="12" />
		<rect class="outline" x="15" y="2" width="1" height="12" />
		<rect class="outline" x="2" y="15" width="12" height="1" />
		<rect class="outline" x="1" y="1" width="1" height="1" />
		<rect class="outline" x="14" y="14" width="1" height="1" />
		<rect class="outline" x="14" y="1" width="1" height="1" />
		<rect class="outline" x="1" y="14" width="1" height="1" />
		<rect class="inner_white" x="1" y="2" width="1" height="12" />
		<rect class="inner_white" x="2" y="1" width="12" height="1" />
		<rect class="grey_base" x="2" y="2" width="12" height="12" />
	</svg>
</div>

<style>
	.inner_white {
		fill: rgb(255, 255, 255);
	}
	.outline {
		fill: rgb(0, 0, 0);
	}
	.grey_base {
		fill: rgb(200, 200, 200);
	}
</style>
