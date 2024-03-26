package api

import "regexp"

func isJsonValid(policy Policy) bool {
	var statements []Statement = policy.PolicyDocument.Statement
	for _, statement := range statements {
		if statement.Resource != nil {
			switch res := statement.Resource.(type) {
			case string:
				if regexp.MustCompile(`^[^*]*\*[^*]*$`).MatchString(res) {
					return false
				}
			case []string:
				for _, r := range res {
					if regexp.MustCompile(`^[^*]*\*[^*]*$`).MatchString(r) {
						return false
					}
				}
			}
		}
	}
	return true
}
