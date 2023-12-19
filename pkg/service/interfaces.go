package service

type SendMail interface {
	Send() error
}

type SaveDBRepository interface {
	Save() error
}
