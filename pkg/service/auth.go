package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/shooosty/rd-app/models"
	"github.com/shooosty/rd-app/pkg/repository"
	"math/rand"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) CreateEmployee(user models.User) (string, error) {
	password := generatePassword()
	SendPasswordToEmployee(password, user.Name, user.Email)
	user.Password = generatePasswordHash(password)

	return s.repo.CreateEmployee(user)
}

func (s *AuthService) GetCurrentUser(username, password string) (models.User, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	signingKey, _ := os.LookupEnv("SIGN_KEY")

	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	signingKey, _ := os.LookupEnv("SIGN_KEY")

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	salt, _ := os.LookupEnv("SALT")

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generatePassword() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := 10
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)

	return str
}

func SendPasswordToEmployee(password string, name string, email string) {
	subject := "Регистрация в личном кабинете Rhinodesign"
	text := "Ваш пароль для входа в кабинет: " + password
	html := "<b>" + name + "</b>" + "<p>" + "Ваш пароль для входа в кабинет: " + password + "<p>" + "</br>" +
		"<p> Рекомендуем сменить пароль при первом входе в кабинет! </p>"
	_ = SendMail(subject, text, html, name, email)
}
