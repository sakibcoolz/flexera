package main

import (
	"flag"
	"flexera/model"
	"flexera/utils"
	"log"
)

var (
	args               model.Args
	dataTranferChannel = make(chan []string, 1000)
)

func init() {
	Arguments()
}

func Arguments() {
	filename := flag.String("filename", "stage/sample-small.csv", "csv filename with path")
	applicationId := flag.String("appid", "374", "application id")

	flag.Parse()
	if *filename == "" || *applicationId == "" {
		log.Fatal("argument missing")
	}

	args.FileName = *filename
	args.AppId = *applicationId
}

func main() {
	go func(filename string, dataTranferChannel chan []string) {
		err := utils.FileReader(filename, dataTranferChannel)
		if err != nil {
			log.Fatal("cant process file due to error", err.Error())
		}
	}(args.FileName, dataTranferChannel)

	store := utils.DataEngine(dataTranferChannel)

	if args.AppId != "0" {
		store.BalancePurchase(args.AppId)
	}
}
