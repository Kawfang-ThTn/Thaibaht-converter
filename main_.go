package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var thaiNum = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var thaiUnit = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}

func main_() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
		// decimal.NewFromFloat(102.75),
		// decimal.NewFromFloat(121.75),
		// decimal.NewFromFloat(1210000.75),
		// decimal.NewFromFloat(12100000.75),
		// decimal.NewFromFloat(1000000),
		// decimal.NewFromFloat(10000000.75),
	}

	for _, input := range inputs {
		fmt.Println(input)

		// แยก บาท กับ สตางค์
		// หาบาท
		baht := input.Floor() // เช่น 33333.75 → 33333

		// หาสตางค์
		satang := input.Sub(baht).Mul(decimal.NewFromInt(100)) // 0.75 × 100 = 75

		// แปลงเป็น text
		if satang.Equal(decimal.NewFromInt(0)) {
			textBaht := convertThaibaht(baht.String()) + "บาทถ้วน"
			fmt.Println(textBaht)
		} else {
			textBaht := convertThaibaht(baht.String()) + "บาท" + convertThaibaht(satang.String()) + "สตางค์"
			fmt.Println(textBaht)
		}

		// result := convertThaibaht(baht.String())
		// fmt.Println(result)
	}

	// fmt.Println(thaiNum)
	// fmt.Println(thaiUnit)
}

func convertThaibaht(baht string) string {
	// หาจำนวนหลัก
	bahtStrLen := len(baht)

	// return fmt.Sprintf("%d\n", bahtStrLen)

	result := ""

	for i, v := range baht {

		// แปลงตัวเลขให้เป้นคำไทย
		num := int(v - '0') // แปลง character → int เช่น '3' → 3 // ASCII

		// หาตำแหน่งหน่วย (สิบ / ร้อย / พัน)
		pos := (bahtStrLen - 1) - i // ตำแหน่งจากขวา

		// หาว่าเกินหลักล้านไหม
		posMillion := pos % 6 // ตำแหน่งในกลุ่มล้าน

		// ถ้าเลขเป็น 0 ข้าม ไม่แปลงเป็นไทย แต่ถ้าเกินหลักล้านให้เติม "ล้าน"
		if num == 0 {
			if pos > 0 && posMillion == 0 {
				result += "ล้าน"
			}
			continue
		}

		// แปลงตัวเลขเป็นไทย
		if posMillion == 0 && num == 1 {
			// จะเป็น "เอ็ด" ได้ก็ต่อเมื่อ มันไม่ใช่ตัวแรกสุดของ String (i > 0) หรือในกลุ่มล้านนั้นๆ มันมีเลขอื่นข้างหน้ามัน (ในกรณีนี้เช็ก i > 0 จะง่ายที่สุด)
			if i > 0 {
				result += "เอ็ด"
			} else {
				result += "หนึ่ง"
			}
		} else if posMillion == 1 && num == 2 {
			result += "ยี่"
		} else if posMillion == 1 && num == 1 {
			result += ""
		} else {
			result += thaiNum[num]
		}

		// result += thaiUnit[pos]

		// ใส่ชื่อหลัก
		if posMillion == 0 {
			result += "ล้าน"
		} else {
			result += thaiUnit[posMillion]
		}

	}

	return result
}
