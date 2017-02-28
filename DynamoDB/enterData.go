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

	// UpdateItem操作の入力を表します。
	param := &dynamodb.UpdateItemInput{
		// Stringは、渡された文字列値へのポインタを返します。
		TableName: aws.String("TableName"), // テーブル名を指定

		// 属性のデータを表します。
		// 各属性値は、名前と値のペアとして記述されます。
		// 名前はデータ型であり、値はデータそのものです。
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("foo"), // キー名を指定
			},
		},

		// 属性の追加。
		ExpressionAttributeNames: map[string]*string{
			"#username": aws.String("username"), // 項目名をプレースホルダに入れる
		},
		// 属性のデータを表します。
		// 各属性値は、名前と値のペアとして記述されます。
		// 名前はデータ型であり、値はデータそのものです。
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":username_value": {
				S: aws.String("hoge"), // 値をプレースホルダに入れる
			},
		},
		// 更新式の詳細については、「Amazon DynamoDB開発者ガイド」の「アイテムと属性の変更」を参照してください。
		UpdateExpression: aws.String("set #username = :username_value"), //プレースホルダを利用して更新の式を書く

		//あとは返してくる情報の種類を指定する
		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	}

	// 既存の項目の属性を編集するか、または新しい項目をテーブルに追加します（存在しない場合）。
	resp, err := ddb.UpdateItem(param) //実行

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp)
}
