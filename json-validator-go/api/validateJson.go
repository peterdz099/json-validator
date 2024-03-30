package api

import (
	"fmt"
	"regexp"
)

func isJsonValid(policy Policy) bool {
	var statements []Statement = policy.PolicyDocument.Statement
	for _, statement := range statements {
		if statement.Resource != nil {
			switch res := statement.Resource.(type) {
			case string:
				if regexp.MustCompile(`\*`).MatchString(res) {
					return false
				}
			case []interface{}:
				for _, r := range res {
					if regexp.MustCompile(`\*`).MatchString(fmt.Sprint(r)) {
						return false
					}
				}
			}
		}
	}
	return true
}
