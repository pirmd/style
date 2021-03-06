package style_test

import (
	"testing"

	"github.com/pirmd/verify"

	"github.com/pirmd/style"
)

func testText(st style.Styler) (s string) {
	H := style.Chain(st.Header(1), st.Blue)
	H2 := style.Chain(st.Header(2), st.Red)

	s = st.Metadata(map[string]string{"title": "github.com/pirmd/style", "mansection": "1", "authors": "pirmd", "date": "2019-07-09"})

	s += H("Introduction")
	s += st.Paragraph("This small piece of text aims at demonstrating and testing package '" + st.Underline("style") + "'.")
	s += st.Paragraph("It is written by a " + st.Bold("non-native") + " English speaker, so pardon any faults.")
	s += st.Paragraph("All details can be found in " + st.Link(st.Img("GoDoc", "https://godoc.org/github.com/pirmd/style?status.svg"), "https://godoc.org/github.com/pirmd/style"))

	s += H("Examples of available styles")

	s += H2("Demonstrating common formatting")
	s += st.Paragraph("Section " + st.Underline("Introduction") + " already demonstrates useful styles from package 'styles', this section completes them with most of the others possibilities.")
	s += st.Paragraph("Notably, package " + st.Underline("style") + " can print in " + st.Red("red") + " or " + st.Bold(st.Green("green bold")) + " (if chosen style supports it).")
	s += st.Paragraph("Several levels of tabulations can be used:")
	s += st.Tab()(
		st.Paragraph("(Level 1) Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.") +
			st.Tab()(
				st.Paragraph("(Level 2) Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.")+
					st.Tab()(
						st.Tab()(
							st.Paragraph("(Level 4) Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						),
					),
			),
	)

	s += H2("Demonstrating lists")
	s += st.Paragraph("It knows how to format " + st.Italic("lists") + ": ")
	s += st.BulletedList()(
		"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped. It should hopefully be so long that it will not fulfill the maximum number of authorized chars per line is reached.",
		"It also can support sub-lists:\n"+st.BulletedList()(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
			"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			"And sub-sub-ub-lists:\n"+st.BulletedList()(
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
				"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			),
		),
		"It is also possible to have a list's item that contains several paragraphs.\n"+
			st.Paragraph("For example, this paragraph that I made artificially long to verify that wrapping is working correctly inside list."),
	)

	s += st.Paragraph("It can build " + st.Bold("ordered"+" "+st.Italic("lists")) + ": ")
	s += st.OrderedList()(
		"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped. It should hopefully be so long that it will not fulfill the maximum number of authorized chars per line is reached.",
		"It also can support sub-lists:\n"+st.OrderedList()(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
			"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		),
		"It is also possible to mix with bullet list if you really want to.\n"+st.BulletedList()(
			"First things usually come first.",
			"Second things should come after the first ones.",
		),
	)

	s += st.Paragraph("It also knows how to " + st.Italic("define") + " terms:")
	s += st.Define("style", "A particular procedure by which something is done; a manner or way.")
	s += st.Paragraph("Even using Tabs:")
	s += st.Tab()(
		st.Define("lorem ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
	)

	s += H2("Demonstrating tables")
	s += st.Paragraph("Package 'style' supports drawing tables for most basic cases. Using Tab to align tables should be done carefully as table way to guess optimum columns size does not take into account the tabulation")
	s += st.Table(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Last but not least shows " + st.Bold("formating") + " within the table"},
		[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a second chance to verify that multi-lines is working"},
	)
	s += st.Paragraph("It is also possible to use tabs with tables:")
	s += st.Tab()(
		st.Table(
			[]string{"Column1", "Column2", "Column3"},
			[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Last but not least shows " + st.Bold("formating") + " within the table"},
			[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a second chance to verify that multi-lines is working"},
		),
	)

	return
}

func TestCoreSyntax(t *testing.T) {
	out := testText(&style.CoreSyntax{})
	if failure := verify.MatchGolden(t.Name(), out); failure != nil {
		t.Errorf("Styling with 'Core' style failed.\n%s", failure)
	}
}

func TestPlainTextSyntax(t *testing.T) {
	out := testText(style.NewPlaintext())
	if failure := verify.MatchGolden(t.Name(), out); failure != nil {
		t.Errorf("Styling with 'Plaintext' style failed.\n%s", failure)
	}
}

func TestTermSyntax(t *testing.T) {
	st := style.NewTerm()
	st.TextWidth = 60 //Fix size for testing purpose otherwise, might have varying resluts
	out := testText(st)
	if failure := verify.MatchGolden(t.Name(), out); failure != nil {
		t.Errorf("Styling with 'Term' style failed.\n%s", failure)
	}
}

func TestColorTermSyntax(t *testing.T) {
	st := style.NewColorterm()
	st.TextWidth = 60 //Fix size for testing purpose otherwise, might have varying resluts
	out := testText(st)
	if failure := verify.MatchGolden(t.Name(), out); failure != nil {
		t.Errorf("Styling with 'Colorterm' style failed.\n%s", failure)
	}
}

func TestMarkdownSyntax(t *testing.T) {
	out := testText(style.NewMarkdown())
	if failure := verify.MatchGolden(t.Name(), out); failure != nil {
		t.Errorf("Styling with 'Markdown' style failed.\n%s", failure)
	}
}

func TestManSyntax(t *testing.T) {
	out := testText(style.NewMan())
	if failure := verify.MatchGolden(t.Name(), out); failure != nil {
		t.Errorf("Styling with 'Man' style failed.\n%s", failure)
	}
}
