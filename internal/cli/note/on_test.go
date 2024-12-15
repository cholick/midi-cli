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

func TestNotOn(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	cmd := note.NewNoteCommand()
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	out := ui.NewOutput(stdOut, stdErr)

	cmd.AddCommand(note.NewOnCommand(fakeOpener, out))

	cmd.SetArgs([]string{
		"on", "--port", "testPort", "--note", "A4", "-o", "123", "-c", "14"},
	)
	flag.Parse()

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	err := cmd.Execute()
	require.NoError(t, err)

	assert.Equal(t, 1, fakeOut.NoteOnCallCount())

	noteName, velocity, channel := fakeOut.NoteOnArgsForCall(0)
	assert.Equal(t, "A4", noteName)
	assert.Equal(t, 123, velocity)
	assert.Equal(t, 14, channel)
}
