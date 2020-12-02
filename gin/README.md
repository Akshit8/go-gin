# Rest Api with Gin

## Inititializing the project
go mod init github.com/Akshit8/gin

## Installing gin-gonic
go get github.com/gin-gonic/gin

## Building application locally before pushing to heroku
go build -o bin/gin-app -v .

## heroku cli deploy
```bash
# test application locally
heroku local

heroku create gin-go-8

# adding heroku remote
heroku git:remote -a gin-go-8

# deploying from local to remote
git push heroku master:master
```