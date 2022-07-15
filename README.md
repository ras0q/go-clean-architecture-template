# go-clean-architecture-template

A template of clean architecture written by Go (with generics)

This template uses following Go frameworks/tools:

- [labstack/echo](https://github.com/labstack/echo) as a web framework
- [ent/ent](https://github.com/ent/ent) as an entity framework (ORM)
- [golang/mock](https://github.com/golang/mock) as a mocking framework
- [google/go-cmp](https://github.com/google/go-cmp) as an assertion tool
- [golangci/golangci-lint](https://github.com/golangci/golangci-lint) as a linters runner
- [cosmtrek/air](https://github.com/cosmtrek/air) as a live reload tool

## Quick Start

NOTE: This project uses Makefile as a task runner.

```bash
# docker is required to set up database.
make db-up
make dev
```

### Test

```bash
make test-all
```

for checking only unit tests, then

```bash
make test-unit
```

for checking only integration tests, then

```bash
make test-integration
```

### Run linters

```bash
make lint
```

## Structure

:construction::construction_worker: Under construction...
