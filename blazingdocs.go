package blazingdocsgo

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/blazingdocs/blazingdocs-go/models"
	"github.com/blazingdocs/blazingdocs-go/parameters"
	"github.com/blazingdocs/blazingdocs-go/utils"
)

type BlazingDocs interface {
	GetAccount() (models.AccountModel, error)

	GetUsage() (models.UsageModel, error)

	GetTemplates(tempatePath string) (models.TemplateModel, error)

	MergeWithFile(data string, fileName string, param parameters.MergeParameters, file utils.FormFile)

	MergeWithGuid(data string, fileName string, param parameters.MergeParameters, guid string)

	MergeWithRelativePath(data string, fileName string, param parameters.MergeParameters, path string)
}

func (client *Client) GetAccount() (acc models.AccountModel, err error) {
	resp, err := client.Get(utils.ACCOUNT_URL)
	var account models.AccountModel
	if err != nil {
		return account, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		var blazingDocsError models.BlazingDocsError
		json.Unmarshal(body, &blazingDocsError)
		return account, &blazingDocsError
	}

	json.Unmarshal(body, &account)
	return account, nil
}

func (client *Client) GetUsage() (us models.UsageModel, err error) {
	resp, err := client.Get(utils.USAGE_URL)
	var usageModel models.UsageModel
	if err != nil {
		return usageModel, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		var blazingDocsError models.BlazingDocsError
		json.Unmarshal(body, &blazingDocsError)
		return usageModel, &blazingDocsError
	}

	json.Unmarshal(body, &usageModel)
	return usageModel, nil
}

func (client *Client) GetTemplates(templatePath string) (temp models.TemplateModel, err error) {
	resp, err := client.Get(utils.USAGE_URL + "/" + templatePath)
	var templateModel models.TemplateModel
	if err != nil {
		return templateModel, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		var blazingDocsError models.BlazingDocsError
		json.Unmarshal(body, &blazingDocsError)
		return templateModel, &blazingDocsError
	}

	json.Unmarshal(body, &templateModel)
	return templateModel, nil
}

func (client *Client) MergeWithFile(data string, fileName string, param parameters.MergeParameters, file utils.FormFile) (op models.OperationModel, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var operationModel models.OperationModel
	_, err = writeDataToField("Data", data, writer)
	if err != nil {
		return operationModel, err
	}

	_, err = writeDataToField("OutputName", fileName, writer)
	if err != nil {
		return operationModel, err
	}

	mergeParam, _ := json.Marshal(param)
	_, err = writeDataToField("MergeParameters", string(mergeParam), writer)
	if err != nil {
		return operationModel, err
	}

	_, err = writeDataToFile("Template", file, writer)
	if err != nil {
		return operationModel, err
	}

	writer.Close()

	resp, err := client.Post(utils.MERGE_URL, writer.FormDataContentType(), bytes.NewReader(body.Bytes()))
	if err != nil {
		return operationModel, err
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var blazingDocsError models.BlazingDocsError
		json.Unmarshal(bodyResp, &blazingDocsError)
		return operationModel, blazingDocsError
	}
	json.Unmarshal(bodyResp, &operationModel)
	return operationModel, err
}

func (client *Client) MergeWithGuid(data string, fileName string, param parameters.MergeParameters, guid string) (op models.OperationModel, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var operationModel models.OperationModel
	_, err = writeDataToField("Data", data, writer)
	if err != nil {
		return operationModel, err
	}

	_, err = writeDataToField("OutputName", fileName, writer)
	if err != nil {
		return operationModel, err
	}

	mergeParam, _ := json.Marshal(param)
	_, err = writeDataToField("MergeParameters", string(mergeParam), writer)
	if err != nil {
		return operationModel, err
	}

	_, err = writeDataToField("Template", guid, writer)
	if err != nil {
		return operationModel, err
	}

	writer.Close()

	resp, err := client.Post(utils.MERGE_URL, writer.FormDataContentType(), bytes.NewReader(body.Bytes()))
	if err != nil {
		return operationModel, err
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var blazingDocsError models.BlazingDocsError
		json.Unmarshal(bodyResp, &blazingDocsError)
		return operationModel, blazingDocsError
	}
	json.Unmarshal(bodyResp, &operationModel)
	return operationModel, err
}

func (client *Client) MergeWithRelativePath(data string, fileName string, param parameters.MergeParameters, path string) (op models.OperationModel, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var operationModel models.OperationModel

	_, err = writeDataToField("Data", data, writer)
	if err != nil {
		return operationModel, err
	}

	_, err = writeDataToField("OutputName", fileName, writer)
	if err != nil {
		return operationModel, err
	}

	mergeParam, _ := json.Marshal(param)
	_, err = writeDataToField("MergeParameters", string(mergeParam), writer)
	if err != nil {
		return operationModel, err
	}

	_, err = writeDataToField("Template", path, writer)
	if err != nil {
		return operationModel, err
	}

	writer.Close()

	resp, err := client.Post(utils.MERGE_URL, writer.FormDataContentType(), bytes.NewReader(body.Bytes()))
	if err != nil {
		return operationModel, err
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var blazingDocsError models.BlazingDocsError
		json.Unmarshal(bodyResp, &blazingDocsError)
		return operationModel, blazingDocsError
	}
	json.Unmarshal(bodyResp, &operationModel)
	return operationModel, err
}

func writeDataToField(name string, data string, writer *multipart.Writer) (models.OperationModel, error) {
	var operationModel models.OperationModel
	fw, err := writer.CreateFormField(name)
	if err != nil {
		return operationModel, err
	}
	_, err = io.Copy(fw, strings.NewReader(data))
	if err != nil {
		return operationModel, err
	}

	return operationModel, nil
}

func writeDataToFile(name string, file utils.FormFile, writer *multipart.Writer) (models.OperationModel, error) {
	var operationModel models.OperationModel
	fw, err := writer.CreateFormFile(name, file.Name)
	if err != nil {
		return operationModel, err
	}
	_, _ = io.Copy(fw, file.Content)
	if err != nil {
		return operationModel, err
	}
	return operationModel, nil
}
