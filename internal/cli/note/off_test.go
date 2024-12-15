package note_test

import (
	"bytes"
	"flag"
	"testing"

	"github.com/cholick/midi-cli/internal/cli/note"
	"github.com/cholick/midi-cli/internal/midifakes"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotOff(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	cmd := note.NewNoteCommand()
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	out := ui.NewOutput(stdOut, stdErr)

	cmd.AddCommand(note.NewOffCommand(fakeOpener, out))

	cmd.SetArgs([]string{
		"off", "--port", "testPort", "--note", "E4", "-o", "99", "-c", "14"},
	)
	flag.Parse()

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	err := cmd.Execute()
	require.NoError(t, err)

	assert.Equal(t, 1, fakeOut.NoteOffCallCount())

	noteName, velocity, channel := fakeOut.NoteOffArgsForCall(0)
	assert.Equal(t, "E4", noteName)
	assert.Equal(t, 99, velocity)
	assert.Equal(t, 14, channel)
}
