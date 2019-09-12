package run

import (
	"fmt"
	"time"
)

var Version string

func printHeader() {
	// fmt.Println("Time      Module  Pipeline  Description")
	// fmt.Println("--------------------------------------------------------------------------------")
}

func printMsg(s *Simulation, msg string) {
	if msg == "" {
		return
	}
	// fmt.Printf("%[1]s  %-[3]*.[4]*[2]s %-8[5]s  %[6]s\n",
	// 	time.Now().Format("15:04:05"), s.Name(), 8, len(, s.Pipeline, msg)
	fmt.Printf("%s [%s] %s\n", time.Now().Format("15:04:05"), s.Name(), msg)
}

func printWelcome(ip string) {
	fmt.Printf(`
AlphaSOC Network Flight Simulator™ %s (https://github.com/alphasoc/flightsim)
The IP address of the network interface is %s
The current time is %s

`, Version, ip, time.Now().Format("02-Jan-06 15:04:05"))
}

func printGoodbye() {
	fmt.Printf("\nAll done! Check your SIEM for alerts using the timestamps and details above.\n")
}
