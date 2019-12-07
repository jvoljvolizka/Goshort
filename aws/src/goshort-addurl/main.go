package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type shortURL struct {
	ID            string `json:"ID"`
	URL           string `json:"URL"`
	DelToken      string `json:"DelToken"`
	MaxClickCount string `json:"MaxClick"`
}

func addUrl(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newURL shortURL
	body := []byte(req.Body)

	err := json.Unmarshal(body, &newURL)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	check, _ := getItem(newURL.ID)

	if check != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "id already exists",
		}, nil
	}

	if newURL.DelToken == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "DelToken can't be empty",
		}, nil
	}

	if newURL.MaxClickCount == "" {
		newURL.MaxClickCount = "null"
	}

	err = putItem(&newURL)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       req.Body,
	}, nil
}

func main() {
	lambda.Start(addUrl)
}
