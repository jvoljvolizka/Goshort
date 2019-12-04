package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-central-1"))

func getItem(urlid string) (*shortURL, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("URLs"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(urlid),
			},
		},
	}

	// Retrieve the item from DynamoDB. If no matching item is found
	// return nil.
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type
	// map[string]*AttributeValue. We can use the UnmarshalMap helper
	// to parse this straight into the fields of a struct. Note:
	// UnmarshalListOfMaps also exists if you are working with multiple
	// items.
	link := new(shortURL)
	err = dynamodbattribute.UnmarshalMap(result.Item, link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func delItem(urlid string) (string, error) {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("URLs"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(urlid),
			},
		},
	}

	_, err := db.DeleteItem(input)
	if err != nil {
		return "", err
	}
	return "Success", nil
}
