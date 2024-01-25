package service

type Service interface {
	User() User
}

func New() (Service, error) {
	return nil, nil
}
