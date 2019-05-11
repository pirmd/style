// +build ignore

//This program generates helpers.go that exposes a set of helpers to `cli/style` end users
//It can be invoked through go generate directive
package main

import (
	"fmt"
	"os"
	"strings"

	"go/ast"
	"go/parser"
	"go/token"
)

const (
	//helpers to be generated for a given style
	//[1]: style name
	//[2]: style constant name
	helperForStyle = `
//%[1]s applies the style '%[1]s' to the formatted string directive.
//The given format and arguments follow fmt.Sprintf format.
func (st Styler) %[1]s(format string, a ...interface{}) string {
    return st.stylef(%[2]s)(format, a...)
}

//%[1]s applies the style '%[1]s' using current Styler.
//The given format and arguments follow fmt.Sprintf format
func %[1]s(format string, a ...interface{}) string {
    return CurrentStyler.%[1]s(format, a...)
}
`

	//filename to write to
	filename = "helpers.go"

	//all styles to iterate over are found in stylelist.go
	//and of constant of type 'style' starting by 'Fmt'
	styleSrc    = "stylelist.go"
	styleType   = "Format"
	stylePrefix = "Fmt"
)

//listStyles reads style.go for all style.style constant
func listStyles(src *ast.File) (ls []string) {

	ast.Inspect(src, func(node ast.Node) bool {
		decl, ok := node.(*ast.GenDecl)
		if !ok || decl.Tok != token.CONST {
			return true
		}

		//we remember the type of const familly declaration
		//that use a ioata multi-element scheme
		curType := ""

		for _, spec := range decl.Specs {
			vspec := spec.(*ast.ValueSpec)
			if vspec.Type == nil && len(vspec.Values) > 0 {
				// Constant with no type but with a value, we are no
				// more in a const ioata familly declaration, break
				curType = ""
				continue
			}

			if vspec.Type != nil {
				ident, ok := vspec.Type.(*ast.Ident)
				if !ok {
					continue //constant with no type definition
				}
				curType = ident.Name
			}

			if curType != styleType {
				continue //not our type
			}

			for _, name := range vspec.Names {
				if name.Name == "_" {
					continue
				}
				ls = append(ls, name.Name)
			}
		}

		return false
	})

	return
}

func main() {
	fset := token.NewFileSet()
	src, err := parser.ParseFile(fset, styleSrc, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Failed to parse %s: %v\n", styleSrc, err)
		os.Exit(1)
	}

	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	styles := listStyles(src)

	fmt.Fprintln(f, `//Code generated by 'go run helpers_generate.go'; DO NOT EDIT.`)
	fmt.Fprintln(f, `package style`)
	for _, st := range styles {
		fmt.Printf("Generate helpers for style '%s'\n", st)
		fmt.Fprintf(f, helperForStyle, strings.TrimPrefix(st, stylePrefix), st)
	}

	//generate a text/template.FuncMap
	fmt.Fprintf(f, "\n//FuncMap provides a text/template FuncMap compatible mapping\n")
	fmt.Fprintf(f, "//to use 'style' functions within templates\n")
	fmt.Fprintf(f, "func (st Styler) FuncMap() map[string]interface{} {\n")
	fmt.Fprintf(f, "    return map[string]interface{} {\n")
	for _, st := range styles {
		fmt.Fprintf(f, "       \"%[1]s\":  st.%[1]s,\n", strings.TrimPrefix(st, stylePrefix))
	}
	fmt.Fprintf(f, "    }\n")
	fmt.Fprintf(f, "}\n")

	//generate a text/template.FuncMap for CurrentStyler
	fmt.Fprintf(f, "\n//FuncMap provides a text/template FuncMap compatible mapping\n")
	fmt.Fprintf(f, "//to use 'style' CurrentStyler's functions within templates\n")
	fmt.Fprintf(f, "func FuncMap() map[string]interface{} {\n")
	fmt.Fprintf(f, "    return map[string]interface{} {\n")
	for _, st := range styles {
		fmt.Fprintf(f, "       \"%[1]s\":  %[1]s,\n", strings.TrimPrefix(st, stylePrefix))
	}
	fmt.Fprintf(f, "    }\n")
	fmt.Fprintf(f, "}\n")
}
