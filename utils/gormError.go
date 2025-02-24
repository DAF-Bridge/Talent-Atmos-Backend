package utils

import "gorm.io/gorm"

func GormErrorAndRowsAffected(result *gorm.DB) error {
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil

}
