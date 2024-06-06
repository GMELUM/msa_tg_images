package config

import (
	"flag"
	"time"
)

var (
	SIZE  = flag.Int("s", 1000, "an int")
	DELAY = flag.Duration("d", time.Minute, "a time")
	TOKEN = flag.String("t", "", "a string")
)
