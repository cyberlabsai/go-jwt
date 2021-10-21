# go-jwt
GoJwt is a Cyberlabs's simple implementation of JWT.
![](https://www.collinsdictionary.com/images/full/gopher_438742864.jpg)
## Generate
```go
package main
import (
	"fmt"
	jwt "github.com/cyberlabsai/jwt-go"
	)

func main() {
	signingKey := []byte("very-very-very-secret")
	payload := map[string]interface{
	}{
		"name": "atila",
	}

	token, _ := jwt.Generate(signingKey, payload)

	fmt.Println(*token)
}
```
## Validate
```go
```
