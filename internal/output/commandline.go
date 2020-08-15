package output

import "os"

type CommandLine struct {
}

func (c *CommandLine) Write(buf []byte) (int, error) {
	return os.Stdout.Write(buf)
}

func (c *CommandLine) Read(buf []byte) (int, error) {
	return os.Stdout.Read(buf)
}
