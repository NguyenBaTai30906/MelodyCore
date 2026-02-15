use sdl2::pixels::Color;

pub struct V {
    pub t: &'static str,
    pub wd: u64,
    pub ld: u64,
    pub c: Color,
    pub d: (i32, i32),
    pub f: &'static str,
    pub s: u16,
}
pub fn font_path(key: &str) -> String {
    let name = match key {
        "b" => "Montserrat-Bold.ttf",
        "sc" => "DancingScript-Bold.ttf",
        "se" => "PlayfairDisplay-Bold.ttf",
        "p" => "Pacifico-Regular.ttf",
        "r" => "Quicksand-Bold.ttf",
        _ => "NotoSans-Bold.ttf",
    };
    format!("resources/fonts/{}", name)
}

pub fn get_verses(w: i32, h: i32) -> Vec<V> {
    let wht = Color::RGB(255, 255, 255);
    let crm = Color::RGB(255, 245, 220);
    let cyn = Color::RGB(200, 240, 255);
    let pnk = Color::RGB(255, 200, 210);
    
    let cx = w / 2;
    let cy = h / 2;

    vec![
        V { t: "Oh baby", wd: 120, ld: 500, c: wht, d: (cx, cy), f: "b", s: 40 },
        V { t: "Nàng cứ hát đi", wd: 120, ld: 500, c: wht, d: (cx, cy), f: "b", s: 40 },
        V { t: "Anh nghe thấy rồi!", wd: 120, ld: 600, c: wht, d: (cx, cy), f: "b", s: 40 },
        V { t: "Anh nghe hết rồi", wd: 180, ld: 1100, c: wht, d: (cx, cy), f: "b", s: 40 }, 
        V { t: "...", wd: 0, ld: 200, c: wht, d: (cx, cy), f: "b", s: 24 }, 
        
        V { t: "Lặng nghe câu hát dịu dàng", wd: 180, ld: 800, c: pnk, d: (cx + 100, cy + 50), f: "sc", s: 55 },
        V { t: "Em đang hát nhẹ nhàng, em muốn quay lưng", wd: 150, ld: 1000, c: crm, d: (cx - 150, cy - 50), f: "r", s: 45 },
        V { t: "Để nhìn thấy vùng trời", wd: 200, ld: 700, c: cyn, d: (cx - 100, cy + 100), f: "se", s: 50 },
        V { t: "Yêu kiều nhất trần đời", wd: 220, ld: 800, c: wht, d: (cx + 50, cy - 100), f: "sc", s: 58 },
        V { t: "Tình treo trên nóc trần nhà", wd: 180, ld: 700, c: wht, d: (cx - 200, cy), f: "b", s: 48 },
        
        V { t: "Nghe không giống thật thà", wd: 180, ld: 800, c: pnk, d: (cx + 150, cy + 80), f: "sc", s: 52 },
        V { t: "Em nói ra đi, em ấy chỉ vì", wd: 160, ld: 600, c: crm, d: (cx - 250, cy - 80), f: "p", s: 46 },
        V { t: "Nghe lý trí thầm thì", wd: 200, ld: 3000, c: cyn, d: (cx, cy + 120), f: "r", s: 50 }, 
    ]
}
