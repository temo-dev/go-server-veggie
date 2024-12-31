package storage

import (
	"fmt"
	commonError "server-veggie/common/error"
	"server-veggie/database/schema"
	"server-veggie/modules/user/model"

	"github.com/google/uuid"
)

func (s *sqlStore) InsertUser(data *model.UserCreationType) error {
	id := uuid.NewString()
	user := schema.User{
		UserID:   id,
		UserName: data.UserName,
		Email:    data.Email,
		Password: data.Password,
	}
	fmt.Println("user", user)
	tx := s.db.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return commonError.ErrDB(err)
	}
	tx.Commit()
	return nil
}
