# RAID//READY — Tarkov Aim Trainer

A warm-up aim trainer for Escape from Tarkov, matched to your in-game sensitivity. Runs as a single Windows exe in its own window.

## Download

Grab **RaidReady.exe** from the [latest release](https://github.com/g-hostproxy/Tarkov-Aim-Trainer/releases/latest) and double-click it. It opens in its own window using Microsoft Edge (falls back to Chrome/Brave).

The exe isn't code-signed, so on first launch Windows SmartScreen may say "Windows protected your PC". Click **More info → Run anyway**.

## Match your sens first

Click **Match my Tarkov sens** and enter your mouse DPI, mouse sensitivity, aiming sensitivity and FOV. The trainer uses raw mouse input and Tarkov's 0.113636°/count rotation, so the same swipe turns you the same distance as in raid. (If BSG changes the formula, the yaw is editable under Advanced.)

## Drills

| # | Drill | What it trains |
|---|-------|----------------|
| 1 | Head, Eyes | One-tap head flicks, Customs containers |
| 2 | Corner Peek | Jiggle- and swing-peeks from cover, shrinking window |
| 3 | ADAD Tracking | Strafing + crouch-spamming PMC, hold fire |
| 4 | Spray Control | M4A1 / AK-74N / AKM / MP7 recoil, grouping plot after |
| 5 | CQB Switch | Three close targets; head = 1 tap, thorax = 2 |
| 6 | Long Shot | 2x–6x optic, sway, hold breath, walkers at 80–200 m |
| — | Raid Warm-Up | Drills 1–4 back to back, 45 s each (~3 min). Start it when you hit "Ready" on the raid loading screen. |

Ranks: TIMMY › SCAV › PMC › RAIDER › ROGUE › BOSS

## Controls

LMB fire · RMB ADS (tap = toggle, hold = hold, like Tarkov) · WASD walk · scroll wheel walk speed · Shift hold breath · R reload · Esc pause · F11 fullscreen · 1–6 quick start

## Walking and barrel movement

Walking can be switched on or off on the screen before each drill. With it on, WASD moves you (slower backwards, sideways and while aiming) and the scroll wheel sets your walk speed, like in raid.

The barrel moves on its own, separate from your view, using the sway model from [TarkovCrosshair](https://github.com/g-hostproxy/TarkovCrosshair) (calibrated against an in-game laser). Keys reach the weapon about 235 ms late. Holding A swings the barrel right and holding D swings it left, and it stays there until you let go. Walking forward dips it, walking back lifts it, and it bobs in a figure-8 at Tarkov's 1.2 Hz step rate. Fast mouse turns make it lag, and every shot kicks it up; on full auto the climb levels off and recovers once you stop. The crosshair, red dot and scope all ride on the barrel, so they always show where the shot will land. Under Settings you can set barrel movement to Tarkov, Reduced or Off, and turn the weapon model off.

## Your data

Settings, PBs and history are stored in `%APPDATA%\RaidReady\data.json`.

## Support the project

RAID//READY is free. If it's helped your raids, you can buy me a coffee:

[![Buy me a coffee](https://img.shields.io/badge/Buy%20me%20a%20coffee-FFDD00?style=for-the-badge&logo=buymeacoffee&logoColor=black)](https://buymeacoffee.com/g.hostproxy)

The **Support** button in the app opens this section.

## Releases

Pushing a change to `VERSION` on `main` builds the exe and publishes a release automatically (`.github/workflows/release.yml`).

## Build from source

Requires Go 1.24+.

```
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui" -o RaidReady.exe .
```

The web UI lives in `web/index.html` and is embedded into the binary. `rsrc_windows_amd64.syso` holds the exe icon/manifest.
