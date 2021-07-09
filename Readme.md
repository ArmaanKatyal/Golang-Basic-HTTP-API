
# Golang-Basic-HTTP-API

It a rest-api made in golang which stores posts in a slice and perform all the CRUD operations on the Data stored.

## Installation

Use the package manager [go get](https://golang.org/cmd/go/) to install the libraries needed for this http-api to work.

```bash
go get -u github.com/gorilla/mux
```

## Essential Libraries

```go
// These are all the imported libraries
import (
	"encoding/json"
	"net/http"
	"strconv"   // This converts strings to integers

	"github.com/gorilla/mux"
)
```

## Execution
This is used to execute the program on a linux platform, It is different for windows and MacOs system.
```bash
go build && ./http-api
``` 
other way to execute the program is 
```bash
go run http-api.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
