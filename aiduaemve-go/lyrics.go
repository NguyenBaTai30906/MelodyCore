package main

import "image/color"

// Word represents a single word with its specific delay
type Word struct {
	Text  string
	Delay int // Ticks to wait AFTER this word appears before showing the next
}

// V is the verse struct - now supports per-word timing
type V struct {
	Words []Word     // List of words with their delays
	LD    int        // LineDelay (ticks to hold entire line after last word)
	C     color.RGBA // Color
	X     int        // X position
	Y     int        // Y position
	S     float64    // Font size
}

func W(t string, d int) Word { return Word{t, d} }

// GetVerses returns lyrics timed to fit ~23 seconds total
// 23s * 60 TPS = 1380 ticks budget (440 + 430 + 510)
func GetVerses() []V {
	wht := color.RGBA{255, 255, 255, 255}
	crm := color.RGBA{255, 245, 220, 255}
	cyn := color.RGBA{200, 240, 255, 255}
	pnk := color.RGBA{255, 200, 210, 255}

	cx := screenWidth / 2
	cy := screenHeight / 2

	return []V{
		// === Verse 1 (Sum: 440) ===
		{
			Words: []Word{W("Đêm", 5), W("nay,", 35), W("sẽ", 5), W("thật", 5), W("dài", 5)}, // sum: 55
			LD:    45,                                                                        // total: 100
			C:     pnk, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Em", 6), W("cũng", 6), W("đã", 6), W("mệt", 6), W("nhoài", 6)}, // sum: 30
			LD:    70,                                                                       // total: 100
			C:     crm, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Sương", 5), W("rơi", 35), W("đêm", 5), W("lạnh", 5), W("đầy", 5)}, // sum: 55
			LD:    45,                                                                          // total: 100
			C:     cyn, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Vậy", 5), W("thì", 5), W("ai", 5), W("đưa", 5), W("em", 5), W("về", 5), W("đêm", 5), W("nay?", 5)}, // sum: 40
			LD:    100,                                                                                                          // total: 140
			C:     wht, X: cx, Y: cy, S: 38,
		},

		// === Chorus (Sum: 430) ===
		{
			Words: []Word{W("Take", 6), W("me", 6), W("back", 6), W("back", 6), W("home", 6)}, // sum: 30
			LD:    70,                                                                         // total: 100
			C:     pnk, X: cx, Y: cy, S: 40,
		},
		{
			Words: []Word{W("Đường", 6), W("về", 6), W("cũng", 6), W("chẳng", 6), W("có", 6), W("xa", 6)}, // sum: 36
			LD:    64,                                                                                     // total: 100
			C:     crm, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đưa", 6), W("em", 6), W("về", 6), W("qua", 6), W("ba", 6), W("ngã", 6), W("năm", 6)}, // sum: 42
			LD:    68,                                                                                             // total: 110
			C:     cyn, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Năm", 8), W("ngã", 8), W("ba", 8), W("là", 8), W("nhà", 8)}, // sum: 40
			LD:    80,                                                                    // total: 120
			C:     wht, X: cx, Y: cy, S: 38,
		},

		// === Chorus 2 (Sum: 510) ===
		{
			Words: []Word{W("Take", 6), W("me", 6), W("back", 6), W("back", 6), W("home", 6)}, // sum: 30
			LD:    70,                                                                         // total: 100
			C:     pnk, X: cx, Y: cy, S: 40,
		},
		{
			Words: []Word{W("Đường", 6), W("về", 6), W("cũng", 6), W("chẳng", 6), W("có", 6), W("xa", 6)}, // sum: 36
			LD:    64,                                                                                     // total: 100
			C:     crm, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đêm", 6), W("khuya", 6), W("rồi", 6), W("sao", 6), W("không", 6), W("có", 6), W("ai?", 6)}, // sum: 42
			LD:    68,                                                                                                   // total: 110
			C:     cyn, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đưa", 8), W("em", 8), W("đi", 8), W("về", 8), W("nhà.", 8)}, // sum: 40
			LD:    70,                                                                    // total: 110
			C:     wht, X: cx, Y: cy, S: 38,
		},
		{
			Words: []Word{W("Là", 10), W("la", 10), W("laa", 10)}, // sum: 30
			LD:    60,                                             // total: 90
			C:     pnk, X: cx, Y: cy, S: 44,
		},
	}
}
