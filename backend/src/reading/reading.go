package reading

import (
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type Reading interface {
	GetReading(text string) string
}

type Tokenizer struct {
	tokenizer *tokenizer.Tokenizer
}

func NewTokenizer() *Tokenizer {
	t, _ := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	return &Tokenizer{tokenizer: t}
}

func (t *Tokenizer) GetReading(text string) string {
	tokens := t.tokenizer.Tokenize(text)

	resp := ""
	for _, token := range tokens {
		reading, b := token.Reading()
		if b {
			resp += reading
		} else {
			resp += token.Surface
		}
	}

	return resp
}
