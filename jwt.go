package jwt

import (
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

func Generate(signingKey []byte, claims map[string]interface{}) (*string, error) {
	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.HS256, Key: signingKey},
		(&jose.SignerOptions{}).WithType("JWT").WithBase64(true),
	)
	if err != nil {
		return nil, err
	}

	builder := jwt.Signed(sig)

	claims["issued_at"] = jwt.NewNumericDate(time.Now())

	builder = builder.Claims(claims)
	JWT, err := builder.CompactSerialize()
	if err != nil {
		return nil, err
	}

	return &JWT, nil
}
