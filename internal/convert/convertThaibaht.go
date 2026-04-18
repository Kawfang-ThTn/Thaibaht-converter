package convert

import (
	"github.com/shopspring/decimal"
)

// ตัวเลขภาษาไทย
var thaiNum = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}

// หน่วย
var thaiUnit = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}

func Convert(value decimal.Decimal) string {
	// 1. จัดการ หน่วย กรณีเป็น 0
	if value.IsZero() {
		return thaiNum[0]
	}

	baht := value.String()
	bahtStrLen := len(baht)
	result := ""

	for i, v := range baht {
		num := int(v - '0')
		pos := (bahtStrLen - 1) - i
		posMillion := pos % 6

		if num == 0 {
			// ถ้าเป็นหลักล้าน แม้เลขจะเป็น 0 ก็ต้องเติมคำว่า "ล้าน" (ยกเว้นหลักหน่วย ตัวเลขหลังสุด)
			if pos > 0 && posMillion == 0 {
				result += "ล้าน"
			}
			continue
		}

		// 2. จัดการคำอ่านตัวเลข
		if posMillion == 0 && num == 1 && i > 0 {
			// เป็น "เอ็ด" เมื่ออยู่ในหลักหน่วย (posMillion == 0) และไม่ใช่ ตัวเลขแรกสุด
			result += "เอ็ด"
		} else if posMillion == 1 && num == 2 {
			result += "ยี่"
		} else if posMillion == 1 && num == 1 {
			// หลักสิบที่เป็นเลข 1 ไม่ต้องอ่าน "หนึ่ง"
			result += ""
		} else {
			result += thaiNum[num]
		}

		// 3. จัดการชื่อหลัก
		if posMillion == 0 {
			// ถ้าไม่ใช่หลักหน่วยที่อยู่หลังสุด (pos > 0) ให้เติม "ล้าน"
			if pos > 0 {
				result += "ล้าน"
			}
		} else {
			result += thaiUnit[posMillion]
		}
	}

	return result
}
