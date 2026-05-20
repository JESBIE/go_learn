# 12 Error Handling Patterns

## Goal

Learn practical error handling in Go services.

## What To Practice

- Wrap errors with context using `fmt.Errorf("...: %w", err)`.
- Check root causes with `errors.Is` and `errors.As`.
- Define sentinel errors for expected business failures.
- Avoid string-compare on errors.

## Mini Tasks

1. Create `readConfig(path string) error` and wrap file errors.
2. Return custom validation errors and classify with `errors.As`.
3. Print user-friendly message by checking `errors.Is`.

## Done Criteria

- You can explain when to use wrapping vs direct return.
- You can classify errors without relying on error text.
