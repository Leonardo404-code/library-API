package logger

import (
	"log"
	"os"
)

type Colorful struct {
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
}

func NewColorful(info, warning, err Color) *Colorful {
	colorful := &Colorful{
		info:    log.New(os.Stdout, string(info)+"[INFO]\t", log.Ldate|log.Ltime),
		warning: log.New(os.Stdout, string(warning)+"[WARNING]\t", log.Ldate|log.Ltime),
		err:     log.New(os.Stdout, string(err)+"[ERROR]\t", log.Ldate|log.Ltime),
	}

	return colorful
}

// Info prints relevant information
func (c *Colorful) Info(format string, v ...interface{}) {
	c.info.Printf(format, v...)
}

// Warning raises a warning
func (c *Colorful) Warning(format string, v ...interface{}) {
	c.warning.Printf(format, v...)
}

// Error alerts about an error
func (c *Colorful) Error(format string, v ...interface{}) {
	c.err.Printf(format, v...)
}

// Error alerts about a fatal error and exits the application
func (c *Colorful) Fatal(format string, v ...interface{}) {
	c.err.Printf(format, v...)
	os.Exit(1)
}
