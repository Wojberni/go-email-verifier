```bash
# GENERATE SPEC -> swagger generate spec --help
# Safe method for merging with old file
go-swagger generate spec -i swagger.yml -m -o swagger.yml
# Overwrite old file
go-swagger generate spec -m -o swagger.yml
# SERVE DOCS -> swagger serve --help
# Port forwading container port 8080 added, two flavours: redoc/swagger awailable
go-swagger serve ./swagger.yml --no-open --port 8080 -F swagger
go-swagger serve ./swagger.yml --no-open --port 8080
```