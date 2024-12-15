package base

import (
	"fmt"

	"github.com/spf13/cobra"
)

type FlagValues struct {
	Port    string
	Channel int
}

func GetFlagValues(cmd *cobra.Command) (*FlagValues, error) {
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		return nil, fmt.Errorf("error getting port flag: %w", err)
	}

	channel, err := cmd.Flags().GetInt("channel")
	if err != nil {
		return nil, fmt.Errorf("error getting channel flag: %w", err)
	}

	return &FlagValues{
		Port:    port,
		Channel: channel,
	}, nil
}
