use sdl2::event::Event;
use sdl2::image::LoadTexture;
use sdl2::keyboard::Keycode;
use sdl2::pixels::Color;
use sdl2::rect::Rect;
use sdl2::render::BlendMode;
use std::path::Path;
use std::time::Duration;
use std::fs::File;
use std::io::BufReader;
use rodio::{Decoder, OutputStream, Sink};

use crate::{gfx, lyrics};

const WIDTH: u32 = 1200;
const HEIGHT: u32 = 675;
const FPS: u32 = 30;
const FADE_DURATION_MS: u32 = 500;

pub struct App;

impl App {
    pub fn run() -> Result<(), String> {
        let sdl = sdl2::init()?;
        let video = sdl.video()?;
        let _image_ctx = sdl2::image::init(sdl2::image::InitFlag::PNG)?;
        let ttf = sdl2::ttf::init().map_err(|e| e.to_string())?;
        let timer = sdl.timer()?;

        // Init Audio
        let (_stream, stream_handle) = OutputStream::try_default().map_err(|e| e.to_string())?;
        let sink = Sink::try_new(&stream_handle).map_err(|e| e.to_string())?;
        
        let file = File::open("resources/-.mp3").ok(); 
        if let Some(f) = file {
            let reader = BufReader::new(f);
            if let Ok(source) = Decoder::new(reader) {
                sink.append(source);
                sink.play();
            }
        } else {
             println!("Warning: specific audio file not found. Playing silent.");
        }

        let window = video
            .window("Phong VSTRA - Open Ver", WIDTH, HEIGHT)
            .position_centered()
            .build()
            .map_err(|e| e.to_string())?;

        let mut canvas = window
            .into_canvas()
            .accelerated()
            .build()
            .map_err(|e| e.to_string())?;

        canvas.set_blend_mode(BlendMode::Blend);
        let texture_creator = canvas.texture_creator();
        let mut event_pump = sdl.event_pump()?;
        
        let verses = lyrics::get_verses(WIDTH as i32, HEIGHT as i32);

        let img_dir = Path::new("img");
        let mut bg_textures = Vec::new();
        for i in 0..20 {
            let path = img_dir.join(format!("{}.png", i));
            match texture_creator.load_texture(&path) {
                Ok(tex) => bg_textures.push(tex),
                Err(e) => {
                    eprintln!("Warning: Failed to load {:?}: {}", path, e);
                    break; 
                }
            }
        }

        let mut current_idx: usize = 0;
        let mut displayed_text = String::new();
        let mut word_index: usize = 0;
        let mut line_done = false;
        
        let mut last_word_time = timer.ticks();
        let mut line_end_time = 0;
        let mut line_start_time = timer.ticks();
        let mut prev_idx: usize = 0;

        let mut words: Vec<&str> = verses[0].t.split_whitespace().collect();

        println!("♪ Playing: Phong - VSTRA (Sliced Assets)");

        'mainloop: loop {
            for event in event_pump.poll_iter() {
                match event {
                    Event::Quit { .. }
                    | Event::KeyDown {
                        keycode: Some(Keycode::Escape),
                        ..
                    } => break 'mainloop,
                    _ => {}
                }
            }

            let now = timer.ticks();
            let verse = &verses[current_idx];

            let fade_elapsed = now.saturating_sub(line_start_time);
            let dst = Rect::new(0, 0, WIDTH, HEIGHT);

            let curr_tex_idx = current_idx % bg_textures.len();
            let prev_tex_idx = prev_idx % bg_textures.len();

            let curr_tex = &bg_textures[curr_tex_idx];
            let prev_tex = &bg_textures[prev_tex_idx];

            if fade_elapsed < FADE_DURATION_MS && current_idx != prev_idx {
                canvas.copy(prev_tex, None, Some(dst)).ok();
                
                let alpha = ((fade_elapsed as f32 / FADE_DURATION_MS as f32) * 255.0) as u8;
                
                canvas.copy(curr_tex, None, Some(dst)).ok();
                canvas.set_draw_color(Color::RGBA(0, 0, 0, 255 - alpha));
                canvas.fill_rect(dst).ok();
            } else {
                canvas.copy(curr_tex, None, Some(dst)).ok();
            }

            if !line_done && (now - last_word_time) as u64 >= verse.wd {
                if word_index < words.len() {
                    if !displayed_text.is_empty() {
                        displayed_text.push(' ');
                    }
                    displayed_text.push_str(words[word_index]);
                    word_index += 1;
                    last_word_time = now;
                    if word_index >= words.len() {
                        line_done = true;
                        line_end_time = now;
                    }
                }
            }

            if line_done && (now - line_end_time) as u64 >= verse.ld {
                prev_idx = current_idx;
                current_idx += 1;
                if current_idx >= verses.len() {
                    break 'mainloop;
                }
                displayed_text.clear();
                words = verses[current_idx].t.split_whitespace().collect();
                word_index = 0;
                line_done = false;
                last_word_time = now;
                line_start_time = now;
            }

            gfx::render_text_wrapped(
                &mut canvas,
                &texture_creator,
                &ttf,
                &displayed_text,
                verse.f,
                verse.s,
                verse.c,
                (verse.d.0 - 300, verse.d.1 - 150),
                400,
            );

            gfx::vignette(&mut canvas);

            gfx::scanlines(&mut canvas, 40);

            canvas.present();
            std::thread::sleep(Duration::from_millis(1000 / FPS as u64));
        }

        Ok(())
    }
}
