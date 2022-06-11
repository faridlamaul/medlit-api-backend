package models

import "gorm.io/gorm"

// ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
// CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	
type Medicine struct {
	gorm.Model
	GenericName string `gorm:"size:255;not null" json:"generic_name"`
	PhotoURL string `gorm:"size:255;not null" json:"photo_url"`
	Purpose string `gorm:"size:255;not null" json:"purpose"`
	SideEffects string `gorm:"size:255;not null" json:"side_effects"`
	Contraindication string `gorm:"size:255;not null" json:"contraindication"`
}

func (m *Medicine) CreateMedicine() (*Medicine, error) {
	err := DB.Create(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetAllMedicine() []Medicine {
	var medicines []Medicine
	DB.Find(&medicines)
	return medicines
}

func GetMedicineByID(id string) (*Medicine, error) {
	m := &Medicine{}
	err := DB.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetMedicineByQuery(query string) ([]Medicine, error) {
	var medicines []Medicine
	DB.Where("generic_name LIKE ?", "%"+query+"%").Find(&medicines)
	return medicines, nil
}