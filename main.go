package total_corp

import (
	"log"

	c "total_corp/constant1"
	"total_corp/model"
	"total_corp/router"
)

func main() {
	log.Println("Intializing DB.....")
	tmp := make(map[int]model.User)
	tmp[1] = model.User{Id: 1, Name: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	tmp[2] = model.User{Id: 2, Name: "Ajay", City: "IN", Phone: 9876543210, Height: 5.10, Married: true}
	tmp[3] = model.User{Id: 3, Name: "Raj", City: "IN", Phone: 9012345678, Height: 6.2, Married: false}
	c.Users = tmp
	log.Println("Starting the server")
	router.Controller()
}
