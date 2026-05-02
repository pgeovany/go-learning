# currency_conversor

Convert an amount in BRL to a target currency using static rates from `rates.json`.

## Run

```sh
go run . <amount> <currency>
```

Example:

```sh
go run . 100 USD
```

Currency is case-insensitive. Amount uses `.` as the decimal separator.

## What I learned

- `os.ReadFile` for whole-file reads (handles file closing), `os.Args` for CLI input
- `encoding/json` with struct tags (`` `json:"base"` ``) — the value side **must** be a quoted string; without quotes `reflect` reports `bad syntax`
- Pointer receivers on methods (`func (e *ExchangeRates) Convert(...)`) so the struct isn't copied per call
- The **comma-ok idiom** for map lookups: `v, ok := m[k]` lets you distinguish "missing" from "zero value"
- `fmt.Errorf` _returns_ an error, it does not print. Calling it and ignoring the return value is a silent bug
- Format verbs: `%s` for strings, `%d` for ints. `%d` on a string compiles but prints garbage
- `string(float64)` is a **compile error** — Go won't let you convert a float to string that way. To format a float as text use `fmt.Sprintf("%f", x)` or `strconv.FormatFloat`. (Related trap: `string(int)` _does_ compile, but interprets the int as a Unicode code point — `string(65)` is `"A"`, not `"65"`. Use `strconv.Itoa` for ints too.)
- Error string convention in Go: lowercase first letter, no trailing punctuation, so they compose with `fmt.Errorf("context: %w", err)`
- stdout vs stderr — `>` only redirects stdout, so diagnostics belong on stderr (`fmt.Fprintln(os.Stderr, ...)`)
- Exit codes matter — a bare `return` from `main` exits with `0` (success), so failure paths must call `os.Exit(1)` (or another non-zero) for shell idioms like `if ./prog; then` to work correctly (this is specially useful for CLI tools)

## Things I left for later

- Brazilian decimal formatting via `golang.org/x/text/message` (did the manual `.` → `,` swap instead)
- The `run() error` pattern in `main` for cleaner exit handling and testability
- `//go:embed` for the rates file so the program isn't tied to its working directory
- Tests
