use sdl2::pixels::Color;
use sdl2::rect::Rect;
use sdl2::render::{BlendMode, Canvas, TextureCreator};
use sdl2::ttf::Sdl2TtfContext;
use sdl2::video::{Window, WindowContext};
use std::path::Path;
use crate::lyrics;

pub fn scanlines(canvas: &mut Canvas<Window>, alpha: u8) {
    canvas.set_blend_mode(BlendMode::Blend);
    canvas.set_draw_color(Color::RGBA(0, 0, 0, alpha));
    let (w, h) = canvas.output_size().unwrap_or((1200, 675));
    for y in (0..h as i32).step_by(3) {
        canvas.draw_line((0, y), (w as i32, y)).ok();
    }
}

pub fn vignette(canvas: &mut Canvas<Window>) {
    canvas.set_blend_mode(BlendMode::Blend);
    let (w, h) = canvas.output_size().unwrap_or((1200, 675));
    for i in 0..100 {
        let alpha = (255.0 * (1.0 - i as f32 / 100.0) * 0.6) as u8;
        canvas.set_draw_color(Color::RGBA(0, 0, 0, alpha));
        canvas.draw_line((0, i), (w as i32, i)).ok();
        canvas.draw_line((0, h as i32 - 1 - i), (w as i32, h as i32 - 1 - i)).ok();
    }
}

fn render_word_shadow(
    canvas: &mut Canvas<Window>,
    texture_creator: &TextureCreator<WindowContext>,
    font: &sdl2::ttf::Font,
    word: &str,
    color: Color,
    x: i32,
    y: i32,
    shadow: bool,
) -> Option<u32> {
    if shadow {
        let shadow_color = Color::RGBA(0, 0, 0, 120);
        if let Ok(surf) = font.render(word).blended(shadow_color) {
            if let Ok(tex) = texture_creator.create_texture_from_surface(&surf) {
                let dst = Rect::new(x + 3, y + 3, surf.width(), surf.height());
                canvas.copy(&tex, None, Some(dst)).ok();
            }
        }
    }

    if let Ok(surf) = font.render(word).blended(color) {
        let w = surf.width();
        if let Ok(tex) = texture_creator.create_texture_from_surface(&surf) {
            let dst = Rect::new(x, y, w, surf.height());
            canvas.copy(&tex, None, Some(dst)).ok();
        }
        Some(w)
    } else {
        None
    }
}

pub fn render_text_wrapped(
    canvas: &mut Canvas<Window>,
    texture_creator: &TextureCreator<WindowContext>,
    ttf: &Sdl2TtfContext,
    text: &str,
    font_key: &str,
    size: u16,
    color: Color,
    pos: (i32, i32),
    max_width: i32,
) {
    let path_str = lyrics::font_path(font_key);
    let font_path = Path::new(&path_str);
    
    let font = match ttf.load_font(font_path, size) {
        Ok(f) => f,
        Err(e) => {
            eprintln!("Failed to load font {:?}: {}", font_path, e);
            return;
        }
    };

    let offsets = [0, 50, -20, 80, 20, 100, 10];
    let (base_x, base_y) = pos;
    let mut x = base_x;
    let mut y = base_y;
    let mut line_idx = 0;

    let space_w = font.size_of(" ").map(|(w, _)| w as i32).unwrap_or(10);
    let line_h = font.height() + 10;

    for word in text.split_whitespace() {
        let word_w = font.size_of(word).map(|(w, _)| w as i32).unwrap_or(50);

        if x + word_w > base_x + max_width {
            line_idx += 1;
            y += line_h;
            x = base_x + offsets[line_idx % offsets.len()];
        }
        
        let use_shadow = true;
        render_word_shadow(canvas, texture_creator, &font, word, color, x, y, use_shadow);
        x += word_w + space_w;
    }
}
