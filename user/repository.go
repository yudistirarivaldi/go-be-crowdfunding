package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error) //balikannya user & error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}