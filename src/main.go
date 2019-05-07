package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[14], 100 * time.Millisecond)
	s.Color("bgBlack", "bold", "fgRed")
	s.FinalMSG = "Complete!\nNew line!\nAnother one!\n"
	s.Start()
	time.Sleep(4 * time.Second)
	s.Stop()
}