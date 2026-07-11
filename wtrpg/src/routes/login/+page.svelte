<script lang="ts">
	import { user } from '$lib/auth.svelte';
	const email = 'Email';
	const username = 'Username';
	const pass = 'Password';
	// to be deleted for later
	const api = 'localhost:8080';
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
	<form>
		Registration
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
	</form>
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
</style>
