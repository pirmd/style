# PACKAGE STYLE
[![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.org/github.com/pirmd/style)&nbsp; 
[![Go Report Card](https://goreportcard.com/badge/github.com/pirmd/style)](https://goreportcard.com/report/github.com/pirmd/style)&nbsp;

`style` provides functions that decorate text (like formatting, changing case
or colors) using several idioms (plain text, (non-)colored term,  mandoc,
markdown). Proposed styles are supposed to be easily extendable.

## EXAMPLE

```go
package main

import (
        "fmt"

        "github.com/pirmd/style"
       )

func Main() {
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
```

A slightly more complete example can be found in [testing file](style_test.go) or in 
[clapp](https://github.com/pirmd/clapp) package's [help](https://github.com/pirmd/clapp/blob/master/help.go)
or [manpage](https://github.com/pirmd/clapp/blob/master/manpage.go) files generation.

## INSTALLATION
Everything should work fine using go standard commands (`build`, `get`,
`install`...).

## USAGE
Running `go doc github.com/pirmd/style` should give you helpful guidelines on
available features.

## CONTRIBUTION
If you feel like to contribute, just follow github guidelines on
[forking](https://help.github.com/articles/fork-a-repo/) then [send a pull
request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
