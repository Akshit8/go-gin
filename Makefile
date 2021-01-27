git:
	git add .
	git commit -m "$(msg)"
	git push origin master

start:
	go run main.go