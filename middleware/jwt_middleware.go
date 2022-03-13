package middleware

import (
	"fmt"
	"net/http"
	"project/constants"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id int, email string, name string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["email"] = email
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user_role string
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		claims, _ := extractClaims(splitToken[1])
		str := fmt.Sprintf("%v", claims["role"])
		user_role = string(str)
		if IsLoggedAdmin(user_role) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(403), 403)
		}
	})
}

func IsLoggedAdmin(role string) bool {
	if strings.Compare(role, "admin") == 0 {
		return true
	} else {
		fmt.Println("Not Admin")
		return false
	}
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := constants.SECRET_JWT
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
func RegisterPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user_id string
		Id := r.URL.Query().Get("id")
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		claims, _ := extractClaims(splitToken[1])
		str := fmt.Sprintf("%v", claims["id"])
		user_id = string(str)
		if user_id == Id {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(403), 403)
		}
	})
}
