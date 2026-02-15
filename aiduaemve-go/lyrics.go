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
// Target Total Ticks: 1380 (23s * 60TPS)
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
			Words: []Word{W("Đêm", 10), W("nay,", 40), W("sẽ", 10), W("thật", 10), W("dài", 10)}, // sum: 80
			LD:    20,                                                                            // total: 100
			C:     pnk, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Em", 12), W("cũng", 12), W("đã", 12), W("mệt", 12), W("nhoài", 12)}, // sum: 60
			LD:    40,                                                                            // total: 100
			C:     crm, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Sương", 10), W("rơi", 40), W("đêm", 10), W("lạnh", 10), W("đầy", 10)}, // sum: 80
			LD:    20,                                                                              // total: 100
			C:     cyn, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Vậy", 10), W("thì", 10), W("ai", 10), W("đưa", 10), W("em", 10), W("về", 10), W("đêm", 10), W("nay?", 10)}, // sum: 80
			LD:    60,                                                                                                                   // total: 140
			C:     wht, X: cx, Y: cy, S: 38,
		},

		// === Chorus (Sum: 430) ===
		{
			Words: []Word{W("Take", 12), W("me", 12), W("back", 12), W("back", 12), W("home", 12)}, // sum: 60
			LD:    40,                                                                              // total: 100
			C:     pnk, X: cx, Y: cy, S: 40,
		},
		{
			Words: []Word{W("Đường", 10), W("về", 10), W("cũng", 10), W("chẳng", 10), W("có", 10), W("xa", 10)}, // sum: 60
			LD:    40,                                                                                           // total: 100
			C:     crm, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đưa", 10), W("em", 10), W("về", 10), W("qua", 10), W("ba", 10), W("ngã", 10), W("năm", 10)}, // sum: 70
			LD:    40,                                                                                                    // total: 110
			C:     cyn, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Năm", 12), W("ngã", 12), W("ba", 12), W("là", 12), W("nhà", 12)}, // sum: 60
			LD:    60,                                                                         // total: 120
			C:     wht, X: cx, Y: cy, S: 38,
		},

		// === Chorus 2 (Sum: 510) ===
		{
			Words: []Word{W("Take", 12), W("me", 12), W("back", 12), W("back", 12), W("home", 12)}, // sum: 60
			LD:    40,                                                                              // total: 100
			C:     pnk, X: cx, Y: cy, S: 40,
		},
		{
			Words: []Word{W("Đường", 10), W("về", 10), W("cũng", 10), W("chẳng", 10), W("có", 10), W("xa", 10)}, // sum: 60
			LD:    40,                                                                                           // total: 100
			C:     crm, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đêm", 10), W("khuya", 10), W("rồi", 10), W("sao", 10), W("không", 10), W("có", 10), W("ai?", 10)}, // sum: 70
			LD:    40,                                                                                                          // total: 110
			C:     cyn, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đưa", 12), W("em", 12), W("đi", 12), W("về", 12), W("nhà.", 12)}, // sum: 60
			LD:    50,                                                                         // total: 110
			C:     wht, X: cx, Y: cy, S: 38,
		},
		{
			Words: []Word{W("Là", 15), W("la", 15), W("laa", 15)}, // sum: 45
			LD:    45,                                             // total: 90
			C:     pnk, X: cx, Y: cy, S: 44,
		},
	}
}
