package main

import "strconv"

func run(p *Player) {
	for {
		PrintAvailableCommands()
		var gotValidInput bool
		for !gotValidInput {
			commandId, err := strconv.Atoi(ReadLine(""))
			if err != nil {
				continue
			}
			handler, ok := ActiveCommandset.Handlers[commandId]
			if !ok {
				continue
			}
			handler(make([]string, 0), p)
			gotValidInput = true
		}
	}
}
