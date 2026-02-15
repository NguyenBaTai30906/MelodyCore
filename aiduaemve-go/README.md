# Ai Đưa Em Về - Bocchi Visualizer

**Status**: Stable | **Tech**: Go + Ebiten

A rhythmic music visualizer written in Go using the Ebiten engine. Features Bocchi the Rock themed visuals and cortisol level tracking.

Một ứng dụng visualizer âm nhạc theo nhịp viết bằng Go sử dụng Ebiten engine. Có tính năng hiển thị hình ảnh chủ đề Bocchi the Rock và theo dõi mức độ Cortisol.

---

## # EN 🇺🇸

### Features
- **Engine**: Go + Ebiten (2D Game Engine)
- **Timing**: High-precision per-word delay system (Ticks)
- **Assets**: 5 Dynamic Bocchi backgrounds + Cortisol Meter overlays
- **Audio**: Sync with `resources/music.mp3`

### Setup & Run (Automated)

**Option 1: Automated Setup**
Run the `install.ps1` script from the repository root.
*Chạy script `install.ps1` từ thư mục gốc.*

**Option 2: Manual Setup**
1.  Install Go compiler.
2.  Run:
    ```bash
    go run .
    ```

### Abbreviation Guide (lyrics.go)
The project uses structs `V` and `Word` for precise timing:

| Key | Full Name | Description |
|---|---|---|
| `Words` | **Words** | Slice of `Word` structs (Text + Delay) |
| `LD` | **Line Delay** | Ticks to wait after the last word finishes |
| `C` | **Color** | `RGBA` color for the text |
| `X`, `Y` | **Position** | Centered coordinates for the verse |
| `S` | **Size** | Font size |

Each `Word` has a `Delay` (ticks to wait AFTER showing that word).

---

## # VI 🇻🇳

### Tính năng
- **Engine**: Go + Ebiten (Engine game 2D)
- **Timing**: Hệ thống trễ từng chữ độ chính xác cao (tính bằng Tick)
- **Tài nguyên**: 5 Background Bocchi thay đổi linh hoạt + Lớp phủ đồng hồ Cortisol
- **Âm thanh**: Đồng bộ với file `resources/music.mp3`

### Hướng dẫn Cài đặt
1.  **Chạy qua Script**: Sử dụng lệnh `irm ... | iex` ở thư mục gốc để tự động cài Go và chạy.
2.  **Chạy thủ công**:
    ```bash
    go run .
    ```
