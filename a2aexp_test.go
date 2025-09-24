// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package a2aexp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/a2a-exp-go"
	"github.com/stainless-sdks/a2a-exp-go/internal/testutil"
	"github.com/stainless-sdks/a2a-exp-go/option"
	"github.com/stainless-sdks/a2a-exp-go/shared"
)

func TestA2aexpInvokeWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := a2aexp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Invoke(context.TODO(), a2aexp.InvokeParams{
		OfSendMessageRequest: &a2aexp.InvokeParamsBodySendMessageRequest{
			ID: a2aexp.InvokeParamsBodySendMessageRequestIDUnion{
				OfString: a2aexp.String("string"),
			},
			Params: shared.MessageSendParams{
				Message: shared.MessageParam{
					MessageID: "messageId",
					Parts: []shared.PartUnionParam{{
						OfPartTextPart: &shared.PartTextPartParam{
							Text: "text",
							Metadata: map[string]any{
								"foo": "bar",
							},
						},
					}},
					Role:       shared.MessageRoleAgent,
					ContextID:  a2aexp.String("contextId"),
					Extensions: []string{"string"},
					Metadata: map[string]any{
						"foo": "bar",
					},
					ReferenceTaskIDs: []string{"string"},
					TaskID:           a2aexp.String("taskId"),
				},
				Configuration: shared.MessageSendParamsConfiguration{
					AcceptedOutputModes: []string{"string"},
					Blocking:            a2aexp.Bool(true),
					HistoryLength:       a2aexp.Int(0),
					PushNotificationConfig: shared.PushNotificationConfigParam{
						URL:   "url",
						ID:    a2aexp.String("id"),
						Token: a2aexp.String("token"),
						Authentication: shared.PushNotificationConfigAuthenticationParam{
							Schemes:     []string{"string"},
							Credentials: a2aexp.String("credentials"),
						},
					},
				},
				Metadata: map[string]any{
					"foo": "bar",
				},
			},
		},
	})
	if err != nil {
		var apierr *a2aexp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
