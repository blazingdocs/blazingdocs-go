package main

import (
	"io/ioutil"
	"log"
	"os"

	blazingdocsgo "github.com/blazingdocs/blazingdocs-go"
	"github.com/blazingdocs/blazingdocs-go/config"
	"github.com/blazingdocs/blazingdocs-go/parameters"
	"github.com/blazingdocs/blazingdocs-go/utils"
)

func main() {
	file, err := ioutil.ReadFile("../PO-Template-Array.json")
	s := string(file)
	params := parameters.MergeParameters{
		DataSourceName: "data",
		DataSourceType: utils.JSON_TYPE,
		Strict:         true,
		ParseColumns:   false,
		Sequence:       true,
	}
	if err != nil {

	}
	ffile, _ := os.Open("../PO-Template-Array.docx")
	formFile := utils.FormFile{
		Name:    "PO-Template-Array.docx",
		Content: ffile,
	}
	config.Default = config.Init("YOUR-API-KEY")
	client := blazingdocsgo.Client{
		Config: *config.Default,
	}
	resp, err := client.MergeWithFile(s, "output.pdf", params, formFile)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp)
}
