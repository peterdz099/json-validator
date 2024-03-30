package handlers

import (
	"errors"
	"reflect"

	"json-validator/internal/messages"
)

/*
This function checks that the JSON contains all required fields,
that Statement.Effect has a valid value, and that the fields declared
as interface{} (Statement.Resource, Statement.Action) are of the specified type,
since these cases are not verified by the JSON decoder.
*/
func verifyJsonFormat(policy Policy) (bool, error) {
	isDocumentValid, err := verifyDocumentFormat(policy.PolicyDocument)
	if policy.PolicyName != "" {
		if isDocumentValid && err == nil {
			isStatementListValid, err := verifyStatementList(policy.PolicyDocument.Statement)
			if !isStatementListValid && err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	} else if policy.PolicyName == "" {
		if !isDocumentValid {
			return false, errors.New(messages.EMPTY_FIELDS_ERR)

		}
		return false, errors.New(messages.EMPTY_POLICYNAME_ERR)
	}
	return true, nil
}

func verifyDocumentFormat(policyDocument PolicyDocument) (bool, error) {
	if policyDocument.Version == "" && len(policyDocument.Statement) == 0 {
		return false, errors.New(messages.EMPTY_DOCUMENT_ERR)

	} else if policyDocument.Version == "" {
		return false, errors.New(messages.EMPTY_VERSION_ERR)

	} else if len(policyDocument.Statement) == 0 {
		return false, errors.New(messages.EMPTY_STATEMENT_ERR)

	}
	return true, nil
}

func verifyStatementList(policyDocumentStatement []Statement) (bool, error) {
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
	isEffectValid, err := verifyEffectField(statement)
	if !isEffectValid && err != nil {
		return false, err
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

func verifyEffectField(statement Statement) (bool, error) {
	if statement.Effect == "" {
		return false, errors.New(messages.EMPTY_EFFECT_ERR)

	} else if statement.Effect != "Allow" && statement.Effect != "Deny" {
		return false, errors.New(messages.INVALID_VALUE_ERR)

	} else {
		return true, nil
	}
}

func verifyActionField(statement Statement) (bool, error) {
	if statement.Action != nil {
		switch action := statement.Action.(type) {
		case string:
			if action == "" {
				return false, errors.New(messages.EMPTY_ACTION_ERR)

			} else {
				return true, nil
			}
		case []interface{}:
			if len(action) > 0 {
				for _, a := range action {
					if reflect.TypeOf(a).Kind() != reflect.String {
						return false, errors.New((messages.INVALID_TYPE_ERR))
					}
				}
			} else {
				return false, errors.New(messages.EMPTY_ACTION_ERR)
			}

		default:
			return false, errors.New((messages.INVALID_TYPE_ERR))
		}
	} else {
		return false, errors.New(messages.EMPTY_ACTION_ERR)
	}

	return true, nil
}

func verifyResourceField(statement Statement) (bool, error) {
	if statement.Resource != nil {
		switch res := statement.Resource.(type) {
		case string:
			if res == "" {
				return false, errors.New(messages.EMPTY_RESOURCE_ERR)
			}
		case []interface{}:
			if len(res) > 0 {
				for _, r := range res {
					if reflect.TypeOf(r).Kind() != reflect.String {
						return false, errors.New((messages.INVALID_TYPE_ERR))
					}
				}
			} else {
				return false, errors.New(messages.EMPTY_RESOURCE_ERR)
			}
		default:
			return false, errors.New(messages.INVALID_TYPE_ERR)
		}
	} else {
		return false, errors.New(messages.EMPTY_RESOURCE_ERR)
	}

	return true, nil
}
