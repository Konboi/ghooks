# ghooks

ghooks is github hooks receiver. inspired by [GitHub::Hooks::Receiver](https://github.com/Songmu/Github-Hooks-Receiver), [octoks](https://github.com/hisaichi5518/octoks)


# Install

```
go get github.com/Konboi/ghooks
```

# Usage

```go
// sample.go
package main

import {
    "fmt"
    "github.com/Konboi/ghooks"
}


function main() {
    port := 8080
    hooks := ghooks.Server(port)

    hooks.on("push", pushHandler)
    hooks.on("pull_request", pullRequestHandler)
    hooks.Run()
}

function pushHandler(payload interface{}) {
    fmt.Printfln("puuuuush")
}

function pullRequestHandler(payload interface{}) {
    fmt.Println("pull_request")
}
```

```
go run sample.go
```

```
curl -H "X-GitHub-Event: push" -d '{"hoge":"fuga"}' http://localhost:8080
> puuuuush
```
