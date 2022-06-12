package models

import "time"

type Medicine struct {
	ID 				 uint32 `gorm:"primary_key;auto_increment" json:"id"`
	GenericName 	 string `gorm:"size:255;not null" json:"generic_name"`
	PhotoURL 		 string `gorm:"size:255;not null" json:"photo_url"`
	Purpose 		 string `gorm:"size:255;not null" json:"purpose"`
	SideEffects 	 string `gorm:"size:255;not null" json:"side_effects"`
	Contraindication string `gorm:"size:255;not null" json:"contraindication"`
	Dosage 			 string `gorm:"size:255;not null" json:"dosage"`
	Ingredients 	 string `gorm:"size:255;not null" json:"ingredients"`
	CreatedAt 		 time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		 time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (medicine *Medicine) CreateMedicine() (*Medicine, error) {
	err := DB.Create(&medicine).Error
	if err != nil {
		return nil, err
	}
	return medicine, nil
}

func GetAllMedicine() []Medicine {
	var medicine []Medicine
	DB.Find(&medicine)
	return medicine
}

func GetMedicineByID(id string) (*Medicine, error) {
	medicine := &Medicine{}
	err := DB.First(&medicine, id).Error
	if err != nil {
		return nil, err
	}
	return medicine, nil
}

func GetMedicineByQuery(query string) ([]Medicine, error) {
	var medicine []Medicine
	err := DB.Where("generic_name LIKE ?", "%"+query+"%").Find(&medicine).Error
	if err != nil {
		return nil, err
	}
	return medicine, nil
}