# Zelda — Hyrule data & Go demo modules

This repository contains small Go modules and data files used to model Hyrule-like game data (players, enemies, bosses, spells, etc.) and a set of example/mod files in `mods/` that demonstrate usage and utilities.

**Quick links**
- Project README: [README.md](zelda/README.md)
- Go module: `go.mod`

**Contents at a glance**
- `base_game/` — core game types and helpers.
- `mods/` — example modules and experimental features (example programs and utilities).
- `json/` — data files used by the project (`bosses.json`, `classes.json`, `enemies.json`, `players.json`, `races.json`, `spells.json`, `traps.json`).

## Requirements

- Go 1.18+ (recommended latest stable Go). Ensure `GOPATH`/modules are set up.
- Git (for cloning and contributions).

## Quick start

1. Build everything (checks compilability):

```sh
go build ./...
```

2. Run an example program (many files in `mods/` are standalone `package main` examples):

```sh
go run ./mods/hyrule_castle.go
```

If a file in `mods/` is not `package main`, import packages from `base_game/` or other libraries and use `go run` on a `main` file that references them.

3. List available packages and modules:

```sh
go list ./...
```

## Project structure

- `base_game/` — core types and functions that model game concepts (characters, races, combat helpers).
- `mods/` — experimental gameplay code and example programs. Use these as starting points to extend or test features.
- `json/` — static data used by the demos and utilities. Each file contains a list of objects relevant to that category.
- `go.mod` — module definition and dependency list.

## JSON data

The `json/` folder contains sample data in JSON format used by the examples:

- `players.json` — sample player characters
- `enemies.json` — generic enemies
- `bosses.json` — boss definitions and stats
- `classes.json` — class or job templates
- `races.json` — race definitions and modifiers
- `spells.json` — spell definitions and effects
- `traps.json` — trap descriptors used in encounters

These files can be loaded with Go's `encoding/json` package. Example snippet:

```go
import (
    "encoding/json"
    "os"
)

f, _ := os.ReadFile("json/players.json")
var players []Player // Player defined in base_game
_ = json.Unmarshal(f, &players)
```

## Development

- Format code: `gofmt -w .` or `go fmt ./...`
- Run builds/tests: `go test ./...` (if any tests are present)
- Add data: update or add JSON files in `json/` and ensure example programs load them by relative path.

## Contributing

Contributions are welcome. Suggested workflow:

1. Fork the repo and create a feature branch.
2. Add tests or an example demonstrating the change.
3. Open a pull request with a brief description.

If you'd like me to add a specific example program, tests, or CI config, tell me which example to add and I'll scaffold it.

## Notes

- Some files in `mods/` are experimental and may be in-progress. Inspect them before running.
- Paths in examples assume you run commands from the repository root (`zelda/`).

## License

No license specified. If you want this project to be open source, add a `LICENSE` file (e.g., MIT, Apache-2.0).

---

If you'd like, I can also:

- add a small `examples/` runner that loads JSON and prints a summary,
- run `go build` to verify compilation in this environment,
- or add a `LICENSE` file.

Tell me which next step you prefer.
