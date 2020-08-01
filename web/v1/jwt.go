package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/MahdiRazaqi/university-system/config"
	"github.com/MahdiRazaqi/university-system/student"
	"github.com/MahdiRazaqi/university-system/teacher"
	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	UID  int    `json:"uid"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var signature = config.Config.SecretKey

// CreateToken generate new token
func CreateToken(UID int, role string) (string, error) {
	claims := new(customClaims)
	claims.UID = UID
	claims.Role = role
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signature))
}

// GetToken get token from header
func GetToken(req *http.Request) string {
	token := req.Header.Get("Authorization")
	return strings.Replace(token, "Bearer ", "", -1)
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signature), nil
	})
}

// LoadByToken load user from token
func LoadByToken(token string) (interface{}, error) {
	t, _ := parseToken(token)

	claims, ok := t.Claims.(*customClaims)
	if !ok {
		return nil, errors.New("converting token failed")
	}

	if claims.Role == "teacher" {
		return teacher.LoadByTeacherID(claims.UID)
	}

	return student.LoadByStudentID(claims.UID)
}

func userBinding(u interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	var mapData map[string]interface{}
	if err := json.Unmarshal(jsonData, &mapData); err != nil {
		return nil, err
	}
	return mapData, nil
}
