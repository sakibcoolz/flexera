package utils

import (
	"encoding/csv"
	"flexera/model"
	"io"
	"log"
	"os"
	"strings"
)

func FileReader(filename string, dataTranferChannel chan []string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)

		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		dataTranferChannel <- records
	}
	close(dataTranferChannel)

	return nil
}

func Spliter(str string) model.Record {
	strs := strings.Split(str, "|")
	return model.Record{
		ComputerID:    strs[0],
		UserID:        strs[1],
		ApplicationID: strs[2],
		ComputerType:  strs[3],
	}
}
