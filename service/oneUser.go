package service

import (
	"log"
	"net/http"
	"strconv"
	c "total_corp/constant1"
	"total_corp/model"

	"github.com/gorilla/mux"
)

func GetUser(r *http.Request) (model.User, model.Error) {
	log.Println("Inside GetUser()")
	req := mux.Vars(r)
	key := req["id"]
	id, _ := strconv.Atoi(key)
	res, present := c.Users[id]
	if present {
		log.Println("Got the data...")
		log.Println("Completed GetUser()")
		return res, model.Error{}
	}
	log.Println("Data not found...")
	log.Println("Completed GetUser()")
	return res, model.Error{Message: "User not found"}
}
