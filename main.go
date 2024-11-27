package main

import (
	"log"

	"github.com/cholick/midi-cli/internal/rtmidi"
)

func main() {
	// Create a new MIDI output
	out, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	count, err := out.PortCount()
	if err != nil {
		log.Fatal(err)
	}
	println(count)

	for i := range count {
		name, err := out.PortName(i)
		if err != nil {
			log.Fatal(err)
		}
		println(i, name)
	}
}
