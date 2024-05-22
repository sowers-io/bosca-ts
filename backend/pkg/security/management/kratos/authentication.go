package kratos

import (
	"context"
	kratos "github.com/ory/kratos-client-go"
)

func NewClient(endpoint string) *kratos.APIClient {
	configuration := kratos.NewConfiguration()
	configuration.Servers = []kratos.ServerConfiguration{
		{
			URL: endpoint,
		},
	}
	return kratos.NewAPIClient(configuration)
}

func Login(ctx context.Context, client *kratos.APIClient, username string, password string) (string, error) {
	flow, _, err := client.FrontendAPI.CreateNativeLoginFlow(ctx).Execute()
	if err != nil {
		return "", err
	}
	success, _, err := client.FrontendAPI.UpdateLoginFlow(ctx).Flow(flow.Id).UpdateLoginFlowBody(kratos.UpdateLoginFlowBody{
		UpdateLoginFlowWithPasswordMethod: &kratos.UpdateLoginFlowWithPasswordMethod{
			Method:             "password",
			Password:           password,
			PasswordIdentifier: &username,
		},
	}).Execute()
	if err != nil {
		return "", err
	}
	return success.GetSessionToken(), nil
}
