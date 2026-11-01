package main

import "os"

// exit ends the process; tests replace it to run main in-process.
var exit = os.Exit
