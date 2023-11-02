package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	c "total_corp/constant1"
	"total_corp/model"
)

func TestGetUser1(t *testing.T) {
	tmp := make(map[int]model.User)
	tmp[1] = model.User{Id: 1, Name: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	tmp[2] = model.User{Id: 2, Name: "Ajay", City: "IN", Phone: 9876543210, Height: 5.10, Married: true}
	tmp[3] = model.User{Id: 3, Name: "Raj", City: "IN", Phone: 9012345678, Height: 6.2, Married: false}
	c.Users = tmp
	url := "http://localhost:8080/getUser/1"
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	response, _ := client.Do(req)
	bytesVal, _ := ioutil.ReadAll(response.Body)
	var res model.User
	json.Unmarshal(bytesVal, &res)
	if res.Id != 1 {
		t.Errorf("Failed during TestGetUser1")
	}
}

func TestGetUser2(t *testing.T) {
	tmp := make(map[int]model.User)
	tmp[1] = model.User{Id: 1, Name: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	tmp[2] = model.User{Id: 2, Name: "Ajay", City: "IN", Phone: 9876543210, Height: 5.10, Married: true}
	tmp[3] = model.User{Id: 3, Name: "Raj", City: "IN", Phone: 9012345678, Height: 6.2, Married: false}
	c.Users = tmp
	url := "http://localhost:8080/getUser/4"
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	response, _ := client.Do(req)
	bytesVal, _ := ioutil.ReadAll(response.Body)
	var res model.User
	json.Unmarshal(bytesVal, &res)
	if res.Id != 0 {
		t.Errorf("Failed during TestGetUser2")
	}
}
