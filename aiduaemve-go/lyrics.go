package main

import "image/color"

// V is the verse struct - mirrors Rust's compact naming
type V struct {
	T  string     // Text
	WD int        // WordDelay (ticks between words, 60 ticks = 1s)
	LD int        // LineDelay (ticks to hold after line finishes)
	C  color.RGBA // Color
	X  int        // X position
	Y  int        // Y position
	S  float64    // Font size
}

// GetVerses returns lyrics timed to fit ~23 seconds total
// 23s * 60 TPS = 1380 ticks budget
func GetVerses() []V {
	wht := color.RGBA{255, 255, 255, 255}
	crm := color.RGBA{255, 245, 220, 255}
	cyn := color.RGBA{200, 240, 255, 255}
	pnk := color.RGBA{255, 200, 210, 255}

	cx := screenWidth / 2
	cy := screenHeight / 2

	return []V{
		// Verse 1 (1380 Ticks total target)
		{"Đêm nay,", 0, 40, pnk, cx + 80, cy + 40, 34},               // 40
		{"sẽ thật dài", 5, 60, pnk, cx + 80, cy + 40, 34},            // 15+60=75  (Sum: 115)
		{"Em cũng đã mệt nhoài", 5, 70, crm, cx - 80, cy - 20, 34},   // 25+70=95  (Sum: 210)
		{"Sương rơi đêm lạnh đầy", 5, 70, cyn, cx + 40, cy + 60, 36}, // 25+70=95  (Sum: 305)
		{"Vậy thì ai đưa em về đêm nay?", 4, 80, wht, cx, cy, 38},    // 32+80=112 (Sum: 417)

		// Chorus
		{"Take me back back home", 4, 70, pnk, cx - 120, cy - 40, 40},   // 20+70=90  (Sum: 507)
		{"Đường về cũng chẳng có xa", 5, 70, crm, cx + 60, cy + 50, 36}, // 30+70=100 (Sum: 607)
		{"Đưa em về qua ba ngã năm", 5, 70, cyn, cx - 40, cy, 34},       // 35+70=105 (Sum: 712)
		{"Năm ngã ba là nhà", 5, 90, wht, cx + 100, cy - 60, 38},        // 25+90=115 (Sum: 827)

		// Chorus 2
		{"Take me back back home", 4, 70, pnk, cx + 80, cy + 30, 40},     // 20+70=90  (Sum: 917)
		{"Đường về cũng chẳng có xa", 4, 70, crm, cx - 100, cy - 50, 36}, // 24+70=94  (Sum: 1011)
		{"Đêm khuya rồi sao không có ai?", 4, 80, cyn, cx, cy + 80, 36},  // 28+80=108 (Sum: 1119)
		{"Đưa em đi về nhà.", 5, 80, wht, cx - 40, cy, 38},               // 25+80=105 (Sum: 1224)

		// Outro: 1380 - 1224 = 156 ticks total
		// Word count 3 * 8 = 24. LD = 156 - 24 = 132.
		{"Là la laa...", 8, 132, pnk, cx, cy, 44}, // 24+132=156 (Total: 1380)
	}
}
