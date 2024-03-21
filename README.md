It works.

# Purpose

POC for using webpack and goja to import javascript methods with 3rd party dependencies in Go.

Main idea is:

- make js bundle using webpack which includes some 3rd party dependency.
- load the bundle in Go using goja and call the js fn.

# Workflow

in `./ts-module`

    pnpm install & pnpm webpack

in `./go-module`

    go run main.go

Webpack bundles the code including dependencies and then goja reads the bundle and calls the js function.

# thoughts

- cgo/v8 wrapper mentioned in goja might fit my usecase better. Sounds like the js crypto stuff will be more performant and im not going to call it a lot https://github.com/dop251/goja
- looks like there is a seperate node compatibility repo. guessing the crypto library from node may not be available in goja... but maybe here? https://github.com/dop251/goja_nodejs
