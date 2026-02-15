package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
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
	lineEndTime    int
	fadeTimer      int

	audioCtx    *audio.Context
	audioPlayer *audio.Player
	bgImages    []*ebiten.Image
}

func NewApp() *App {
	g := &App{
		verses: GetVerses(),
	}

	// Generate 3 gradient backgrounds (1 high cortisol, 2 low cortisol)
	gradients := []struct {
		r1, g1, b1 uint8
		r2, g2, b2 uint8
	}{
		{180, 40, 80, 40, 10, 60},  // High cortisol: vibrant pink
		{20, 30, 80, 10, 15, 40},   // Low cortisol 1: deep night blue
		{40, 25, 60, 15, 10, 35},   // Low cortisol 2: soft purple
	}

	imgDir := "img"
	os.MkdirAll(imgDir, 0755)

	for i, grad := range gradients {
		img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
		for y := 0; y < screenHeight; y++ {
			t := float64(y) / float64(screenHeight)
			r := uint8(float64(grad.r1)*(1-t) + float64(grad.r2)*t)
			g := uint8(float64(grad.g1)*(1-t) + float64(grad.g2)*t)
			b := uint8(float64(grad.b1)*(1-t) + float64(grad.b2)*t)
			for x := 0; x < screenWidth; x++ {
				img.Set(x, y, color.RGBA{r, g, b, 255})
			}
		}
		// Save as PNG
		path := filepath.Join(imgDir, fmt.Sprintf("%d.png", i))
		f, _ := os.Create(path)
		if f != nil {
			png.Encode(f, img)
			f.Close()
		}
		g.bgImages = append(g.bgImages, ebiten.NewImageFromImage(img))
	}

	// Init audio
	g.audioCtx = audio.NewContext(sampleRate)
	audioFile := "resources/no stress, dance pill.mp3"
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
	verse := g.verses[g.currentIdx]

	words := strings.Fields(verse.T)
	totalWords := len(words)

	if !g.lineDone {
		if verse.WD == 0 {
			g.displayedWords = totalWords
			g.lineDone = true
			g.lineEndTime = g.timer
		} else if g.timer%verse.WD == 0 && g.displayedWords < totalWords {
			g.displayedWords++
			if g.displayedWords >= totalWords {
				g.lineDone = true
				g.lineEndTime = g.timer
			}
		}
	}

	if g.lineDone && (g.timer-g.lineEndTime) >= verse.LD {
		g.prevIdx = g.currentIdx
		g.currentIdx++
		g.timer = 0
		g.displayedWords = 0
		g.lineDone = false
		g.lineEndTime = 0
		g.fadeTimer = 0
	}

	return nil
}

func (g *App) Draw(screen *ebiten.Image) {
	// Draw background with crossfade
	if len(g.bgImages) > 0 {
		currBg := g.currentIdx % len(g.bgImages)
		prevBg := g.prevIdx % len(g.bgImages)
		fadeDur := 20

		if g.fadeTimer < fadeDur && g.currentIdx != g.prevIdx {
			op1 := &ebiten.DrawImageOptions{}
			screen.DrawImage(g.bgImages[prevBg], op1)
			op2 := &ebiten.DrawImageOptions{}
			op2.ColorScale.ScaleAlpha(float32(g.fadeTimer) / float32(fadeDur))
			screen.DrawImage(g.bgImages[currBg], op2)
		} else {
			screen.DrawImage(g.bgImages[currBg], nil)
		}
	} else {
		screen.Fill(color.RGBA{20, 20, 40, 255})
	}

	// Draw scanlines effect
	drawScanlines(screen)

	// Draw lyrics
	if g.currentIdx < len(g.verses) && fontSource != nil {
		verse := g.verses[g.currentIdx]
		words := strings.Fields(verse.T)

		displayed := ""
		for i := 0; i < g.displayedWords && i < len(words); i++ {
			if i > 0 {
				displayed += " "
			}
			displayed += words[i]
		}

		if displayed != "" {
			face := &text.GoTextFace{
				Source: fontSource,
				Size:   verse.S,
			}

			// Measure text for centering
			w, h := text.Measure(displayed, face, 0)
			x := float64(verse.X) - w/2
			y := float64(verse.Y) - h/2

			// Shadow
			shadowOp := &text.DrawOptions{}
			shadowOp.GeoM.Translate(x+3, y+3)
			shadowOp.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 150})
			text.Draw(screen, displayed, face, shadowOp)

			// Main text
			mainOp := &text.DrawOptions{}
			mainOp.GeoM.Translate(x, y)
			mainOp.ColorScale.ScaleWithColor(verse.C)
			text.Draw(screen, displayed, face, mainOp)
		}
	}
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
