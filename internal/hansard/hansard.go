package hansard

type HansardMetaData struct {
	hasPossibleQuestionNo bool
	possibleQuestionNo    string
	questionNumberSnippet string
}

type HansardPage struct {
	pageNo            int
	isStartOfQuestion bool
	questionNo        int
	plainTextContent  string
	hansardMetaData   *HansardMetaData
}

func ExtractQuestionNo(snippet string) (int, error) {
	var questionNo int
	// What are the Guard checks ..?

	// If questionNo is still 0; something is wrong!!

	return questionNo, nil
}

func Split(t string, c string) []string {
	return []string{"bob"}
}
