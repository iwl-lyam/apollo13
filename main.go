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

	fmt.Print("\n")
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

func scene2() bool {
	time.Sleep(2 * time.Second)
	write("T+00:00:05 "+color.Ize(color.Green, "Commander: Yaw program."))
	write("T+00:00:12 "+color.Ize(color.Green, "CM Pilot: Clear the tower."))
	write("T+00:00:14 "+color.Ize(color.Green, "Commander: Yaw complete. Roll program."))
	write("T+00:00:30 "+color.Ize(color.Yellow, "CAPCOM: 13, Houston. GO at 30 seconds."))
	write("T+00:00:34 "+color.Ize(color.Green, "Commander: Roll complete, and we are pitching."))
	wait()
	write("T+00:02:02 "+color.Ize(color.Yellow, "CAPCOM: GO for staging."))
	write("T+00:02:16 "+color.Ize(color.Green, "Commander: Inboard."))
	write("T+00:02:27 "+color.Ize(color.Yellow, "CAPCOM: We confirm Inboard out, 13. You're looking good."))
	wait()
	write("T+00:02:48 "+color.Ize(color.Green, "Commander: S-II ignition"))
	write("T+00:03:00 "+color.Ize(color.Yellow, "CAPCOM: 13, Houston. Trajectory is good; thrust is good."))
	wait()
	write("T+00:05:36 "+color.Ize(color.Green, "Commander: Inboard."))
	write("T+00:05:44 "+color.Ize(color.Yellow, "CAPCOM: Roger. We confirm Inboard out."))
	wait()
	write("T+00:06:10 "+color.Ize(color.Green, "Commander: And, Houston, What's the story on engine 5?"))
	write("T+00:06:14 "+color.Ize(color.Yellow, "CAPCOM: Jim, Houston. We don't have the story on why the inboard out was early, but the other engines are GO and you are GO."))
	wait()
	write("T+00:09:19 "+color.Ize(color.Yellow, "CAPCOM: 13, Houston. You are GO for staging."))
	write("T+00:09:48 "+color.Ize(color.Yellow, "CAPCOM: MARK"))
	write("T+00:09:50 "+color.Ize(color.Green, "Commander: Roger. Staging."))
	write("T+00:10:00 "+color.Ize(color.Green, "Commander: And S-IV ignition, Houston."))
	write("T+00:10:17 "+color.Ize(color.Yellow, "CAPCOM: 13, Houston. You're looking good. Trajectory, guidance, CMC are all GO."))
	write("T+00:11:35 "+color.Ize(color.Yellow, "CAPCOM: Apollo 13, Houston. You're GO at 11-1/2, and predicted cut-off time is 12 plus 34. Over."))
	wait()
	write("T+00:12:31 "+color.Ize(color.Green, "Commander: SECO."))
	write("T+00:12:32 "+color.Ize(color.Yellow, "CAPCOM: Copy, SECO."))
	wait()
	wait()
	write("T+00:21:44 "+color.Ize(color.Yellow, "CAPCOM: Okay. Couple minutes to LOS, Jim. Everything is looking real good. Your AOS time at Carnarvon will be 52:36, and we don't have too much of a handle on why the inboard cut off early except that it apparently was an engine problem and not a switch-select function. But we are certain that you'll be able to make TLI based on what we are looking at now."))
	wait()

	fmt.Print("\n")

	fmt.Print("\033[H\033[2J")
	wait()
	write("One hour later...")
	wait()
	fmt.Print("\033[H\033[2J")
	wait()
	write("T+01:48:16 "+color.Ize(color.Yellow, "CAPCOM: And you are GO for TLI. Huntsville reperts that you have a 6-second propellant pad which is 3 seconds more than a 3-sigma case; so you're good on consumabies . The IU is so good that we're not going to update it. The only change we have for you is in the TLI checkiist. At 57 minutes where you slew the FDAI to 18 degrees, we recommend 20 degrees there, and we recommend that you look for 8 degrees instead of 6 degrees at ignition. The S-IVB is riding on the top of its deadband."))

	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")
}

func scene3() {
	time.Sleep(2 * time.Second)
	write(color.Ize(color.Blue, "4 hours later, all is going well, until...")) 
	wait()
	write("T+05:52:32 "+color.Ize(color.Red, "Master caution and warning triggered by low hydrogen pressure in tank #1."))
	write("T+05:53:22 "+color.Ize(color.Gray, "11.1 amp “spike” recorded in fuel cell #3"))
	write("T+05:53:22 "+color.Ize(color.Gray, "22.9-amp “spike” recorded in fuel cell #3"))
	write("T+05:54:51 "+color.Ize(color.Gray, "Oxygen tank #2 quantity jumped to off-scale high and then started to drop until the time of telemetry loss, indicating a failed sensor."))
	write("T+05:54:53 "+color.Ize(color.Red, "A loud bang was heard"))
	write("T+05:55:20 "+color.Ize(color.Green, "Commander: I believe we've had a problem here."))
	write("T+05:55:28 "+color.Ize(color.Yellow, "CAPCOM: This is Houston. Say again, please."))
	write("T+05:55:35 "+color.Ize(color.Green, "Commander: Houston, we've had a problem. We've had a MAIN B BUS UNDERVOLT."))
	write("T+05:55:58 "+color.Ize(color.Yellow, "CAPCOM: Okay, stand by, 13. We're looking at it."))
	write("T+05:56:10 "+color.Ize(color.Green, "LM pilot: Okay. Right now, Houston, the voltage is - is looking good. And we had a pretty large bang associated with the CAUTION AND WARNING there. And as I recall, MAIN B was the one that had had an amp spike on it once before."))
	wait()
	write("T+05:58:25 "+color.Ize(color.Green, "LM pilot: We got a MAIN BUS A UNDERVOLT now, too, showing."))

}

func write(text string) {
	print := strings.Split(text, "")
	for i, s := range print {
		fmt.Print(s)
		i=i
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
	fmt.Print("\n")
}

func wait() {
	time.Sleep(time.Duration(1000) * time.Millisecond)
}

func main() {
	// start()
	// fmt.Println("Chapter 1: Liftoff")
	// scene1()
	// fmt.Println("Chapter 2: Problems arising")
	// scene2()
	fmt.Println("Chapter 3: Houston, we've had a problem")
	scene3()
}
