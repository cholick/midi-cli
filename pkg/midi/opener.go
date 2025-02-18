package midi

import (
	"fmt"

	"github.com/cholick/midi-cli/pkg/rtmidi"
)

//go:generate go tool counterfeiter -o ../../internal/midifakes/fake_opener.go . Opener
type Opener interface {
	NewDefaultOut() (Out, error)
	NewOutForPort(port string) (Out, error)
}

func NewOpener() Opener {
	return &opener{}
}

type opener struct{}

func (o *opener) NewDefaultOut() (Out, error) {
	def, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		return nil, fmt.Errorf("error opening default out: %w", err)
	}
	return &out{
		midiOut: def,
	}, nil
}

func (o *opener) NewOutForPort(port string) (Out, error) {
	midiOut, err := o.NewDefaultOut()
	if err != nil {
		return nil, err
	}

	err = midiOut.OpenPort(port)
	if err != nil {
		return nil, fmt.Errorf("error opening port: %w", err)
	}

	return midiOut, nil
}
