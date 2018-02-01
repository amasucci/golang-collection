package main

import (
	"./models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//docker run --rm -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=db -e MYSQL_USER=user -e MYSQL_PASSWORD=pass mysql:latest

/*
CREATE TABLE `credentials` (
        `uid` INT(10) NOT NULL AUTO_INCREMENT,
        `username` VARCHAR(64) NOT NULL,
        `password` VARCHAR(64) NOT NULL,
        `tenant` VARCHAR(64) NOT NULL,
        `created` DATETIME NOT NULL,
        `modified` DATETIME NOT NULL,
        PRIMARY KEY (`uid`)
    );
*/

func main() {
	db := models.InitDB("user:pass@tcp(localhost:3306)/db?charset=utf8")
	hash, _ := HashPassword("Password1")
	db.Insert("antonio", hash, "site1")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
