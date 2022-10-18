package service

type Content struct {
	Name  string
	Words []string
}

var contentMap map[string]Content

func init() {
	contentMap = map[string]Content{
		"A": {"A", []string{"A"}},
	}
}

func GetWords(name string) []string {
	return contentMap[name].Words
}
