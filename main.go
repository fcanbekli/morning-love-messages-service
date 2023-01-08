package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_ = SendMessage("Canim babam")
}
