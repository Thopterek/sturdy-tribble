import{$ as e,A as t,B as n,F as r,I as i,J as a,P as o,S as s,U as c,V as l,W as u,Y as d,a as f,d as p,et as m,g as h,h as g,w as _,z as v}from"../chunks/BzihiBaM.js";import"../chunks/xihTtKlq.js";function y(e){return Object.prototype.toString.call(e)===`[object Date]`}function b(e){return e}function x(e,t){if(e===t||e!==e)return()=>e;let n=typeof e;if(n!==typeof t||Array.isArray(e)!==Array.isArray(t))throw Error(`Cannot interpolate values of different type`);if(Array.isArray(e)){let n=t.map((t,n)=>x(e[n],t));return e=>n.map(t=>t(e))}if(n===`object`){if(!e||!t)throw Error(`Object cannot be null`);if(y(e)&&y(t)){let n=e.getTime(),r=t.getTime()-n;return e=>new Date(n+e*r)}let n=Object.keys(t),r={};return n.forEach(n=>{r[n]=x(e[n],t[n])}),e=>{let t={};return n.forEach(n=>{t[n]=r[n](e)}),t}}if(n===`number`){let n=t-e;return t=>e+t*n}return()=>t}var S=class e{#e;#t;#n;#r=null;constructor(e,t={}){this.#e=u(e),this.#t=u(e),this.#n=t}static of(t,n){let r=new e(t(),n);return o(()=>{r.set(t())}),r}set(e,t){c(this.#t,e);let{delay:n=0,duration:r=400,easing:i=b,interpolate:a=x}={...this.#n,...t};if(r===0)return this.#r?.abort(),c(this.#e,e),Promise.resolve();let o=h.now()+n,s,l=!1,u=this.#r;return this.#r=g(t=>{if(t<o)return!0;if(!l){l=!0;let t=this.#e.v;s=a(t,e),typeof r==`function`&&(r=r(t,e)),u?.abort()}let n=t-o;return n>r?(c(this.#e,e),!1):(c(this.#e,s(i(n/r))),!0)}),this.#r.promise}get current(){return t(this.#e)}get target(){return t(this.#t)}set target(e){this.set(e)}},C=_(`<h1>Terms and Conditions</h1> <div><p class="check_term svelte-1uha8ag">Here is some text that goes and goes. Idea is is to check if we can handle the change of value.
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
		section of the license agreement and contacted the company; it took over 3,000 downloads before</p></div> <div class="topbar svelte-1uha8ag"><progress class="svelte-1uha8ag"></progress> <button><a href="/login">I assume you want to login inside anyway</a></button></div>`,1);function w(o,h){d(h,!0);let g=new S(0),_=u(0);i(()=>{let e=document.body.scrollHeight-window.innerHeight;g.set(t(_)/e)});var y=C(),b=l(n(y),4),x=v(b);e(2),m(b),r(()=>p(x,g.current)),f(`y`,()=>t(_),e=>c(_,e,!0)),s(o,y),a()}export{w as component};