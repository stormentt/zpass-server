package db

type Password struct {
	ID       int
	Selector string `gorm:"unique_index;size:20"`

	Data []byte `gorm:"type:longblob"`

	User   User
	UserID int
}

func FindPassword(sel string) (*Password, error) {
	pass := Password{}
	err := Con.Where("selector = ?", sel).Find(&pass).Error
	if err != nil {
		return nil, err
	}

	return &pass, nil
}
