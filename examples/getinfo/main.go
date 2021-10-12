package main

import (
	"log"

	blazingdocsgo "github.com/blazingdocs/blazingdocs-go"
	"github.com/blazingdocs/blazingdocs-go/config"
)

func main() {
	config.Default = config.Init("YOUR-API-KEY")
	client := blazingdocsgo.Client{
		Config: *config.Default,
	}
	resp, err := client.GetAccount()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp)
	var s string
	tempResp, tempErr := client.GetTemplates(s)
	if tempErr != nil {
		log.Println(tempErr.Error())
	}
	log.Println(tempResp)
	useageResp, usageErr := client.GetUsage()
	if usageErr != nil {
		log.Println(usageErr.Error())
	}
	log.Println(useageResp)
}
