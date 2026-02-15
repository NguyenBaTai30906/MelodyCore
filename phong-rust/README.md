# Phong VSTRA - Lyric Video

**Status**: Stable | **Tech**: Rust + SDL2

A high-performance lyric video application written in Rust using SDL2. featuring custom typography, transitions, sliced asset loading, and audio synchronization.

Một ứng dụng video lyric hiệu năng cao viết bằng Rust sử dụng SDL2. Có tính năng hiển thị chữ tùy chỉnh, hiệu ứng chuyển cảnh, tải ảnh cắt lớp (sliced assets) và đồng bộ âm thanh.

---

## # EN 🇺🇸

### Features
- **Engine**: Rust + SDL2 (Hardware Accelerated)
- **Assets**: Dynamic background loading from `img/`
- **Audio**: Sync verification with `resources/-.mp3`
- **Effects**: Scanlines, Vignette, Text Shadows

### Setup & Run (Automated)

**Option 1: Automated Setup**
Run the `setup.ps1` script from the repository root (see [Main README](../README.md)).
*Chạy script `setup.ps1` từ thư mục gốc (xem hướng dẫn ở README chính).*

**Option 2: Manual Setup**
1.  Double-click `setup.bat`.
2.  Run:
    ```bash
    cargo run
    ```

### Abbreviation Guide (src/lyrics.rs)
The project uses a compact struct `V` for easy editing:

| Key | Full Name | Description |
|---|---|---|
| `t` | **Text** | The lyric line content |
| `wd` | **Word Delay** | Time (ms) between each word appearing |
| `ld` | **Line Delay** | Time (ms) to wait after line finishes before next line |
| `c` | **Color** | RGB Color struct |
| `d` | **Destination** | `(x, y)` coordinates for text position |
| `f` | **Font** | `b`(bold), `sc`(script), `se`(serif), `p`(playful), `r`(round) |
| `s` | **Size** | Font size in points |

---

## # VI 🇻🇳

### Tính năng
- **Engine**: Rust + SDL2 (Tăng tốc phần cứng)
- **Tài nguyên**: Tự động tải background từ thư mục `img/`
- **Âm thanh**: Đồng bộ với file `resources/-.mp3`
- **Hiệu ứng**: Scanlines (quét dòng), Vignette (tối góc), Bóng chữ

### Hướng dẫn Cài đặt (Tự động)
1.  **Chạy Setup**: Nhấn đúp vào file `setup.bat` ở thư mục gốc `MelodyCore`.
    *   Tool sẽ tự động tải thư viện SDL2 và copy vào đúng chỗ.
2.  **Chạy Visualizer**:
    ```bash
    cargo run
    ```
