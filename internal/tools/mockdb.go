package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"calvin": {
		AuthToken: "123ABD",
		Username:  "calvin",
	},
	"andrew": {
		AuthToken: "456DEF",
		Username:  "andrew",
	},
	"gerry": {
		AuthToken: "789GHI",
		Username:  "gerry",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"calvin": {
		Coins:    100,
		Username: "calvin",
	},
	"andrew": {
		Coins:    200,
		Username: "andrew",
	},
	"gerry": {
		Coins:    300,
		Username: "gerry",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (c *mockDB) SetupDatabase() error {
	return nil
}
