# BlazingDocs GO client
High-performance document generation API. Generate documents and reports from СSV, JSON, XML with 99,9% uptime and 24/7 monitoring.

## Installation

```go
import (
    blazingdocsgo "github.com/blazingdocs/blazingdocs-go"
    "github.com/blazingdocs/blazingdocs-go/config"
)
```

```
go get -u github.com/blazingdocs/blazingdocs-go
```

## Integration basics

### Setup

You can get your API Key at https://app.blazingdocs.com

```go
config.Default = config.Init("YOUR-API-KEY")
client := blazingdocsgo.Client{
    Config: *config.Default,
}
```

### Getting account info

```go
resp, err := client.GetAccount()
```

### Getting merge templates list

```go
var s string
tempResp, tempErr := client.GetTemplates(s)
```

### Getting usage info

```go
usageResp, usageErr := client.GetUsage()
```

### Executing merge

```go
file, _ := ioutil.ReadFile("../PO-Template.json")
s := string(file)

params := parameters.MergeParameters{
    DataSourceName: "data",
    DataSourceType: utils.JSON_TYPE, // data in json format
    Strict:         true, // keep json types
    ParseColumns:   false,
    Sequence:       false, // data is object
}

ffile, _ := os.Open("../PO-Template.docx")
formFile := utils.FormFile{
    Name:    "PO-Template.docx",
    Content: ffile,
}

config.Default = config.Init("YOUR-API-KEY")
client := blazingdocsgo.Client{
    Config: *config.Default,
}

resp, err := client.MergeWithFile(s, "output.pdf", params, formFile)
```

## Documentation

See more details here https://docs.blazingdocs.com
