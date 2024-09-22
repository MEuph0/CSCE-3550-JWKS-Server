# CSCE 3550 JWKS Server
Mason Willy

## Server
A server with 2 endpoints: a RESTful JWKS endpoint that serves public keys in JWKS format and a /auth endpoint that returns an unexpired, signed JWT on a POST request.

## Client
A client with 2 functionalities: makes http GET requests to recieve a public key from the server in JWKS format and makes http POST requests to recieve a signed JWT.

## Running the Porgrams
The server and client must be ran in separate software terminals.
The command is as follows, replacing "program" with the deisred program: go run program.go