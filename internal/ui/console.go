package ui

import (
	"bytes"
	"fmt"
	"io"

	"github.com/samber/lo"
)

type Console interface {
	SetDebug(bool)

	Debug(v ...any)
	Debugf(format string, v ...any)

	Info(v ...any)
	Infof(format string, v ...any)

	Error(v ...any)
	Errorf(format string, v ...any)
}

type ConsoleForTesting struct {
	Console

	StdOut io.Writer
	ErrOut io.Writer
}

type console struct {
	out    io.Writer
	errOut io.Writer
	debug  bool
}

func NewOutput(out, errOut io.Writer) Console {
	return &console{
		out:    out,
		errOut: errOut,
	}
}

func NewOutputForTesting() *ConsoleForTesting {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	return &ConsoleForTesting{
		Console: NewOutput(stdOut, stdErr),
		StdOut:  stdOut,
		ErrOut:  stdErr,
	}
}

func (c *console) SetDebug(debug bool) {
	c.debug = debug
}

func (c *console) Debug(v ...any) {
	if !c.debug {
		return
	}
	c.withNewline(c.out, v...)
}

func (c *console) Debugf(format string, v ...any) {
	if !c.debug {
		return
	}
	c.withNewlineF(c.out, format, v...)
}

func (c *console) Info(v ...any) {
	c.withNewline(c.out, v...)
}

func (c *console) Infof(format string, v ...any) {
	c.withNewlineF(c.out, format, v...)
}

func (c *console) Error(v ...any) {
	c.withNewline(c.errOut, v...)
}

func (c *console) Errorf(format string, v ...any) {
	c.withNewlineF(c.errOut, format, v...)
}

func (c *console) withNewline(w io.Writer, items ...any) {
	msg := lo.Reduce(items[1:], func(agg string, item any, i int) string {
		return fmt.Sprintf("%s %v", agg, item)
	}, fmt.Sprintf("%v", items[0]))

	_, _ = fmt.Fprint(w, msg, "\n")
}

func (c *console) withNewlineF(w io.Writer, format string, items ...any) {
	_, _ = fmt.Fprintf(w, format+"\n", items...)
}
