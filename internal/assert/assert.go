package assert

import (
	"log"

	"github.com/prafitradimas/interpreter/internal/token"
)

var tokenData []token.Token = []token.Token{}

func AddTokenData(token token.Token) {
	tokenData = append(tokenData, token)
}

func ClearTokenData() {
	if len(tokenData) > 0 {
		tokenData = []token.Token{}
	}
}

func printTokenData() {
	if len(tokenData) > 0 {
		log.Printf("tokens: %+v\n", tokenData)
	}
}

func InvalidToken(cond bool, token token.Token, msg string) {
	if cond {
		printTokenData()
		log.Fatalf("Invalid Token. Type: %v, Literal: %s\n%s", token.Type, token.Literal, msg)
	}
}
