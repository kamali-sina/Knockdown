package main

type Commandset struct {
	Names    map[int]string
	Handlers map[int]CommandHandler
}

var ActiveCommandset Commandset

var AttackingCommandset = Commandset{
	map[int]string{
		1: "Jab High",
		2: "Jab Mid",
		3: "Jab Low",
		4: "Punch High",
		5: "Punch Mid",
		6: "Punch Low",
		7: "Grab",
	},
	map[int]CommandHandler{
		1: HandleBlock,
		2: HandleDodge,
	},
}

var DefendingCommandset = Commandset{
	map[int]string{
		1: "Block High",
		2: "Block Mid",
		3: "Block Low",
		4: "Dodge",
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
