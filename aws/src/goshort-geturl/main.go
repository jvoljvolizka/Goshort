package main

import (
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type shortURL struct {
	ID            string `json:"ID"`
	URL           string `json:"URL"`
	DelToken      string `json:"DelToken"`
	MaxClickCount string `json:"MaxClick"`
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

	if link.MaxClickCount != "null" {
		Click, err := strconv.Atoi(link.MaxClickCount)

		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       err.Error(),
			}, nil
		}

		if int32(Click) > 1 {
			link.MaxClickCount = strconv.Itoa(Click - 1)
			putItem(link)
		} else {
			_, err = delItem(link.ID)

		}
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       err.Error(),
			}, nil
		}
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
