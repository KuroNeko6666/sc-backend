package helper

import (
	"errors"

	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/golang-jwt/jwt/v4"
)

func GetTokenFromUser(user model.User, key string, token *string) error {
	keyToken := []byte(key)
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
	})
	tkn, err := rawToken.SignedString(keyToken)

	if err != nil {
		return err
	}

	*token = tkn

	return nil
}

func GetTokenFromAdmin(admin model.Admin, key string, token *string) error {
	keyToken := []byte(key)
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       admin.ID,
		"name":     admin.Name,
		"username": admin.Username,
		"email":    admin.Email,
	})
	tkn, err := rawToken.SignedString(keyToken)

	if err != nil {
		return err
	}

	*token = tkn

	return nil
}

func GetUserFromToken(tkn string, key string, user *model.User) error {
	var response model.User
	var claims jwt.MapClaims

	if err := getClaims(tkn, key, &claims); err != nil {
		return err
	}

	response.ID = claims["id"].(string)
	response.Username = claims["username"].(string)
	response.Email = claims["email"].(string)

	*user = response

	return nil
}

func GetAdminFromToken(tkn string, key string, admin *model.Admin) error {
	var response model.Admin
	var claims jwt.MapClaims

	if err := getClaims(tkn, key, &claims); err != nil {
		return err
	}

	response.ID = claims["id"].(string)
	response.Username = claims["username"].(string)
	response.Email = claims["email"].(string)

	*admin = response

	return nil
}

func getClaims(tokenSTR string, key string, claims *jwt.MapClaims) error {
	secret := []byte(key)

	token, err := jwt.Parse(tokenSTR, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return err
	}

	if clm, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		*claims = clm
		return nil
	} else {
		return errors.New("error extract token")
	}
}
