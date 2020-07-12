### boxclock

High intensity interval training workout clock

## Parts
- Raspberry Pi 
- Monitor (optional: HDMI-CEC support with remote)
- Buttons (Red, Green, Yellow, Blue, White)

## Dependencies
- Go (https://golang.org/project/)
- go-rpio (https://github.com/stianeikeland/go-rpio)
- OpenVG (https://github.com/ajstarks/openvg)
- testing (https://golang.org/pkg/testing/)
- cec (https://github.com/chbmuc/cec)

## Design Principles
- no modes (https://en.wikipedia.org/wiki/Larry_Tesler)
- high data-ink ratio (https://infovis-wiki.net/wiki/Data-Ink_Ratio)
- large readable typography
- white button for allstart,allstop
- colors for individual rep counting and stops

## Core functionality
- Wallclock
- Stopwatch (reset/start/multi-user-stop/timecap)
- Countdown timer (set/reset/start/end)
- Intervals (set work time,rest time,# cycles)
- Rep Counter (optional:rep target/each button has its own timer)

## Feature composition
- For time (stopwatch + timecap)
- For reps (wallclock + rep counter)
- AMRAP (Countdown timer: set=timecap, + rep counter)
- EMOM (Intervals: 1m-work,0m rest,n cycles)
- Tabata (Intervals: 20s-work,10sec rest,8 cycles)


<!--
**boxclock/boxclock** is a ✨ _special_ ✨ repository because its `README.md` (this file) appears on your GitHub profile.

Here are some ideas to get you started:

- 🔭 I’m currently working on ...
- 🌱 I’m currently learning ...
- 👯 I’m looking to collaborate on ...
- 🤔 I’m looking for help with ...
- 💬 Ask me about ...
- 📫 How to reach me: ...
- 😄 Pronouns: ...
- ⚡ Fun fact: ...
-->
