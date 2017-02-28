package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// GetItem操作の入力を表します。
	params := &dynamodb.GetItemInput{
		TableName: aws.String("TableName"), // テーブル名

		Key: map[string]*dynamodb.AttributeValue{
			"id": { // キー名
				S: aws.String("foo"), // 持ってくるキーの値
			},
		},
		AttributesToGet: []*string{
			aws.String("username"), // 欲しいデータの名前
		},
		ConsistentRead: aws.Bool(true), // 常に最新を取得するかどうか

		//返ってくるデータの種類
		ReturnConsumedCapacity: aws.String("NONE"),
	}

	// GetItem操作は、指定された主キーを持つ項目の属性のセットを返します。
	// respはGetItem操作の出力を表します。
	resp, err := ddb.GetItem(params)

	if err != nil {
		fmt.Println(err.Error())
	}

	//resp.Item[項目名].型 でデータへのポインタを取得
	fmt.Println(*resp.Item["username"].S)
	fmt.Printf("%T", *resp.Item["username"].S)

}
