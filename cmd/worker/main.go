package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func getCurrency() {

	url := "http://api.fixer.io/latest?base=EUR"
	content, err := http.Get(url) //Gets the api's content based on url
	if err != nil {
		fmt.Printf("something went wrong with getcurrency: %v", err.Error())
	}
	body, err := ioutil.ReadAll(content.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	c := database.Currency{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Printf("Check Internett connection. getCurrency failed %v", err)
	}
	db := myDB()
	db.Addcurrency(c)

}
	func myDB() database.MgoDB {
	//mongodb://<dbuser>:<dbpassword>@ds141274.mlab.com:41274/cloudimt2681
		mydb := database.MgoDB {
			"mongodb://admin:imt2681@dds141274.mlab.com:41274/cloudimt2681",
				"cloudimt2681",
				"currency",
				"webhooks",
	}
	return mydb

}