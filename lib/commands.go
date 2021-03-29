package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch_btc_data(){
	resp, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The request failed with error: %s", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("price of bitcoin is currently: %s", data)
}