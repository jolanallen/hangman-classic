package hangman



type HANGMAN struct {
	IsRunning     bool
	InGame        bool
	Win           bool
	Loose         bool
	wordIsGood    bool
	TabMots       []string
	TabRune       []rune
	Mot           string
	MotAdeviner   []string
	motIconnu     []string
	randomNb      int
	flag           bool
	erreur       int
	lettreIsGood bool
	lettre       string
	
	
	
}

