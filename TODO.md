# TODOs

Yeah yeah yeah. Markdown list for task tracking. W/e I'm the only one working on
this right now so this is sufficient for tracking purposes.

## Backend

[ ] Generally need to add pagination across the different "List" APIs
	[ ] Recipes
	[ ] Users
	[ ] Comments
[ ] Filtering
	[ ] Recipes
	[ ] Users (?)
[ ] Authentication
[ ] Authorization Middleware
	[ ] Protect Edits to owned...
		[ ] Recipes
		[ ] User settings
		[ ] Comments
	[ ] Rate Limiting
[ ] E2E API Tests
	[ ] Need to figure out a pattern that works for everything. 
		- Best plan might end up being having two executions that talk to each
		  other. The main binary, and a test executor. Better still would be if
		  I can get this spun up in docker containers so we're executing in a
		  prod-like environment.

## Frontend
