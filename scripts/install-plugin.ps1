#!/usr/bin/pwsh

$version = '0.2.0'

$arch = if ([Environment]::Is64BitOperatingSystem) { 'amd64' } else { '386' }
if ($IsWindows) {
    $os = 'windows'
    $ext = '.exe'
    $pluginPath = Join-Path $env:APPDATA "terraform.d\plugins\$($os)_$($arch)"
    $pluginDest = Join-Path $pluginPath "terraform-provider-infoblox_v$($version)$($ext)"
}
else {
    $os = if ($IsMacOs) { 'darwin' } else { 'linux' }
    $ext = ''
    $pluginPath = Join-Path $env:HOME ".terraform.d/plugins/$($os)_$($arch)"
    $pluginDest = Join-Path $pluginPath "terraform-provider-infoblox_v$($version)$($ext)"
}
$pluginSource = Join-Path './dist' "$($os)_$($arch)_terraform-provider-infoblox_v$($version)$($ext)"

if (-not (Test-Path $pluginPath)) {
    New-Item $pluginPath -ItemType Directory | Out-Null
}
Write-Host "Copying $pluginSource to $pluginDest"
Copy-Item $pluginSource $pluginDest -Force
