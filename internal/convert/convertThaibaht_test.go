package convert

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestConvert(t *testing.T) {
	// กำหนดเคสที่ต้องการทดสอบ (Test Cases)
	tests := []struct {
		input    decimal.Decimal
		expected string
	}{
		{decimal.NewFromFloat(1234), "หนึ่งพันสองร้อยสามสิบสี่"},
		{decimal.NewFromFloat(33333), "สามหมื่นสามพันสามร้อยสามสิบสาม"},
		{decimal.NewFromFloat(75), "เจ็ดสิบห้า"},
		{decimal.NewFromFloat(0), "ศูนย์"},
		{decimal.NewFromFloat(10), "สิบ"},
		{decimal.NewFromFloat(21), "ยี่สิบเอ็ด"},
		{decimal.NewFromFloat(1000000), "หนึ่งล้าน"},
	}

	for _, tt := range tests {
		t.Run(tt.input.String(), func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
