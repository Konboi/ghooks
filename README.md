# ghooks

ghooks is github evnet hooks server

# Useage

```
go get <url>
```

```golang

import "path/ghooks"


function main() {
  hooks := Ghooks.NewHooks()

  hooks.on("push", pushHandler)
  hooks.on("pull_request", pullRequestHandler)
  hooks.Run()
}

function pushHandler(event ghooks.Evnet) {
  if event.Message == nil {
     fmt.Printf("Write commit message")
  }
}

function pullRequestHandler(event ghooks.Evnet) {
  if event.Message == nil {
     fmt.Printf("Write commit message")
  }
}

```

# Option


* `--port` Set port number
