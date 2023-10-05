# GREST

A Simple golang Rest API builder

## Quick Start

- install the library

```go
go get github.com/felixa1243/grest
```

- add starter code

```go
package main

import (
	"net/http"

	"github.com/felixa1243/grest/libs"
)

type Example struct {
	Message string
}

func main() {
	r := &libs.Router{}
	r.Route("GET", "/", r.Json(Example{Message: "Hello"}, 200))
	http.ListenAndServe(":5000", r)
}

```

- run the code

```
go run main.go
```
