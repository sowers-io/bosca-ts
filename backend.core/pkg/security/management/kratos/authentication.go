/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
