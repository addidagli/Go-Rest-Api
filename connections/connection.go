package connections

import (
	"log"
	"main/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/test?charset=utf8")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrating....")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Wallet{})
}
