install:
	which godep || go get github.com/tools/godep
	godep restore

db/webvuln.db:
	go run cmd/initdb.go db/sql/setup.sql db/webvuln.db

server: db/webvuln.db
	go run webvuln.go

static/vue.js:
	curl https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js > static/vue.js
