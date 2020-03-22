# CI sandbox

## Usage

```
A simple tool to give a sign of welcome when meeting someone

Usage:
  greet [name] [flags]
  greet [command]

Available Commands:
  help        Help about any command
  version     Print the version number

Flags:
  -h, --help          help for greet
  -l, --lang string   Language to use for the greeting (default "en")

Use "greet [command] --help" for more information about a command.
```

## Developer's HOW TO

```
# run tests
make test

# prepare release artifacts
make release

# github release
RELEASE=v0.2.0
git tag "${RELEASE}"
git push origin "${RELEASE}"
```
