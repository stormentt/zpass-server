package db

type User struct {
	ID       int
	Selector string `gorm:"unique_index;size:20"`

	//Devices []Device
}

func FindUser(sel string) (*User, error) {
	user := User{}
	err := Con.Where("selector = ?", sel).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
