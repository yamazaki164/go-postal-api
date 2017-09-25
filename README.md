go-postal-api
====

Overview

郵便番号API

## Description

リクエストされた郵便番号7桁に該当する住所情報をJSON形式でレスポンスするAPI

## Requirement

- Go: 1.9+
- github.com/BurntSushi/toml
- github.com/yamazaki164/go-postal/

## Installation

```
go get github.com/yamazaki164/go-postal-api
```

## Usage

```shell
$ ./go-postal-api -c /path/to/config
```

## Configration

|param|value type|description|
|:--|:--|:--|
| port | int | bind port of this api |
| endpoint | string | URI of this api |
| json_dir | string | directory of fetching json data |

## API Params

|param|value type|description|
|:--|:--|:--|
| code | string | 7 number charactors  |

### Response of API
```
{
      "result":[
            {
                  "jis_code":"07447",
                  "postal_code":"9696031",
                  "kana_prefecture":"ﾌｸｼﾏｹﾝ",
                  "kana_address1":"ｵｵﾇﾏｸﾞﾝｱｲﾂﾞﾐｻﾄﾏﾁ",
                  "kana_address2":"ﾋﾛﾂﾞﾗ",
                  "prefecture":"福島県",
                  "address1":"大沼郡会津美里町",
                  "address2":"広面",
                  "flag1":false,
                  "flag2":false,
                  "flag3":false,
                  "flag4":false
            }
      ],
      "status":200
}
```

|fields|description|
|:--|:--|
| flag1 | true: 一町域が二以上の郵便番号で表される場合の表示 |
| flag2 | true: 小字毎に番地が起番されている町域の表示 |
| flag3 | true: 丁目を有する町域の場合の表示 |
| flag4 | true: 一つの郵便番号で二以上の町域を表す場合の表示 |



## Licence

MIT