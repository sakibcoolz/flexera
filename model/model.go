package model

type Record struct {
	ComputerID, UserID, ApplicationID, ComputerType string
}

type Args struct {
	FileName string
	AppId    string
}
