# golang-logging-library-for-data-platform

golang-logging-library-for-data-platform は データ連携基盤 のリソースにおいて、 Go ランタイム の マイクロサービス の ログ を出力する際に、ログの json フォーマットを統一するためのライブラリです。

## 動作環境

動作には以下の環境であることを前提とします。

* OS: Linux OS    
* CPU: ARM/AMD/Intel   
* Golang Runtime

## golang-logging-library-for-data-platform のソースコードベース
golang-logging-library-for-data-platform は、[golang-logging-library](https://github.com/latonaio/avis-golang-logging-library) のソースコードをベースとして作成されています。  

## golang-logging-library-for-data-platform の機能
golang-logging-library-for-data-platform は、[golang-logging-library](https://github.com/latonaio/avis-golang-logging-library) の機能に加えて、次の機能を有しています。

* OPEN API(ODATA) における日付/時間項目のJSONフォーマットを一般的な日付時間のJSONフォーマットに変換

golang-logging-library が扱うJSONフォーマットに、特定の日付/時間書式である項目が含まれる場合、次の例にある項目値の変換を行います。  
なお、特定の日付/時間項目であるかどうかの判断基準は、項目値が <"/Date("で始まり>、かつ、<")/"で終わり>、かつ、<これらの間の値が全て数字である>、の条件となっています。  

変換前のフォーマット例：
（※下記の "CreationDate"、"LastChangeDate"が変換前の項目と項目値の例です）
```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-product-master-reads/SAP_API_Caller/caller.go#L98",
	"function": "sap-api-integrations-product-master-reads/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"Product": "A001",
			"IndustrySector": "M",
			"ProductType": "SERV",
			"BaseUnit": "AU",
			"ValidityStartDate": "",
			"ProductGroup": "A001",
			"Division": "00",
			"GrossWeight": "0.000",
			"WeightUnit": "KG",
			"SizeOrDimensionText": "",
			"ProductStandardID": "",
			"CreationDate": "/Date(1466380800000)/",
			"LastChangeDate": "/Date(1491868800000)/",
			"IsMarkedForDeletion": false,
			"NetWeight": "0.000",
			"ChangeNumber": "",
			"to_Description": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCT_SRV/A_Product('A001')/to_Description"
		}
	],
	"time": "2022-01-08T10:38:23.402329+09:00"
}
```

変換後のフォーマット例：
（※下記の "CreationDate"、"LastChangeDate"が変換後の項目と項目値の例です）
```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-product-master-reads/SAP_API_Caller/caller.go#L98",
	"function": "sap-api-integrations-product-master-reads/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"Product": "A001",
			"IndustrySector": "M",
			"ProductType": "SERV",
			"BaseUnit": "AU",
			"ValidityStartDate": "",
			"ProductGroup": "A001",
			"Division": "00",
			"GrossWeight": "0.000",
			"WeightUnit": "KG",
			"SizeOrDimensionText": "",
			"ProductStandardID": "",
			"CreationDate": "2016-06-20T09:00:00+09:00",
			"LastChangeDate": "2017-04-11T09:00:00+09:00",
			"IsMarkedForDeletion": false,
			"NetWeight": "0.000",
			"ChangeNumber": "",
			"to_Description": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCT_SRV/A_Product('A001')/to_Description"
		}
	],
	"time": "2022-01-08T10:38:23.402329+09:00"
}
```

## golang-logging-library-for-data-platform の Output sample.json
golang-logging-library-for-data-platform は、Output として、本レポジトリ内の sample.json にある、次のような json を出力します。
```
{
	"cursor": "/home/ampamman/go/src/sap-api-integrations-product-master-reads/SAP_API_Caller/caller.go#L108",
	"function": "sap-api-integrations-product-master-reads/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"Product": "A001",
			"IndustrySector": "M",
			"ProductType": "SERV",
			"BaseUnit": "AU",
			"ValidityStartDate": "",
			"ProductGroup": "A001",
			"Division": "00",
			"GrossWeight": "0.000",
			"WeightUnit": "KG",
			"SizeOrDimensionText": "",
			"ProductStandardID": "",
			"CreationDate": "2016-06-20T09:00:00+09:00",
			"LastChangeDate": "2017-04-11T09:00:00+09:00",
			"IsMarkedForDeletion": false,
			"NetWeight": "0.000",
			"ChangeNumber": "",
			"to_Description": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCT_SRV/A_Product('A001')/to_Description"
		}
	],
	"time": "2022-01-26T14:51:52.138052513+09:00"
}
```

## 利用方法

#### 本リポジトリを2通りの方法のいずれかでインストールしてください。

【インストール方法①】  
go getでインストールしてください。  

```sh
go get github.com/latonaio/golang-logging-library-for-data-platform/logger@v1.0.0
```

【インストール方法②】  
各マイクロサービスのgo.modに以下のように定義してから、go mod downloadでインストールしてください。  

```
module MODULE-NAME

go 1.18

require (
	github.com/latonaio/golang-logging-library-for-data-platform v1.0.0
)
```

```
go mod download   #全てインストールする場合
go mod download github.com/latonaio/golang-logging-library-for-data-platform@v1.0.0   #一部のみインストールする場合
```

#### 各マイクロサービスのソース内に以下を配置してください。

```go
import "github.com/latonaio/golang-logging-library-for-data-platform/logger"
```

#### インスタンスの作成は下記のように実行します。

```go
l := logger.NewLogger()
```

#### 出力の形式

- log.Fatal(msg) 
- log.Error(msg)
- log.Warn (msg)
- log.Info (msg)
- log.Debug(msg)
 
#### パラメーター

- msg: interface型、文字列で渡した場合とJSONにマッピングできるmapや構造体の形式で渡した場合で挙動が異なります。

#### ログ出力例は以下の通りです。

```go
// 引数に文字列を渡した場合
logging.Debug("This is Test")
{"cursor":"/Users/xxx/test/test.go#L35","level":"DEBUG","message":"This is Test","time":"2021-11-05T18:33:49.495918+09:00"}

// 引数を文字列に渡した場合、フォーマット指定することも可能です
var A = "test"
var B = 111
logging.Debug("%v %v", A, B)
{"cursor":"/Users/xxx/test/test.go#L36","level":"DEBUG","message":"test 111","time":"2021-11-05T18:33:49.496388+09:00"}

// JSONにマッピング可能なマップを渡した場合
var testJson1 = map[string]interface{}{
    "message": "debug",
    "params": []string{
        "nested", "world",
    },
}
logging.Debug(testJson1)
{"cursor":"/Users/xxx/test/test.go#L39","level":"DEBUG","message":"debug","params":["nested","world"],"time":"2021-11-05T18:33:49.496445+09:00"}

// JSONにマッピング可能な構造体を渡した場合
type testStruct struct {
	This string `json:"message"`
	Is   string `json:"is"`
	Test int    `json:"test"`
}

var testJson2 = testStruct{
    "hello", "world", 100,
}
logging.Debug(testJson2)
{"cursor":"/Users/xxx/test/test.go#L40","is":"world","level":"DEBUG","message":"hello","test":100,"time":"2021-11-05T18:33:49.496519+09:00"}

```
