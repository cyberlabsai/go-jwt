# go-jwt
GoJwt is a Cyberlabs's simple implementation of JWT.
![](https://emojis.slackmojis.com/emojis/images/1507931630/3036/gopher_dance.gif?1507931630)
## Example
```go
package main

import (
	"fmt"
	"log"

	jwt "github.com/cyberlabsai/jwt-go"
)

func main() {
	signingKey := []byte("very-very-very-secret")
	claims := map[string]interface{}{
		"name": "atila",
	}
	// Generate token
	token, err := jwt.Generate(signingKey, claims)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Tk: ", *token)

	// Validate token
	payload, err := jwt.Validate(signingKey, *token)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Payload: ", *payload)
}
```
