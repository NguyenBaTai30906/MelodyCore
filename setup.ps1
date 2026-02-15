$ErrorActionPreference = "Stop"

# Detect if running from script file or web content
$scriptPath = $MyInvocation.MyCommand.Path
$BaseDir = Get-Location

if ($scriptPath) {
    # Running via File -> Use script directory
    $BaseDir = Split-Path -Parent $scriptPath
}

# Check if we are in the project root (look for phong-rust)
$phongCheck = Join-Path $BaseDir "phong-rust"

if (-not (Test-Path $phongCheck)) {
    Write-Host "Project not found. Bootstrapping MelodyCore..." -ForegroundColor Cyan
    
    # Try cloning into current directory
    try {
        git clone https://github.com/NguyenBaTai30906/MelodyCore.git .
        if ($LASTEXITCODE -ne 0) {
            # If failed (non-empty dir), clone into subdir
            Write-Host "Directory not empty, cloning into 'MelodyCore'..." -ForegroundColor Yellow
            git clone https://github.com/NguyenBaTai30906/MelodyCore.git
            $BaseDir = Join-Path $BaseDir "MelodyCore"
        }
    }
    catch {
        Write-Error "Git clone failed. Ensure Git is installed."
    }
}

$depsDir = Join-Path $BaseDir "deps"
$phongDir = Join-Path $BaseDir "phong-rust"

# URLs for SDL2 VC development libraries
$sdl2_url = "https://github.com/libsdl-org/SDL/releases/download/release-2.30.1/SDL2-devel-2.30.1-VC.zip"
$sdl2_image_url = "https://github.com/libsdl-org/SDL_image/releases/download/release-2.8.2/SDL2_image-devel-2.8.2-VC.zip"
$sdl2_ttf_url = "https://github.com/libsdl-org/SDL_ttf/releases/download/release-2.22.0/SDL2_ttf-devel-2.22.0-VC.zip"

function Download-And-Extract ($url, $dest) {
    if (-not (Test-Path $dest)) {
        New-Item -ItemType Directory -Force -Path $dest | Out-Null
    }
    
    $filename = ([System.IO.Path]::GetFileName($url))
    $zipPath = Join-Path $env:TEMP $filename
    
    Write-Host "Downloading $filename..." -ForegroundColor Cyan
    Invoke-WebRequest -Uri $url -OutFile $zipPath
    
    Write-Host "Extracting to $dest..." -ForegroundColor Cyan
    Expand-Archive -Path $zipPath -DestinationPath $dest -Force
    
    # Move inner folder content to $dest if it's nested (standard SDL zip structure)
    $subDir = Get-ChildItem -Path $dest -Directory | Select-Object -First 1
    if ($subDir) {
        Get-ChildItem -Path $subDir.FullName | Move-Item -Destination $dest -Force
        Remove-Item -Path $subDir.FullName -Recurse -Force
    }
    
    Remove-Item $zipPath -Force
}

Write-Host "=== MelodyCore Setup ===" -ForegroundColor Green

# 1. Create deps directory
if (-not (Test-Path $depsDir)) {
    New-Item -ItemType Directory -Force -Path $depsDir | Out-Null
}

# 2. Download Libs
Download-And-Extract $sdl2_url (Join-Path $depsDir "SDL2")
Download-And-Extract $sdl2_image_url (Join-Path $depsDir "SDL2_image")
Download-And-Extract $sdl2_ttf_url (Join-Path $depsDir "SDL2_ttf")

# 3. Copy DLLs to Project
Write-Host "Copying DLLs..." -ForegroundColor Cyan
$dlls = @(
    (Join-Path $depsDir "SDL2\lib\x64\SDL2.dll"),
    (Join-Path $depsDir "SDL2_image\lib\x64\SDL2_image.dll"),
    (Join-Path $depsDir "SDL2_ttf\lib\x64\SDL2_ttf.dll")
)

# Target dirs
$targetDirs = @(
    $phongDir,
    (Join-Path $phongDir "target\debug"),
    (Join-Path $phongDir "target\release")
)

foreach ($dir in $targetDirs) {
    if (-not (Test-Path $dir)) {
        New-Item -ItemType Directory -Force -Path $dir | Out-Null
    }
    foreach ($dll in $dlls) {
        Copy-Item -Path $dll -Destination $dir -Force
    }
}

Write-Host "Setup Complete! You can now run 'cargo run' in phong-rust." -ForegroundColor Green
Read-Host "Press Enter to exit"
