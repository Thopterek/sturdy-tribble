<script lang="ts">
	import { user } from '$lib/auth.svelte';
	const email = 'Email';
	const username = 'Username';
	const pass = 'Password';
	// to be deleted for later
	const api: string = $state('localhost:8080');
</script>

<h1>Login and Registration</h1>

<div class="container">
	<form>
		Login
		<div>
			<input type="email" placeholder={email} bind:value={user.email} />
		</div>
		<div>
			<input type="password" placeholder={pass} bind:value={user.password} />
		</div>
		<div>
			<button>
				<a href="/dashboard">Yaaaay login</a>
			</button>
		</div>
	</form>
	<div>
		<div>
			<input type="email" placeholder={email} bind:value={user.email} />
		</div>
		<div>
			<input type="text" placeholder={username} bind:value={user.name} />
		</div>
		<div>
			<input type="password" placeholder={pass} bind:value={user.password} />
		</div>
		<div>
			<button>
				<a href="/dashboard">Here is Registration</a>
			</button>
		</div>
	</div>
</div>
<div>
	<button>
		<a href="/">Are you sure you don't want to read terms and conditions first?</a>
	</button>
</div>

<style>
	.container {
		display: flex;
		gap: 1em;
	}
	input[type='text'] {
		flex: 1;
		padding: 0.5em;
		margin: -0.2em 0;
		border: none;
	}
</style>
