package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
)
type Data struct {
	//Key   string `json:"key"`
	ItemName  string `json:"itemname"`
	NameOfPerson string `json:"nameofperson"`
	SellingValue string `json:"sellingvalue"`
}

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	fmt.Println("Check1")
	if r.FormValue("submitted") == "true" {
		fmt.Println("Check2")
		key := r.FormValue("key")
		productData := Data{}
		productData.ItemName = r.FormValue("itemname")
		productData.NameOfPerson = r.FormValue("nameofperson")
		productData.SellingValue = r.FormValue("sellingvalue")
        RequestData, _ := json.Marshal(productData)
		txid, err := app.Fabric.InvokeHello(key,string(RequestData))
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
		fmt.Println(data)
	}
	renderTemplate(w, r, "request.html", data)
}
