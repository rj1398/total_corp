In This assignment the user details are fetched on the basis of user ids.

use command:
go run main.go

to run the project


Port Address = 8080

The End Points are:


Request 1: To get single User

url

http://localhost:8080/getUser/1
Response

{
    "id": 1,
    "fname": "Steve",
    "city": "LA",
    "phone": 1234567890,
    "height": 5.8,
    "Married": true
}
Request 2: To get multiple Users

url

http://localhost:8080/getUsers/1,2,3
Response

[
    {
        "id": 1,
        "fname": "Steve",
        "city": "LA",
        "phone": 1234567890,
        "height": 5.8,
        "Married": true
    },
    {
        "id": 2,
        "fname": "Ajay",
        "city": "IN",
        "phone": 9876543210,
        "height": 5.1,
        "Married": true
    },
    {
        "id": 3,
        "fname": "Raj",
        "city": "IN",
        "phone": 9012345678,
        "height": 6.2,
        "Married": false
    }
]
