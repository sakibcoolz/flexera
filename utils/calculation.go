package utils

import (
	"flexera/constants"
	"fmt"
	"log"
	"math"
	"slices"
)

func (s *Store) BalancePurchase(appid string) {
	laptopApp := fmt.Sprintf("%s_%s", appid, constants.LAPTOP)
	desktopApp := fmt.Sprintf("%s_%s", appid, constants.DESKTOP)

	if !slices.Contains(s.inventryQueue, appid) {
		return
	}
	laptopcount := s.applicationInventory[laptopApp]
	desktopcount := s.applicationInventory[desktopApp]

	if laptopcount > desktopcount {
		final := math.Round(float64(laptopcount-desktopcount) / 2)
		log.Println("we need to purchase applications : ", final)
		log.Println("extra laptops are ", laptopcount-desktopcount)
	} else if laptopcount == desktopcount {
		log.Println("no need to purchase applications")
	} else if laptopcount < desktopcount {
		log.Println("desktop:", desktopcount, " count greater than laptop ", laptopcount, "count for appid", appid)
		log.Println("we need to purchase", desktopcount-laptopcount, "application for balance desktop", desktopcount-laptopcount)
	}
}
