package cli_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/cholick/midi-cli/internal/cli"
	"github.com/cholick/midi-cli/internal/midifakes"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPanic(t *testing.T) {
	fakeOut := &midifakes.FakeOut{}
	fakeOut.ListPortsReturns([]string{"device1", "device2"}, nil)
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	fakeOpener.NewDefaultOutReturns(fakeOut, nil)

	out := ui.NewOutputForTesting()

	cmd := cli.NewPanicCommand(fakeOpener, out)
	cmd.SetOut(out.StdOut)
	cmd.SetErr(out.ErrOut)

	cmd.SetArgs([]string{})
	flag.Parse()

	err := cmd.Execute()
	require.NoError(t, err)

	require.Equal(t, 2, fakeOut.PanicAllCallCount())

	assert.Equal(t, "device1", fakeOpener.NewOutForPortArgsForCall(0))
	assert.Equal(t, "device2", fakeOpener.NewOutForPortArgsForCall(1))
}

func TestPanicCallsOtherDevicesAfterError(t *testing.T) {
	fakeOut := &midifakes.FakeOut{}
	fakeOut.ListPortsReturns([]string{"device1", "device2"}, nil)
	fakeOpener := &midifakes.FakeOpener{}
	// fail on the first to test that we move to the second
	fakeOpener.NewOutForPortReturnsOnCall(0, nil, errors.New("first device failed"))
	fakeOpener.NewOutForPortReturnsOnCall(1, fakeOut, nil)
	fakeOpener.NewDefaultOutReturns(fakeOut, nil)

	out := ui.NewOutputForTesting()

	cmd := cli.NewPanicCommand(fakeOpener, out)
	cmd.SetOut(out.StdOut)
	cmd.SetErr(out.ErrOut)

	cmd.SetArgs([]string{})
	flag.Parse()

	err := cmd.Execute()
	require.Error(t, err)

	assert.Equal(t, "device1", fakeOpener.NewOutForPortArgsForCall(0))
	assert.Equal(t, "device2", fakeOpener.NewOutForPortArgsForCall(1))

	require.Equal(t, 1, fakeOut.PanicAllCallCount())
}
