
# Media Remote Control App

## What

Simple Server-Client App that allows you to use media control (or keys) remotelly using mobile application (in future web page)

## Why
- Stop media player of desktop from longer than bluetooth range
- Using Desktop as TV more pleasant
- etc.

## How

### Client

Noting special just UI that send api calls.
#### Implmented

- In Kotlin multiplatform

### Server

- Volume 
Checks using volume-go #expanation

- Media Keys

Windows/Linux: Simulation of keyboard presses
Mac: Simulates AUX keys (like in wired headphone)

#### Tested/Developed on

MacBook Pro 2019 (touchbar)

#### Possible to explore

- Media keys on Mac (F8)
Doesn't worked assume because of Touchbar

- Using bluetooth canal
Using AVRCP (if used probably could be used for excange server IP?)

- Pushes for syncing IP?

- Static Domain for IP

- Fully Remote Server

- "Go to sleep mode" command

- Rewrite server logic in C
Will look fun

Android (Kotlin) \                         [Desktop as Server]           ?
                  |->  Mobile logic (KMP) - > Server (GO) -> Logic (C)
iOS     (Swift)  /                     [OR]      ^              \-> macOS (swift) 
                                        \        |
                       Front (htmx)  - - \       |
                                          \      |
                       Arduino (C) - - - - \     |
                                            v    |
         (FULL REMOTE) Desktop (GO) <= = =>  Server (GO)

# Future

- [ ] Info about currently playing in Desktop
    - [ ] Transfer currently playing to mobile (cutomizable actions where to lead e.g. Spotify/SoundCloud/Etc)
- [ ] Seeing what is playing on mobile from Desktop
    - [ ] Send currently playing to the desktop
- [ ] Some animation (loading) shared between client and server (rendered differently)
    - [ ] OpenGL in mobile
    - [ ] ACII   in Server CLI / Build CLI
    - [ ] WebGPU in Web