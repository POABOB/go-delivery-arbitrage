package main

import (
	"fmt"
	"log"
	"strconv"
    "context"

	"github.com/adshao/go-binance/v2"
)

// 取得現貨價格
func getSpotPrice(client *binance.Client, symbol string) (float64, error) {
	price, err := client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(price[0].Price, 64)
}

// 取得交割合約價格
func getDeliveryPrice(client *binance.Client, symbol string) (float64, error) {
	price, err := client.NewFuturesDeliveryListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(price[0].Price, 64)
}

func main() {
	// 初始化 Binance API 客戶端（無需 API Key）
	client := binance.NewClient("", "")

	// 現貨 & 交割合約交易對
	spotSymbol := "ETHUSDT"
	deliverySymbol := "ETHUSDT_250627"

	// 取得 ETH 現貨價格
	spotPrice, err := getSpotPrice(client, spotSymbol)
	if err != nil {
		log.Fatalf("獲取現貨價格失敗: %v", err)
	}

	// 取得 ETH 交割合約價格
	deliveryPrice, err := getDeliveryPrice(client, deliverySymbol)
	if err != nil {
		log.Fatalf("獲取交割合約價格失敗: %v", err)
	}

	// 計算價差
	priceDiff := deliveryPrice - spotPrice

	// 顯示價格資訊
	fmt.Printf("ETH 現貨價格 (%s): %.2f USDT\n", spotSymbol, spotPrice)
	fmt.Printf("ETH 交割合約價格 (%s): %.2f USDT\n", deliverySymbol, deliveryPrice)
	fmt.Printf("價差: %.2f USDT\n", priceDiff)
}