package business

import (
	"fmt"
	commonError "server-veggie/common/error"
	"server-veggie/modules/user/model"
	utils "server-veggie/src/utils"
)

type CreateUserStorage interface {
	InsertUser(data *model.UserCreationType) error
}

type createUsersBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUsersBiz {
	return &createUsersBiz{store: store}
}

func (biz *createUsersBiz) CreateNewUser(data *model.UserCreationType) error {
	fmt.Println("data---2", data)
	if err := biz.store.InsertUser(data); err != nil {
		return commonError.ErrCannotCreateUser(model.EntityName, err)
	}
	// Cấu hình thông tin email
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := "tuananh.ngta@gmail.com"
	senderPassword := "Tuananh181993" // Khuyến nghị sử dụng biến môi trường thay vì hardcode
	recipientEmail := "anh.ngt18@gmail.com"
	subject := "Hello from Golang"
	body := "This is a test email sent using Go!"
	// Gửi email
	err := utils.SendEmail(smtpHost, smtpPort, senderEmail, senderPassword, recipientEmail, subject, body)
	if err != nil {
		fmt.Println("Lỗi khi gửi email:", err)
	} else {
		fmt.Println("Email đã được gửi thành công!")
	}
	return nil
}
