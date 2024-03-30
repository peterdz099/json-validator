package messages

const (
	VALID     string = "JSON from file %s is VALID"
	NOT_VALID string = "JSON from file %s is NOT VALID"

	EMPTY_FIELDS_ERR     string = "invalid format: empty PolicyName and PolicyDocument fields"
	EMPTY_POLICYNAME_ERR string = "invalid format: empty PolicyName field"
	EMPTY_DOCUMENT_ERR   string = "invalid format: empty PolicyDocument field"
	EMPTY_VERSION_ERR    string = "invalid format: empty Version field"
	EMPTY_STATEMENT_ERR  string = "invalid format: empty Statement field"
	EMPTY_EFFECT_ERR     string = "invalid format: found empty Effect field"
	EMPTY_ACTION_ERR     string = "invalid format: found empty Action field"
	EMPTY_RESOURCE_ERR   string = "invalid format: found empty Resource field"

	INVALID_TYPE_ERR     string = "invalid type: found invalid field type"
	INVALID_VALUE_ERR    string = "invalid value: found invalid Effect value"
	INVALID_FORMAT_ERR   string = "invalid format: invalid JSON format"
	UNSUPPORTED_FILE_ERR string = "unsupported media type: wrong file extension"
	INTERNAL_ERR         string = "Internal Server Error - Generating Response"
)
