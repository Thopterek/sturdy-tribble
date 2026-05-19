# Sturdy-Tribble

ft_transcendence, README for notes

## Roles

* Developers: Everyone, actually Stelio (building the code)

* Product Owner: Nate (bigger picture), actually Stelio

* Product Manager: Gianni, actually Stelio (distributed tasks)

* Technical Lead: Dimi, actually Stelio (overseeing technical decision)

The roles were distributed through the knowledge of individual skills.

### Rules

* Version Control: Branching and forcing to always use pull requests

* Documentation: Left out out for questions how to do it

### Tree of features

TBD: showcase who did which ones and on which worked / implemented / helped

Basic:

Make it into tasks, so Gianni can decide what the hell, presenting the IDEA

* Map making: usage of reusable components (walls/door/gate/bush/tree/)
* Chat: people can write on it
* Dice rolling: When someone rolles the dice can choose which dice to use
* Moving around: On the map with other players
* People can be invited to games/kicked out and so on (roles who sees what)

### Modules

* TRPG, better Roll20
  * Major modules: Real-Time/WebSockets, Chat, Public API, User Managment
  * Minor modules: Frontend_framework, Notifications, Filtering, Stats, OAuth2.0

Maybe:

* Major:
* Minor: ORM, SSR, PWA, File Upload


### Infrastructure

* AWS Services
  * S3 Bucket: Cloud storage for the frontend

    SPA does not run on the server. The S3 bucket is not executing any code; it is completely dumb. It just hands out files.

    The actual "launch" happens entirely inside the player's CPU and RAM. Here is the exact millisecond-by-millisecond breakdown:

    The Request: The player types www.my-dnd-game.com and hits Enter. Their browser follows the DNS signpost to your S3 bucket.

    The Handshake: The browser knocks on the S3 bucket and asks, "Do you have a default file?" Because you configured Static Website Hosting, S3 says, "Yes, here is index.html."

    The Skeleton: The index.html file in an SPA is practically empty. It usually just contains a single, empty <div> tag (e.g., <div id="root"></div>) and a link to your compiled JavaScript file.

    The Download: The browser reads the HTML, sees the link to the JavaScript file, and asks S3 to download it, along with your CSS and pixel-art assets.

    The Execution (The "Launch"): Once the JavaScript is downloaded, the browser's JavaScript Engine (like V8 in Chrome) executes it. Your React/Vue code wakes up, looks for that empty <div id="root">, and injects the entire user interface—the login screen, the buttons, the maps—directly into the browser window.

    From that moment on, the app is fully launched. When the player clicks between the Login Page and the Map Builder, they are not downloading new HTML pages from S3. The JavaScript code already living in their browser is simply erasing the login UI and drawing the map UI instantly.
  
  * Elastic Beanstalk: Backend server

    * Maintains dictionary mapping (Room IDs <-> Connected Users):
      * When player attempts to establish a WebSocket conn, provide JWT (handshake)
      * When player attempts to join a Room, server checks DB for invitation (validation)
      * When player attempts to play, server checks move (Game engine - authority)

    * API Controller
        TODO: Establish Endpoints @Thopterek @itsiros

    * Websocket Gateway

  * RDS (PostgreSQL): Central DB
    
    * Contains cold assets: TODO configure tables' format ('...' = unfinished)
      * Users: Basic User profile info
        [Username, Password, PlayerName, UUID, ...]
      * Player Stats: In-game AND/OR general profile statistics 
        [UUID, SessionsPlayed, ObjectsInteracted, ...]
      * Maps: Same format as cache DB
      * Characters:  
  
  * ElastiCache (Redis): In-game cache DB

    * Contains game session state. Periodically copies state back to the central DB. In case of a server drop, central DB requests, game state from cache. In game the session gets interrupted, a state is saved on the central DB.
    
    * TODO Configure format 
  
  * 