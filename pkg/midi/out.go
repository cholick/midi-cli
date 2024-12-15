package midi

import (
	"fmt"
	"strings"

	"github.com/cholick/midi-cli/pkg/rtmidi"
	"github.com/samber/lo"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ../../internal/midifakes/fake_out.go . Out
type Out interface {
	ListPorts() ([]string, error)
	OpenPort(name string) error

	NoteOn(noteName string, velocity, channel int) error
	NoteOff(noteName string, velocity, channel int) error

	Close()
}

type out struct {
	midiOut rtmidi.MIDIOut
}

func (o *out) ListPorts() ([]string, error) {
	ports, err := o.getPorts()
	if err != nil {
		return nil, err
	}

	return lo.Keys(ports), nil
}

func (o *out) OpenPort(name string) error {
	ports, err := o.getPorts()
	if err != nil {
		return fmt.Errorf("error listing ports %w", err)
	}

	//exact match
	for key, val := range ports {
		if key == name {
			err := o.midiOut.OpenPort(val, "midi-cli")
			if err != nil {
				return fmt.Errorf("error opening port %s (%v)", name, val)
			}
			return nil
		}

	}

	//partial match (just take the first)
	for key, val := range ports {
		if strings.Contains(key, name) {
			err := o.midiOut.OpenPort(val, "midi-cli")
			if err != nil {
				return fmt.Errorf("error opening port %s (%v)", name, val)
			}
			return nil
		}
	}

	//no match
	return fmt.Errorf("no match for port '%s' in %v", name, lo.Keys(ports))
}

func (o *out) NoteOn(noteName string, velocity, channel int) error {
	num, err := NoteNumberFromName(noteName)
	if err != nil {
		return err
	}

	//1000nnnn, nnnn = 0-15 for channels 1-16
	status := byte(0b10010000 + channel - 1)
	err = o.midiOut.SendMessage([]byte{
		status,
		byte(num),
		byte(velocity),
	})
	if err != nil {
		return fmt.Errorf("error sending note on message: %w", err)
	}

	return nil
}

func (o *out) NoteOff(noteName string, velocity, channel int) error {
	num, err := NoteNumberFromName(noteName)
	if err != nil {
		return err
	}

	//1001nnnn, nnnn = 0-15 for channels 1-16
	status := byte(0b10000000 + channel - 1)
	err = o.midiOut.SendMessage([]byte{
		status,
		byte(num),
		byte(velocity),
	})
	if err != nil {
		return fmt.Errorf("error sending note on message: %w", err)
	}

	return nil
}

func (o *out) Close() {
	_ = o.midiOut.Close()
}

func (o *out) getPorts() (map[string]int, error) {
	count, err := o.midiOut.PortCount()
	if err != nil {
		return nil, fmt.Errorf("err getting port count: %w", err)
	}

	ports := map[string]int{}
	for i := range count {
		name, err := o.midiOut.PortName(i)
		if err != nil {
			return nil, fmt.Errorf("err getting port name: %w", err)
		}

		ports[name] = i
	}

	return ports, nil
}
