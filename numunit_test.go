package numunit_test

import (
	"testing"

	"numunit"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	t.Parallel()
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, numunit.Analyzer, "a")
}

func TestConvert(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		input  string
		expect string
	}{
		{
			input:  "1234",
			expect: "1_234",
		},
		{
			input:  "123",
			expect: "123",
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			v := numunit.ConvertLiteral(tc.input)
			if v != tc.expect {
				t.Fatalf("ConvertLiteral(%s)[%s]!= %s", tc.input, v, tc.expect)
			}
		})
	}
}
