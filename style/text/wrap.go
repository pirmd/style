package text

import (
	"strings"
	"unicode"
)

//Indent indents a string (add prefix at the begining and before any new line)
func Indent(s string, prefix []byte) string {
	b := []byte{}
	lastRune := len(s) - 1

	for i, c := range s {
		b = append(b, byte(c))
		if c == '\n' && i != lastRune {
			b = append(b, prefix...)
		}
	}

	return string(append(prefix, b...))
}

//Wrap wraps a text by ensuring that no text's line will be longer than the provided
//limit.
//
//Wrap calculates the visual width of a string, discarding non printable rune or ansi
//coloring/styling escape sequence.
//
//Wrap is pretty bad when called multiple times as it cannot differenciates line breaks
//coming from previous wrapping from line breaks introduced by the end-user in the
//first place. It is therfore better to avoid multiple calls to Wrap
//
//BUG: Wrap works at word limits (space) so it does not behaves properly for words
//longer than the limit (it does not split words so it feedbacks  line longer than
//limit)
func Wrap(txt string, limit int) string {
	return strings.Join(wrap(txt, limit), "\n")
}

//Tab wraps then indent
//Tab calculates the correct wraping limits taking indent's prefix length
//It does not work if prefix is made of tabs as indent's prefix length is
//unknown (like '\t')
func Tab(s string, prefix []byte, limit int) string {
	r := Wrap(s, limit-len(prefix))
	return Indent(r, prefix)
}

//BUG(pirmd): wrap blindly assumes that word length is smaller than
//the limit can we do better than that (like truncating the word)?
func wrap(txt string, limit int) []string {
	var c rune
	var ws []string
	var line, word string
	var linelen int

	for _, c = range txt {
		switch {
		case c == '\n':
			wordlen := visualLen(word)
			if linelen+wordlen > limit {
				ws = append(ws, line)
				line, linelen = "", 0
			}

			ws = append(ws, line+word)
			word, line, linelen = "", "", 0

		case unicode.IsSpace(c):
			wordlen := visualLen(word)
			switch l := linelen + wordlen; {
			case linelen == 0 && l >= limit:
				ws = append(ws, word)
			case l > limit:
				ws = append(ws, line)
				line = word + string(c)
				linelen = wordlen + 1
			case l == limit:
				ws = append(ws, line+word)
				line, linelen = "", 0
			default:
				line += word + string(c)
				linelen += wordlen + 1
			}
			word = ""

		default:
			word = word + string(c)
		}
	}

	wordlen := visualLen(word)
	switch l := linelen + wordlen; {
	case l == 0 && c == '\n':
		return append(ws, "")
	case l == 0:
		return ws
	case l > limit && linelen == 0:
		return append(ws, word)
	case l > limit:
		return append(ws, line, word)
	default:
		return append(ws, line+word)
	}
}