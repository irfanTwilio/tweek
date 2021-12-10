package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/piquette/finance-go/quote"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func sms(w http.ResponseWriter, req *http.Request) {

	symbol := req.FormValue("Body")
	fmt.Println(symbol)
	a, _ := quote.Get(symbol)
	fmt.Printf(" --%v--\n", a.ShortName)
	fmt.Printf("Current Price: $%v\n", a.Ask)
	text := fmt.Sprintf("Current Price: $%v\n%s", a.Ask, a.ShortName)
	xml := twiml(text)

	//Twilio syntax

	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("You requested a quote.")

	fmt.Fprintf(w, xml)
	params.SetBody(text)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}

}

func twiml(input string) string {
	begining := `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Message>`
	middle := input
	end := `</Message>
</Response>
`
	out := fmt.Sprint(begining + middle + end)
	return out
}

/*
func () {
		flag.Parse()

		if len(flag.Args()) == 0 {
			logrus.Fatalf("Input one stock symbol", os.Args[0])
		}
		symbol := flag.Args()[0]
		a, _ := quote.Get(symbol)
		fmt.Printf(" --%v--\n", a.ShortName)
		fmt.Printf("Current Price: $%v\n", a.Ask)
	}

*/

func main() {

	http.HandleFunc("/sms", sms)
	http.ListenAndServe(":8090", nil)
}
