package jwt

import (
	"fmt"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"jwt_server/config"
)

func init() {
	getDbProvider().init()
}

// Create creates JWT token based on user id and stores it in cache and database
func Create(userID string) (string, time.Time, error) {
	cfg := config.GetConfig()

	authID, _ := uuid.NewRandom()
	now := time.Now()
	expiry := now.Add(cfg.JwtExpiryDuration)

	// Add claims
	atClaims := jwtgo.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["auth_id"] = authID.String()
	atClaims["iat"] = now.Unix()
	atClaims["nbf"] = now.Unix()
	atClaims["exp"] = expiry.Unix()

	// Sign the token
	at := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(cfg.JwtSigningSecret))
	if err != nil {
		return "", expiry, err
	}

	// Save the token entry
	e := newEntry(userID, authID.String(), expiry)
	if err = e.save(getCacheProvider(), getDbProvider()); err != nil {
		return "", expiry, err
	}

	return token, expiry, err
}

func verify(tokenString string) (*jwtgo.Token, error) {
	cfg := config.GetConfig()

	// Parse internally verifies the token (checks iat, nbf, exp)
	return jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JwtSigningSecret), nil
	})
}

func getEntry(token *jwtgo.Token) *entry {
	claims := token.Claims.(jwtgo.MapClaims)
	userID := claims["user_id"].(string)
	authID := claims["auth_id"].(string)
	expiry := time.Unix(int64(claims["exp"].(float64)), 0)

	return newEntry(userID, authID, expiry)
}

// Verify checks validity of access token
func Verify(tokenString string) (bool, error) {
	token, err := verify(tokenString)

	if err != nil {
		return false, err
	}

	e := getEntry(token)
	return e.isValid(getCacheProvider(), getDbProvider())
}

// Delete removes JWT from database and cache
func Delete(tokenString string) error {
	token, err := verify(tokenString)

	if err != nil {
		return err
	}

	e := getEntry(token)
	return e.delete(getCacheProvider(), getDbProvider())
}

// Responsible for instantiating cache provider
func getCacheProvider() cacheProvider {
	return &redisCacheProvider{}
}

// Responsible for instantiating database provider
func getDbProvider() dbProvider {
	return &postgresDbProvider{}
}
