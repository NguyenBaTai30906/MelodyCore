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
		// Verse 1 ~6s
		{"Đêm nay, sẽ thật dài", 6, 20, pnk, cx + 80, cy + 40, 34},
		{"Em cũng đã mệt nhoài", 6, 20, crm, cx - 80, cy - 20, 34},
		{"Sương rơi đêm lạnh đầy", 6, 25, cyn, cx + 40, cy + 60, 36},
		{"Vậy thì ai đưa em về đêm nay?", 6, 40, wht, cx, cy, 38},

		// Chorus ~7s
		{"Take me back back home", 6, 30, pnk, cx - 120, cy - 40, 40},
		{"Đường về cũng chẳng có xa", 7, 30, crm, cx + 60, cy + 50, 36},
		{"Đưa em về qua ba ngã năm", 7, 30, cyn, cx - 40, cy, 34},
		{"Năm ngã ba là nhà", 8, 40, wht, cx + 100, cy - 60, 38},

		// Chorus 2 ~6s
		{"Take me back back home", 6, 25, pnk, cx + 80, cy + 30, 40},
		{"Đường về cũng chẳng có xa", 6, 25, crm, cx - 100, cy - 50, 36},
		{"Đêm khuya rồi sao không có ai?", 6, 30, cyn, cx, cy + 80, 36},
		{"Đưa em đi về nhà.", 8, 30, wht, cx - 40, cy, 38},

		// Outro ~1s
		{"Là la laa...", 10, 60, pnk, cx, cy, 44},
	}
}
