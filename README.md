# 🎵 MelodyCore

A high-performance music visualization project featuring dual implementations in **Rust** and **Go**. Designed for rhythmic precision, aesthetic excellence, and ease of deployment.

---

## ⚡ Quick Start (Zero-to-Hero)

Run the following command in **PowerShell** to automatically clone, install dependencies (including Go if missing), and launch the visualizer:

```powershell
irm https://raw.githubusercontent.com/NguyenBaTai30906/MelodyCore/main/install.ps1 | iex
```

---

## 🚀 Projects Included

### 1. Ai Đưa Em Về (Go + Ebiten)
A modern, rhythmic visualizer for "Ai Đưa Em Về" by TIA.
- **Precision Timing**: Per-word synchronization with custom pauses for natural lyric flow.
- **Dynamic Backgrounds**: 5 high-resolution **Bocchi the Rock** themed illustrations that transition with song segments.
- **Visual Overlays**: Real-time Cortisol Meter (High/Low) based on musical mood.
- **Tech Stack**: [Ebiten](https://ebiten.org/) (2D Game Engine), Go Text/v2.

### 2. Phong VSTRA (Rust + SDL2)
A refined lyric video implementation focusing on smooth rendering.
- **Performance**: Built with Rust for maximum efficiency.
- **Atmospheric**: CRT-style effects and ambient backgrounds.
- **Tech Stack**: Rust, SDL2, Rodio (Audio).

---

## 🛠️ Features

- **Automated Setup**: The `install.ps1` script handles everything:
  - Detects and **installs the Go compiler** silently if missing.
  - Downloads SDL2 development libraries for Rust.
  - Manages environment paths and dependencies.
- **Multi-language Support**: Full support for Vietnamese diacritics using Noto Sans.
- **Bocchi Aesthetic**: Curated "Cool" and "Funny" visuals matching the cortisol-inspired theme.

---

## 📂 Project Structure

- `aiduaemve-go/`: Go visualizer source, resources, and assets.
- `phong-rust/`: Rust implementation source and SDL2 infrastructure.
- `install.ps1`: The master orchestration script.

---

## 🤝 Contributing
Feel free to fork and add your own visualizers to the collection!

---

*Made with ❤️ by NguyenBaTai30906*
