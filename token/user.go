package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"visitor-management-system/config"
	"visitor-management-system/types"
)

func GenerateUserTokens(email string, id int, usertype string, company_id int) (signed_token string, signed_refreshtoken string, err error) {

	claims := &types.SignedUserDetails{
		Email:     email,
		Id:        id,
		UserType:  usertype,
		CompanyId: company_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &types.SignedUserDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetConfig().SecretKey))
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(config.GetConfig().SecretKey))

	return token, refresh_token, nil
}
