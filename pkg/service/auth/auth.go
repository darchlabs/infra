package authservice

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/darchlabs/infra/pkg/auth"
	authstorage "github.com/darchlabs/infra/pkg/storage/auth"
	"github.com/darchlabs/infra/pkg/users"
	"github.com/dgrijalva/jwt-go"
)

// AuthService ...
type AuthService struct {
	AuthStore   *authstorage.AuthStore
	UserService users.Service
}

// NewService ...
func NewService(store *authstorage.AuthStore, us users.Service) *AuthService {
	return &AuthService{
		AuthStore:   store,
		UserService: us,
	}
}

// GetByToken ...
func (as *AuthService) GetByToken(token string) (*auth.Auth, error) {
	// validate token param
	if token == "" {
		return nil, errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.AuthStore.Get(&auth.Query{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Login ...
func (as *AuthService) Login(email, password string) (*auth.Response, error) {
	// validate email param
	if email == "" {
		return nil, errors.New("invalid email")
	}

	// validate password param
	if password == "" {
		return nil, errors.New("invalid password")
	}

	// verify if password is valid
	err := as.UserService.VerifyPassword(email, password)
	if err != nil {
		return nil, err
	}

	// get user by email
	user, err := as.UserService.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// create token difinition
	t := auth.Token{
		UserID: user.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(8760 * time.Hour).Unix(), // one year of expiration
		},
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("env variable JWT_SECRET must be defined")
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	// create auth definition
	a := &auth.Auth{
		Token:     tokenStr,
		Kind:      auth.KindUser,
		Blacklist: false,
		UserID:    user.ID,
	}

	// save auth store
	err = as.AuthStore.Create(a)
	if err != nil {
		return nil, err
	}

	// prepare metatoken
	mt := &auth.MetaToken{
		Token: tokenStr,
	}

	// prepare response
	res := &auth.Response{
		Data: user,
		Meta: mt,
	}

	return res, nil
}

// Signup ...
func (as *AuthService) Signup(u *users.User) (*auth.Response, error) {
	// validate user existence
	if u == nil {
		return nil, errors.New("invalid user")
	}

	// validate user params
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return nil, errors.New("invalid user params")
	}

	// create new user
	err := as.UserService.Create(u)
	if err != nil {
		return nil, err
	}

	// create token definition
	t := auth.Token{
		UserID: u.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(4320 * time.Hour).Unix(), // six months of expiration
		},
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("env variable JWT_SECRET must be defined")
	}

	// generate login jwt token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	// create signup auth definition
	a := &auth.Auth{
		Token:     tokenStr,
		Kind:      auth.KindUser,
		Blacklist: false,
		UserID:    u.ID,
	}

	// save auth store
	err = as.AuthStore.Create(a)
	if err != nil {
		return nil, err
	}

	// create verify token definition
	vt := auth.Token{
		UserID: u.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(2160 * time.Hour).Unix(), // 3 months of expiration
		},
	}

	// generate verify jwt token
	verificationToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), vt)
	verificationTokenStr, err := verificationToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	// create verification auth definition
	va := &auth.Auth{
		Token:     verificationTokenStr,
		Kind:      auth.KindVerifyPassword,
		Blacklist: false,
		UserID:    u.ID,
	}

	// save auth store
	err = as.AuthStore.Create(va)
	if err != nil {
		return nil, err
	}

	// prepare metatoken
	mt := &auth.MetaToken{
		Token:             tokenStr,
		VerificationToken: verificationTokenStr,
	}

	// prepare response
	res := &auth.Response{
		Data: u,
		Meta: mt,
	}

	return res, nil
}

// VerifyToken ...
func (as *AuthService) VerifyToken(token string, kind string) (*auth.Token, error) {
	// validate token param
	if token == "" {
		return nil, errors.New("invalid token")
	}

	// validate kind param
	if kind == "" {
		return nil, errors.New("invalid kind")
	}

	// get Auth by token from store
	a, err := as.AuthStore.Get(&auth.Query{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	// validate auth kind
	if a.Kind != kind {
		return nil, errors.New("invalid kind")
	}

	// check if token is blacklisted
	if a.Blacklist {
		return nil, errors.New("token is blacklisted")
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("env variable JWT_SECRET must be defined")
	}

	// decode token
	// validate token is valid with JWT_SECRET
	// validate token is not expired
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// parser map to struct
	data, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}
	decode := new(auth.Token)
	err = json.Unmarshal(data, decode)
	if err != nil {
		return nil, err
	}

	return decode, nil
}

// VerifyEmail ...
func (as *AuthService) VerifyEmail(token string) error {
	// validate token param
	if token == "" {
		return errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.AuthStore.Get(&auth.Query{
		Token: token,
	})
	if err != nil {
		return err
	}

	// validate auth kind
	if a.Kind != auth.KindVerifyPassword {
		return errors.New("invalid kind")
	}

	// check if token is blacklisted
	if a.Blacklist {
		return errors.New("token is blacklisted")
	}

	// update verified to true
	err = as.UserService.Update(&users.User{
		ID:       a.UserID,
		Verified: true,
	})
	if err != nil {
		return err
	}

	// update blacklist to true
	a.Blacklist = true
	err = as.AuthStore.Update(a)
	if err != nil {
		return err
	}

	return nil
}

// Logout ...
func (as *AuthService) Logout(token string) error {
	// validate token param
	if token == "" {
		return errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.AuthStore.Get(&auth.Query{
		Token: token,
	})
	if err != nil {
		return err
	}

	// check if token is blacklisted
	if a.Blacklist {
		return errors.New("token is blacklisted")
	}

	// validate auth kind is user
	if a.Kind != auth.KindUser {
		return errors.New("invalid auth kind")
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Env variable JWT_SECRET must be defined")
	}

	// decode token
	// validate token is valid with JWT_SECRET
	// validate token is not expired
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return err
	}

	// parser map to struct
	data, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	var decode auth.Token
	err = json.Unmarshal(data, &decode)
	if err != nil {
		return err
	}

	// update blacklist to true
	a.Blacklist = true
	err = as.AuthStore.Update(a)
	if err != nil {
		return err
	}

	return nil
}

// ForgotPassword ...
func (as *AuthService) ForgotPassword(e string) (string, error) {
	// validate email param
	if e == "" {
		return "", errors.New("invalid email")
	}

	// check if email exist on users service
	user, err := as.UserService.GetByEmail(e)
	if err != nil {
		return "", err
	}

	// create temporal token difinition
	t := auth.Token{
		UserID: user.ID,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), // 2 hours of expiration
		},
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Env variable JWT_SECRET must be defined")
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	// create auth definition
	a := &auth.Auth{
		Token:     tokenStr,
		Kind:      auth.KindForgotPassword,
		Blacklist: false,
		UserID:    user.ID,
	}

	// save auth store
	err = as.AuthStore.Create(a)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// RecoverPassword ...
func (as *AuthService) RecoverPassword(newPassword, token string) error {
	// validate newPassword param
	if newPassword == "" {
		return errors.New("invalid newPassword")
	}

	// validate token param
	if token == "" {
		return errors.New("invalid token")
	}

	// get Auth by token from store
	a, err := as.AuthStore.Get(&auth.Query{
		Token: token,
	})
	if err != nil {
		return err
	}

	// check if token is blacklisted
	if a.Blacklist {
		return errors.New("token is blacklisted")
	}

	// validate auth kind is forgot-password
	if a.Kind != auth.KindForgotPassword {
		return errors.New("invalid auth kind")
	}

	// get jwt secret env value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Env variable JWT_SECRET must be defined")
	}

	// decode token
	// validate token is valid with JWT_SECRET
	// validate token is not expired
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return err
	}

	// parser map to struct
	data, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	var decode auth.Token
	err = json.Unmarshal(data, &decode)
	if err != nil {
		return err
	}

	user := &users.User{
		ID:       decode.UserID,
		Password: newPassword,
	}

	// update user password
	err = as.UserService.Update(user)
	if err != nil {
		return err
	}

	// update blacklist to true
	a.Blacklist = true
	err = as.AuthStore.Update(a)
	if err != nil {
		return err
	}

	return nil
}
