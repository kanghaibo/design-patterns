package simple

type Translator interface {
	Translate(string) string
}

type GermanTranslator struct{}

func (t GermanTranslator) Translate(s string) string {
	return "german"
}

type EnglishTranslator struct{}

func (t EnglishTranslator) Translate(s string) string {
	return "english"
}

type ChineseTranslator struct{}

func (t ChineseTranslator) Translate(s string) string {
	return "chinese"
}

func NewTranslator(name string) Translator {
	var t Translator
	switch name {
	case "german":
		t = GermanTranslator{}
	case "english":
		t = EnglishTranslator{}
	case "chinese":
		t = ChineseTranslator{}
	default:
		panic("no such translator")
	}
	return t
}
