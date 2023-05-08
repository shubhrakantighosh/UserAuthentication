package main

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

func main() {
	//initializers.LoadEnvVariables()
	//db := initializers.ConnectToDB()
	//initializers.SyncDatabase(db)
	//ctrl := user.Wire(db)
	//router := router.Router(ctrl)
	//server := &http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//}
	//server.ListenAndServe()

	phoneNumber := "+917044063694"
	regionCode := "US"

	parsedNumber, err := phonenumbers.Parse(phoneNumber, regionCode)
	if err != nil {

	}

	isValid := phonenumbers.IsValidNumberForRegion(parsedNumber, regionCode)
	if isValid {
		fmt.Println("Phone number is valid!")
	} else {
		fmt.Println("Phone number is not valid!")
	}

	//formattedNum := phonenumbers.GetExampleNumber("SA")
	//f := phonenumbers.GetExampleNumber("IN")
	//fmt.Println(len(strconv.FormatUint(formattedNum.GetNationalNumber(), 10)))
	//fmt.Println(len(strconv.FormatUint(f.GetNationalNumber(), 10)))
}
