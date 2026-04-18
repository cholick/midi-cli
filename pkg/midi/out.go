package midi

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cholick/midi-cli/pkg/rtmidi"
	"github.com/samber/lo"
)

//go:generate go tool counterfeiter -o ../../internal/midifakes/fake_out.go . Out
type Out interface {
	ListPorts() ([]string, error)
	OpenPort(name string) error

	NoteOn(noteName string, velocity, channel int) error
	NoteOff(noteName string, velocity, channel int) error
	ProgramChange(programNumber, channel int) error
	ControlChange(controllerNumber, controllerValue, channel int) error

	PanicAll() error
	Panic(channel int) error

	Close()
}

type out struct {
	midiOut  rtmidi.MIDIOut
	portName string
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
			debugf("open exact match requested=%q index=%d", name, val)
			err := o.midiOut.OpenPort(val, "midi-cli")
			if err != nil {
				return fmt.Errorf("error opening port %s (%v)", name, val)
			}
			o.portName = key
			debugf("open success port=%q index=%d", key, val)
			return nil
		}

	}

	//partial match (just take the first)
	for key, val := range ports {
		if strings.Contains(key, name) {
			debugf("open partial match requested=%q matched=%q index=%d", name, key, val)
			err := o.midiOut.OpenPort(val, "midi-cli")
			if err != nil {
				return fmt.Errorf("error opening port %s (%v)", name, val)
			}
			o.portName = key
			debugf("open success port=%q index=%d", key, val)
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

	// status byte: 1000nnnn, nnnn = 0-15 for channels 1-16
	// data byte 1: 0kkkkkkk note
	// data byte 2: 0vvvvvvv velocity
	status := byte(0b10010000 + channel - 1)
	err = o.sendMessage([]byte{
		status,
		byte(num),
		byte(velocity),
	}, "note_on")
	if err != nil {
		return fmt.Errorf("error sending note 'on' message: %w", err)
	}

	return nil
}

func (o *out) NoteOff(noteName string, velocity, channel int) error {
	num, err := NoteNumberFromName(noteName)
	if err != nil {
		return err
	}

	// status byte: 1001nnnn, nnnn = 0-15 for channels 1-16
	// data byte 1: 0kkkkkkk note
	// data byte 2: 0vvvvvvv velocity
	status := byte(0b10000000 + channel - 1)
	err = o.sendMessage([]byte{
		status,
		byte(num),
		byte(velocity),
	}, "note_off")
	if err != nil {
		return fmt.Errorf("error sending note 'off' message: %w", err)
	}

	return nil
}

func (o *out) ProgramChange(programNumber int, channel int) error {
	// status byte: 1100nnnn, nnnn = 0-15 for channels 1-16
	// data byte: 0ppppppp
	status := byte(0b11000000 + channel - 1)
	err := o.sendMessage([]byte{
		status,
		byte(programNumber),
	}, "program_change")
	if err != nil {
		return fmt.Errorf("error sending program change message: %w", err)
	}

	return nil
}

func (o *out) ControlChange(controllerNumber int, controllerValue int, channel int) error {
	// status byte: 1011nnnn, nnnn = 0-15 for channels 1-16
	// data byte 1: 0ccccccc controller number
	// data byte 2: 0vvvvvvv controller value
	status := byte(0b10110000 + channel - 1)
	err := o.sendMessage([]byte{
		status,
		byte(controllerNumber),
		byte(controllerValue),
	}, "control_change")
	if err != nil {
		return fmt.Errorf("error sending program change message: %w", err)
	}

	return nil
}

func (o *out) PanicAll() error {
	// For Panic, collect errors rather than immediately returning to silence as much as possible
	errs := make([]error, 0)
	for c := range 16 {
		err := o.Panic(c + 1)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (o *out) Panic(channel int) error {
	return o.ControlChange(120, 0, channel)
}

func (o *out) Close() {
	// debugf("close delaying")
	// time.Sleep(100 * time.Millisecond)
	debugf("close port=%q", o.portName)
	_ = o.midiOut.Close()
}

func (o *out) sendMessage(message []byte, kind string) error {
	debugf("send start port=%q kind=%s bytes=%v", o.portName, kind, message)
	err := o.midiOut.SendMessage(message)
	if err != nil {
		debugf("send error port=%q kind=%s err=%v", o.portName, kind, err)
		return err
	}
	debugf("send success port=%q kind=%s bytes=%v", o.portName, kind, message)
	return nil
}

func debugf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(os.Stderr, "midi-cli debug %s %s\n", time.Now().Format(time.RFC3339Nano), message)
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
