package utils

import (
    "os"
)

/**
IsStdinPresent function is used
to check if an argument is passed 
from stdin using pipe.
*/
func IsStdinPresent() bool {
    fi, err := os.Stdin.Stat()
    if err != nil {
        panic(err)
    }

    return fi.Mode() & os.ModeCharDevice == 0
}
