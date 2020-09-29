package main

import (
	"flag"
	"mustache-proxy/src"
)

func main() {
	host := flag.String("--host", "127.0.0.1", "Host")
	port := flag.String("--port", "5555", "Port")
	allowedTargets := flag.String("--targets", "./allowed_targets.ini", "Allowed Targets Config")
	debugMode := flag.Bool("--debug", false, "Enable debug mode")

	src.RunService(host, port, allowedTargets, debugMode)
}
