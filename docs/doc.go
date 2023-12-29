// Package docs Email Verifier API.
//
// The purpose of this application is to provide an application that allows to:
// <ul>
// <li>validate email and its domain</li>
// <li>create clean mailing lists</li>
// <li>blacklist domains and invalid emails</li>
// </ul>
//
// Terms Of Service:
//
// There are no TOS at this moment, use at your own risk we take no responsibility
//
//	Schemes: http, https
//	Host: localhost:1323
//	BasePath: /v2
//	Version: 0.0.1
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: Wojciech Bernatek<wojciech.bernatek@gmail.com>
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- api_key:
//
//	SecurityDefinitions:
//	api_key:
//	     type: apiKey
//	     name: KEY
//	     in: header
//	oauth2:
//	    type: oauth2
//	    authorizationUrl: /oauth2/auth
//	    tokenUrl: /oauth2/token
//	    in: header
//	    scopes:
//	      bar: foo
//	    flow: accessCode
//
//	Extensions:
//	x-meta-value: value
//	x-meta-array:
//	  - value1
//	  - value2
//	x-meta-array-obj:
//	  - name: obj
//	    value: field
//
// swagger:meta
package docs
