package output

type UserOutputDTO struct {
	Hash     string `json:"hash"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
