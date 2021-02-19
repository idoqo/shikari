package tokens

import "gitlab.com/idoko/shikari/models"

type JWT struct {

}

func New() JWT {
	return JWT{}
}

func (j JWT) GenerateToken(user models.User) (string, error) {
	return "hello-world", nil
}