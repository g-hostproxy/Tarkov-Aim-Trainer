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

LMB fire · RMB ADS (tap = toggle, hold = hold, like Tarkov) · Shift hold breath · R reload · Esc pause · F11 fullscreen · 1–6 quick start

## Your data

Settings, PBs and history are stored in `%APPDATA%\RaidReady\data.json`.

## Releases

Pushing a change to `VERSION` on `main` builds the exe and publishes a release automatically (`.github/workflows/release.yml`).

## Build from source

Requires Go 1.24+.

```
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui" -o RaidReady.exe .
```

The web UI lives in `web/index.html` and is embedded into the binary. `rsrc_windows_amd64.syso` holds the exe icon/manifest.
