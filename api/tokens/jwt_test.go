package tokens

import (
	"gitlab.com/idoko/shikari/models"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	jwt := JWT{}
	user := models.User{
		ID: 1,
		Email: "test@test.com",
		Password: "test",
	}
	token, err := jwt.GenerateToken(user)
	if err != nil {
		t.Errorf("failed to generate token: %v", err)
	}
	if token == "" {
		t.Error("token generator returned an empty token")
	}
	//todo: decode and check token fields for email - or do that in separate func
}