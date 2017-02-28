package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"), // プライマリキー名
				AttributeType: aws.String("S"),  // データ型(String:S, Number:N, Binary:B の三種)
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),   // インデックス名
				KeyType:       aws.String("HASH"), // インデックスの型(HASH または RANGE)
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
			ReadCapacityUnits:  aws.Int64(1), // 読み込みキャパシティーユニット（デフォルト：１）
			WriteCapacityUnits: aws.Int64(1), // 書き込みキャパシティーユニット（デフォルト：１）
		},
		TableName: aws.String("TableName"), // テーブル名
	}

	resp, err := ddb.CreateTable(params)

	if err != nil {
		fmt.Println(err.Error()) // エラー処理
	}

	fmt.Println(resp)
}
