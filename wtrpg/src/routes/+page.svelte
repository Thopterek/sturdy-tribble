<script lang="ts">
	import { Tween } from 'svelte/motion';
	let progress = new Tween(0);
</script>

<h1>Terms and Conditions</h1>
<div>
	<textarea>
		Here is some text that goes and goes. Idea is is to check if we can handle the change of value.
		Like in between this text that would be actual terms and conditions. But for now it's just
		couple of lines, that are being handwritten. Okay here is some term and conditions: Amazon
		Lumberyard’s Zombie Apocalypse Clause: The EULA explicitly states that restrictions on using the
		software for safety-critical systems do not apply in the event of a widespread viral infection
		causing human corpses to reanimate, allowing users to run medical equipment during a zombie
		outbreak without legal repercussion. GameStation’s Soul Clause: In an April Fool’s Day prank,
		the UK video game store included a clause granting them the non-transferable option to claim
		your immortal soul, which over 7,500 people accidentally agreed to before the company nullified
		it the next day. Apple’s Nuclear Weapon Ban: The iTunes EULA explicitly forbids using their
		products for the development, design, manufacture, or production of nuclear, missile, or
		chemical or biological weapons, settling the long-standing debate on whether iTunes can launch
		nukes. PC Pitstop’s $1,000 Prize: A hidden clause offered $1,000 to anyone who read a specific
		section of the license agreement and contacted the company; it took over 3,000 downloads before
		someone claimed the prize. Social Media Data Rights: Platforms like Facebook and Instagram grant
		companies a royalty-free, worldwide license to use, modify, and display user content (including
		photos) for any purpose, even after account deletion, while also forcing users into binding
		arbitration that waives the right to class-action lawsuits. Unexpected Physical Clauses: Other
		reports include daycare forms granting permission to take children to Canada, gym memberships
		requiring doctor’s notes for injury-based cancellation, and rock climbing waivers specifically
		excluding liability for stolen underwear.
	</textarea>
</div>

<div>
	<button>
		<a href="/login">I assume you want to login inside anyway</a>
	</button>

	<progress value={progress.current}></progress>
</div>

<style>
	progress {
		display: block;
		width: 100%;
	}
</style>
