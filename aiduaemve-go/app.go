package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const sampleRate = 44100

var fontSource *text.GoTextFaceSource

func init() {
	// Load Vietnamese-supporting font
	fontData, err := os.ReadFile("resources/fonts/NotoSans-Bold.ttf")
	if err != nil {
		fmt.Println("⚠ Font not found, text may not render properly")
		return
	}
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fontData))
	if err != nil {
		fmt.Printf("⚠ Failed to parse font: %v\n", err)
		return
	}
	fontSource = source
}

type App struct {
	verses         []V
	currentIdx     int
	prevIdx        int
	timer          int
	displayedWords int
	lineDone       bool
	fadeTimer      int
	totalTicks     int

	audioCtx    *audio.Context
	audioPlayer *audio.Player
	bgImages    []*ebiten.Image
	meters      []*ebiten.Image
}

func NewApp() *App {
	g := &App{
		verses: GetVerses(),
	}

	// Load 5 Bocchi backgrounds
	for i := 0; i < 5; i++ {
		path := fmt.Sprintf("resources/img/bg%d.png", i)
		data, err := os.ReadFile(path)
		if err == nil {
			img, _, _ := image.Decode(bytes.NewReader(data))
			g.bgImages = append(g.bgImages, ebiten.NewImageFromImage(img))
		} else {
			fmt.Printf("⚠ Failed to load bg: %s\n", path)
		}
	}

	// Load 2 Cortisol meters
	meterPaths := []string{"resources/img/meter_high.png", "resources/img/meter_low.png"}
	for _, p := range meterPaths {
		data, err := os.ReadFile(p)
		if err == nil {
			img, _, _ := image.Decode(bytes.NewReader(data))
			g.meters = append(g.meters, ebiten.NewImageFromImage(img))
		} else {
			fmt.Printf("⚠ Failed to load meter: %s\n", p)
		}
	}

	// Init audio
	g.audioCtx = audio.NewContext(sampleRate)
	audioFile := "resources/music.mp3"
	data, err := os.ReadFile(audioFile)
	if err != nil {
		fmt.Printf("⚠ Audio file not found: %s\n", audioFile)
	} else {
		decoded, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(data))
		if err != nil {
			fmt.Printf("⚠ Failed to decode audio: %v\n", err)
		} else {
			player, err := g.audioCtx.NewPlayer(decoded)
			if err != nil {
				fmt.Printf("⚠ Failed to create player: %v\n", err)
			} else {
				g.audioPlayer = player
				g.audioPlayer.Play()
				fmt.Println("♪ Playing: Ai Đưa Em Về - TIA")
			}
		}
	}

	return g
}

func (g *App) Update() error {
	if g.currentIdx >= len(g.verses) {
		return ebiten.Termination // End after all lyrics
	}

	g.timer++
	g.fadeTimer++
	g.totalTicks++

	verse := g.verses[g.currentIdx]
	totalWords := len(verse.Words)

	if !g.lineDone {
		if g.displayedWords < totalWords {
			currentWord := verse.Words[g.displayedWords]
			if g.timer >= currentWord.Delay {
				g.displayedWords++
				g.timer = 0
				if g.displayedWords >= totalWords {
					g.lineDone = true
				}
			}
		}
	} else {
		if g.timer >= verse.LD {
			g.prevIdx = g.currentIdx
			g.currentIdx++
			g.timer = 0
			g.displayedWords = 0
			g.lineDone = false
			g.fadeTimer = 0
		}
	}

	return nil
}

func (g *App) Draw(screen *ebiten.Image) {
	// BG Logic:
	// V1 (0-3) = bg0 (Scared/Cool Bocchi)
	// C1 (4-7) = bg1, bg2 (Group shots)
	// C2 (8-12) = bg3, bg4 (Group shots)
	getBgIdx := func(idx int) int {
		if idx < 4 {
			return 0
		} else if idx < 6 {
			return 1
		} else if idx < 8 {
			return 2
		} else if idx < 10 {
			return 3
		}
		return 4
	}

	// Draw background with crossfade
	if len(g.bgImages) > 0 {
		currBg := getBgIdx(g.currentIdx) % len(g.bgImages)
		prevBg := getBgIdx(g.prevIdx) % len(g.bgImages)
		fadeDur := 30

		if g.fadeTimer < fadeDur && g.currentIdx != g.prevIdx && currBg != prevBg {
			op1 := &ebiten.DrawImageOptions{}
			w1, h1 := g.bgImages[prevBg].Bounds().Dx(), g.bgImages[prevBg].Bounds().Dy()
			op1.GeoM.Scale(float64(screenWidth)/float64(w1), float64(screenHeight)/float64(h1))
			screen.DrawImage(g.bgImages[prevBg], op1)

			op2 := &ebiten.DrawImageOptions{}
			w2, h2 := g.bgImages[currBg].Bounds().Dx(), g.bgImages[currBg].Bounds().Dy()
			op2.GeoM.Scale(float64(screenWidth)/float64(w2), float64(screenHeight)/float64(h2))
			op2.ColorScale.ScaleAlpha(float32(g.fadeTimer) / float32(fadeDur))
			screen.DrawImage(g.bgImages[currBg], op2)
		} else {
			op := &ebiten.DrawImageOptions{}
			w, h := g.bgImages[currBg].Bounds().Dx(), g.bgImages[currBg].Bounds().Dy()
			op.GeoM.Scale(float64(screenWidth)/float64(w), float64(screenHeight)/float64(h))
			screen.DrawImage(g.bgImages[currBg], op)
		}
	} else {
		screen.Fill(color.RGBA{20, 20, 40, 255})
	}

	// Draw Cortisol Meter in Top-Left (Increased size)
	if len(g.meters) > 0 {
		meterIdx := 0 // High
		if g.currentIdx >= 4 {
			meterIdx = 1 // Low
		}
		img := g.meters[meterIdx%len(g.meters)]
		op := &ebiten.DrawImageOptions{}
		// Scaled to 320x240 for high clarity
		mw, mh := 320.0, 240.0
		op.GeoM.Scale(mw/float64(img.Bounds().Dx()), mh/float64(img.Bounds().Dy()))
		op.GeoM.Translate(20, 20)
		screen.DrawImage(img, op)
	}

	// Draw scanlines effect
	drawScanlines(screen)

	// Draw lyrics
	if g.currentIdx < len(g.verses) && fontSource != nil {
		verse := g.verses[g.currentIdx]
		displayed := ""
		for i := 0; i < g.displayedWords && i < len(verse.Words); i++ {
			if i > 0 {
				displayed += " "
			}
			displayed += verse.Words[i].Text
		}

		if displayed != "" {
			face := &text.GoTextFace{
				Source: fontSource,
				Size:   verse.S,
			}
			w, h := text.Measure(displayed, face, 0)
			x := float64(verse.X) - w/2
			y := float64(verse.Y) - h/2

			shadowOp := &text.DrawOptions{}
			shadowOp.GeoM.Translate(x+3, y+3)
			shadowOp.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 150})
			text.Draw(screen, displayed, face, shadowOp)

			mainOp := &text.DrawOptions{}
			mainOp.GeoM.Translate(x, y)
			mainOp.ColorScale.ScaleWithColor(verse.C)
			text.Draw(screen, displayed, face, mainOp)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Ai Dua Em Ve - TIA | MelodyCore (Go)\nTicks: %d", g.totalTicks))
}

func (g *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// === Effects ===

var scanlineImg *ebiten.Image

func drawScanlines(screen *ebiten.Image) {
	if scanlineImg == nil {
		scanlineImg = ebiten.NewImage(screenWidth, screenHeight)
		for y := 0; y < screenHeight; y += 3 {
			for x := 0; x < screenWidth; x++ {
				scanlineImg.Set(x, y, color.RGBA{0, 0, 0, 30})
			}
		}
	}
	screen.DrawImage(scanlineImg, nil)
}
