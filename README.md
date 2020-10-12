<!-- 

MIT License

Copyright (c) 2020 boxclock

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

-->
### boxclock

High intensity interval training workout clock

## Note
Current features include a stopwatch function and a persistent 12 hour wall clock. 
Everything else is still a work in progress.

## Materials
- Raspberry Pi 
- Monitor (optional: HDMI-CEC support with remote)
- Buttons (Red, Green, Yellow, Blue, White)

## Dependencies
- Go (https://golang.org/project/)
- OpenVG (https://github.com/ajstarks/openvg)
- periph.io (https://periph.io/)
- testing (https://golang.org/pkg/testing/)
- (pending) cec (https://github.com/chbmuc/cec) 

## Design Principles
- no modes (https://en.wikipedia.org/wiki/Larry_Tesler)
- high data-ink ratio (https://infovis-wiki.net/wiki/Data-Ink_Ratio)
- large readable typography
- use colors for rep counting and set timing in partner workouts

## Core functions
- Wallclock (implemented)
- Stopwatch (partially implemented - reset/start/multi-user-stop/timecap)
- Countdown timer (not implemented - set/reset/start/end)
- Intervals (not implemented - set work time,rest time,# cycles)
- Rep Counter (not implemented - count to rep target/each button tied to its own counter and timer)

## Features to compose from core functions
- For time: stopwatch + timecap
- For reps: wallclock + rep counter
- As many rounds as possible (AMRAP):  countdown timer + rep counter
- Every minute on the minute (EMOM): intervals set at 1m work, 0m rest, n cycles
- Tabata: intervals set at 20s work, 10s rest, 8 cycles