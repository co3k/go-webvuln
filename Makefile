install:
	which godep || go get github.com/tools/godep
	godep restore

db/webvuln.db:
	go run cmd/initdb.go db/sql/setup.sql db/webvuln.db

server: db/webvuln.db
	go run webvuln.go

npm-install:
	npm install

static/app.js: npm-install
	npx webpack-cli --mode development --module-bind 'jsx=babel-loader' client-src/app.jsx -o static/app.js
