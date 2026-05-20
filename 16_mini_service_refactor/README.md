# 16 Mini Service Refactor

## Goal

Combine previous topics in one small project-style exercise.

## Scope

- Use package split (`cmd`, `internal` or simple equivalent).
- Add error classification.
- Add context-aware workflow.
- Add unit tests for service layer.

## Suggested Steps

1. Pick one old exercise and refactor folder layout.
2. Add interface-based dependency injection.
3. Add tests for success/failure/timeout.
4. Keep `main` as wiring only.

## Done Criteria

- You can explain each package responsibility in one sentence.
- You have at least one tested business flow end-to-end (without real network).
