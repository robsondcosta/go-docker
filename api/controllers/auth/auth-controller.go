package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"docker/api/models"
	"docker/api/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

type JwtToken struct {
	AcessoToken string `json:"acesso-token"`
}

var jwt_secret = os.Getenv("jwt_secret")

func Login(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	Administrador := &models.Administrador{}
	utils.ParseBody(r, Administrador)

	login := Administrador.Login
	senha := Administrador.Senha

	result := models.SelectAdministradorByLogin(login)
	if result == nil {
		utils.SetJsonReturn(w, false, 400, "Error", "Suas credenciais não correspondem aos nossos registros!", nil)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Senha), []byte(senha))
	if err != nil {
		fmt.Println(err)
		utils.SetJsonReturn(w, false, 400, "Error", "Suas credenciais não correspondem aos nossos registros!", nil)
		return
	}

	ttl := 2 * time.Minute
	accessTokenExpire := os.Getenv("access_token_expire")
	min, err := strconv.Atoi(accessTokenExpire)

	if err != nil {
		log.Println(err)
	}

	if accessTokenExpire != "" {
		ttl = time.Duration(min) * time.Minute
	}

	CreateToken(w, login, senha, ttl)
}

func CreateToken(w http.ResponseWriter, login string, senha string, ttl time.Duration) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   ttl,
	})

	tokenString, error := token.SignedString([]byte(jwt_secret))

	if error != nil {
		fmt.Println(error)
	}

	utils.SetJsonReturn(w, true, 200, "Success", "Login efetuado com sucesso!", JwtToken{AcessoToken: tokenString})
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Ocorreu um erro durante a requisição")
					}
					return []byte(jwt_secret), nil
				})

				if error != nil {
					utils.SetJsonReturn(w, false, 400, "System Error", error.Error(), nil)
					return
				}

				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					utils.SetJsonReturn(w, false, 400, "Error", "O token informado é inválido", nil)
					return
				}
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "É necessário enviar o token de autorização", nil)
			return
		}
	})
}
