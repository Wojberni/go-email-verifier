## Useful commands
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

## Email verification schema

1. Email address validation
   * syntax check 
   * special characters check
   * whitespaces check
2. Domain Existence check
   1. DNS records lookup
   2. Ping domain
   3. MX records validation
3. Disposable Email address detection
   * Using opensource blacklists from GitHub with disposable domain names
   * Blacklists managed as plain text files, every malicious domain is seperated with newline character
4. Email Server interactions
   * SMTP Connection: Establish a connection to the mail server specified in the MX records using a SMTP client.
   * Recipient Validation: Attempt to send a test email to the provided address using the SMTP connection. If the server responds with a successful delivery confirmation, the address is valid.

## TODOS

* Send test email with various providers - DONE WP and Gmail (for now)
* Custom parser for whois
* Change context for different emails