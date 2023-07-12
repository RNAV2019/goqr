package model

func GetAllGoqrs() ([]Goqr, error) {
	var goqrs []Goqr
	tx := db.Find(&goqrs)
	if tx.Error != nil {
		return []Goqr{}, tx.Error
	}
	return goqrs, nil
}

func CreateGoqr(goqr Goqr) error {
	tx := db.Create(&goqr)
	return tx.Error
}

func UpdateGoqr(goqr Goqr) error {
	tx := db.Save(&goqr)
	return tx.Error
}

func DeleteGoqr(id int) error {
	tx := db.Unscoped().Delete(&Goqr{}, id)
	return tx.Error
}
