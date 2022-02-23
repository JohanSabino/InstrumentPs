.PHONY: build production staging squad

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/v1 v1/*.go

production: build
	../../node_modules/.bin/serverless --stage production create_domain
	../../node_modules/.bin/serverless --stage production deploy

squad: build
	../../node_modules/.bin/serverless --stage squad create_domain
	../../node_modules/.bin/serverless --stage squad deploy
