package style_test

import (
	"fmt"

	"github.com/pirmd/style"
)

func Example() {
	st := style.NewColorterm()
	//Can also be st := style.NewMarkdown()

	H := style.Chain(st.Header(1), st.Blue)
	H2 := style.Chain(st.Header(2), st.Red)

	s := H("Introduction")
	s += st.Paragraph("This small piece of text aims at demonstrating and testing package '" + st.Underline("style") + "'.")
	s += st.Paragraph("It is written by a " + st.Bold("non-native") + " English speaker, so pardon any faults.")

	s += H2("Demonstrating lists")
	s += st.Paragraph("It knows how to format " + st.Italic("lists") + ": ")
	s += st.BulletedList()(
		"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped. It should hopefully be so long that it will not fulfill the maximum number of authorized chars per line is reached.",
		"It also can support sub-lists:\n"+st.BulletedList()(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		),
	)

	s += H2("Demonstrating tables")
	s += st.Paragraph("Package 'style' supports drawing tables for most basic cases.")
	s += st.Table(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Last but not least shows " + st.Bold("formating") + " within the table"},
		[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a second chance to verify that multi-lines is working"},
	)

	fmt.Print(s)
}
