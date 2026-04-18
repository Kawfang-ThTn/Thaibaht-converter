package main

import (
	"fmt"

	"github.com/Kawfang-ThTn/Thaibaht-converter/internal/convert"
	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}

	for _, input := range inputs {
		fmt.Println(input)

		// แยก บาท กับ สตางค์
		baht := input.Floor()
		satang := input.Sub(baht).Mul(decimal.NewFromInt(100))

		// ทำการ print ผลลัพธ์
		var result string
		if satang.IsZero() {
			result = convert.Convert(baht) + "บาทถ้วน"
		} else {
			result = convert.Convert(baht) + "บาท" + convert.Convert(satang) + "สตางค์"
		}
		fmt.Println(result)
	}
}
