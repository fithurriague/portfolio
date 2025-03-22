all:run

run: 
	go run cmd/main.go

test:
	go test -v cmd/main.go

formatdoc:
	swag fmt 

builddoc:
	swag init -q -g  cmd/main.go --parseDependency
