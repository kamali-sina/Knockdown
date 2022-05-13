package main

type Commandset struct {
	Names    map[int]string
	Handlers map[int]CommandHandler
}

var ActiveCommandset Commandset

var DefendingCommandset = Commandset{
	map[int]string{
		1: "Block",
		2: "Dodge",
	},
	map[int]CommandHandler{
		1: HandleBlock,
		2: HandleDodge,
	},
}

var MainMenuCommandset = Commandset{
	map[int]string{
		1: "Fight",
		2: "Level Up",
		3: "Exit",
	},
	map[int]CommandHandler{
		1: HandleFight,
		2: HandleLevelup,
		3: HandleExit,
	},
}
