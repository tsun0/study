package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// セッションを持つDynamoDBクライアントの新しいインスタンスを作成
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// DeleteTable操作の入力を表します。
	params := &dynamodb.DeleteTableInput{
		TableName: aws.String("TableName"), // 削除するテーブルの名前
	}

	// DeleteTableは、テーブルとそのすべてのアイテムを削除します。
	resp, err := ddb.DeleteTable(params)

	if err != nil {
		fmt.Println(err.Error()) // エラー処理
	}

	fmt.Println(resp)
}
