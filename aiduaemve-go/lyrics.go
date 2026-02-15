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

// GetVerses returns lyrics timed to fit ~22.5 seconds total to avoid "feeling slow"
// 1350 ticks target (reduced from 1380 slightly for snappier feel)
func GetVerses() []V {
	wht := color.RGBA{255, 255, 255, 255}
	crm := color.RGBA{255, 245, 220, 255}
	cyn := color.RGBA{200, 240, 255, 255}
	pnk := color.RGBA{255, 200, 210, 255}

	cx := screenWidth / 2
	cy := screenHeight / 2

	return []V{
		// === Verse 1 (Sum: 430) ===
		{
			Words: []Word{W("Đêm", 4), W("nay,", 30), W("sẽ", 4), W("thật", 4), W("dài", 4)}, // sum: 46
			LD:    44,                                                                        // total: 90
			C:     pnk, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Em", 5), W("cũng", 5), W("đã", 5), W("mệt", 5), W("nhoài", 5)}, // sum: 25
			LD:    65,                                                                       // total: 90
			C:     crm, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Sương", 4), W("rơi", 30), W("đêm", 4), W("lạnh", 4), W("đầy", 4)}, // sum: 46
			LD:    44,                                                                          // total: 90
			C:     cyn, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Vậy", 4), W("thì", 4), W("ai", 4), W("đưa", 4), W("em", 4), W("về", 4), W("đêm", 4), W("nay?", 4)}, // sum: 32
			LD:    128,                                                                                                          // total: 160 (extra wait for transition to chorus)
			C:     wht, X: cx, Y: cy, S: 38,
		},

		// === Chorus (Sum: 420) ===
		{
			Words: []Word{W("Take", 5), W("me", 5), W("back", 5), W("back", 5), W("home", 5)}, // sum: 25
			LD:    65,                                                                         // total: 90
			C:     pnk, X: cx, Y: cy, S: 40,
		},
		{
			Words: []Word{W("Đường", 5), W("về", 5), W("cũng", 5), W("chẳng", 5), W("có", 5), W("xa", 5)}, // sum: 30
			LD:    60,                                                                                     // total: 90
			C:     crm, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đưa", 5), W("em", 5), W("về", 5), W("qua", 5), W("ba", 5), W("ngã", 5), W("năm", 5)}, // sum: 35
			LD:    65,                                                                                             // total: 100
			C:     cyn, X: cx, Y: cy, S: 34,
		},
		{
			Words: []Word{W("Năm", 7), W("ngã", 7), W("ba", 7), W("là", 7), W("nhà", 7)}, // sum: 35
			LD:    105,                                                                   // total: 140
			C:     wht, X: cx, Y: cy, S: 38,
		},

		// === Chorus 2 (Sum: 500) ===
		{
			Words: []Word{W("Take", 5), W("me", 5), W("back", 5), W("back", 5), W("home", 5)}, // sum: 25
			LD:    65,                                                                         // total: 90
			C:     pnk, X: cx, Y: cy, S: 40,
		},
		{
			Words: []Word{W("Đường", 5), W("về", 5), W("cũng", 5), W("chẳng", 5), W("có", 5), W("xa", 5)}, // sum: 30
			LD:    60,                                                                                     // total: 90
			C:     crm, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đêm", 5), W("khuya", 5), W("rồi", 5), W("sao", 5), W("không", 5), W("có", 5), W("ai?", 5)}, // sum: 35
			LD:    65,                                                                                                   // total: 100
			C:     cyn, X: cx, Y: cy, S: 36,
		},
		{
			Words: []Word{W("Đưa", 7), W("em", 7), W("đi", 7), W("về", 7), W("nhà.", 7)}, // sum: 35
			LD:    75,                                                                    // total: 110
			C:     wht, X: cx, Y: cy, S: 38,
		},
		{
			Words: []Word{W("Là", 8), W("la", 8), W("laa", 8)}, // sum: 24
			LD:    86,                                          // total: 110
			C:     pnk, X: cx, Y: cy, S: 44,
		},
	}
}
