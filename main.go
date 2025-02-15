package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	// start, endのデフォルト値を今日の日付に設定
	today := time.Now().Format("20060102")

	// コマンドライン引数の取得
	cityName := flag.String("cityname", "東京都千代田区", "都市名")
	start := flag.String("start", today, "取得開始年月日 (YYYYMMDD) 例: 20250208")
	end := flag.String("end", today, "取得終了年月日 (YYYYMMDD) 例: 20250214")
	flag.Parse()

	// city codeを取得
	cityCode, err := getCityCodeFuzzy(*cityName)
	if err != nil {
		log.Fatalf("エラー: %v", err)
	}

	// APIからデータを取得
	data, err := getPollenData(cityCode, *start, *end)
	if err != nil {
		log.Fatalf("エラー: %v", err)
	}

	// 取得したデータを表示
  currentTime := time.Now()
	for _, d := range data {
    if d.Date.After(currentTime) {
        continue
    }
		fmt.Printf(
			"日時: %s, 花粉飛散数: %d\n",
			d.Date.Format("2006-01-02 15:04:05"),
			d.Pollen,
		)
	}
}
