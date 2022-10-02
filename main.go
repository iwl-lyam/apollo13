package main

import (
	"fmt"
	"strings"
	"time"
	"github.com/TwiN/go-color"
)

func start() {
	fmt.Println(" ________  ________  ________  ___       ___       ________           _____  ________     ")
	fmt.Println(`|\   __  \|\   __  \|\   __  \|\  \     |\  \     |\   __  \         / __  \|\_____  \    `)
	fmt.Println(`\ \  \|\  \ \  \|\  \ \  \|\  \ \  \    \ \  \    \ \  \|\  \       |\/_|\  \|____|\ /_   `)
	fmt.Println(` \ \   __  \ \   ____\ \  \\\  \ \  \    \ \  \    \ \  \\\  \      \|/ \ \  \    \|\  \  `)
	fmt.Println(`  \ \  \ \  \ \  \___|\ \  \\\  \ \  \____\ \  \____\ \  \\\  \          \ \  \  __\_\  \ `)
	fmt.Println(`   \ \__\ \__\ \__\    \ \_______\ \_______\ \_______\ \_______\          \ \__\|\_______\`)
	fmt.Println(`    \|__|\|__|\|__|     \|_______|\|_______|\|_______|\|_______|           \|__|\|_______| - the game`)
	write(color.Ize(color.Blue, "\n--- Press enter to start ---"))
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")
}

func scene1() {
	time.Sleep(2 * time.Second)
	write(color.Ize(color.Blue, "Date: 10th April 1970"))
	write(color.Ize(color.Blue, "T-28:00:00"))
	write(color.Ize(color.Blue, "On the pad"))
	fmt.Print("\n")
	write("Go for terminal count.")
	write("Copy, terminal count.")
	time.Sleep(2 * time.Second)

	time.Sleep(2 * time.Second)
	write(color.Ize(color.Blue, "Date: 11th April 1970"))
	write(color.Ize(color.Blue, "T-10:00"))
	fmt.Print("\n")
	write("Cleared for launch")
	wait()

	time.Sleep(2 * time.Second)
	fmt.Print("\n")
	write("T-10")
	write("T-9")
	write("Main engine start")
	write("T-6")
	write("T-5")
	write("T-4")
	write("T-3")
	write("T-2")
	write("T-1")
	write("Liftoff, we have a liftoff!")
	wait()
	fmt.Print("\033[H\033[2J")
}

func scene2() {
	time.Sleep(2 * time.Second)
	write("T+12 Roll program")
	write("T+1:21 Max-Q")
	write("T+2:15 Center engine shutdown")
	write("T+2:43 Outer engine shutdown")
	write("T+2:44 Stage separation")
	fmt.Print("\n")

	write("T+2:46 S-II Ignition")
	write("T+3:14 Interstage jettison")
	wait()
	write(color.Ize(color.Red, "T+5:30 Inboard engine cutoff"))
	write(color.Ize(color.Yellow, "Pilot (Swigert): Inboard."))
	write(color.Ize(color.Green, "CAPCOM (Joe Kerwin): We confirm Inboard out."))
	write(color.Ize(color.Yellow, "Commander (Lovell): That shouldn't have happened."))
	write(color.Ize(color.Yellow, "Pilot: No, that's 5:30. That's two minutes early."))
	wait()
	write("T+7:42 Inboard engine cutoff sent by flight computer")
	write("T+9:52 Outboard engine cutoff (34.53 seconds late)")
	write("T+9:53 Stage separation")
	fmt.Print("\n")

	write("T+9:56 S-IVB ignition")
	write("T+12:29 S-IVB cutoff (9 seconds late)")
	write("Orbital insertion.")
	fmt.Print("\n")

	wait()
	write("T+2:35:46 S-IVB ignition")
	write("T+2:41:37 S-IVB shutdown")
	write("Translunar injection.")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")
}

func write(text string) {
	print := strings.Split(text, "")
	for i, s := range print {
		fmt.Print(s)
		i=i
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
	fmt.Scanln()
}

func wait() {
	time.Sleep(time.Duration(1000) * time.Millisecond)
}

func main() {
	start()
	fmt.Println("Chapter 1: Liftoff")
	scene1()
	fmt.Println("Chapter 2: Flight")
	scene2()
	fmt.Println("Chapter 3: Problems arising")
}
