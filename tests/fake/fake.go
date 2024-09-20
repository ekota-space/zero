package fake

import (
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/go-faker/faker/v4"
)

func GenerateRandomRegisterDao() authDao.RegisterDao {
	return authDao.RegisterDao{
		Username:  faker.Username(),
		Email:     faker.Email(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Password:  faker.Password(),
	}
}
