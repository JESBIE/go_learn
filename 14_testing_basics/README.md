# 14 Testing Basics

## Goal

Write reliable tests for pure logic and service logic.

## What To Practice

- Table-driven tests.
- Subtests with `t.Run`.
- Compare expected/actual output clearly.
- Test error paths, not only happy path.

## Mini Tasks

1. Write tests for one existing pure function.
2. Add table cases: normal, edge, invalid input.
3. Test one service function with mocked dependency.

## Commands

- `go test ./...`
- `go test -run TestName ./...`

## Done Criteria

- You can add new case in table quickly.
- Tests cover both success and failure branches.
