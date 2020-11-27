# golang

## golang runtime in docker
```bash
image used = golang:1.15-alpine

# installing dependencies manually
go get -u github.com/gorilla/mux

# installation path
/go/src/github.com/gorilla/mux
```

## Intellisense with docker and vscode
install <b>Remote-Container</b> extension in vscode
<br>
mount go-dump folder to GOPATH of docker-container for dependency persistence [Apparently this is not working due to permission issues]
<br>
install and configure vscode <b>go extension</b> with it's analytics tool
<br>
<a href="https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos">Configuring Local Environment</a>

## Starting vscode
click remote icon<br>
attatch to a running container<br>
navigate to source dir

## Setting the editor
```bash
create a file .vscode > setting.json

# for auto-import package
"go.autocompleteUnimportedPackages": true,

# for go docs
# guru is doc package
"go.docsTool": "guru"
```

## References
https://code.visualstudio.com/docs/languages/go