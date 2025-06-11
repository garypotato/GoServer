package store

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStore struct {
	users []User
}

func (us *UserStore) NewUserStore() *UserStore {
	return &UserStore{
		users: []User{
			{ID: 1, Username: "john_doe", Email: "john_doe@john.com.au", Password: "password123"},
			{ID: 2, Username: "jane_doe", Email: "jane_doe@john.com.au", Password: "password123"},
		},
	}
}

func (us *UserStore) GetAll() ([]User, error) {
	return us.users, nil
}
