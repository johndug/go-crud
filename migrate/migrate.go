package main

import (
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()

}
func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
}
