<script lang="ts">
	import Dsix from '$lib/assets/dsix.svelte';
	import { user } from '$lib/auth.svelte';
	const email = 'Email';
	const username = 'Username';
	const pass = 'Password';
	// to be deleted for later
	const api = 'localhost:8080';
	// not binding in to the same things
	let c_email = $state('');
	let c_pass = $state('');
	let c_name = $state('');
	// adding the update versions, changing user name is not handled
	let u_email = $state('');
	let u_name = $state('');
	let u_pass = $state('');
</script>

<h1><Dsix width={24} height={24} /> Login and Registration</h1>
<div class="main">
	<form>
		<legend>Login</legend>
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
			<input type="email" placeholder={email} bind:value={c_email} />
		</div>
		<div>
			<input type="text" placeholder={username} bind:value={c_name} />
		</div>
		<div>
			<input type="password" placeholder={pass} bind:value={c_pass} />
		</div>
		<div>
			<button>
				<a href="/dashboard">Here is Registration</a>
			</button>
		</div>
	</form>
	<form>
		Update User
		<div>
			<input type="email" placeholder="new email" bind:value={u_email} />
		</div>
		<div>
			<input type="test" placeholder="new user" bind:value={u_name} />
		</div>
		<div>
			<input type="password" placeholder="new password" bind:value={u_pass} />
		</div>
		<div><button>Just click it</button></div>
	</form>
</div>
<div class="testing">
	<form>
		JWT related things
		<div><button>Show the JWT Access Token</button></div>
		<div><button>DELETE the JWT Access Token</button></div>
	</form>
	<form>
		Admin without admin
		<div><button>Get Metrics</button></div>
		<div><button>RESET DB</button></div>
	</form>
</div>
Terms & Conditions
<div>
	<button>
		<a href="/">Are you sure you don't want to read terms and conditions first?</a>
	</button>
</div>

<style>
	.main {
		display: flex;
		gap: 1em;
	}
	.testing {
		display: flex;
		gap: 1em;
	}
</style>
