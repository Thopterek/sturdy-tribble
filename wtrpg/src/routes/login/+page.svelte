<script lang="ts">
	let username = $state('Username please');
	let pass = $state('Password also');
</script>

<h1>Login</h1>

<div bind:innerHTML={username} contenteditable></div>
<div>
	<input type="text" placeholder={pass} />
</div>

<div>
	<button>
		<a href="/dashboard">My login is my registration, let's goo</a>
	</button>
</div>

<div>
	<button>
		<a href="/">Are you sure you don't want to read terms and conditions first?</a>
	</button>
</div>

<style>
	[contenteditable] {
		padding: 0.5em;
		border: 1px solid #eee;
		border-radius: 4px;
	}
	input[type='text'] {
		flex: 1;
		padding: 0.5em;
		margin: -0.2em 0;
		border: none;
	}
</style>
