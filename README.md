## Forum written in Go and Vue
Allows users to create account, read and create posts.

### ENDPOINTS:
- AUTH:
	- /register
	- /login
	- /token/refresh
-  USER POST:
	- /post POST
	- /post/{id} GET
    - /post/{id} DELETE
    - /post/{id} PUT/PATCH
- USER ACCOUNT:
	- /account/{id} GET
    - /account/update PUT/PATCH
    - /account/delete DELETE
    - /account/password/reset PUT/PATCH
- FORUM FEED
    - /posts?page={i} GET

### MODELS:
- User model:
{id: uuid/ulid, first_name: string, last_name: string, email: string, password: string}
- Post model:
{user: uuid/ulid, title: string, body: string, timestamp: Timestamp}