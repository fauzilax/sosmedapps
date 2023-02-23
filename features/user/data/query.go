package data

import (
	"errors"
	"log"
	"sosmedapps/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// Register implements user.UserData
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	// validasi cek duplicate email
	dupeEmail := CoreToData(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupeEmail).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email duplicated")
	}
	// validasi cek duplicate username
	dupeUN := CoreToData(newUser)
	err = uq.db.Where("user_name = ?", newUser.UserName).First(&dupeUN).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("username duplicated")
	}
	// proses query
	qry := CoreToData(newUser)
	err = uq.db.Create(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("query error")
	}
	newUser.ID = qry.ID
	return newUser, nil

}

// Login implements user.UserData
func (uq *userQuery) Login(username string) (user.Core, error) {
	qry := User{}
	err := uq.db.Where("email = ? OR user_name = ?", username, username).First(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("query error")
	}
	return DataToCore(qry), nil
}

// Update implements user.UserData
func (uq *userQuery) Update(id int, updateData user.Core) (user.Core, error) {

	if updateData.Email != "" {
		// Proses validasi cek duplicate email
		dupe := CoreToData(updateData)
		err := uq.db.Where("email = ?", dupe.Email).First(&dupe).Error
		if err == nil {
			log.Println("duplicated")
			return user.Core{}, errors.New("email duplicated")
		}
	}
	if updateData.UserName != "" {
		// Proses validasi cek duplicate username
		dupe := CoreToData(updateData)
		err := uq.db.Where("user_name = ?", dupe.UserName).First(&dupe).Error
		if err == nil {
			log.Println("duplicated")
			return user.Core{}, errors.New("username duplicated")
		}
	}
	data := CoreToData(updateData)
	qry := uq.db.Where("id = ?", id).Updates(&data)
	if qry.RowsAffected <= 0 {
		log.Println("update error : no rows affected")
		return user.Core{}, errors.New("update error : no rows updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update error")
		return user.Core{}, errors.New("query error,update fail")
	}
	return DataToCore(data), nil
}

// Profile implements user.UserData
func (uq *userQuery) Profile(id int) (interface{}, error) {
	res := User{}
	err := uq.db.Preload("Content.Comment").Preload("Content").Where("id = ?", id).First(&res).Error
	if err != nil {
		log.Println("query err", err.Error())
		return user.Core{}, nil
	}
	result := make(map[string]interface{})
	result["id"] = res.ID
	result["username"] = res.UserName
	result["name"] = res.Name
	result["bio"] = res.Bio
	result["profilepicture"] = res.Image
	result["content"] = make([]map[string]interface{}, len(res.Content))
	z := 0
	for i, element := range res.Content {
		m := make(map[string]interface{})
		m["id"] = element.ID
		m["content"] = element.Content
		m["content_image"] = element.ContentImage
		m["create_at"] = element.CreatedAt
		m["comments"] = len(res.Content[z].Comment)
		z++
		result["content"].([]map[string]interface{})[i] = m
	}

	return result, nil
}

// Delete implements user.UserData
func (uq *userQuery) Delete(id int) error {
	qry := uq.db.Delete(&User{}, id)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}
	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}
	return nil
}

// Searching implements user.UserData
func (uq *userQuery) Searching(quote string) ([]user.Core, error) {
	find := []User{}
	err := uq.db.Where("email LIKE ?", "%"+quote+"%").Or("user_name LIKE ?", "%"+quote+"%").Find(&find).Error
	if err != nil {
		log.Println("no data processed", err.Error())
		return []user.Core{}, errors.New("no user found")
	}
	res := []user.Core{}
	for i := 0; i < len(find); i++ {
		res = append(res, DataToCore(find[i]))
	}
	return res, nil
}
