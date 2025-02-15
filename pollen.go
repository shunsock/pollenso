package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// PollenData はAPIから取得する花粉データの構造体です。
type PollenData struct {
	CityCode string
	Date     time.Time
	Pollen   int
}

// getPollenData はウェザーニュースのAPIから花粉データを取得します。
// 取得するデータはウェザーニュースのポールンロボで観測されたデータです。
// cityCode: 市区町村コード（例: 13103）
// start: 取得開始年月日 (YYYYMMDD) 例: 20250208
// end: 取得終了年月日 (YYYYMMDD) 例: 20250214
func getPollenData(cityCode, start, end string) ([]PollenData, error) {
	// APIエンドポイントのURL作成
	url := fmt.Sprintf("https://wxtech.weathernews.com/opendata/v1/pollen?citycode=%s&start=%s&end=%s", cityCode, start, end)
	fmt.Println("リクエストURL:", url)

	// HTTP GETリクエスト
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTPリクエスト失敗: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("リクエストが失敗しました。ステータスコード: %d", resp.StatusCode)
	}

	// CSVリーダーの作成
	reader := csv.NewReader(resp.Body)
	var data []PollenData

	// 最初の行を読み込み、ヘッダー行かどうか確認
	firstRow, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("CSVの読み込みに失敗: %w", err)
	}
	if len(firstRow) > 0 && firstRow[0] == "citycode" {
		// ヘッダー行の場合は、次の行から読み込み開始
	} else {
		// ヘッダーがない場合は、最初の行もデータとして処理
		record, err := parseRecord(firstRow)
		if err != nil {
			log.Printf("行のパースに失敗: %v", err)
		} else {
			data = append(data, record)
		}
	}

	// 残りのレコードを読み込み
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("CSV読み込み中のエラー: %v", err)
			continue
		}
		record, err := parseRecord(row)
		if err != nil {
			log.Printf("行のパースに失敗: %v", err)
			continue
		}
		data = append(data, record)
	}
	return data, nil
}

// CSVの1行を受け取り PollenData に変換
func parseRecord(record []string) (PollenData, error) {
	if len(record) < 3 {
		return PollenData{}, fmt.Errorf("行の要素数が不足しています: %v", record)
	}

	// 1列目: 市区町村コード
	citycode := record[0]

	// 2列目: 日時（ISO8601形式: 例 "2025-02-03T12:00:00+09:00"）
	parsedDate, err := time.Parse(time.RFC3339, record[1])
	if err != nil {
		return PollenData{}, fmt.Errorf("日付のパースに失敗: %s, エラー: %v", record[1], err)
	}

	// 3列目: 花粉飛散数（整数に変換、欠測値は -9999 として扱う）
	pollen, err := strconv.Atoi(record[2])
	if err != nil {
		return PollenData{}, fmt.Errorf("花粉飛散数のパースに失敗: %s, エラー: %v", record[2], err)
	}

	return PollenData{
		CityCode: citycode,
		Date:     parsedDate,
		Pollen:   pollen,
	}, nil
}
