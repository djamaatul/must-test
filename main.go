package main

import (
	tests "must-test/tests"
	server "must-test/tests/Practical"
)

func main() {
	tests.Test()
	server.StartServer()
}
