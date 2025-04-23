package db

import "github.com/atyagi012/go-user-auth/models"

func SyncDatabase() {
	Database.AutoMigrate(&models.User{})
}
