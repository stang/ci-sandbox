# CI sandbox

## Usage

```
usage:
    hello <name>
    hello --version
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
