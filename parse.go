package engine

import (
	"bytes"
)

func Execute(query string, args any) (string, []any) {
	engine := NewEngine()

	tmpl, err := engine.tempalte.Parse(query)
	if err != nil {
		panic(err)
	}

	out := new(bytes.Buffer)
	if err := tmpl.Execute(out, args); err != nil {
		panic(err)
	}

	return out.String(), engine.parser.args
}
