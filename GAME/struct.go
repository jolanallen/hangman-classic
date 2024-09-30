package hangman



type HANGMAN struct {
	IsRunning     bool
	InGame        bool
	Win           bool
	Loose         bool
	wordIsGood    bool
	TabMots       []string
	Mot           string
	lMot          int
	MotAdeviner   string
	randomNb      int
	TabByte       []byte
	flag           bool
	
	
}

