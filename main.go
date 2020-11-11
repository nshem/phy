package main

import "github.com/nshem/phy/internal/session"

func main() {
	currSession := session.NewSession()
	currSession.Start()
}
