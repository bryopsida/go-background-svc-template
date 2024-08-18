# GO Background SVC Starter Template

## What is this?

This is a starter template intended for kick starting a simple background service.

It provides

1) A top level main that listens for SIGINT and SIGTERM and gracefully closes its database connection and signals to sub routines to exit
2) Has 2 go routines [increment](./incrementor/increment.go) and [print](./incrementor/print.go), that accept cancellation contexts
3) Has a simple embedded data store [badger](https://dgraph.io/docs/badger)
4) Has a starting point for configuration with a [interface wrapper](./config//config.go) around [viper](https://github.com/spf13/viper)
5) Has a [github workflow](https://github.com/curium-rocks/flows/blob/main/.github/workflows/golang.yml)
6) Has a [dev container](./.devcontainer/devcontainer.json)
7) Has a [Dockerfile](./Dockerfile)
8) Has a [Makefile](./Makefile) for common tasks such as building, testing, linting
9) Is automatically updated with renovate [renovate.json](./renovate.json)
