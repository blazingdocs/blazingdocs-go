package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/blazingdocs/blazingdocs-go"
	"github.com/blazingdocs/blazingdocs-go/config"
	"github.com/blazingdocs/blazingdocs-go/parameters"
	"github.com/blazingdocs/blazingdocs-go/utils"
)

func main() {
	file, _ := ioutil.ReadFile("../PO-Template.json")
	s := string(file)
	params := parameters.MergeParameters{
		DataSourceName: "data",
		DataSourceType: utils.JSON_TYPE,
		Strict:         true,
		ParseColumns:   false,
		Sequence:       false,
	}
	ffile, _ := os.Open("../PO-Template.docx")
	formFile := utils.FormFile{
		Name:    "PO-Template.docx",
		Content: ffile,
	}
	config.Default = config.Init("YOUR-API-KEY")
	client := blazingdocs.Client{
		Config: *config.Default,
	}
	resp, err := client.MergeWithFile(s, "output.pdf", params, formFile)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp)
}
