package jwt

import (
	"log"
	commonError "server-veggie/common/error"
	"server-veggie/config"
	"server-veggie/middleware/auth/business"
	"server-veggie/middleware/auth/model"
	"server-veggie/middleware/auth/strorage"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func EnCodeToken(authHeader string, db *gorm.DB) error {
	config, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	//check-bear-token
	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		return commonError.ErrInvalidTokenFormat(model.EntityName, model.ErrInvalidTokenFormat)
	}

	tokenString := authToken[1]
	//encode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, commonError.ErrAuthorizationMissing(model.EntityName, model.ErrAuthorizationMissing)
		}
		return []byte(config.Secret), nil
	})

	if err != nil || !token.Valid {
		return commonError.ErrInvalidToken(model.EntityName, model.ErrInvalidToken)
	}

	//claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return commonError.ErrInvalidClaims(model.EntityName, model.ErrInvalidClaims)
	}
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return commonError.ErrInvalidExpToken(model.EntityName, model.ErrInvalidExpToken)
	}
	//create store
	store := strorage.NewSQLStore(db)
	//calculate logic
	business := business.NewTokenBiz(store)
	if err := business.TokenUser(claims["name_account"].(string)); err != nil {
		return err
	}
	return nil
}
