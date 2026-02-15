use std::path::PathBuf;

fn main() {
    let manifest_dir = PathBuf::from(std::env::var("CARGO_MANIFEST_DIR").unwrap());
    
    // Check for auto-downloaded deps in ../deps
    let deps_root = manifest_dir.parent().unwrap().join("deps");
    
    // Fallback to legacy local folder if deps not found
    let sdl_root = if deps_root.exists() {
        deps_root
    } else {
        manifest_dir.join("SDL2-devel")
    };

    if deps_root.exists() {
        // New structure (from setup.ps1)
        println!("cargo:rustc-link-search=native={}", sdl_root.join("SDL2").join("lib").join("x64").display());
        println!("cargo:rustc-link-search=native={}", sdl_root.join("SDL2_image").join("lib").join("x64").display());
        println!("cargo:rustc-link-search=native={}", sdl_root.join("SDL2_ttf").join("lib").join("x64").display());
    } else {
        // Legacy structure
        println!("cargo:rustc-link-search=native={}", sdl_root.join("SDL2-2.30.12").join("lib").join("x64").display());
        println!("cargo:rustc-link-search=native={}", sdl_root.join("SDL2_image-2.8.4").join("lib").join("x64").display());
        println!("cargo:rustc-link-search=native={}", sdl_root.join("SDL2_ttf-2.22.0").join("lib").join("x64").display());
    }
}
