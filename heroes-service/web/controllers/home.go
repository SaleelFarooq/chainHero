package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHello()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type HelloData struct {
		Key    string `json:"key"`
		Record struct{
			ItemName   string `json:"itemname"`
			NameOfPerson string `json:"nameofperson"`
			SellingValue string `json:"Sellingvalue"`
	} `json:"Record"`
   }
	var data []HelloData
	json.Unmarshal([]byte(helloValue), &data)
	//fmt.Println(data)
	fmt.Println("JIVE")

	returnData := &struct {
		ResponseData []HelloData
	}{
		ResponseData: data,
	}

	renderTemplate(w, r, "home.html", returnData)
}
