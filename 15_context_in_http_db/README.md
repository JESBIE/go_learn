# 15 Context in HTTP and DB

## Goal

Use `context` to control timeout, cancellation, and request scope.

## What To Practice

- Pass `ctx` from HTTP handler to downstream calls.
- Use timeout with `context.WithTimeout`.
- Stop work early when client disconnects or timeout happens.
- Avoid storing context in struct fields.

## Mini Tasks

1. Add timeout to a simulated DB call.
2. In handler, cancel child context when request ends.
3. Return proper status for timeout vs internal failure.

## Done Criteria

- Long-running call exits quickly after cancel/timeout.
- Handler returns deterministic behavior under timeout.
