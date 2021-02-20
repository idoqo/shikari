package guard

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/idoko/shikari/models"
	"time"
)


type Guard struct {
	algo jwt.SigningMethod
	signingKey []byte
	ttl time.Duration
}

func Configure(signingKey string, ttl time.Duration) (Guard, error) {
	minSecretLength := 64
	if len(signingKey)< minSecretLength {
		return Guard{}, fmt.Errorf("signing key is too short, minimum length is %d", minSecretLength)
	}
	return Guard{
		algo: jwt.SigningMethodHS256,
		signingKey: []byte(signingKey),
		ttl: ttl,
	}, nil
}

func (g Guard) GenerateToken(user models.User) (string, error) {
	return jwt.NewWithClaims(g.algo, jwt.MapClaims{
		"email": user.Email,
		"user_id": user.ID,
		"exp": time.Now().Add(g.ttl).Unix(),

	}).SignedString(g.signingKey)
}