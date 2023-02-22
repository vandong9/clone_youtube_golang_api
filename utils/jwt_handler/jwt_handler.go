package jwt_handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide") // Should read from evn, secret

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

// var sampleSecretKey = Message{Status: "true", Info: "value"}

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

// comment these
func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte("You're Unauthorized"))
					if err != nil {
						return nil, err
					}
				}
				return "", nil

			})
			// parsing errors result
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err2 := writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
				if err2 != nil {
					return
				}

			}
			// if there's a token
			if token.Valid {
				endpointHandler(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}
		// response for if there's no token header
	})
}

func ExtractClaims(_ http.ResponseWriter, request *http.Request) (string, error) {
	if request.Header["Token"] != nil {
		tokenString := request.Header["Token"][0]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return sampleSecretKey, nil
		})
		if err != nil {
			return "Error Parsing Token: ", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username := claims["username"].(string)
			return username, nil
		}
	}

	return "", nil // unable to extract claims
}
