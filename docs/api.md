This document describes the REST API for the server.  It should be possible to
make most actions via REST, though there will also be a websocket for events
like move updates.

The objects are:

- [ ] players (aka users)
- [ ] sessions
- [ ] challenges
- [ ] games
- [ ] moves (this is really a property of the games)
- [ ] Possibly in the future: tournaments, puzzles, etc.

## Players

- [ ] GET /players/:id
- [ ] POST /players (registration)
- [ ] GET /players/me

## Sessions

- [ ] POST /sessions (login)
- [ ] DELETE /sessions (logout)

## Challenges

- [ ] POST /challenges
- [ ] GET /challenges
- [ ] GET /challenges/:id
- [ ] POST /challenges/:id/accept
- [ ] POST /challenges/:id/cancel

## Games

- [ ] GET /games/:id
- [ ] GET /games
- [ ] POST /games/:id/moves
