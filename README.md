# Pollenso

[WeatherNewsの花粉飛散情報API](https://wxtech.weathernews.com/pollen/index.html#format)を利用して、ウェザーニュースのポールンロボで観測された花粉飛散数を取得するツールです。

## :rocket: Installation

### Build from source

```
$ git clone git@github.com:shunsock/pollenso.git
$ cd pollenso
$ go build .
$ ./pollenso -h
```

### Download binary

Releasesからバイナリをダウンロードしてください。

## Usage

```
$ pollenso -h
Usage of pollenso:
  -cityname string
        都市名 (default "東京都千代田区")
  -end string
        取得終了年月日 (YYYYMMDD) 例: 20250214 (default "20250215")
  -start string
        取得開始年月日 (YYYYMMDD) 例: 20250208 (default "20250215")
```

### Details

- `-cityname`: 都市名を指定します。デフォルトは「東京都千代田区」です。
    - Fuzzy Searchに対応しています。例: `-cityname 港区` で「東京都港区」を指定できます。
    - なお、東京都以外の地域は指定できません。追加する場合は、`citycode.go` を編集し、Pull Requestをお願いします。
- `-start`: 取得開始年月日を指定します。デフォルトはコマンド実行日です。
- `-end`: 取得終了年月日を指定します。デフォルトはコマンド実行日です。


### Example

東京都港区の2025年2月15日の花粉飛散数を取得する例です。

```
$ pollenso -cityname 東京都港区
リクエストURL: https://wxtech.weathernews.com/opendata/v1/pollen?citycode=13103&start=20250215&end=20250215
日時: 2025-02-15 00:00:00, 花粉飛散数: 0
日時: 2025-02-15 01:00:00, 花粉飛散数: 0
日時: 2025-02-15 02:00:00, 花粉飛散数: 0
日時: 2025-02-15 03:00:00, 花粉飛散数: 0
日時: 2025-02-15 04:00:00, 花粉飛散数: 0
日時: 2025-02-15 05:00:00, 花粉飛散数: 0
日時: 2025-02-15 06:00:00, 花粉飛散数: 0
日時: 2025-02-15 07:00:00, 花粉飛散数: 1
日時: 2025-02-15 08:00:00, 花粉飛散数: 0
日時: 2025-02-15 09:00:00, 花粉飛散数: 1
日時: 2025-02-15 10:00:00, 花粉飛散数: 1
日時: 2025-02-15 11:00:00, 花粉飛散数: 0
日時: 2025-02-15 12:00:00, 花粉飛散数: 0
日時: 2025-02-15 13:00:00, 花粉飛散数: 1
日時: 2025-02-15 14:00:00, 花粉飛散数: 1
```

## :warning: 注意

- 本ツールは、非公式です。
- 本ツールは、WeatherNewsの花粉飛散情報APIを利用しています。APIの仕様変更により、動作しなくなる可能性があります。
    - なお、Pull Requestを歓迎します。気軽にお知らせください。
- 本ツールを用いて大量のリクエストを送信することは、WeatherNewsのサーバに負荷をかけることになります。適切な利用をお願いします。
- 本ツールを用いて発生したいかなる問題についても、一切の責任を負いません。

## :key: License

[MIT License](./LICENSE)
