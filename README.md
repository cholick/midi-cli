## midi-cli

Probably use https://github.com/gbevin/SendMIDI instead of this; it has considerably more features. I only
created this because I wanted to send program changes and a few other MIDI commands using a Stream Deck, and used
that as an excuse to do some MIDI work in Go (and get myself a bit more familiar with the messages for
some hardware work).

### Use


```shell
$ midi-cli --help
Send MIDI messages

Usage:
  midi-cli [command]

Available Commands:
  bs          Send bank select (cc 0) message
  cc          Send control change messages
  help        Help about any command
  note        Send note messages
  panic       Send all notes off on all channels to all visible ports
  pc          Send program change messages
  port        Manage MIDI ports

Flags:
  -h, --help      help for midi-cli
  -v, --verbose   Verbose output
```

Example use that selects the third preset in the second playlist in Pigments (the choice
of counting starting from 0 or 1 is implementation dependent, the cli starts at 0):
```shell
midi-cli bs --value 1 --port midi-cli
midi-cli pc --number 2 --port midi-cli
```

### Setup

See [detailed-setup.md](docs/detailed-setup.md) for some detailed setup instructions for a Mac. It amounts to
making sure the built-in IAC driver is on in MIDI Studio, sending signals to that driver, and making sure the
stand-alone instrument / DAW is listening.
