package usecases

import (
	"code-space-backend-api/app/User/entities"
	"code-space-backend-api/app/User/interfaces/dto/input"
	"code-space-backend-api/app/User/interfaces/repository"
	"code-space-backend-api/common/errors"
	"context"

	"gorm.io/gorm"
)

const createUserContext = "usecase/create-user"

var (
	invalidUserInput = errors.InvalidField.
				WithContext(createUserContext).
				WithName("invalid_field_send").
				WithTemplate("field send is either required or invalid: %v")

	policyPrivacyNotAccepted = errors.BusinessRule.
					WithContext(createUserContext).
					WithName("policy_not_accepted").
					WithTemplate("policy_not_accepted")

	userAlreadyExists = errors.BusinessRule.
				WithContext(createUserContext).
				WithName("user_already_exist").
				WithTemplate("user is already created")

	failedToHashPassword = errors.BusinessRule.
				WithContext(createUserContext).
				WithName("failed_to_hash_password").
				WithTemplate("could not hash given password: %v")

	couldNotCreateUser = errors.Unknown.
				WithContext(createUserContext).
				WithName("could_not_create_user").
				WithTemplate("could not create user: %v")
)

type CreateUser interface {
	Execute(ctx context.Context, userDTO input.CreateUserInputDTO) error
}

type createUser struct {
	repository repository.UserRepository
}

func NewCreateUser(repository repository.UserRepository) CreateUser {
	return &createUser{
		repository: repository,
	}
}

func (usecase *createUser) Execute(ctx context.Context, userDTO input.CreateUserInputDTO) error {
	if err := userDTO.ValidateDTO(); err != nil {
		return invalidUserInput.WithArgs(err.Error())
	}

	var user entities.UserDomain = userDTO.ToDomain()

	if !user.IsPrivacyPolicyAccepted() {
		return policyPrivacyNotAccepted
	}

	if _, err := usecase.repository.GetUserByEmail(ctx, user.Email); err == nil || err != gorm.ErrRecordNotFound {
		return userAlreadyExists
	}

	if err := user.HashPassword(); err != nil {
		return failedToHashPassword.WithArgs(err.Error())
	}

	user.CreateHashIfEmpty()
	if err := usecase.repository.Create(ctx, user); err != nil {
		return couldNotCreateUser.WithArgs(err.Error())
	}

	return nil
}
