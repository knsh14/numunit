package numunit

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "numunit checks int literal is grouped into 3 digits by _"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "numunit",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BasicLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.BasicLit:
			if n.Kind != token.INT {
				return
			}
			if len(n.Value) < 4 {
				return
			}
			if !strings.Contains(n.Value, "_") {
				pass.Reportf(n.Pos(), "use %s", ConvertLiteral(n.Value))
				return
			}
			if !checkLength(n.Value) {
				s := strings.ReplaceAll(n.Value, "_", "")
				pass.Reportf(n.Pos(), "use %s", ConvertLiteral(s))
				return
			}
		}
	})

	return nil, nil
}

func ConvertLiteral(original string) string {
	var s []string
	m := len(original) % 3
	if m != 0 {
		s = append(s, original[0:m])
	}
	for i := m; i < len(original)-1; i += 3 {
		s = append(s, original[i:i+3])
	}
	return strings.Join(s, "_")
}

func checkLength(s string) bool {
	v := strings.Split(s, "_")
	if len(v) < 2 {
		return true
	}
	for _, s := range v[1:] {
		if len(s) != 3 {
			return false
		}
	}
	return true
}
