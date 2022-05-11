package numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestConvertToRoman(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("Converting %d to %q", tt.arabic, tt.roman), func(t *testing.T) {
			if got := ConvertToRoman(tt.arabic); got != tt.roman {
				t.Errorf("ConvertToRoman() = %v, want %v", got, tt.roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("Converting %q to %d", tt.roman, tt.arabic), func(t *testing.T) {
			if got := ConvertToArabic(tt.roman); got != tt.arabic {
				t.Errorf("ConvertToArabic() = %v, want %v", got, tt.arabic)
			}
		})
	}
}

var cases = []struct {
	arabic uint16
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{41, "XLI"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{2014, "MMXIV"},
	{1984, "MCMLXXXIV"},
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}
