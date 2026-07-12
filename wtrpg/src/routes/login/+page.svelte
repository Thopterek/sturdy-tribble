<script lang="ts">
	import Dsix from '$lib/assets/dsix.svelte';
	import { user, create_user, login } from '$lib/auth.svelte';
	const email = 'Email';
	const username = 'Username';
	const pass = 'Password';
	// to be deleted for later
	// not binding in to the same things
	let c_email = $state('');
	let c_pass = $state('');
	let c_name = $state('');
	// adding the update versions, changing user name is not handled
	let u_email = $state('');
	let u_name = $state('');
	let u_pass = $state('');
	// some logic to be changed surely but for testing? HMMM?!
	// like using the Cookie Token and so on
</script>

<h1><Dsix width={24} height={24} /> Login and Registration</h1>
<div><a href="/dashboard">Just a link to a generic dashboard</a></div>
<div class="main">
	<form
		onsubmit={async (e) => {
			e.preventDefault();
			login(user.email, user.password);
		}}
	>
		<legend>Login</legend>
		<div><input type="email" placeholder={email} bind:value={user.email} required /></div>
		<div><input type="password" placeholder={pass} bind:value={user.password} required /></div>
		<div><button type="submit">Actually sending a login</button></div>
	</form>
	<form
		onsubmit={async (e) => {
			e.preventDefault();
			user.access_token = await create_user(c_email, c_name, c_pass);
		}}
	>
		<legend>Registration</legend>
		<div><input type="email" placeholder={email} bind:value={c_email} required /></div>
		<div><input type="text" placeholder={username} bind:value={c_name} required /></div>
		<div><input type="password" placeholder={pass} bind:value={c_pass} required /></div>
		<div><button type="submit">Register sending a post</button></div>
	</form>
	<form>
		<legend>Update User</legend>
		<div><input type="email" placeholder="new email" bind:value={u_email} /></div>
		<div><input type="test" placeholder="new user" bind:value={u_name} /></div>
		<div><input type="password" placeholder="new password" bind:value={u_pass} /></div>
		<div><button>Just click it</button></div>
	</form>
</div>
<div class="testing">
	<form>
		<legend>JWT related things</legend>
		<div><button>Show the JWT Access Token</button></div>
		<div><button>DELETE the JWT Access Token</button></div>
	</form>
	<form>
		<legend>Admin without admin</legend>
		<div><button>Get Metrics</button></div>
		<div><button>RESET DB</button></div>
	</form>
</div>
<legend>Terms & Conditions</legend>
<div><button><a href="/">Would you like to read T&C first?</a></button></div>

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
