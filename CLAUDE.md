# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is an Exercism solutions repository containing Go programming exercises. The repository structure follows Exercism's format with solutions organized by language and exercise name.

## Architecture

- `solutions/go/` - Contains all Go exercise solutions
- Each exercise has its own directory (e.g., `hello-world/`, `hamming/`, `leap/`)
- Each exercise may have multiple iterations numbered as `1/`, `2/`, `3/`, etc.
- Solutions are single Go files named after the exercise (e.g., `hello_world.go`, `hamming.go`)

## Working with Go Solutions

### Running Individual Solutions
Since this is an Exercism repository, each solution is a standalone Go package. To work with a specific exercise:

```bash
cd solutions/go/[exercise-name]/[iteration]
go run .
```

### Testing Solutions
Exercism exercises typically come with test files. If working on solutions, you would run:

```bash
cd solutions/go/[exercise-name]/[iteration]
go test
```

### Package Structure
Each solution file:
- Belongs to a package named after the exercise (e.g., `package greeting` for hello-world)
- Implements the required functions as specified by the exercise
- Uses standard Go conventions and error handling patterns

## Development Notes

- Solutions follow Go naming conventions (exported functions start with capital letters)
- Error handling uses Go's idiomatic `error` return pattern
- Some exercises have multiple iterations showing evolution of solutions
- No build system or dependency management is present - this is a simple collection of exercise solutions