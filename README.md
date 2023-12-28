```bash
go-swagger generate spec -i swagger.yml -m -o swagger.yml
go-swagger serve ./swagger.yml --no-open --port 8080 -F swagger
go-swagger serve ./swagger.yml --no-open --port 8080
```