test: lint

	go test -v ./...

lint:

	go tool vet -all -printfuncs=Criticalf,Infof,Warningf,Debugf,Tracef ./
