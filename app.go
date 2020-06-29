package jawc

import (
	"github.com/ikawaha/kagome/tokenizer"
)

// App represents this application itself.
type App struct {
	Mode      tokenizer.TokenizeMode
	Tokenizer tokenizer.Tokenizer
}

// NewApp creates a new application.
func NewApp(mode tokenizer.TokenizeMode) *App {
	t := tokenizer.New()
	return &App{Mode: mode, Tokenizer: t}
}

// analyze analyzes input string with Kagome.
func (a App) analyze(input string) []tokenizer.Token {
	return a.Tokenizer.Analyze(input, a.Mode)
}

// CountWords counts the number of words.
// By default dummy tokens are excluded from counting.
// TODO: Exclude symbols such as punctuation and line breaks.
func (a App) CountWords(input string) int {
	c := 0
	for _, t := range a.analyze(input) {
		if t.Class == tokenizer.DUMMY {
			continue
		}
		c++
	}

	return c
}
