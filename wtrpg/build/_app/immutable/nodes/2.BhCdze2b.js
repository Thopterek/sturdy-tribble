import{B as e,F as t,H as n,J as r,N as i,P as a,Q as o,R as s,S as c,U as l,Z as u,a as d,d as f,g as p,h as m,k as h,q as g,w as _,z as v}from"../chunks/CFHTQXBM.js";import"../chunks/xihTtKlq.js";function y(e){return Object.prototype.toString.call(e)===`[object Date]`}function b(e){return e}function x(e,t){if(e===t||e!==e)return()=>e;let n=typeof e;if(n!==typeof t||Array.isArray(e)!==Array.isArray(t))throw Error(`Cannot interpolate values of different type`);if(Array.isArray(e)){let n=t.map((t,n)=>x(e[n],t));return e=>n.map(t=>t(e))}if(n===`object`){if(!e||!t)throw Error(`Object cannot be null`);if(y(e)&&y(t)){let n=e.getTime(),r=t.getTime()-n;return e=>new Date(n+e*r)}let n=Object.keys(t),r={};return n.forEach(n=>{r[n]=x(e[n],t[n])}),e=>{let t={};return n.forEach(n=>{t[n]=r[n](e)}),t}}if(n===`number`){let n=t-e;return t=>e+t*n}return()=>t}var S=class e{#e;#t;#n;#r=null;constructor(e,t={}){this.#e=l(e),this.#t=l(e),this.#n=t}static of(t,n){let r=new e(t(),n);return i(()=>{r.set(t())}),r}set(e,t){n(this.#t,e);let{delay:r=0,duration:i=400,easing:a=b,interpolate:o=x}={...this.#n,...t};if(i===0)return this.#r?.abort(),n(this.#e,e),Promise.resolve();let s=p.now()+r,c,l=!1,u=this.#r;return this.#r=m(t=>{if(t<s)return!0;if(!l){l=!0;let t=this.#e.v;c=o(t,e),typeof i==`function`&&(i=i(t,e)),u?.abort()}let r=t-s;return r>i?(n(this.#e,e),!1):(n(this.#e,c(a(r/i))),!0)}),this.#r.promise}get current(){return h(this.#e)}get target(){return h(this.#t)}set target(e){this.set(e)}},C=_(`<h1>Terms and Conditions</h1> <div><p class="check_term svelte-1uha8ag">Here is some text that goes and goes. Idea is is to check if we can handle the change of value.
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
		excluding liability for stolen underwear. Here is some text that goes and goes. Idea is is to
		check if we can handle the change of value. Like in between this text that would be actual terms
		and conditions. But for now it's just couple of lines, that are being handwritten. Okay here is
		some term and conditions: Amazon Lumberyard’s Zombie Apocalypse Clause: The EULA explicitly
		states that restrictions on using the software for safety-critical systems do not apply in the
		event of a widespread viral infection causing human corpses to reanimate, allowing users to run
		medical equipment during a zombie outbreak without legal repercussion. GameStation’s Soul
		Clause: In an April Fool’s Day prank, the UK video game store included a clause granting them
		the non-transferable option to claim your immortal soul, which over 7,500 people accidentally
		agreed to before the company nullified it the next day. Apple’s Nuclear Weapon Ban: The iTunes
		EULA explicitly forbids using their products for the development, design, manufacture, or
		production of nuclear, missile, or chemical or biological weapons, settling the long-standing
		debate on whether iTunes can launch nukes. PC Pitstop’s $1,000 Prize: A hidden clause offered
		$1,000 to anyone who read a specific section of the license agreement and contacted the company;
		it took over 3,000 downloads before someone claimed the prize. Social Media Data Rights:
		Platforms like Facebook and Instagram grant companies a royalty-free, worldwide license to use,
		modify, and display user content (including photos) for any purpose, even after account
		deletion, while also forcing users into binding arbitration that waives the right to
		class-action lawsuits. Unexpected Physical Clauses: Other reports include daycare forms granting
		permission to take children to Canada, gym memberships requiring doctor’s notes for injury-based
		cancellation, and rock climbing waivers specifically. Here is some text that goes and goes. Idea
		is is to check if we can handle the change of value. Like in between this text that would be
		actual terms and conditions. But for now it's just couple of lines, that are being handwritten.
		Okay here is some term and conditions: Amazon Lumberyard’s Zombie Apocalypse Clause: The EULA
		explicitly states that restrictions on using the software for safety-critical systems do not
		apply in the event of a widespread viral infection causing human corpses to reanimate, allowing
		users to run medical equipment during a zombie outbreak without legal repercussion.
		GameStation’s Soul Clause: In an April Fool’s Day prank, the UK video game store included a
		clause granting them the non-transferable option to claim your immortal soul, which over 7,500
		people accidentally agreed to before the company nullified it the next day. Apple’s Nuclear
		Weapon Ban: The iTunes EULA explicitly forbids using their products for the development, design,
		manufacture, or production of nuclear, missile, or chemical or biological weapons, settling the
		long-standing debate on whether iTunes can launch nukes. PC Pitstop’s $1,000 Prize: A hidden
		clause offered $1,000 to anyone who read a specific section of the license agreement and
		contacted the company; it took over 3,000 downloads before someone claimed the prize. Social
		Media Data Rights: Platforms like Facebook and Instagram grant companies a royalty-free,
		worldwide license to use, modify, and display user content (including photos) for any purpose,
		even after account deletion, while also forcing users into binding arbitration that waives the
		right to class-action lawsuits. Unexpected Physical Clauses: Other reports include daycare forms
		granting permission to take children to Canada, gym memberships requiring doctor’s notes for
		injury-based cancellation, and rock climbing waivers specifically Here is some text that goes
		and goes. Idea is is to check if we can handle the change of value. Like in between this text
		that would be actual terms and conditions. But for now it's just couple of lines, that are being
		handwritten. Okay here is some term and conditions: Amazon Lumberyard’s Zombie Apocalypse
		Clause: The EULA explicitly states that restrictions on using the software for safety-critical
		systems do not apply in the event of a widespread viral infection causing human corpses to
		reanimate, allowing users to run medical equipment during a zombie outbreak without legal
		repercussion. GameStation’s Soul Clause: In an April Fool’s Day prank, the UK video game store
		included a clause granting them the non-transferable option to claim your immortal soul, which
		over 7,500 people accidentally agreed to before the company nullified it the next day. Apple’s
		Nuclear Weapon Ban: The iTunes EULA explicitly forbids using their products for the development,
		design, manufacture, or production of nuclear, missile, or chemical or biological weapons,
		settling the long-standing debate on whether iTunes can launch nukes. PC Pitstop’s $1,000 Prize:
		A hidden clause offered $1,000 to anyone who read a specific section of the license agreement
		and contacted the company; it took over 3,000 downloads before someone claimed the prize. Social
		Media Data Rights: Platforms like Facebook and Instagram grant companies a royalty-free,
		worldwide license to use, modify, and display user content (including photos) for any purpose,
		even after account deletion, while also forcing users into binding arbitration that waives the
		right to class-action lawsuits. Unexpected Physical Clauses: Other reports include daycare forms
		granting permission to take children to Canada, gym memberships requiring doctor’s notes for
		injury-based cancellation, and rock climbing waivers specifically excluding liability for stolen
		underwear. Here is some text that goes and goes. Idea is is to check if we can handle the change
		of value. Like in between this text that would be actual terms and conditions. But for now it's
		just couple of lines, that are being handwritten. Okay here is some term and conditions: Amazon
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
		requiring doctor’s notes for injury-based cancellation, and rock climbing waivers specifically.
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
		section of the license agreement and contacted the company; it took over 3,000 downloads before</p></div> <div class="topbar svelte-1uha8ag"><progress class="svelte-1uha8ag"></progress> <button><a href="/login">I assume you want to login inside anyway</a></button></div>`,1);function w(i,p){r(p,!0);let m=new S(0),_=l(0);t(()=>{let e=document.body.scrollHeight-window.innerHeight;m.set(h(_)/e)});var y=C(),b=e(v(y),4),x=s(b);u(2),o(b),a(()=>f(x,m.current)),d(`y`,()=>h(_),e=>n(_,e,!0)),c(i,y),g()}export{w as component};