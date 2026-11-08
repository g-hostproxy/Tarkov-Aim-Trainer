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

Defaults match Tarkov, and every key can be rebound under **Settings → Key binds** (keyboard keys or mouse buttons).

| Action | Default |
|---|---|
| Fire / aim down sights | LMB / RMB (tap = toggle, hold = hold) |
| Move | WASD (scroll wheel = walk speed) |
| Crouch / prone | C / X |
| Lean left / right | Q / E (hold, or toggle in settings) |
| Slow lean | Hold Alt + Q/E or A/D to lean gradually; it stays until you tap Q/E |
| Fire mode | B (semi / full auto) |
| Scope zoom | Hold Alt + scroll |
| Hold breath | Shift |
| Reload | R |
| Pause / fullscreen | Esc / F11 |

## How it plays

- **Hip fire has no crosshair**, like Tarkov. Aim down sights to get a reticle; it sits on the barrel, so it always shows where the shot lands.
- **Scopes are picture-in-picture**: the world outside the eyepiece stays at normal zoom and only the view inside is magnified, with the rifle and mount visible underneath. Alt + scroll zooms the scope smoothly.
- **Barrel sway** uses the model from [TarkovCrosshair](https://github.com/g-hostproxy/TarkovCrosshair), calibrated against an in-game laser: keys reach the weapon about 235 ms late, holding A swings the barrel right and D swings it left, walking forward dips it and walking back lifts it, plus a heavy figure-8 walk bob that carries into the scope.
- **Stances**: crouching and going prone lower your eye height, slow you down and steady the gun. You can't shoot while getting down or up from prone, or while crawling.
- **Bullet holes** stay where your misses land (up to 45 at a time, oldest removed first).
- **Leaning** rolls the view and moves your head about 40 cm out from cover.
- **Full auto** climbs hard for the first few rounds, levels off, wobbles side to side, and pulls back down once you let go.

Walking can be switched on or off on the screen before each drill. Recoil, sway, stance and lean values are tuned to feel like Tarkov; they are not pulled from the game files.

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
