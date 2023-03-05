package jwt_handler

import (
	"fmt"
	"net/http"
	"time"

	commonModels "com.vandong9.clone_youtube_golang_api/common/models"
	"github.com/golang-jwt/jwt/v5"
)

var sampleSecretKey = []byte("SecretYouShouldHide") // Should read from evn, secret

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

type MyCustomClaims struct {
	Token commonModels.LoginToken `json:"token"`
	jwt.RegisteredClaims
}

type ValidClaimError struct {
}

func (valid ValidClaimError) Error() string {
	return "ValidClaimError"
}

func (my MyCustomClaims) Valid() error {
	return ValidClaimError{}
}

// var sampleSecretKey = Message{Status: "true", Info: "value"}

func GenerateJWT(loginInfo commonModels.LoginToken) (string, error) {
	// mapToken, _ := json.Marshal(&token)
	// var m jwt.MapClaims
	// _ = json.Unmarshal(mapToken, &m)

	// value := jwt.NewWithClaims(jwt.SigningMethodHS256, m)

	// // token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// // 	"foo": "bar",
	// // 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// // })

	// // Sign and get the complete encoded token as a string using the secret
	// tokenString, err := value.SignedString(sampleSecretKey)
	// if err != nil {
	// 	return "Signing Error", err
	// }

	// return tokenString, nil

	claims := MyCustomClaims{
		loginInfo,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "Signing Error", err
	}
	return ss, nil

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

func ExtractClaims(_ http.ResponseWriter, request *http.Request) (*commonModels.LoginToken, error) {
	if request.Header["Token"] != nil {
		tokenString := request.Header["Token"][0]
		return ParseToken(tokenString)
	}

	fmt.Println("There is no token in header")
	return nil, nil // unable to extract claims
}

func ParseToken(tokenString string) (*commonModels.LoginToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})

	if token == nil || err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Token, claims.RegisteredClaims.Issuer)
		return &claims.Token, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}
