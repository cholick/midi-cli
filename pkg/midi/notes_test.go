package midi_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	testCases := []struct {
		noteName    string
		expectedNum int
	}{
		{"C4", 60},
		{"c4", 60},
		{"C#4", 61},
		{"Db4", 61},
		{"D♭4", 61},
	}
	for _, tc := range testCases {
		t.Run(tc.noteName, func(t *testing.T) {
			num, err := midi.NoteNumberFromName(tc.noteName)
			require.NoError(t, err)

			assert.Equal(t, num, tc.expectedNum)
		})
	}
}
