package comments

import (
	"bufio"
	"os"
	"regexp"
)

type CommentPatter struct {
	SingleLine *regexp.Regexp
	MultiLine  *regexp.Regexp
}

var predefinedPatterns = map[string]CommentPatter{
	".go": {
		SingleLine: regexp.MustCompile(`^\s*//`),
		MultiLine:  regexp.MustCompile(`(?s)/\*.*?\*/`),
	},
	".java": {
		SingleLine: regexp.MustCompile(`^\s*//`),
		MultiLine:  regexp.MustCompile(`(?s)/\*.*?\*/`),
	},
	".py": {
		SingleLine: regexp.MustCompile(`^\s*#`),
		MultiLine:  nil,
	},
}

func ExtractCommentsFromFile(filepath string, extension string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var comments []string
	scanner := bufio.NewScanner(file)

	patterns, ok := predefinedPatterns[extension]
	if !ok {
		return nil, nil
	}

	for scanner.Scan() {
		line := scanner.Text()

		if patterns.SingleLine != nil && patterns.SingleLine.MatchString(line) {
			comments = append(comments, line)
		}

		if patterns.MultiLine != nil {
			multilineComments := patterns.MultiLine.FindAllString(line, -1)
			if len(multilineComments) > 0 {
				comments = append(comments, multilineComments...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
