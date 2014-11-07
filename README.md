# ghooks

ghooks is github evnet hooks server

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

}

function pullRequestHandler(payload interface{}) {
  fmt.Println("pull_request")
}
```
