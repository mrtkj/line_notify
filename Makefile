# .PHONY: deps gen lint error build dev test

GOOS = linux
GOARCH = amd64
GOPATH = ${shell go env GOPATH}

dev:
	docker build -t line_notify .
	docker run --env-file ./.env -p 3000:3000 -it --rm line_notify

heroku_login:
	heroku login
	heroku container:login

release:
	heroku container:push web
	heroku container:release web
	osascript -e 'display notification "Deploy Done" with title "Heroku Deploy" subtitle "My Go Server" sound name "Submarine"'
