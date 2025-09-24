// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package a2aexp

import (
	"encoding/json"

	"github.com/stainless-sdks/a2a-exp-go/internal/apijson"
	"github.com/stainless-sdks/a2a-exp-go/packages/param"
	"github.com/stainless-sdks/a2a-exp-go/packages/respjson"
	"github.com/stainless-sdks/a2a-exp-go/shared"
	"github.com/stainless-sdks/a2a-exp-go/shared/constant"
)

// InvokeResponseUnion contains all possible properties and values from
// [InvokeResponseJsonrpcErrorResponse],
// [InvokeResponseSendMessageSuccessResponse],
// [InvokeResponseSendStreamingMessageSuccessResponse],
// [InvokeResponseGetTaskSuccessResponse],
// [InvokeResponseCancelTaskSuccessResponse],
// [InvokeResponseSetTaskPushNotificationConfigSuccessResponse],
// [InvokeResponseGetTaskPushNotificationConfigSuccessResponse],
// [InvokeResponseListTaskPushNotificationConfigSuccessResponse],
// [InvokeResponseDeleteTaskPushNotificationConfigSuccessResponse],
// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InvokeResponseUnion struct {
	// This field is a union of [InvokeResponseJsonrpcErrorResponseIDUnion],
	// [InvokeResponseSendMessageSuccessResponseIDUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseIDUnion],
	// [InvokeResponseGetTaskSuccessResponseIDUnion],
	// [InvokeResponseCancelTaskSuccessResponseIDUnion],
	// [InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion],
	// [InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion],
	// [InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion],
	// [InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion],
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion]
	ID InvokeResponseUnionID `json:"id"`
	// This field is from variant [InvokeResponseJsonrpcErrorResponse].
	Error InvokeResponseJsonrpcErrorResponseErrorUnion `json:"error"`
	// This field is from variant [InvokeResponseJsonrpcErrorResponse].
	Jsonrpc constant.String2_0 `json:"jsonrpc"`
	// This field is a union of [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion], [shared.Task],
	// [shared.TaskPushNotificationConfig], [[]shared.TaskPushNotificationConfig],
	// [any], [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult]
	Result InvokeResponseUnionResult `json:"result"`
	JSON   struct {
		ID      respjson.Field
		Error   respjson.Field
		Jsonrpc respjson.Field
		Result  respjson.Field
		raw     string
	} `json:"-"`
}

func (u InvokeResponseUnion) AsInvokeResponseJsonrpcErrorResponse() (v InvokeResponseJsonrpcErrorResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseSendMessageSuccessResponse() (v InvokeResponseSendMessageSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseSendStreamingMessageSuccessResponse() (v InvokeResponseSendStreamingMessageSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseGetTaskSuccessResponse() (v InvokeResponseGetTaskSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseCancelTaskSuccessResponse() (v InvokeResponseCancelTaskSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseSetTaskPushNotificationConfigSuccessResponse() (v InvokeResponseSetTaskPushNotificationConfigSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseGetTaskPushNotificationConfigSuccessResponse() (v InvokeResponseGetTaskPushNotificationConfigSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseListTaskPushNotificationConfigSuccessResponse() (v InvokeResponseListTaskPushNotificationConfigSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseDeleteTaskPushNotificationConfigSuccessResponse() (v InvokeResponseDeleteTaskPushNotificationConfigSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseUnion) AsInvokeResponseGetAuthenticatedExtendedCardSuccessResponse() (v InvokeResponseGetAuthenticatedExtendedCardSuccessResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseUnionID is an implicit subunion of [InvokeResponseUnion].
// InvokeResponseUnionID provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [InvokeResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseUnionID struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (r *InvokeResponseUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseUnionResult is an implicit subunion of [InvokeResponseUnion].
// InvokeResponseUnionResult provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [InvokeResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTaskPushNotificationConfigArray
// OfInvokeResponseDeleteTaskPushNotificationConfigSuccessResponseResult]
type InvokeResponseUnionResult struct {
	// This field will be present if the value is a
	// [[]shared.TaskPushNotificationConfig] instead of an object.
	OfTaskPushNotificationConfigArray []shared.TaskPushNotificationConfig `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfInvokeResponseDeleteTaskPushNotificationConfigSuccessResponseResult any `json:",inline"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion], [shared.Task].
	ID        string `json:"id"`
	ContextID string `json:"contextId"`
	Kind      string `json:"kind"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion], [shared.Task].
	Status shared.TaskStatus `json:"status"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion], [shared.Task].
	Artifacts []shared.Artifact `json:"artifacts"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion], [shared.Task].
	History  []shared.Message `json:"history"`
	Metadata any              `json:"metadata"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	MessageID string `json:"messageId"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	Parts []shared.PartUnion `json:"parts"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	Role shared.MessageRole `json:"role"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	Extensions []string `json:"extensions"`
	// This field is from variant
	// [InvokeResponseSendMessageSuccessResponseResultUnion],
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	ReferenceTaskIDs []string `json:"referenceTaskIds"`
	TaskID           string   `json:"taskId"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	Final bool `json:"final"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	Artifact shared.Artifact `json:"artifact"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	Append bool `json:"append"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultUnion].
	LastChunk bool `json:"lastChunk"`
	// This field is from variant [shared.TaskPushNotificationConfig].
	PushNotificationConfig shared.PushNotificationConfig `json:"pushNotificationConfig"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Capabilities InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilities `json:"capabilities"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	DefaultInputModes []string `json:"defaultInputModes"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	DefaultOutputModes []string `json:"defaultOutputModes"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Description string `json:"description"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Name string `json:"name"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	ProtocolVersion string `json:"protocolVersion"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Skills []InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSkill `json:"skills"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	URL string `json:"url"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Version string `json:"version"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	AdditionalInterfaces []InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultAdditionalInterface `json:"additionalInterfaces"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	DocumentationURL string `json:"documentationUrl"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	IconURL string `json:"iconUrl"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	PreferredTransport string `json:"preferredTransport"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Provider InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultProvider `json:"provider"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Security []map[string][]string `json:"security"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	SecuritySchemes map[string]InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion `json:"securitySchemes"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	Signatures []InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSignature `json:"signatures"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult].
	SupportsAuthenticatedExtendedCard bool `json:"supportsAuthenticatedExtendedCard"`
	JSON                              struct {
		OfTaskPushNotificationConfigArray                                     respjson.Field
		OfInvokeResponseDeleteTaskPushNotificationConfigSuccessResponseResult respjson.Field
		ID                                                                    respjson.Field
		ContextID                                                             respjson.Field
		Kind                                                                  respjson.Field
		Status                                                                respjson.Field
		Artifacts                                                             respjson.Field
		History                                                               respjson.Field
		Metadata                                                              respjson.Field
		MessageID                                                             respjson.Field
		Parts                                                                 respjson.Field
		Role                                                                  respjson.Field
		Extensions                                                            respjson.Field
		ReferenceTaskIDs                                                      respjson.Field
		TaskID                                                                respjson.Field
		Final                                                                 respjson.Field
		Artifact                                                              respjson.Field
		Append                                                                respjson.Field
		LastChunk                                                             respjson.Field
		PushNotificationConfig                                                respjson.Field
		Capabilities                                                          respjson.Field
		DefaultInputModes                                                     respjson.Field
		DefaultOutputModes                                                    respjson.Field
		Description                                                           respjson.Field
		Name                                                                  respjson.Field
		ProtocolVersion                                                       respjson.Field
		Skills                                                                respjson.Field
		URL                                                                   respjson.Field
		Version                                                               respjson.Field
		AdditionalInterfaces                                                  respjson.Field
		DocumentationURL                                                      respjson.Field
		IconURL                                                               respjson.Field
		PreferredTransport                                                    respjson.Field
		Provider                                                              respjson.Field
		Security                                                              respjson.Field
		SecuritySchemes                                                       respjson.Field
		Signatures                                                            respjson.Field
		SupportsAuthenticatedExtendedCard                                     respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (r *InvokeResponseUnionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a JSON-RPC 2.0 Error Response object.
type InvokeResponseJsonrpcErrorResponse struct {
	// The identifier established by the client.
	ID InvokeResponseJsonrpcErrorResponseIDUnion `json:"id,required"`
	// An object describing the error that occurred.
	Error InvokeResponseJsonrpcErrorResponseErrorUnion `json:"error,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Error       respjson.Field
		Jsonrpc     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponse) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseJsonrpcErrorResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseJsonrpcErrorResponseIDUnion contains all possible properties and
// values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseJsonrpcErrorResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseJsonrpcErrorResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseJsonrpcErrorResponseIDUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseJsonrpcErrorResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseJsonrpcErrorResponseErrorUnion contains all possible properties
// and values from [InvokeResponseJsonrpcErrorResponseErrorJsonrpcError],
// [InvokeResponseJsonrpcErrorResponseErrorJsonParseError],
// [InvokeResponseJsonrpcErrorResponseErrorInvalidRequestError],
// [InvokeResponseJsonrpcErrorResponseErrorMethodNotFoundError],
// [InvokeResponseJsonrpcErrorResponseErrorInvalidParamsError],
// [InvokeResponseJsonrpcErrorResponseErrorInternalError],
// [InvokeResponseJsonrpcErrorResponseErrorTaskNotFoundError],
// [InvokeResponseJsonrpcErrorResponseErrorTaskNotCancelableError],
// [InvokeResponseJsonrpcErrorResponseErrorPushNotificationNotSupportedError],
// [InvokeResponseJsonrpcErrorResponseErrorUnsupportedOperationError],
// [InvokeResponseJsonrpcErrorResponseErrorContentTypeNotSupportedError],
// [InvokeResponseJsonrpcErrorResponseErrorInvalidAgentResponseError],
// [InvokeResponseJsonrpcErrorResponseErrorAuthenticatedExtendedCardNotConfiguredError].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InvokeResponseJsonrpcErrorResponseErrorUnion struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	JSON    struct {
		Code    respjson.Field
		Message respjson.Field
		Data    respjson.Field
		raw     string
	} `json:"-"`
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorJsonrpcError() (v InvokeResponseJsonrpcErrorResponseErrorJsonrpcError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorJsonParseError() (v InvokeResponseJsonrpcErrorResponseErrorJsonParseError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorInvalidRequestError() (v InvokeResponseJsonrpcErrorResponseErrorInvalidRequestError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorMethodNotFoundError() (v InvokeResponseJsonrpcErrorResponseErrorMethodNotFoundError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorInvalidParamsError() (v InvokeResponseJsonrpcErrorResponseErrorInvalidParamsError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorInternalError() (v InvokeResponseJsonrpcErrorResponseErrorInternalError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorTaskNotFoundError() (v InvokeResponseJsonrpcErrorResponseErrorTaskNotFoundError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorTaskNotCancelableError() (v InvokeResponseJsonrpcErrorResponseErrorTaskNotCancelableError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorPushNotificationNotSupportedError() (v InvokeResponseJsonrpcErrorResponseErrorPushNotificationNotSupportedError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorUnsupportedOperationError() (v InvokeResponseJsonrpcErrorResponseErrorUnsupportedOperationError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorContentTypeNotSupportedError() (v InvokeResponseJsonrpcErrorResponseErrorContentTypeNotSupportedError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorInvalidAgentResponseError() (v InvokeResponseJsonrpcErrorResponseErrorInvalidAgentResponseError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseJsonrpcErrorResponseErrorUnion) AsInvokeResponseJsonrpcErrorResponseErrorAuthenticatedExtendedCardNotConfiguredError() (v InvokeResponseJsonrpcErrorResponseErrorAuthenticatedExtendedCardNotConfiguredError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseJsonrpcErrorResponseErrorUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseJsonrpcErrorResponseErrorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a JSON-RPC 2.0 Error object, included in an error response.
type InvokeResponseJsonrpcErrorResponseErrorJsonrpcError struct {
	// A number that indicates the error type that occurred.
	Code int64 `json:"code,required"`
	// A string providing a short description of the error.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorJsonrpcError) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseJsonrpcErrorResponseErrorJsonrpcError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error indicating that the server received invalid JSON.
type InvokeResponseJsonrpcErrorResponseErrorJsonParseError struct {
	// The error code for a JSON parse error.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorJsonParseError) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseJsonrpcErrorResponseErrorJsonParseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error indicating that the JSON sent is not a valid Request object.
type InvokeResponseJsonrpcErrorResponseErrorInvalidRequestError struct {
	// The error code for an invalid request.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorInvalidRequestError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorInvalidRequestError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error indicating that the requested method does not exist or is not
// available.
type InvokeResponseJsonrpcErrorResponseErrorMethodNotFoundError struct {
	// The error code for a method not found error.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorMethodNotFoundError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorMethodNotFoundError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error indicating that the method parameters are invalid.
type InvokeResponseJsonrpcErrorResponseErrorInvalidParamsError struct {
	// The error code for an invalid parameters error.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorInvalidParamsError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorInvalidParamsError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error indicating an internal error on the server.
type InvokeResponseJsonrpcErrorResponseErrorInternalError struct {
	// The error code for an internal server error.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorInternalError) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseJsonrpcErrorResponseErrorInternalError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An A2A-specific error indicating that the requested task ID was not found.
type InvokeResponseJsonrpcErrorResponseErrorTaskNotFoundError struct {
	// The error code for a task not found error.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorTaskNotFoundError) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseJsonrpcErrorResponseErrorTaskNotFoundError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An A2A-specific error indicating that the task is in a state where it cannot be
// canceled.
type InvokeResponseJsonrpcErrorResponseErrorTaskNotCancelableError struct {
	// The error code for a task that cannot be canceled.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorTaskNotCancelableError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorTaskNotCancelableError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An A2A-specific error indicating that the agent does not support push
// notifications.
type InvokeResponseJsonrpcErrorResponseErrorPushNotificationNotSupportedError struct {
	// The error code for when push notifications are not supported.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorPushNotificationNotSupportedError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorPushNotificationNotSupportedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An A2A-specific error indicating that the requested operation is not supported
// by the agent.
type InvokeResponseJsonrpcErrorResponseErrorUnsupportedOperationError struct {
	// The error code for an unsupported operation.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorUnsupportedOperationError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorUnsupportedOperationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An A2A-specific error indicating an incompatibility between the requested
// content types and the agent's capabilities.
type InvokeResponseJsonrpcErrorResponseErrorContentTypeNotSupportedError struct {
	// The error code for an unsupported content type.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorContentTypeNotSupportedError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorContentTypeNotSupportedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An A2A-specific error indicating that the agent returned a response that does
// not conform to the specification for the current method.
type InvokeResponseJsonrpcErrorResponseErrorInvalidAgentResponseError struct {
	// The error code for an invalid agent response.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorInvalidAgentResponseError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorInvalidAgentResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An A2A-specific error indicating that the agent does not have an Authenticated
// Extended Card configured
type InvokeResponseJsonrpcErrorResponseErrorAuthenticatedExtendedCardNotConfiguredError struct {
	// The error code for when an authenticated extended card is not configured.
	Code int64 `json:"code,required"`
	// The error message.
	Message string `json:"message,required"`
	// A primitive or structured value containing additional information about the
	// error. This may be omitted.
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseJsonrpcErrorResponseErrorAuthenticatedExtendedCardNotConfiguredError) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseJsonrpcErrorResponseErrorAuthenticatedExtendedCardNotConfiguredError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the `message/send` method.
type InvokeResponseSendMessageSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseSendMessageSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result, which can be a direct reply Message or the initial Task object.
	Result InvokeResponseSendMessageSuccessResponseResultUnion `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseSendMessageSuccessResponse) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseSendMessageSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseSendMessageSuccessResponseIDUnion contains all possible properties
// and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseSendMessageSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseSendMessageSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseSendMessageSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseSendMessageSuccessResponseIDUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseSendMessageSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseSendMessageSuccessResponseResultUnion contains all possible
// properties and values from [shared.Task], [shared.Message].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InvokeResponseSendMessageSuccessResponseResultUnion struct {
	// This field is from variant [shared.Task].
	ID        string `json:"id"`
	ContextID string `json:"contextId"`
	Kind      string `json:"kind"`
	// This field is from variant [shared.Task].
	Status shared.TaskStatus `json:"status"`
	// This field is from variant [shared.Task].
	Artifacts []shared.Artifact `json:"artifacts"`
	// This field is from variant [shared.Task].
	History  []shared.Message `json:"history"`
	Metadata any              `json:"metadata"`
	// This field is from variant [shared.Message].
	MessageID string `json:"messageId"`
	// This field is from variant [shared.Message].
	Parts []shared.PartUnion `json:"parts"`
	// This field is from variant [shared.Message].
	Role shared.MessageRole `json:"role"`
	// This field is from variant [shared.Message].
	Extensions []string `json:"extensions"`
	// This field is from variant [shared.Message].
	ReferenceTaskIDs []string `json:"referenceTaskIds"`
	// This field is from variant [shared.Message].
	TaskID string `json:"taskId"`
	JSON   struct {
		ID               respjson.Field
		ContextID        respjson.Field
		Kind             respjson.Field
		Status           respjson.Field
		Artifacts        respjson.Field
		History          respjson.Field
		Metadata         respjson.Field
		MessageID        respjson.Field
		Parts            respjson.Field
		Role             respjson.Field
		Extensions       respjson.Field
		ReferenceTaskIDs respjson.Field
		TaskID           respjson.Field
		raw              string
	} `json:"-"`
}

func (u InvokeResponseSendMessageSuccessResponseResultUnion) AsTask() (v shared.Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseSendMessageSuccessResponseResultUnion) AsMessage() (v shared.Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseSendMessageSuccessResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseSendMessageSuccessResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the `message/stream` method. The
// server may send multiple response objects for a single request.
type InvokeResponseSendStreamingMessageSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseSendStreamingMessageSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result, which can be a Message, Task, or a streaming update event.
	Result InvokeResponseSendStreamingMessageSuccessResponseResultUnion `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseSendStreamingMessageSuccessResponse) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseSendStreamingMessageSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseSendStreamingMessageSuccessResponseIDUnion contains all possible
// properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseSendStreamingMessageSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseSendStreamingMessageSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseSendStreamingMessageSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseSendStreamingMessageSuccessResponseIDUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseSendStreamingMessageSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseSendStreamingMessageSuccessResponseResultUnion contains all
// possible properties and values from [shared.Task], [shared.Message],
// [InvokeResponseSendStreamingMessageSuccessResponseResultTaskStatusUpdateEvent],
// [InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InvokeResponseSendStreamingMessageSuccessResponseResultUnion struct {
	// This field is from variant [shared.Task].
	ID        string `json:"id"`
	ContextID string `json:"contextId"`
	Kind      string `json:"kind"`
	// This field is from variant [shared.Task].
	Status shared.TaskStatus `json:"status"`
	// This field is from variant [shared.Task].
	Artifacts []shared.Artifact `json:"artifacts"`
	// This field is from variant [shared.Task].
	History  []shared.Message `json:"history"`
	Metadata any              `json:"metadata"`
	// This field is from variant [shared.Message].
	MessageID string `json:"messageId"`
	// This field is from variant [shared.Message].
	Parts []shared.PartUnion `json:"parts"`
	// This field is from variant [shared.Message].
	Role shared.MessageRole `json:"role"`
	// This field is from variant [shared.Message].
	Extensions []string `json:"extensions"`
	// This field is from variant [shared.Message].
	ReferenceTaskIDs []string `json:"referenceTaskIds"`
	TaskID           string   `json:"taskId"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultTaskStatusUpdateEvent].
	Final bool `json:"final"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent].
	Artifact shared.Artifact `json:"artifact"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent].
	Append bool `json:"append"`
	// This field is from variant
	// [InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent].
	LastChunk bool `json:"lastChunk"`
	JSON      struct {
		ID               respjson.Field
		ContextID        respjson.Field
		Kind             respjson.Field
		Status           respjson.Field
		Artifacts        respjson.Field
		History          respjson.Field
		Metadata         respjson.Field
		MessageID        respjson.Field
		Parts            respjson.Field
		Role             respjson.Field
		Extensions       respjson.Field
		ReferenceTaskIDs respjson.Field
		TaskID           respjson.Field
		Final            respjson.Field
		Artifact         respjson.Field
		Append           respjson.Field
		LastChunk        respjson.Field
		raw              string
	} `json:"-"`
}

func (u InvokeResponseSendStreamingMessageSuccessResponseResultUnion) AsTask() (v shared.Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseSendStreamingMessageSuccessResponseResultUnion) AsMessage() (v shared.Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseSendStreamingMessageSuccessResponseResultUnion) AsInvokeResponseSendStreamingMessageSuccessResponseResultTaskStatusUpdateEvent() (v InvokeResponseSendStreamingMessageSuccessResponseResultTaskStatusUpdateEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseSendStreamingMessageSuccessResponseResultUnion) AsInvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent() (v InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseSendStreamingMessageSuccessResponseResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InvokeResponseSendStreamingMessageSuccessResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An event sent by the agent to notify the client of a change in a task's status.
// This is typically used in streaming or subscription models.
type InvokeResponseSendStreamingMessageSuccessResponseResultTaskStatusUpdateEvent struct {
	// The context ID associated with the task.
	ContextID string `json:"contextId,required"`
	// If true, this is the final event in the stream for this interaction.
	Final bool `json:"final,required"`
	// The type of this event, used as a discriminator. Always 'status-update'.
	Kind constant.StatusUpdate `json:"kind,required"`
	// The new status of the task.
	Status shared.TaskStatus `json:"status,required"`
	// The ID of the task that was updated.
	TaskID string `json:"taskId,required"`
	// Optional metadata for extensions.
	Metadata map[string]any `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContextID   respjson.Field
		Final       respjson.Field
		Kind        respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseSendStreamingMessageSuccessResponseResultTaskStatusUpdateEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseSendStreamingMessageSuccessResponseResultTaskStatusUpdateEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An event sent by the agent to notify the client that an artifact has been
// generated or updated. This is typically used in streaming models.
type InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent struct {
	// The artifact that was generated or updated.
	Artifact shared.Artifact `json:"artifact,required"`
	// The context ID associated with the task.
	ContextID string `json:"contextId,required"`
	// The type of this event, used as a discriminator. Always 'artifact-update'.
	Kind constant.ArtifactUpdate `json:"kind,required"`
	// The ID of the task this artifact belongs to.
	TaskID string `json:"taskId,required"`
	// If true, the content of this artifact should be appended to a previously sent
	// artifact with the same ID.
	Append bool `json:"append"`
	// If true, this is the final chunk of the artifact.
	LastChunk bool `json:"lastChunk"`
	// Optional metadata for extensions.
	Metadata map[string]any `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Artifact    respjson.Field
		ContextID   respjson.Field
		Kind        respjson.Field
		TaskID      respjson.Field
		Append      respjson.Field
		LastChunk   respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseSendStreamingMessageSuccessResponseResultTaskArtifactUpdateEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the `tasks/get` method.
type InvokeResponseGetTaskSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseGetTaskSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result, containing the requested Task object.
	Result shared.Task `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetTaskSuccessResponse) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseGetTaskSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseGetTaskSuccessResponseIDUnion contains all possible properties and
// values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseGetTaskSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseGetTaskSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseGetTaskSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseGetTaskSuccessResponseIDUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseGetTaskSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the `tasks/cancel` method.
type InvokeResponseCancelTaskSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseCancelTaskSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result, containing the final state of the canceled Task object.
	Result shared.Task `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseCancelTaskSuccessResponse) RawJSON() string { return r.JSON.raw }
func (r *InvokeResponseCancelTaskSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseCancelTaskSuccessResponseIDUnion contains all possible properties
// and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseCancelTaskSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseCancelTaskSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseCancelTaskSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseCancelTaskSuccessResponseIDUnion) RawJSON() string { return u.JSON.raw }

func (r *InvokeResponseCancelTaskSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the
// `tasks/pushNotificationConfig/set` method.
type InvokeResponseSetTaskPushNotificationConfigSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result, containing the configured push notification settings.
	Result shared.TaskPushNotificationConfig `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseSetTaskPushNotificationConfigSuccessResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseSetTaskPushNotificationConfigSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InvokeResponseSetTaskPushNotificationConfigSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the
// `tasks/pushNotificationConfig/get` method.
type InvokeResponseGetTaskPushNotificationConfigSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result, containing the requested push notification configuration.
	Result shared.TaskPushNotificationConfig `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetTaskPushNotificationConfigSuccessResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetTaskPushNotificationConfigSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InvokeResponseGetTaskPushNotificationConfigSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the
// `tasks/pushNotificationConfig/list` method.
type InvokeResponseListTaskPushNotificationConfigSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result, containing an array of all push notification configurations for the
	// task.
	Result []shared.TaskPushNotificationConfig `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseListTaskPushNotificationConfigSuccessResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseListTaskPushNotificationConfigSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InvokeResponseListTaskPushNotificationConfigSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the
// `tasks/pushNotificationConfig/delete` method.
type InvokeResponseDeleteTaskPushNotificationConfigSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result is null on successful deletion.
	Result any `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseDeleteTaskPushNotificationConfigSuccessResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseDeleteTaskPushNotificationConfigSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion contains
// all possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InvokeResponseDeleteTaskPushNotificationConfigSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a successful JSON-RPC response for the
// `agent/getAuthenticatedExtendedCard` method.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponse struct {
	// The identifier established by the client.
	ID InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion `json:"id,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The result is an Agent Card object.
	Result InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Jsonrpc     respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion contains all
// possible properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result is an Agent Card object.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult struct {
	// A declaration of optional capabilities supported by the agent.
	Capabilities InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilities `json:"capabilities,required"`
	// Default set of supported input MIME types for all skills, which can be
	// overridden on a per-skill basis.
	DefaultInputModes []string `json:"defaultInputModes,required"`
	// Default set of supported output MIME types for all skills, which can be
	// overridden on a per-skill basis.
	DefaultOutputModes []string `json:"defaultOutputModes,required"`
	// A human-readable description of the agent, assisting users and other agents in
	// understanding its purpose.
	Description string `json:"description,required"`
	// A human-readable name for the agent.
	Name string `json:"name,required"`
	// The version of the A2A protocol this agent supports.
	ProtocolVersion string `json:"protocolVersion,required"`
	// The set of skills, or distinct capabilities, that the agent can perform.
	Skills []InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSkill `json:"skills,required"`
	// The preferred endpoint URL for interacting with the agent. This URL MUST support
	// the transport specified by 'preferredTransport'.
	URL string `json:"url,required"`
	// The agent's own version number. The format is defined by the provider.
	Version string `json:"version,required"`
	// A list of additional supported interfaces (transport and URL combinations). This
	// allows agents to expose multiple transports, potentially at different URLs.
	//
	// Best practices:
	//
	// - SHOULD include all supported transports for completeness
	// - SHOULD include an entry matching the main 'url' and 'preferredTransport'
	// - MAY reuse URLs if multiple transports are available at the same endpoint
	// - MUST accurately declare the transport available at each URL
	//
	// Clients can select any interface from this list based on their transport
	// capabilities and preferences. This enables transport negotiation and fallback
	// scenarios.
	AdditionalInterfaces []InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultAdditionalInterface `json:"additionalInterfaces"`
	// An optional URL to the agent's documentation.
	DocumentationURL string `json:"documentationUrl"`
	// An optional URL to an icon for the agent.
	IconURL string `json:"iconUrl"`
	// The transport protocol for the preferred endpoint (the main 'url' field). If not
	// specified, defaults to 'JSONRPC'.
	//
	// IMPORTANT: The transport specified here MUST be available at the main 'url'.
	// This creates a binding between the main URL and its supported transport
	// protocol. Clients should prefer this transport and URL combination when both are
	// supported.
	PreferredTransport string `json:"preferredTransport"`
	// Information about the agent's service provider.
	Provider InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultProvider `json:"provider"`
	// A list of security requirement objects that apply to all agent interactions.
	// Each object lists security schemes that can be used. Follows the OpenAPI 3.0
	// Security Requirement Object. This list can be seen as an OR of ANDs. Each object
	// in the list describes one possible set of security requirements that must be
	// present on a request. This allows specifying, for example, "callers must either
	// use OAuth OR an API Key AND mTLS."
	Security []map[string][]string `json:"security"`
	// A declaration of the security schemes available to authorize requests. The key
	// is the scheme name. Follows the OpenAPI 3.0 Security Scheme Object.
	SecuritySchemes map[string]InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion `json:"securitySchemes"`
	// JSON Web Signatures computed for this AgentCard.
	Signatures []InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSignature `json:"signatures"`
	// If true, the agent can provide an extended agent card with additional details to
	// authenticated users. Defaults to false.
	SupportsAuthenticatedExtendedCard bool `json:"supportsAuthenticatedExtendedCard"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities                      respjson.Field
		DefaultInputModes                 respjson.Field
		DefaultOutputModes                respjson.Field
		Description                       respjson.Field
		Name                              respjson.Field
		ProtocolVersion                   respjson.Field
		Skills                            respjson.Field
		URL                               respjson.Field
		Version                           respjson.Field
		AdditionalInterfaces              respjson.Field
		DocumentationURL                  respjson.Field
		IconURL                           respjson.Field
		PreferredTransport                respjson.Field
		Provider                          respjson.Field
		Security                          respjson.Field
		SecuritySchemes                   respjson.Field
		Signatures                        respjson.Field
		SupportsAuthenticatedExtendedCard respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A declaration of optional capabilities supported by the agent.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilities struct {
	// A list of protocol extensions supported by the agent.
	Extensions []InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilitiesExtension `json:"extensions"`
	// Indicates if the agent supports sending push notifications for asynchronous task
	// updates.
	PushNotifications bool `json:"pushNotifications"`
	// Indicates if the agent provides a history of state transitions for a task.
	StateTransitionHistory bool `json:"stateTransitionHistory"`
	// Indicates if the agent supports Server-Sent Events (SSE) for streaming
	// responses.
	Streaming bool `json:"streaming"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Extensions             respjson.Field
		PushNotifications      respjson.Field
		StateTransitionHistory respjson.Field
		Streaming              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilities) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A declaration of a protocol extension supported by an Agent.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilitiesExtension struct {
	// The unique URI identifying the extension.
	Uri string `json:"uri,required"`
	// A human-readable description of how this agent uses the extension.
	Description string `json:"description"`
	// Optional, extension-specific configuration parameters.
	Params map[string]any `json:"params"`
	// If true, the client must understand and comply with the extension's requirements
	// to interact with the agent.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		Description respjson.Field
		Params      respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilitiesExtension) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultCapabilitiesExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a distinct capability or function that an agent can perform.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSkill struct {
	// A unique identifier for the agent's skill.
	ID string `json:"id,required"`
	// A detailed description of the skill, intended to help clients or users
	// understand its purpose and functionality.
	Description string `json:"description,required"`
	// A human-readable name for the skill.
	Name string `json:"name,required"`
	// A set of keywords describing the skill's capabilities.
	Tags []string `json:"tags,required"`
	// Example prompts or scenarios that this skill can handle. Provides a hint to the
	// client on how to use the skill.
	Examples []string `json:"examples"`
	// The set of supported input MIME types for this skill, overriding the agent's
	// defaults.
	InputModes []string `json:"inputModes"`
	// The set of supported output MIME types for this skill, overriding the agent's
	// defaults.
	OutputModes []string `json:"outputModes"`
	// Security schemes necessary for the agent to leverage this skill. As in the
	// overall AgentCard.security, this list represents a logical OR of security
	// requirement objects. Each object is a set of security schemes that must be used
	// together (a logical AND).
	Security []map[string][]string `json:"security"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Tags        respjson.Field
		Examples    respjson.Field
		InputModes  respjson.Field
		OutputModes respjson.Field
		Security    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSkill) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSkill) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Declares a combination of a target URL and a transport protocol for interacting
// with the agent. This allows agents to expose the same functionality over
// multiple transport mechanisms.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultAdditionalInterface struct {
	// The transport protocol supported at this URL.
	Transport string `json:"transport,required"`
	// The URL where this interface is available. Must be a valid absolute HTTPS URL in
	// production.
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transport   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultAdditionalInterface) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultAdditionalInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the agent's service provider.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultProvider struct {
	// The name of the agent provider's organization.
	Organization string `json:"organization,required"`
	// A URL for the agent provider's website or relevant documentation.
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultProvider) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion
// contains all possible properties and values from
// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme],
// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme],
// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme],
// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOpenIDConnectSecurityScheme],
// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeMutualTlsSecurityScheme].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion struct {
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme].
	In string `json:"in"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme].
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme].
	Scheme string `json:"scheme"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme].
	BearerFormat string `json:"bearerFormat"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme].
	Flows InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlows `json:"flows"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme].
	Oauth2MetadataURL string `json:"oauth2MetadataUrl"`
	// This field is from variant
	// [InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOpenIDConnectSecurityScheme].
	OpenIDConnectURL string `json:"openIdConnectUrl"`
	JSON             struct {
		In                respjson.Field
		Name              respjson.Field
		Type              respjson.Field
		Description       respjson.Field
		Scheme            respjson.Field
		BearerFormat      respjson.Field
		Flows             respjson.Field
		Oauth2MetadataURL respjson.Field
		OpenIDConnectURL  respjson.Field
		raw               string
	} `json:"-"`
}

func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion) AsInvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme() (v InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion) AsInvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme() (v InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion) AsInvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme() (v InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion) AsInvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOpenIDConnectSecurityScheme() (v InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOpenIDConnectSecurityScheme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion) AsInvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeMutualTlsSecurityScheme() (v InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeMutualTlsSecurityScheme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a security scheme using an API key.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme struct {
	// The location of the API key.
	//
	// Any of "cookie", "header", "query".
	In string `json:"in,required"`
	// The name of the header, query, or cookie parameter to be used.
	Name string `json:"name,required"`
	// The type of the security scheme. Must be 'apiKey'.
	Type constant.APIKey `json:"type,required"`
	// An optional description for the security scheme.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		In          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeAPIKeySecurityScheme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a security scheme using HTTP authentication.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme struct {
	// The name of the HTTP Authentication scheme to be used in the Authorization
	// header, as defined in RFC7235 (e.g., "Bearer"). This value should be registered
	// in the IANA Authentication Scheme registry.
	Scheme string `json:"scheme,required"`
	// The type of the security scheme. Must be 'http'.
	Type constant.HTTP `json:"type,required"`
	// A hint to the client to identify how the bearer token is formatted (e.g.,
	// "JWT"). This is primarily for documentation purposes.
	BearerFormat string `json:"bearerFormat"`
	// An optional description for the security scheme.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scheme       respjson.Field
		Type         respjson.Field
		BearerFormat respjson.Field
		Description  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeHTTPAuthSecurityScheme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a security scheme using OAuth 2.0.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme struct {
	// An object containing configuration information for the supported OAuth 2.0
	// flows.
	Flows InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlows `json:"flows,required"`
	// The type of the security scheme. Must be 'oauth2'.
	Type constant.Oauth2 `json:"type,required"`
	// An optional description for the security scheme.
	Description string `json:"description"`
	// URL to the oauth2 authorization server metadata
	// [RFC8414](https://datatracker.ietf.org/doc/html/rfc8414). TLS is required.
	Oauth2MetadataURL string `json:"oauth2MetadataUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flows             respjson.Field
		Type              respjson.Field
		Description       respjson.Field
		Oauth2MetadataURL respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecurityScheme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing configuration information for the supported OAuth 2.0
// flows.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlows struct {
	// Configuration for the OAuth Authorization Code flow. Previously called
	// accessCode in OpenAPI 2.0.
	AuthorizationCode InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsAuthorizationCode `json:"authorizationCode"`
	// Configuration for the OAuth Client Credentials flow. Previously called
	// application in OpenAPI 2.0.
	ClientCredentials InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsClientCredentials `json:"clientCredentials"`
	// Configuration for the OAuth Implicit flow.
	Implicit InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsImplicit `json:"implicit"`
	// Configuration for the OAuth Resource Owner Password flow.
	Password InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsPassword `json:"password"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationCode respjson.Field
		ClientCredentials respjson.Field
		Implicit          respjson.Field
		Password          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlows) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the OAuth Authorization Code flow. Previously called
// accessCode in OpenAPI 2.0.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsAuthorizationCode struct {
	// The authorization URL to be used for this flow. This MUST be a URL and use TLS.
	AuthorizationURL string `json:"authorizationUrl,required"`
	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string `json:"scopes,required"`
	// The token URL to be used for this flow. This MUST be a URL and use TLS.
	TokenURL string `json:"tokenUrl,required"`
	// The URL to be used for obtaining refresh tokens. This MUST be a URL and use TLS.
	RefreshURL string `json:"refreshUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationURL respjson.Field
		Scopes           respjson.Field
		TokenURL         respjson.Field
		RefreshURL       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsAuthorizationCode) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsAuthorizationCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the OAuth Client Credentials flow. Previously called
// application in OpenAPI 2.0.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsClientCredentials struct {
	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string `json:"scopes,required"`
	// The token URL to be used for this flow. This MUST be a URL.
	TokenURL string `json:"tokenUrl,required"`
	// The URL to be used for obtaining refresh tokens. This MUST be a URL.
	RefreshURL string `json:"refreshUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		TokenURL    respjson.Field
		RefreshURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsClientCredentials) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsClientCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the OAuth Implicit flow.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsImplicit struct {
	// The authorization URL to be used for this flow. This MUST be a URL.
	AuthorizationURL string `json:"authorizationUrl,required"`
	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string `json:"scopes,required"`
	// The URL to be used for obtaining refresh tokens. This MUST be a URL.
	RefreshURL string `json:"refreshUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationURL respjson.Field
		Scopes           respjson.Field
		RefreshURL       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsImplicit) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsImplicit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the OAuth Resource Owner Password flow.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsPassword struct {
	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string `json:"scopes,required"`
	// The token URL to be used for this flow. This MUST be a URL.
	TokenURL string `json:"tokenUrl,required"`
	// The URL to be used for obtaining refresh tokens. This MUST be a URL.
	RefreshURL string `json:"refreshUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scopes      respjson.Field
		TokenURL    respjson.Field
		RefreshURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsPassword) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOAuth2SecuritySchemeFlowsPassword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a security scheme using OpenID Connect.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOpenIDConnectSecurityScheme struct {
	// The OpenID Connect Discovery URL for the OIDC provider's metadata.
	OpenIDConnectURL string `json:"openIdConnectUrl,required"`
	// The type of the security scheme. Must be 'openIdConnect'.
	Type constant.OpenIDConnect `json:"type,required"`
	// An optional description for the security scheme.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenIDConnectURL respjson.Field
		Type             respjson.Field
		Description      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOpenIDConnectSecurityScheme) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeOpenIDConnectSecurityScheme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a security scheme using mTLS authentication.
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeMutualTlsSecurityScheme struct {
	// The type of the security scheme. Must be 'mutualTLS'.
	Type constant.MutualTls `json:"type,required"`
	// An optional description for the security scheme.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeMutualTlsSecurityScheme) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSecuritySchemeMutualTlsSecurityScheme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentCardSignature represents a JWS signature of an AgentCard. This follows the
// JSON format of an RFC 7515 JSON Web Signature (JWS).
type InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSignature struct {
	// The protected JWS header for the signature. This is a Base64url-encoded JSON
	// object, as per RFC 7515.
	Protected string `json:"protected,required"`
	// The computed signature, Base64url-encoded.
	Signature string `json:"signature,required"`
	// The unprotected JWS header values.
	Header map[string]any `json:"header"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Protected   respjson.Field
		Signature   respjson.Field
		Header      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSignature) RawJSON() string {
	return r.JSON.raw
}
func (r *InvokeResponseGetAuthenticatedExtendedCardSuccessResponseResultSignature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvokeParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `message/send` method.
	OfSendMessageRequest *InvokeParamsBodySendMessageRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `message/stream` method.
	OfSendStreamingMessageRequest *InvokeParamsBodySendStreamingMessageRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `tasks/get` method.
	OfGetTaskRequest *InvokeParamsBodyGetTaskRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `tasks/cancel` method.
	OfCancelTaskRequest *InvokeParamsBodyCancelTaskRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/set` method.
	OfSetTaskPushNotificationConfigRequest *InvokeParamsBodySetTaskPushNotificationConfigRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/get` method.
	OfGetTaskPushNotificationConfigRequest *InvokeParamsBodyGetTaskPushNotificationConfigRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `tasks/resubscribe` method, used to resume
	// a streaming connection.
	OfTaskResubscriptionRequest *InvokeParamsBodyTaskResubscriptionRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/list`
	// method.
	OfListTaskPushNotificationConfigRequest *InvokeParamsBodyListTaskPushNotificationConfigRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/delete`
	// method.
	OfDeleteTaskPushNotificationConfigRequest *InvokeParamsBodyDeleteTaskPushNotificationConfigRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Represents a JSON-RPC request for the `agent/getAuthenticatedExtendedCard`
	// method.
	OfGetAuthenticatedExtendedCardRequest *InvokeParamsBodyGetAuthenticatedExtendedCardRequest `json:",inline"`

	paramObj
}

func (u InvokeParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSendMessageRequest,
		u.OfSendStreamingMessageRequest,
		u.OfGetTaskRequest,
		u.OfCancelTaskRequest,
		u.OfSetTaskPushNotificationConfigRequest,
		u.OfGetTaskPushNotificationConfigRequest,
		u.OfTaskResubscriptionRequest,
		u.OfListTaskPushNotificationConfigRequest,
		u.OfDeleteTaskPushNotificationConfigRequest,
		u.OfGetAuthenticatedExtendedCardRequest)
}
func (r *InvokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a JSON-RPC request for the `message/send` method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodySendMessageRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodySendMessageRequestIDUnion `json:"id,omitzero,required"`
	// The parameters for sending a message.
	Params shared.MessageSendParams `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'message/send'.
	//
	// This field can be elided, and will marshal its zero value as "message/send".
	Method constant.MessageSend `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodySendMessageRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodySendMessageRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodySendMessageRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodySendMessageRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodySendMessageRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodySendMessageRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodySendMessageRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Represents a JSON-RPC request for the `message/stream` method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodySendStreamingMessageRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodySendStreamingMessageRequestIDUnion `json:"id,omitzero,required"`
	// The parameters for sending a message.
	Params shared.MessageSendParams `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'message/stream'.
	//
	// This field can be elided, and will marshal its zero value as "message/stream".
	Method constant.MessageStream `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodySendStreamingMessageRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodySendStreamingMessageRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodySendStreamingMessageRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodySendStreamingMessageRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodySendStreamingMessageRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodySendStreamingMessageRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodySendStreamingMessageRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Represents a JSON-RPC request for the `tasks/get` method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodyGetTaskRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodyGetTaskRequestIDUnion `json:"id,omitzero,required"`
	// The parameters for querying a task.
	Params InvokeParamsBodyGetTaskRequestParams `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'tasks/get'.
	//
	// This field can be elided, and will marshal its zero value as "tasks/get".
	Method constant.TasksGet `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodyGetTaskRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyGetTaskRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyGetTaskRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyGetTaskRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyGetTaskRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodyGetTaskRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyGetTaskRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The parameters for querying a task.
//
// The property ID is required.
type InvokeParamsBodyGetTaskRequestParams struct {
	// The unique identifier (e.g. UUID) of the task.
	ID string `json:"id,required"`
	// The number of most recent messages from the task's history to retrieve.
	HistoryLength param.Opt[int64] `json:"historyLength,omitzero"`
	// Optional metadata associated with the request.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r InvokeParamsBodyGetTaskRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyGetTaskRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyGetTaskRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a JSON-RPC request for the `tasks/cancel` method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodyCancelTaskRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodyCancelTaskRequestIDUnion `json:"id,omitzero,required"`
	// The parameters identifying the task to cancel.
	Params shared.TaskIDParams `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'tasks/cancel'.
	//
	// This field can be elided, and will marshal its zero value as "tasks/cancel".
	Method constant.TasksCancel `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodyCancelTaskRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyCancelTaskRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyCancelTaskRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyCancelTaskRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyCancelTaskRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodyCancelTaskRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyCancelTaskRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/set` method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodySetTaskPushNotificationConfigRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodySetTaskPushNotificationConfigRequestIDUnion `json:"id,omitzero,required"`
	// The parameters for setting the push notification configuration.
	Params shared.TaskPushNotificationConfigParam `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'tasks/pushNotificationConfig/set'.
	//
	// This field can be elided, and will marshal its zero value as
	// "tasks/pushNotificationConfig/set".
	Method constant.TasksPushNotificationConfigSet `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodySetTaskPushNotificationConfigRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodySetTaskPushNotificationConfigRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodySetTaskPushNotificationConfigRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodySetTaskPushNotificationConfigRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodySetTaskPushNotificationConfigRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodySetTaskPushNotificationConfigRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodySetTaskPushNotificationConfigRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/get` method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodyGetTaskPushNotificationConfigRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodyGetTaskPushNotificationConfigRequestIDUnion `json:"id,omitzero,required"`
	// The parameters for getting a push notification configuration.
	Params InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'tasks/pushNotificationConfig/get'.
	//
	// This field can be elided, and will marshal its zero value as
	// "tasks/pushNotificationConfig/get".
	Method constant.TasksPushNotificationConfigGet `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodyGetTaskPushNotificationConfigRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyGetTaskPushNotificationConfigRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyGetTaskPushNotificationConfigRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyGetTaskPushNotificationConfigRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyGetTaskPushNotificationConfigRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodyGetTaskPushNotificationConfigRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyGetTaskPushNotificationConfigRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion struct {
	OfTaskIDs                                                                                  *shared.TaskIDParams                                                                           `json:",omitzero,inline"`
	OfInvokesBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams *InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTaskIDs, u.OfInvokesBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams)
}
func (u *InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion) asAny() any {
	if !param.IsOmitted(u.OfTaskIDs) {
		return u.OfTaskIDs
	} else if !param.IsOmitted(u.OfInvokesBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams) {
		return u.OfInvokesBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion) GetPushNotificationConfigID() *string {
	if vt := u.OfInvokesBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams; vt != nil && vt.PushNotificationConfigID.Valid() {
		return &vt.PushNotificationConfigID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion) GetID() *string {
	if vt := u.OfTaskIDs; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfInvokesBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's Metadata property, if present.
func (u InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsUnion) GetMetadata() map[string]any {
	if vt := u.OfTaskIDs; vt != nil {
		return vt.Metadata
	} else if vt := u.OfInvokesBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams; vt != nil {
		return vt.Metadata
	}
	return nil
}

// Defines parameters for fetching a specific push notification configuration for a
// task.
//
// The property ID is required.
type InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams struct {
	// The unique identifier (e.g. UUID) of the task.
	ID string `json:"id,required"`
	// The ID of the push notification configuration to retrieve.
	PushNotificationConfigID param.Opt[string] `json:"pushNotificationConfigId,omitzero"`
	// Optional metadata associated with the request.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyGetTaskPushNotificationConfigRequestParamsGetTaskPushNotificationConfigParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a JSON-RPC request for the `tasks/resubscribe` method, used to resume
// a streaming connection.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodyTaskResubscriptionRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodyTaskResubscriptionRequestIDUnion `json:"id,omitzero,required"`
	// The parameters identifying the task to resubscribe to.
	Params shared.TaskIDParams `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'tasks/resubscribe'.
	//
	// This field can be elided, and will marshal its zero value as
	// "tasks/resubscribe".
	Method constant.TasksResubscribe `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodyTaskResubscriptionRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyTaskResubscriptionRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyTaskResubscriptionRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyTaskResubscriptionRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyTaskResubscriptionRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodyTaskResubscriptionRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyTaskResubscriptionRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/list`
// method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodyListTaskPushNotificationConfigRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodyListTaskPushNotificationConfigRequestIDUnion `json:"id,omitzero,required"`
	// The parameters identifying the task whose configurations are to be listed.
	Params InvokeParamsBodyListTaskPushNotificationConfigRequestParams `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'tasks/pushNotificationConfig/list'.
	//
	// This field can be elided, and will marshal its zero value as
	// "tasks/pushNotificationConfig/list".
	Method constant.TasksPushNotificationConfigList `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodyListTaskPushNotificationConfigRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyListTaskPushNotificationConfigRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyListTaskPushNotificationConfigRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyListTaskPushNotificationConfigRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyListTaskPushNotificationConfigRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodyListTaskPushNotificationConfigRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyListTaskPushNotificationConfigRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The parameters identifying the task whose configurations are to be listed.
//
// The property ID is required.
type InvokeParamsBodyListTaskPushNotificationConfigRequestParams struct {
	// The unique identifier (e.g. UUID) of the task.
	ID string `json:"id,required"`
	// Optional metadata associated with the request.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r InvokeParamsBodyListTaskPushNotificationConfigRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyListTaskPushNotificationConfigRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyListTaskPushNotificationConfigRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a JSON-RPC request for the `tasks/pushNotificationConfig/delete`
// method.
//
// The properties ID, Jsonrpc, Method, Params are required.
type InvokeParamsBodyDeleteTaskPushNotificationConfigRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodyDeleteTaskPushNotificationConfigRequestIDUnion `json:"id,omitzero,required"`
	// The parameters identifying the push notification configuration to delete.
	Params InvokeParamsBodyDeleteTaskPushNotificationConfigRequestParams `json:"params,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'tasks/pushNotificationConfig/delete'.
	//
	// This field can be elided, and will marshal its zero value as
	// "tasks/pushNotificationConfig/delete".
	Method constant.TasksPushNotificationConfigDelete `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodyDeleteTaskPushNotificationConfigRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyDeleteTaskPushNotificationConfigRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyDeleteTaskPushNotificationConfigRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyDeleteTaskPushNotificationConfigRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyDeleteTaskPushNotificationConfigRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodyDeleteTaskPushNotificationConfigRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyDeleteTaskPushNotificationConfigRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The parameters identifying the push notification configuration to delete.
//
// The properties ID, PushNotificationConfigID are required.
type InvokeParamsBodyDeleteTaskPushNotificationConfigRequestParams struct {
	// The unique identifier (e.g. UUID) of the task.
	ID string `json:"id,required"`
	// The ID of the push notification configuration to delete.
	PushNotificationConfigID string `json:"pushNotificationConfigId,required"`
	// Optional metadata associated with the request.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r InvokeParamsBodyDeleteTaskPushNotificationConfigRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyDeleteTaskPushNotificationConfigRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyDeleteTaskPushNotificationConfigRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a JSON-RPC request for the `agent/getAuthenticatedExtendedCard`
// method.
//
// The properties ID, Jsonrpc, Method are required.
type InvokeParamsBodyGetAuthenticatedExtendedCardRequest struct {
	// The identifier for this request.
	ID InvokeParamsBodyGetAuthenticatedExtendedCardRequestIDUnion `json:"id,omitzero,required"`
	// The version of the JSON-RPC protocol. MUST be exactly "2.0".
	//
	// This field can be elided, and will marshal its zero value as "2.0".
	Jsonrpc constant.String2_0 `json:"jsonrpc,required"`
	// The method name. Must be 'agent/getAuthenticatedExtendedCard'.
	//
	// This field can be elided, and will marshal its zero value as
	// "agent/getAuthenticatedExtendedCard".
	Method constant.AgentGetAuthenticatedExtendedCard `json:"method,required"`
	paramObj
}

func (r InvokeParamsBodyGetAuthenticatedExtendedCardRequest) MarshalJSON() (data []byte, err error) {
	type shadow InvokeParamsBodyGetAuthenticatedExtendedCardRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvokeParamsBodyGetAuthenticatedExtendedCardRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvokeParamsBodyGetAuthenticatedExtendedCardRequestIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u InvokeParamsBodyGetAuthenticatedExtendedCardRequestIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *InvokeParamsBodyGetAuthenticatedExtendedCardRequestIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvokeParamsBodyGetAuthenticatedExtendedCardRequestIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}
