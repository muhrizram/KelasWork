package user

import (
	"crypto/cipher"
	"go-restaurant/m/internal/model"

	"gorm.io/gorm"
)

type userRepo struct {
	db      *gorm.DB
	gcm     cipher.AEAD
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func GetRepository(
	db *gorm.DB,
	time uint32,
	memory uint32,
	threads uint8,
	keyLen uint32,
) Repository {
	return &userRepo{
		db:      db,
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

func (u *userRepo) RegisterUser(userData model.User) (model.User, error) {
	if err := u.db.Create(&userData).Error; err != nil {
		return userData, err
	}
	return userData, nil
}

func (u *userRepo) CheckRegisetered(username string) (bool, error) {
	var userData model.User

	if err := u.db.Where(model.User{Username: username}).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}

	}

	return userData.Username != "", nil
}
