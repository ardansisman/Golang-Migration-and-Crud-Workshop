package seed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ardansisman/Golang-Migration-and-Crud-Workshop/model"
)

func Seed(us model.UserStore) {

	file, err := ioutil.ReadFile("db/seed/users.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []model.User
	if err := json.Unmarshal(file, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range users {
		_, err := us.Create(&item)
		if err != nil {
			fmt.Println(err)
		}
	}
}
