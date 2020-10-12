/*

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

*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ajstarks/openvg"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

// globals
var cmd chan string
var start time.Time
var elapsed_prior time.Duration
var state string
var width, height int
var w, w2, h2, hb openvg.VGfloat
var sw_text string
var countdown_index int

func clock() {

	// initialize OpenGL and prepare width and height
	width, height = openvg.Init()
	w2 = openvg.VGfloat(width / 2)
	h2 = openvg.VGfloat(height / 2)
	w = openvg.VGfloat(width)
	hb = openvg.VGfloat(height / 6)

	// load local time zone location
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}

	// begin at a stopped state, initialize local color settings
	state = "stopped"
	sw_text = "00:00:00"
	bg_color := "black"
	sw_fg_color := "gray"
	wc_fg_color := "white"

	for {
		// get wall clock text string in a 12 hour format
		// for the current time zone
		t := time.Now().In(loc)
		wc_text := t.Format("3:04:05 PM")

		// prepare display colors and strings based on current state:
		// stopped, countdown, or running
		if state == "stopped" {
			bg_color = "black"
			sw_fg_color = "gray"
			wc_fg_color = "white"
			countdown_index = 10
		} else if state == "countdown" {
			sw_fg_color = "white"
			wc_fg_color = "gray"
			// 10,9,8,7,6,5,4,3,2,1,Go!
			if countdown_index > 0 {
				sw_text = fmt.Sprintf("%d\n", countdown_index)
				countdown_index -= 1
			} else {
				// start clock, enter running state, reset countdown
				start = time.Now().In(loc)
				state = "running"
				sw_text = "Go!"
				bg_color = "green"
				countdown_index = 10
			}
		} else if state == "running" {
			// get elapsed hours, minutes, seconds since start
			elapsed := t.Sub(start)
			h := int(elapsed.Hours())
			m := int(elapsed.Minutes()) % 60
			s := int(elapsed.Seconds()) % 60
			// get stop watch text string with elapsed time
			sw_text = fmt.Sprintf("%02d:%02d:%02d\n", h, m, s)

			// set display colors
			bg_color = "black"
			sw_fg_color = "white"
			wc_fg_color = "gray"
		}

		// update display
		openvg.Start(width, height)
		openvg.BackgroundColor(bg_color)
		openvg.FillColor(sw_fg_color)
		openvg.TextMid(w2, h2, sw_text, "serif", width/10)
		openvg.FillColor(wc_fg_color)
		openvg.TextMid(w2, hb, wc_text, "serif", width/20)
		openvg.End()

		// wait one second before updating display again
		time.Sleep(time.Second)
	}
}

func button(gpio_name string) {

	// get the specified GPIO pin
	p := gpioreg.ByName(gpio_name)
	if p == nil {
		log.Fatal("Failed to find " + gpio_name)
	}

	// set GPIO pin as input
	err := p.In(gpio.PullUp, gpio.FallingEdge)
	if err != nil {
		log.Fatal(err)
	}

	// for debouncing, track last button press time
	lastButtonTime := time.Now()

	// wait for falling edge, and act on the button press
	for {
		p.WaitForEdge(-1)

		// quick and dirty button debounce logic ...
		// ignore edges within 500ms of last button time
		currentButtonTime := time.Now()
		if currentButtonTime.After(lastButtonTime.Add(500 * time.Millisecond)) {
			lastButtonTime = currentButtonTime
			if state == "running" {
				// button pushed while running -> stop running clock
				state = "stopped"
				cmd <- "stop"
			} else if state == "countdown" {
				// button pushed during countdown -> restart countdown
				countdown_index = 10
				cmd <- "restart_countdown"
			} else if state == "stopped" {
				// button pushed while stopped -> start countdown
				state = "countdown"
				cmd <- "countdown"
			}
		}
	}
}

func main() {

	// load periph.io drivers for button input
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	go clock()

	// handle any of the five buttons equally at this point
	go button("GPIO2")
	go button("GPIO3")
	go button("GPIO4")
	go button("GPIO5")
	go button("GPIO6")

	cmd = make(chan string, 1)

	for {
		select {
		case res := <-cmd:
			fmt.Println(res)
			if res == "quit" {
				return
			}
		}
	}

}
