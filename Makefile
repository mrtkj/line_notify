# .PHONY: deps gen lint error build dev test

GOOS = linux
GOARCH = amd64
GOPATH = ${shell go env GOPATH}

dev:
	docker-compose build --no-cache
	docker-compose up -d

heroku_login:
	heroku login
	heroku container:login

release:
	heroku container:push web
	heroku container:release web
	osascript -e 'display notification "Deploy Done" with title "Heroku Deploy" subtitle "My Go Server" sound name "Submarine"'
