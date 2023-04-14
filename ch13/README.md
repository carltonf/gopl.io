As this chapter involves `cgo`, it requires some extra step in building and
linting.

# How to build

## Deps

```sh
sudo apt install build-essential libbz2-dev
```

## Enable `CGO`

Check `go env CGO_ENABLED`, if not yet 1, set it to 1.

# Please `gopls` in VSCode

Click the button `regenerate cgo definitions` above the line
> import "C"

[Ref](https://github.com/golang/go/issues/35721)

## Re-linting

Somehow the above regeneration wouldn't trigger re-linting of the package.

To mitigate the issue, two options exist:

1. Execute `Go: Lint current package`
2. Change some files.