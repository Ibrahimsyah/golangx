package domain

type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type IUserRepository interface {
	Get() (*[]User, error)
	GetById(id string) *User
	Insert(user *User) (string, error)
	Delete(id string) error
}

type IUserUseCase interface {
	GetAllUser() (*[]User, error)
	GetUserById(id string) *User
	InsertUser(user *User) (string, error)
	DeleteUserById(id string) error
}