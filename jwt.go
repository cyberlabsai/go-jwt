package jwt

import (
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)
type Payload struct {
	jwt.Claims
	Data map[string]interface{}
}

func Generate(signingKey interface{}, alg jose.SignatureAlgorithm, claims Payload) (*string, error) {
	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: alg, Key: signingKey},
		(&jose.SignerOptions{}).WithType("JWT").WithBase64(true),
	)
	if err != nil {
		return nil, err
	}

	builder := jwt.Signed(sig)

	builder = builder.Claims(claims)
	JWT, err := builder.CompactSerialize()
	if err != nil {
		return nil, err
	}

	return &JWT, nil
}

func Validate(signingKey interface{}, token string) (*Payload, error) {
	parsedJWT, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, err
	}

	var resultClaims Payload
	err = parsedJWT.Claims(signingKey, &resultClaims)
	if err != nil {
		return nil, err
	}

	return &resultClaims, nil
}
