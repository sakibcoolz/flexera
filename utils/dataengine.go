package utils

import (
	"fmt"
	"slices"
	"strings"
)

type Store struct {
	applicationInventory map[string]int
	inventryQueue        []string
}

var (
	uniqueDataMapper = make(map[string]bool)
)

func DataEngine(dataTranferChannel chan []string) *Store {
	store := &Store{
		applicationInventory: make(map[string]int),
		inventryQueue:        make([]string, 0),
	}
	for data := range dataTranferChannel {
		key := strings.Join(
			[]string{strings.ToUpper(data[0]),
				strings.ToUpper(data[1]),
				strings.ToUpper(data[2]),
				strings.ToUpper(data[3])}, "|")

		_, ok := uniqueDataMapper[key]
		if !ok {
			uniqueDataMapper[key] = true

			store.StoreApplicationInvetory(key)

			continue
		}
	}

	uniqueDataMapper = make(map[string]bool)

	return store
}

func (s *Store) StoreApplicationInvetory(str string) {
	record := Spliter(str)

	appIdWithType := fmt.Sprintf("%s_%s", record.ApplicationID, record.ComputerType)

	if !slices.Contains(s.inventryQueue, record.ApplicationID) {
		s.inventryQueue = append(s.inventryQueue, record.ApplicationID)
	}

	s.applicationInventory[appIdWithType]++
}
