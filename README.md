# echo-template
echo-template to develop Control Panel in Go using Echo framework.

Latest release: template-0.1.0

### feature
- Create new project template.
- Sample demo CRUD
- Connect to postgresSQL
- Demo using go-dal as helper to CRUD with database
- Support Logger with elastic search
- Support authenticate with JWT
- Support i18n
- Support throw performance log to statsd
- Support docker build image
- Support kubernetes stack

### Getting Started
#### Start demo
1. Checkout https://github.com/LTNB/go-echo-template.git
2. Run `data/data.sql` in postgresSQL to create database
2. Install bower to get JS library
3. Run ```brower install```
4. Run ```go build```
5. Run ```go run server.go```
6. http://localhost:9000
#### Docker
