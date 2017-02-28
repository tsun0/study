// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitosync

const (

	// ErrCodeAlreadyStreamedException for service response error code
	// "AlreadyStreamedException".
	//
	// An exception thrown when a bulk publish operation is requested less than
	// 24 hours after a previous bulk publish operation completed successfully.
	ErrCodeAlreadyStreamedException = "AlreadyStreamedException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Thrown if there are parallel requests to modify a resource.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeDuplicateRequestException for service response error code
	// "DuplicateRequestException".
	//
	// An exception thrown when there is an IN_PROGRESS bulk publish operation for
	// the given identity pool.
	ErrCodeDuplicateRequestException = "DuplicateRequestException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// Indicates an internal service error.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidConfigurationException for service response error code
	// "InvalidConfigurationException".
	ErrCodeInvalidConfigurationException = "InvalidConfigurationException"

	// ErrCodeInvalidLambdaFunctionOutputException for service response error code
	// "InvalidLambdaFunctionOutputException".
	//
	// The AWS Lambda function returned invalid output or an exception.
	ErrCodeInvalidLambdaFunctionOutputException = "InvalidLambdaFunctionOutputException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// Thrown when a request parameter does not comply with the associated constraints.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeLambdaThrottledException for service response error code
	// "LambdaThrottledException".
	//
	// AWS Lambda throttled your account, please contact AWS Support
	ErrCodeLambdaThrottledException = "LambdaThrottledException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Thrown when the limit on the number of objects or operations has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// Thrown when a user is not authorized to access the requested resource.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// Thrown if an update can't be applied because the resource was changed by
	// another call and this would result in a conflict.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Thrown if the resource doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Thrown if the request is throttled.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)
