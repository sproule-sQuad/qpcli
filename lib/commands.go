package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func FetchBtcData(){
	resp, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The request failed with error: %s", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("price of bitcoin is currently: %s", data)
}

func InstallDev() int32 {
	cmd := &exec.Cmd{
		Path: "./install.sh",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err := cmd.Start()
	cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	return 0

}