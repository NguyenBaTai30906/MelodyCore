$ErrorActionPreference = "Stop"
Write-Host "DEBUG: v2.1 (Fix Change-Dir)" -ForegroundColor Magenta

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
} else {
    Write-Host "Project found. Updating source code..." -ForegroundColor Cyan
    git pull
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

# ... (Bootstrap logic remains above)

function Setup-PhongRust {
    Write-Host "`n=== Setting up Phong VSTRA (Rust) ===" -ForegroundColor Cyan
    
    # 2. Download Libs
    Download-And-Extract $sdl2_url (Join-Path $depsDir "SDL2")
    Download-And-Extract $sdl2_image_url (Join-Path $depsDir "SDL2_image")
    Download-And-Extract $sdl2_ttf_url (Join-Path $depsDir "SDL2_ttf")

    # 3. Copy DLLs to Project
    Write-Host "Copying DLLs to phong-rust..." -ForegroundColor Cyan
    $dlls = @(
        (Join-Path $depsDir "SDL2\lib\x64\SDL2.dll"),
        (Join-Path $depsDir "SDL2_image\lib\x64\SDL2_image.dll"),
        (Join-Path $depsDir "SDL2_ttf\lib\x64\SDL2_ttf.dll")
    )

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
    Write-Host "Phong VSTRA setup complete!" -ForegroundColor Green
    Write-Host "Building and Running (Release Mode)..." -ForegroundColor Cyan
    
    Set-Location $phongDir
    # Clean debug artifacts if they exist to save space
    if (Test-Path "target\debug") { Remove-Item -Recurse -Force "target\debug" }
    
    Write-Host "Run 'cargo run' inside 'phong-rust' folder to start." -ForegroundColor Yellow
}

function Install-Go {
    Write-Host "Go compiler not found. Automatically installing Go..." -ForegroundColor Yellow
    
    $goMsiUrl = "https://go.dev/dl/go1.22.4.windows-amd64.msi"
    $goMsiPath = Join-Path $env:TEMP "go_installer.msi"
    
    Write-Host "Downloading Go installer (MSI)..." -ForegroundColor Cyan
    Invoke-WebRequest -Uri $goMsiUrl -OutFile $goMsiPath
    
    Write-Host "Installing Go silently (requires Admin privileges)..." -ForegroundColor Cyan
    $process = Start-Process msiexec.exe -ArgumentList "/i `"$goMsiPath`" /quiet /qn /norestart" -Wait -PassThru
    
    if ($process.ExitCode -eq 0) {
        Write-Host "Go installed successfully!" -ForegroundColor Green
        # Refresh session PATH
        $env:PATH += ";C:\Program Files\Go\bin"
    } else {
        Write-Host "Go installation failed with ExitCode: $($process.ExitCode)" -ForegroundColor Red
        Write-Host "Please install Go manually from https://go.dev/dl/"
    }
    
    Remove-Item $goMsiPath -Force
}

function Setup-AiDuaEmVeGo {
    Write-Host "`n=== Setting up Ai dua em ve (Go) ===" -ForegroundColor Cyan
    
    # Check if go is available in session
    if (-not (Get-Command "go" -ErrorAction SilentlyContinue)) {
        # Check standard installation path
        if (Test-Path "C:\Program Files\Go\bin\go.exe") {
             $env:PATH += ";C:\Program Files\Go\bin"
        } else {
            # Not found anywhere, install it
            Install-Go
        }
    }
    
    $goPath = Get-Command "go" -ErrorAction SilentlyContinue
    if (-not $goPath) {
        Write-Error "Go compiler still not found! Setup aborted."
        return
    }

    $goDir = Join-Path $BaseDir "aiduaemve-go"
    if (-not (Test-Path $goDir)) {
        Write-Error "Project directory 'aiduaemve-go' not found!"
        return
    }
    
    Set-Location $goDir
    Write-Host "Installing dependencies..." -ForegroundColor Cyan
    & go mod tidy
    
    Write-Host "Running Ai dua em ve..." -ForegroundColor Green
    & go run .
}

# === Main Menu ===
Clear-Host
Write-Host "=== MelodyCore Installer ===" -ForegroundColor Magenta
Write-Host "1. Phong VSTRA (Rust) - Lyric Video"
Write-Host "2. Ai dua em ve (Go) - *Ebiten Visualizer*"
Write-Host "3. Exit"

$choice = Read-Host "Select a project to setup [1-3]"

switch ($choice) {
    "1" { Setup-PhongRust }
    "2" { Setup-AiDuaEmVeGo }
    default { Write-Host "Exiting." }
}

Read-Host "Press Enter to exit"
