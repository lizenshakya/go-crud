package initializers

import "github.com/lizenshakya/go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
