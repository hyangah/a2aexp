// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/hyangah/a2a-exp/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type String2_0 string                         // Always "2.0"
type AgentGetAuthenticatedExtendedCard string // Always "agent/getAuthenticatedExtendedCard"
type APIKey string                            // Always "apiKey"
type ArtifactUpdate string                    // Always "artifact-update"
type Data string                              // Always "data"
type File string                              // Always "file"
type HTTP string                              // Always "http"
type Message string                           // Always "message"
type MessageSend string                       // Always "message/send"
type MessageStream string                     // Always "message/stream"
type MutualTls string                         // Always "mutualTLS"
type Oauth2 string                            // Always "oauth2"
type OpenIDConnect string                     // Always "openIdConnect"
type StatusUpdate string                      // Always "status-update"
type Task string                              // Always "task"
type TasksCancel string                       // Always "tasks/cancel"
type TasksGet string                          // Always "tasks/get"
type TasksPushNotificationConfigDelete string // Always "tasks/pushNotificationConfig/delete"
type TasksPushNotificationConfigGet string    // Always "tasks/pushNotificationConfig/get"
type TasksPushNotificationConfigList string   // Always "tasks/pushNotificationConfig/list"
type TasksPushNotificationConfigSet string    // Always "tasks/pushNotificationConfig/set"
type TasksResubscribe string                  // Always "tasks/resubscribe"
type Text string                              // Always "text"

func (c String2_0) Default() String2_0 { return "2.0" }
func (c AgentGetAuthenticatedExtendedCard) Default() AgentGetAuthenticatedExtendedCard {
	return "agent/getAuthenticatedExtendedCard"
}
func (c APIKey) Default() APIKey                 { return "apiKey" }
func (c ArtifactUpdate) Default() ArtifactUpdate { return "artifact-update" }
func (c Data) Default() Data                     { return "data" }
func (c File) Default() File                     { return "file" }
func (c HTTP) Default() HTTP                     { return "http" }
func (c Message) Default() Message               { return "message" }
func (c MessageSend) Default() MessageSend       { return "message/send" }
func (c MessageStream) Default() MessageStream   { return "message/stream" }
func (c MutualTls) Default() MutualTls           { return "mutualTLS" }
func (c Oauth2) Default() Oauth2                 { return "oauth2" }
func (c OpenIDConnect) Default() OpenIDConnect   { return "openIdConnect" }
func (c StatusUpdate) Default() StatusUpdate     { return "status-update" }
func (c Task) Default() Task                     { return "task" }
func (c TasksCancel) Default() TasksCancel       { return "tasks/cancel" }
func (c TasksGet) Default() TasksGet             { return "tasks/get" }
func (c TasksPushNotificationConfigDelete) Default() TasksPushNotificationConfigDelete {
	return "tasks/pushNotificationConfig/delete"
}
func (c TasksPushNotificationConfigGet) Default() TasksPushNotificationConfigGet {
	return "tasks/pushNotificationConfig/get"
}
func (c TasksPushNotificationConfigList) Default() TasksPushNotificationConfigList {
	return "tasks/pushNotificationConfig/list"
}
func (c TasksPushNotificationConfigSet) Default() TasksPushNotificationConfigSet {
	return "tasks/pushNotificationConfig/set"
}
func (c TasksResubscribe) Default() TasksResubscribe { return "tasks/resubscribe" }
func (c Text) Default() Text                         { return "text" }

func (c String2_0) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c AgentGetAuthenticatedExtendedCard) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c APIKey) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c ArtifactUpdate) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Data) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c HTTP) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Message) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c MessageSend) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c MessageStream) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c MutualTls) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Oauth2) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c OpenIDConnect) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c StatusUpdate) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c Task) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c TasksCancel) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c TasksGet) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c TasksPushNotificationConfigDelete) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c TasksPushNotificationConfigGet) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c TasksPushNotificationConfigList) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c TasksPushNotificationConfigSet) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c TasksResubscribe) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                              { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
