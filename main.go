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
	cityCode := flag.String("citycode", "13103", "市区町村コード（例: 13103）")
	start := flag.String("start", today, "取得開始年月日 (YYYYMMDD) 例: 20250208")
	end := flag.String("end", today, "取得終了年月日 (YYYYMMDD) 例: 20250214")
	flag.Parse()

	// APIからデータを取得
	data, err := getPollenData(*cityCode, *start, *end)
	if err != nil {
		log.Fatalf("エラー: %v", err)
	}

	// 取得したデータを表示
	for _, d := range data {
		fmt.Printf(
      "日時: %s, 花粉飛散数: %d\n",
			d.Date.Format("2006-01-02 11:45:14"),
      d.Pollen,
    )
	}
}

