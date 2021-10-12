package models

type OperationModel struct {
	Id                  string
	Type                string
	PageCount           int
	ElapsedMilliseconds int
	RemoteIpAddress     int
	Files               []FileModel
}
