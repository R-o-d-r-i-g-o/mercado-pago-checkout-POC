package entities

import (
	"code-space-backend-api/app/User/interfaces/dto/output"
	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/crypt"
	"code-space-backend-api/common/uuid"
	"code-space-backend-api/infra/database/models"
)

type UserDomain struct {
	Id                    uint
	Hash                  string
	Name                  string
	Email                 string
	Phone                 string
	Password              string
	PolicyPrivacyAccepted bool
}

func (u *UserDomain) IsUserAuthorized(hashedPassword string) bool {
	return crypt.CheckPasswordHash(u.Password, hashedPassword)
}

func (u *UserDomain) CreateHashIfEmpty() {
	if u.Hash == str.EMPTY_STRING {
		u.Hash = uuid.New()
	}
}

func (u *UserDomain) HashPassword() error {
	hashedPassword, err := crypt.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}

func (u *UserDomain) AcceptPrivacyPolicy() {
	u.PolicyPrivacyAccepted = true
}

func (u *UserDomain) IsPrivacyPolicyAccepted() bool {
	return u.PolicyPrivacyAccepted
}

func (u *UserDomain) ToModel() models.User {
	return models.User{
		Name:                  u.Name,
		Hash:                  u.Hash,
		Phone:                 u.Phone,
		Email:                 u.Email,
		Password:              u.Password,
		PolicyPrivacyAccepted: u.PolicyPrivacyAccepted,
	}
}

func (u *UserDomain) ToDomain(model models.User) UserDomain {
	return UserDomain{
		Id:                    model.ID,
		Hash:                  model.Hash,
		Name:                  model.Name,
		Email:                 model.Email,
		Phone:                 model.Phone,
		Password:              model.Password,
		PolicyPrivacyAccepted: model.PolicyPrivacyAccepted,
	}
}

func (u *UserDomain) ToDTO() output.UserOutputDTO {
	return output.UserOutputDTO{
		Hash:  u.Hash,
		Name:  u.Name,
		Email: u.Email,
	}
}
