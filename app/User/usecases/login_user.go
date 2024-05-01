package usecases

import (
	"code-space-backend-api/app/User/entities"
	"code-space-backend-api/app/User/interfaces/dto/input"
	"code-space-backend-api/app/User/interfaces/repository"
	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/errors"
	"code-space-backend-api/common/token"
	"code-space-backend-api/common/uuid"
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const loginUserContext = "usecase/login-user"

var (
	invalidLoginInput = errors.InvalidField.
				WithContext(loginUserContext).
				WithName("invalid_field_send").
				WithTemplate("field send is either required or invalid: %v")

	couldNotFindUserWithGivenEmail = errors.NotFound.
					WithContext(loginUserContext).
					WithName("could_not_find_user").
					WithTemplate("could not find user with given email: %v")

	unauthorizedUser = errors.InvalidField.
				WithContext(loginUserContext).
				WithName("unauthorized_user").
				WithTemplate("send user credentials does not match")
)

type LoginUser interface {
	Execute(ctx context.Context, loginDTO input.LoginUserInputDTO) (token string, err error)
}

type loginUser struct {
	repository repository.UserRepository
}

func NewLoginUser(repository repository.UserRepository) LoginUser {
	return &loginUser{
		repository: repository,
	}
}

func (usecase *loginUser) Execute(ctx context.Context, loginDTO input.LoginUserInputDTO) (token string, err error) {
	if err = loginDTO.ValidateDTO(); err != nil {
		return str.EMPTY_STRING, invalidLoginInput.WithArgs(err.Error())
	}

	var user entities.UserDomain = loginDTO.ToDomain()

	storedUser, err := usecase.repository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return str.EMPTY_STRING, couldNotFindUserWithGivenEmail.WithArgs(err.Error())
	}

	if !user.IsUserAuthorized(storedUser.Password) {
		return str.EMPTY_STRING, unauthorizedUser
	}

	token, err = usecase.createUserToken(storedUser.Hash, storedUser.Email, storedUser.Name)
	if err != nil {
		return str.EMPTY_STRING, err
	}

	return
}

func (usecase *loginUser) createUserToken(hash, email, name string) (string, error) {
	claims := createTokenClaims(hash, email, name)
	return token.CreateToken(claims)
}

func createTokenClaims(hash, email, name string) token.CustomClaims {
	return token.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New(),
			Issuer:    "codespace",
			Subject:   "authentication",
			Audience:  "codespace_users",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		CustomKeys: map[string]interface{}{
			"user_hash": hash,
			"email":     email,
			"name":      name,
		},
	}
}
