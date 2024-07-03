package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockUserHeart = map[string]UserHearts{
	"alex": {
		Heart:    "Black",
		Username: "alex",
	},
	"jason": {
		Heart:    "Red",
		Username: "jason",
	},
	"marie": {
		Heart:    "Red",
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserHearts(username string) *UserHearts {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = UserHearts{}
	clientData, ok := mockUserHeart[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
