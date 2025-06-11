package store

type Store struct {
	Users interface {
		GetAll() ([]User, error)
	}
}

func NewStore() *Store {
	return &Store{
		Users: (&UserStore{}).NewUserStore(),
	}
}
