package api

import (
	"errors"
	"reflect"
)

// this function checks if json contains all required fields
func verifyJsonFormat(policy Policy) (bool, error) {
	isDocumentValid, err := checkDocumentFormat(policy.PolicyDocument)
	if policy.PolicyName != "" {
		if isDocumentValid && err == nil {
			isStatementListValid, err := checkStatementList(policy.PolicyDocument.Statement)
			if !isStatementListValid && err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	} else if policy.PolicyName == "" {
		if !isDocumentValid {
			return false, errors.New("invalid format: invalid JSON format")
		}
		return false, errors.New("invalid format: empty PolicyName field")
	}
	return true, nil
}

func checkDocumentFormat(policyDocument PolicyDocument) (bool, error) {
	if policyDocument.Version == "" && len(policyDocument.Statement) == 0 {
		return false, errors.New("invalid format: empty PolicyDocument")
	} else if policyDocument.Version == "" {
		return false, errors.New("invalid format: empty Version field")
	} else if len(policyDocument.Statement) == 0 {
		return false, errors.New("invalid format: empty Statement field")
	}
	return true, nil
}

func checkStatementList(policyDocumentStatement []Statement) (bool, error) {
	for _, statement := range policyDocumentStatement {
		isStatmentValid, err := verifySingleStatement(statement)
		if !isStatmentValid && err != nil {
			return false, err
		}
	}
	return true, nil
}

// SID is optional field, so may be empty, type checked by decoder
func verifySingleStatement(statement Statement) (bool, error) {
	if statement.Effect == "" {
		return false, errors.New("invalid format: found empty Effect field")
	}

	isActionValid, err := verifyActionField(statement)
	if !isActionValid && err != nil {
		return false, err
	}

	isResourceValid, err := verifyResourceField(statement)
	if !isResourceValid && err != nil {
		return false, err
	}
	return true, nil
}

func verifyActionField(statement Statement) (bool, error) {
	if statement.Action != nil {
		switch action := statement.Action.(type) {
		case string:
			if action == "" {
				return false, errors.New("invalid format: found empty Action field")
			} else {
				return true, nil
			}
		case []interface{}:
			if len(action) > 0 {
				for _, a := range action {
					if reflect.TypeOf(a).Kind() != reflect.String {
						return false, errors.New("invalid format: found invalid Action type")
					}
				}
			} else {
				return false, errors.New("invalid format: found empty Action field")
			}

		default:
			return false, errors.New("invalid format: found invalid Action type")
		}
	} else {
		return false, errors.New("invalid format: Action in Statement is required")
	}
	return true, nil
}

func verifyResourceField(statement Statement) (bool, error) {
	if statement.Resource != nil {

		switch res := statement.Resource.(type) {
		case string:
			if res == "" {
				return false, errors.New("invalid format: found empty Resource field")
			}
		case []interface{}:
			if len(res) > 0 {
				for _, r := range res {
					if reflect.TypeOf(r).Kind() != reflect.String {
						return false, errors.New("invalid format: found invalid Resource type")
					}
				}
			} else {
				return false, errors.New("invalid format: found empty Resource field")
			}
		default:
			return false, errors.New("invalid format: found invalid Resource type")
		}
	} else {
		return false, errors.New("invalid format: Resource in Statement is required")
	}
	return true, nil
}
