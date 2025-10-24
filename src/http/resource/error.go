package resource

import "gbase/src/core/http/request"

type ErrorValidation struct {
	Errors []request.ErrorBag `json:"errors"`
} //@name ErrorValidationResource
