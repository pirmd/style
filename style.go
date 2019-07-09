package style

//Styler represents any type that knows how to style text
type Styler interface {
	//Upper changes a string case to upper case
	Upper(string) string

	//Lower changes a string case to lower case
	Lower(string) string

	//TitleCase changes all letters that begin words to their title case.
	TitleCase(string) string

	//Black changes a string foreground color to black
	Black(string) string

	//Red changes a string foreground color to red
	Red(string) string

	//Green changes a string foreground color to green
	Green(string) string

	//Yellow changes a string foreground color to yellow
	Yellow(string) string

	//Blue changes a string foreground color to blue
	Blue(string) string

	//Magenta changes a string foreground color to magenta
	Magenta(string) string

	//Cyan changes a string foreground color to cyan
	Cyan(string) string

	//White changes a string foreground color to white
	White(string) string

	//Bold changes a string case to bold
	Bold(string) string

	//Italic changes a string case to italic
	Italic(string) string

	//Underline changes a string to be underlined
	Underline(string) string

	//Inverse changes a string by inverting its fore- and back-ground
	//colors
	Inverse(string) string

	//Strike changes a string to be strikethrough
	Strike(string) string

	//Tab indents then wraps the provided text by the given indent level.
	Tab(int) func(string) string

	//Header returns text as a chapter's header
	//Header of level 0 corresponds to document header.
	Header(int) func(string) string

	//Paragraph returns text as a new paragraph.
	Paragraph(string) string

	//List returns a new bullet list with the given nested level.
	List(int) func(...string) string

	//Define returns a term definition.
	Define(string, string) string

	//Table draws a table.
	Table(...[]string) string

	//Escape escapes the provided text. Chaining Escapes with Styler formatting
	//functions can lead to unexpected results.
	Escape(string) string
}
