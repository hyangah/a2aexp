// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package a2aexp_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/a2a-exp-go"
	"github.com/stainless-sdks/a2a-exp-go/internal/testutil"
	"github.com/stainless-sdks/a2a-exp-go/option"
	"github.com/stainless-sdks/a2a-exp-go/shared"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Invoke(context.TODO(), a2aexp.InvokeParamsSendMessageRequest{
		ID: a2aexp.InvokeParamsSendMessageRequestIDUnion{
			OfString: a2aexp.String("string"),
		},
		Jsonrpc: "2.0",
		Method:  "message/send",
		Params: shared.MessageSendParams{
			Message: shared.MessageParam{
				MessageID: "messageId",
				Parts: []shared.PartUnionParam{{
					OfPartTextPart: &shared.PartTextPartParam{
						Text: "text",
					},
				}},
				Role: shared.MessageRoleAgent,
			},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response)
}
