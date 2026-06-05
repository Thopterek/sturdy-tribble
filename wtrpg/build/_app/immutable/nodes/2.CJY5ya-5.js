import{H as e,I as t,L as n,M as r,N as i,O as a,P as o,V as s,a as c,b as l,d as u,f as d,k as f,l as p,q as m,v as h,w as g}from"../chunks/CwmggkRd.js";import"../chunks/xihTtKlq.js";import"../chunks/DRxzPdee.js";function _(e){return Object.prototype.toString.call(e)===`[object Date]`}function v(e){return e}function y(e,t){if(e===t||e!==e)return()=>e;let n=typeof e;if(n!==typeof t||Array.isArray(e)!==Array.isArray(t))throw Error(`Cannot interpolate values of different type`);if(Array.isArray(e)){let n=t.map((t,n)=>y(e[n],t));return e=>n.map(t=>t(e))}if(n===`object`){if(!e||!t)throw Error(`Object cannot be null`);if(_(e)&&_(t)){let n=e.getTime(),r=t.getTime()-n;return e=>new Date(n+e*r)}let n=Object.keys(t),r={};return n.forEach(n=>{r[n]=y(e[n],t[n])}),e=>{let t={};return n.forEach(n=>{t[n]=r[n](e)}),t}}if(n===`number`){let n=t-e;return t=>e+t*n}return()=>t}var b=class e{#e;#t;#n;#r=null;constructor(e,t={}){this.#e=n(e),this.#t=n(e),this.#n=t}static of(t,n){let r=new e(t(),n);return a(()=>{r.set(t())}),r}set(e,n){t(this.#t,e);let{delay:r=0,duration:i=400,easing:a=v,interpolate:o=y}={...this.#n,...n};if(i===0)return this.#r?.abort(),t(this.#e,e),Promise.resolve();let s=d.now()+r,c,l=!1,f=this.#r;return this.#r=u(n=>{if(n<s)return!0;if(!l){l=!0;let t=this.#e.v;c=o(t,e),typeof i==`function`&&(i=i(t,e)),f?.abort()}let r=n-s;return r>i?(t(this.#e,e),!1):(t(this.#e,c(a(r/i))),!0)}),this.#r.promise}get current(){return g(this.#e)}get target(){return g(this.#t)}set target(e){this.set(e)}},x=l(`<h1>Terms and Conditions</h1> <div><textarea>
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
	</textarea></div> <div><button><a href="/login">I assume you want to login inside anyway</a></button> <progress class="svelte-1uha8ag"></progress></div>`,1);function S(t,n){e(n,!1);let a=new b(0);c();var l=x(),u=o(i(l),4),d=o(r(u),2);m(u),f(()=>p(d,a.current)),h(t,l),s()}export{S as component};