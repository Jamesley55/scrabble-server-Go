package constant

type MessageType string

const (
	Message_Player   MessageType = "player"
	Message_Opponent MessageType = "opponent"
	Message_System   MessageType = "system"
	Message_Reserve  MessageType = "reserve"
	Message_Help     MessageType = "help"
)

type SCORE_PERCENTILE float64

const (
	EIGHTIETH_PERCENTILE_SCORE  SCORE_PERCENTILE = 0.8
	SEVENTIETH_PERCENTILE_SCORE SCORE_PERCENTILE = 0.7
	SIXTIETH_PERCENTILE_SCORE   SCORE_PERCENTILE = 0.6
)

type CommandName string

const (
	command_Place    CommandName = "placer"
	command_Exchange CommandName = "échanger"
	command_Skip     CommandName = "passer"
	command_Reserve  CommandName = "réserve"
	command_Help     CommandName = "aide"
	command_Hint     CommandName = "indice"
)

type CommandError string

const (
	Syntax     CommandError = "syntax"
	Impossible CommandError = "impossible"
	Invalid    CommandError = "invalid"
	GameDone   CommandError = "gameDone"
)

const (
	SETTING_DICTIONNARY           = "Dictionnaire"
	SETTING_JV_NAMES              = "Nom des Joueur Virtuel"
	SETTING_HISTORIQUE_DES_PARTIS = "Historique des Parties"
	SETTING_HIGHSCORE             = "HighScore"
)

type GameMode string

const (
	Classic GameMode = "Classic"
	Log2990 GameMode = "Log2990"
)

var SETTINGS_TYPE = []string{SETTING_DICTIONNARY, SETTING_JV_NAMES, SETTING_HISTORIQUE_DES_PARTIS, SETTING_HIGHSCORE}

type TileType int

const (
	BASIC TileType = iota
	DOUBLELETTER
	TRIPLELETTER
	DOUBLEWORD
	TRIPLEWORD
	STAR
	ROW
	COL
	EMPTY
)

type ObjectiveName string

const (
	Palindrome       ObjectiveName = "Palindrome"
	FirstTo50        ObjectiveName = "50 points"
	TripleSameLetter ObjectiveName = "3 times same letters"
	TripleVowels     ObjectiveName = "3 vowels"
	StartEnd         ObjectiveName = "Start finish same letter"
	Word20Points     ObjectiveName = "Form 20 point word"
	FormScrabble     ObjectiveName = "Form word Scrabble"
	WordOnGrid       ObjectiveName = "Form word already on board"
)

type VirtualPlayerType string

const (
	Debutant VirtualPlayerType = "Debutant"
	Expert   VirtualPlayerType = "Expert"
)
const (
	CHOOSE_ACTION_DELAY_MS                              int     = 3000
	SKIP_TURN_DELAY_MS                                  int     = 17000
	MAX_RECURSION_CALL                                  int     = 50
	PROBABILITY_WORD_SCORE_BETWEEN_1_AND_6              float64 = 0.4
	PROBABILITY_WORD_SCORE_BETWEEN_7_AND_12_LOWER_BOUND float64 = 0.4
	PROBABILITY_WORD_SCORE_BETWEEN_7_AND_12_UPPER_BOUND float64 = 0.7
	PROBABILITY_WORD_SCORE_BETWEEN_13_AND_18            float64 = 0.3
	WORD_SCORE_RANGE                                    int     = 6
)
const MAX_SKIP_COUNT = 6
