<script lang="ts">
	let username = $state('Email please');
	let pass = $state('Password also');
</script>

<h1>Login and Registration</h1>

<div class="container">
	<div>
		<div bind:innerHTML={username} contenteditable></div>
		<div>
			<input type="text" placeholder={pass} />
		</div>

		<div>
			<button>
				<a href="/dashboard">Yaaaay login</a>
			</button>
		</div>
	</div>
	<div>
		<div bind:innerHTML={username} contenteditable></div>
		<div>
			<input type="text" placeholder="Put your password" />
		</div>
		<div>
			<input type="text" placeholder="Now please repeat it" />
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
