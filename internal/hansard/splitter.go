package hansard

import "fmt"

type SplitPlan struct {
	QuestionNum  string
	PageNumStart int
	PageNumEnd   int
}

// NewSplitPlan will use a Reader (better!) to extract out the plan
func NewSplitPlan(planFilename string) []SplitPlan {
	// TODO: Read the plan file
	return nil
}

// NewMockSplitPlan returns
// 	portion of actual test case file PDF
func NewMockSplitPlan() []SplitPlan {
	return []SplitPlan{
		{"1", 2, 3},
		{"2", 4, 4},
		{"3", 5, 6},
		{"4", 7, 8},
		{"5", 9, 9},
		{"6", 10, 12},
	}
}

func (sp *SplitPlan) ExecuteSplit() {
	outputFilename := fmt.Sprintf("soalan-%s-bukanlisan", sp.QuestionNum)
	fmt.Println("====== ", outputFilename, " =======")
}
