package users_postgres_repository

import "github.com/moond0wner/todo-nilchan/internal/core/domain"

type UserModel struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func UserDomainFromModel(user UserModel) domain.User {
	return domain.NewUser(
		user.ID,
		user.Version,
		user.FullName,
		user.PhoneNumber,
	)
}

func UserDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = UserDomainFromModel(user)
	}
	return userDomains
}
