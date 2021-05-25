package filters

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type ChosenInlineResult func(cir *gotgbot.ChosenInlineResult) bool

func (f ChosenInlineResult) And(f2 ChosenInlineResult) ChosenInlineResult {
	return func(cir *gotgbot.ChosenInlineResult) bool {
		return f(cir) && f2(cir)
	}
}

func (f ChosenInlineResult) Or(f2 ChosenInlineResult) ChosenInlineResult {
	return func(cir *gotgbot.ChosenInlineResult) bool {
		return f(cir) || f2(cir)
	}
}

func (f ChosenInlineResult) Not() ChosenInlineResult {
	return func(cir *gotgbot.ChosenInlineResult) bool {
		return !f(cir)
	}
}

func ChosenResultUserID(id int64) ChosenInlineResult {
	return func(cir *gotgbot.ChosenInlineResult) bool {
		return cir.From.Id == id
	}
}

func ChosenResultQuery(q string) ChosenInlineResult {
	return func(cir *gotgbot.ChosenInlineResult) bool {
		return cir.Query == q
	}
}

func ChosenResultQueryPrefix(prefix string) ChosenInlineResult {
	return func(cir *gotgbot.ChosenInlineResult) bool {
		return strings.HasPrefix(cir.Query, prefix)
	}
}

func ChosenResultQuerySuffix(suffix string) ChosenInlineResult {
	return func(cir *gotgbot.ChosenInlineResult) bool {
		return strings.HasSuffix(cir.Query, suffix)
	}
}