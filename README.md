# ghooks

ghooks is github hooks receiver. inspired by [GitHub::Hooks::Receiver](https://github.com/Songmu/Github-Hooks-Receiver), [octoks](https://github.com/hisaichi5518/octoks)

# Useage

```
go get <url>
```

```golang

import "path/ghooks"


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
