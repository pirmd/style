//Code generated by 'go run helpers_generate.go'; DO NOT EDIT.
package style

//Upper applies the style 'Upper' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Upper(format string, a ...interface{}) string {
	return st.stylef(FmtUpper)(format, a...)
}

//Upper applies the style 'Upper' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Upper(format string, a ...interface{}) string {
	return CurrentStyler.Upper(format, a...)
}

//Lower applies the style 'Lower' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Lower(format string, a ...interface{}) string {
	return st.stylef(FmtLower)(format, a...)
}

//Lower applies the style 'Lower' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Lower(format string, a ...interface{}) string {
	return CurrentStyler.Lower(format, a...)
}

//Black applies the style 'Black' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Black(format string, a ...interface{}) string {
	return st.stylef(FmtBlack)(format, a...)
}

//Black applies the style 'Black' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Black(format string, a ...interface{}) string {
	return CurrentStyler.Black(format, a...)
}

//Red applies the style 'Red' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Red(format string, a ...interface{}) string {
	return st.stylef(FmtRed)(format, a...)
}

//Red applies the style 'Red' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Red(format string, a ...interface{}) string {
	return CurrentStyler.Red(format, a...)
}

//Green applies the style 'Green' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Green(format string, a ...interface{}) string {
	return st.stylef(FmtGreen)(format, a...)
}

//Green applies the style 'Green' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Green(format string, a ...interface{}) string {
	return CurrentStyler.Green(format, a...)
}

//Yellow applies the style 'Yellow' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Yellow(format string, a ...interface{}) string {
	return st.stylef(FmtYellow)(format, a...)
}

//Yellow applies the style 'Yellow' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Yellow(format string, a ...interface{}) string {
	return CurrentStyler.Yellow(format, a...)
}

//Blue applies the style 'Blue' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Blue(format string, a ...interface{}) string {
	return st.stylef(FmtBlue)(format, a...)
}

//Blue applies the style 'Blue' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Blue(format string, a ...interface{}) string {
	return CurrentStyler.Blue(format, a...)
}

//Magenta applies the style 'Magenta' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Magenta(format string, a ...interface{}) string {
	return st.stylef(FmtMagenta)(format, a...)
}

//Magenta applies the style 'Magenta' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Magenta(format string, a ...interface{}) string {
	return CurrentStyler.Magenta(format, a...)
}

//Cyan applies the style 'Cyan' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Cyan(format string, a ...interface{}) string {
	return st.stylef(FmtCyan)(format, a...)
}

//Cyan applies the style 'Cyan' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Cyan(format string, a ...interface{}) string {
	return CurrentStyler.Cyan(format, a...)
}

//White applies the style 'White' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) White(format string, a ...interface{}) string {
	return st.stylef(FmtWhite)(format, a...)
}

//White applies the style 'White' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func White(format string, a ...interface{}) string {
	return CurrentStyler.White(format, a...)
}

//Bold applies the style 'Bold' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Bold(format string, a ...interface{}) string {
	return st.stylef(FmtBold)(format, a...)
}

//Bold applies the style 'Bold' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Bold(format string, a ...interface{}) string {
	return CurrentStyler.Bold(format, a...)
}

//Italic applies the style 'Italic' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Italic(format string, a ...interface{}) string {
	return st.stylef(FmtItalic)(format, a...)
}

//Italic applies the style 'Italic' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Italic(format string, a ...interface{}) string {
	return CurrentStyler.Italic(format, a...)
}

//Underline applies the style 'Underline' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Underline(format string, a ...interface{}) string {
	return st.stylef(FmtUnderline)(format, a...)
}

//Underline applies the style 'Underline' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Underline(format string, a ...interface{}) string {
	return CurrentStyler.Underline(format, a...)
}

//Inverse applies the style 'Inverse' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Inverse(format string, a ...interface{}) string {
	return st.stylef(FmtInverse)(format, a...)
}

//Inverse applies the style 'Inverse' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Inverse(format string, a ...interface{}) string {
	return CurrentStyler.Inverse(format, a...)
}

//Strike applies the style 'Strike' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Strike(format string, a ...interface{}) string {
	return st.stylef(FmtStrike)(format, a...)
}

//Strike applies the style 'Strike' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Strike(format string, a ...interface{}) string {
	return CurrentStyler.Strike(format, a...)
}

//Wrap applies the style 'Wrap' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Wrap(format string, a ...interface{}) string {
	return st.stylef(FmtWrap)(format, a...)
}

//Wrap applies the style 'Wrap' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Wrap(format string, a ...interface{}) string {
	return CurrentStyler.Wrap(format, a...)
}

//Tab applies the style 'Tab' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Tab(format string, a ...interface{}) string {
	return st.stylef(FmtTab)(format, a...)
}

//Tab applies the style 'Tab' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Tab(format string, a ...interface{}) string {
	return CurrentStyler.Tab(format, a...)
}

//Tab2 applies the style 'Tab2' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Tab2(format string, a ...interface{}) string {
	return st.stylef(FmtTab2)(format, a...)
}

//Tab2 applies the style 'Tab2' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Tab2(format string, a ...interface{}) string {
	return CurrentStyler.Tab2(format, a...)
}

//DocHeader applies the style 'DocHeader' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) DocHeader(format string, a ...interface{}) string {
	return st.stylef(FmtDocHeader)(format, a...)
}

//DocHeader applies the style 'DocHeader' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func DocHeader(format string, a ...interface{}) string {
	return CurrentStyler.DocHeader(format, a...)
}

//Header applies the style 'Header' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Header(format string, a ...interface{}) string {
	return st.stylef(FmtHeader)(format, a...)
}

//Header applies the style 'Header' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Header(format string, a ...interface{}) string {
	return CurrentStyler.Header(format, a...)
}

//Paragraph applies the style 'Paragraph' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Paragraph(format string, a ...interface{}) string {
	return st.stylef(FmtParagraph)(format, a...)
}

//Paragraph applies the style 'Paragraph' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Paragraph(format string, a ...interface{}) string {
	return CurrentStyler.Paragraph(format, a...)
}

//NewLine applies the style 'NewLine' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) NewLine(format string, a ...interface{}) string {
	return st.stylef(FmtNewLine)(format, a...)
}

//NewLine applies the style 'NewLine' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func NewLine(format string, a ...interface{}) string {
	return CurrentStyler.NewLine(format, a...)
}

//List applies the style 'List' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) List(format string, a ...interface{}) string {
	return st.stylef(FmtList)(format, a...)
}

//List applies the style 'List' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func List(format string, a ...interface{}) string {
	return CurrentStyler.List(format, a...)
}

//DefTerm applies the style 'DefTerm' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) DefTerm(format string, a ...interface{}) string {
	return st.stylef(FmtDefTerm)(format, a...)
}

//DefTerm applies the style 'DefTerm' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func DefTerm(format string, a ...interface{}) string {
	return CurrentStyler.DefTerm(format, a...)
}

//DefDesc applies the style 'DefDesc' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) DefDesc(format string, a ...interface{}) string {
	return st.stylef(FmtDefDesc)(format, a...)
}

//DefDesc applies the style 'DefDesc' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func DefDesc(format string, a ...interface{}) string {
	return CurrentStyler.DefDesc(format, a...)
}

//Code applies the style 'Code' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Code(format string, a ...interface{}) string {
	return st.stylef(FmtCode)(format, a...)
}

//Code applies the style 'Code' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Code(format string, a ...interface{}) string {
	return CurrentStyler.Code(format, a...)
}

//Escape applies the style 'Escape' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Escape(format string, a ...interface{}) string {
	return st.stylef(FmtEscape)(format, a...)
}

//Escape applies the style 'Escape' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Escape(format string, a ...interface{}) string {
	return CurrentStyler.Escape(format, a...)
}

//Auto applies the style 'Auto' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) Auto(format string, a ...interface{}) string {
	return st.stylef(FmtAuto)(format, a...)
}

//Auto applies the style 'Auto' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func Auto(format string, a ...interface{}) string {
	return CurrentStyler.Auto(format, a...)
}

//FuncMap provides a text/template FuncMap compatible mapping
//to use 'style' functions within templates
func (st Styler) FuncMap() map[string]interface{} {
	return map[string]interface{}{
		"Upper":     st.Upper,
		"Lower":     st.Lower,
		"Black":     st.Black,
		"Red":       st.Red,
		"Green":     st.Green,
		"Yellow":    st.Yellow,
		"Blue":      st.Blue,
		"Magenta":   st.Magenta,
		"Cyan":      st.Cyan,
		"White":     st.White,
		"Bold":      st.Bold,
		"Italic":    st.Italic,
		"Underline": st.Underline,
		"Inverse":   st.Inverse,
		"Strike":    st.Strike,
		"Wrap":      st.Wrap,
		"Tab":       st.Tab,
		"Tab2":      st.Tab2,
		"DocHeader": st.DocHeader,
		"Header":    st.Header,
		"Paragraph": st.Paragraph,
		"NewLine":   st.NewLine,
		"List":      st.List,
		"DefTerm":   st.DefTerm,
		"DefDesc":   st.DefDesc,
		"Code":      st.Code,
		"Escape":    st.Escape,
		"Auto":      st.Auto,
	}
}

//FuncMap provides a text/template FuncMap compatible mapping
//to use 'style' CurrentStyler's functions within templates
func FuncMap() map[string]interface{} {
	return map[string]interface{}{
		"Upper":     Upper,
		"Lower":     Lower,
		"Black":     Black,
		"Red":       Red,
		"Green":     Green,
		"Yellow":    Yellow,
		"Blue":      Blue,
		"Magenta":   Magenta,
		"Cyan":      Cyan,
		"White":     White,
		"Bold":      Bold,
		"Italic":    Italic,
		"Underline": Underline,
		"Inverse":   Inverse,
		"Strike":    Strike,
		"Wrap":      Wrap,
		"Tab":       Tab,
		"Tab2":      Tab2,
		"DocHeader": DocHeader,
		"Header":    Header,
		"Paragraph": Paragraph,
		"NewLine":   NewLine,
		"List":      List,
		"DefTerm":   DefTerm,
		"DefDesc":   DefDesc,
		"Code":      Code,
		"Escape":    Escape,
		"Auto":      Auto,
	}
}
