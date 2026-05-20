# 13 Structs, Methods, and Package Design

## Goal

Organize code into reusable types and small packages.

## What To Practice

- Design structs with clear responsibility.
- Attach behavior via methods.
- Keep package API small and intentional.
- Use interfaces at boundaries (consumer side).

## Mini Tasks

1. Build a `UserService` struct with constructor and methods.
2. Move storage logic to another package (for example `store`).
3. Inject dependency via interface, not concrete type.

## Done Criteria

- `main` only wires dependencies and starts flow.
- Business logic is testable without starting real I/O.
