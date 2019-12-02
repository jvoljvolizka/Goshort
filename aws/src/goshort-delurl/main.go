package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type shortURL struct {
	ID       string `json:"ID"`
	URL      string `json:"URL"`
	DelToken string `json:"DelToken"`
}

func getUrl(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	path := req.Path[1:]

	link, err := getItem(path)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	if link == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusTeapot,
			Body:       "{'Oh all praise to the' : 'G̝̠͓̰ͦͫ͑ͪ͒͞o͎͔̝̦̠͉̊̃ͨ͝L̵̡̰̘̜̫̻̙͖̣ͣ͑d͚̝̥̳̦̳̮̗͑ͩ́̊ͧ̈́ͤ̀͢Ę̛͈̯̦̝͈̥̽ͥ̎ṇ̶̢͕͚̠̱͈̝̫ͬ̓̇̈ͤ ̴̸͇̲̂ͮ̑͑ͩ̓Ď̷̲̤̠̝̘̻̹͇̈ͧ̇̅ͭ̀ͩ͠å̷͖̖̻̩̺̠̘̝͌ͪ̃͠ͅR̨̘̬̦̱̩̖̙͌ͫ͂ͣ̾k̶͙͉̿̇̊N̝͎̈̽͊ͅě̶̢̬̬̺͖͔͍̪̺̠ͩ̐̊̎̚S̷̛̠͖̣͈̖̦̫͍ͤ̈̓̉̋ͮͭ̀̈s̲̙̝̬͍̩͊͗͝͠ !!! '}",
		}, nil
	}

	retmap := map[string]string{
		"Location": link.URL,
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusPermanentRedirect,
		Headers:    retmap,
	}, nil
}

func main() {
	lambda.Start(getUrl)
}
