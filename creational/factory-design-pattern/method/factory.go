package method

type Translator interface {
	Translate(string) string
}

type Factory interface {
	Create() Translator
}

type GermanTranslator struct{}
type GermanTranslatorFactory struct{}

func (t GermanTranslator) Translate(s string) string {
	return "german"
}
func (f GermanTranslatorFactory) Create() Translator {
	return GermanTranslator{}
}

type EnglishTranslator struct{}
type EnglishTranslatorFactory struct{}

func (t EnglishTranslator) Translate(s string) string {
	return "english"
}
func (f EnglishTranslatorFactory) Create() Translator {
	return EnglishTranslator{}
}

type ChineseTranslator struct{}
type ChineseTranslatorFactory struct{}

func (t ChineseTranslator) Translate(s string) string {
	return "chinese"
}
func (f ChineseTranslatorFactory) Create() Translator {
	return ChineseTranslator{}
}

//func NewTranslator(name string) Translator {
//	var t Translator
//	switch name {
//	case "german":
//		t = GermanTranslator{}
//	case "english":
//		t = EnglishTranslator{}
//	case "chinese":
//		t = ChineseTranslator{}
//	default:
//		panic("no such translator")
//	}
//	return t
//}

// todo: cache factory
func NewTranslatorFactory(name string) Factory {
	var f Factory
	switch name {
	case "german":
		f = GermanTranslatorFactory{}
	case "english":
		f = EnglishTranslatorFactory{}
	case "chinese":
		f = ChineseTranslatorFactory{}
	default:
		panic("no such translator factory")
	}
	return f
}
