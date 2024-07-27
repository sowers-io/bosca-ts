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

package supplementary

import (
	"bosca.io/api/protobuf/bosca"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"bosca.io/pkg/util"
	"context"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"net/url"
	"os"
)

var Command = &cobra.Command{
	Use:   "supplementary [id] [key] [file]",
	Short: "Download a metadata supplementary",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		u, err := client.GetMetadataSupplementaryDownloadUrl(ctx, &bosca.SupplementaryIdRequest{
			Id:  args[0],
			Key: args[1],
		})
		if err != nil {
			return err
		}

		urlValue, _ := url.Parse(u.Url)
		request := &http.Request{
			Method: u.Method,
			Header: map[string][]string{},
			URL:    urlValue,
		}

		for _, header := range u.Headers {
			if request.Header[header.Name] == nil {
				request.Header[header.Name] = []string{header.Value}
			} else {
				request.Header[header.Name] = append(request.Header[header.Name], header.Value)
			}
		}

		httpClient := util.NewDefaultHttpClient()
		response, err := httpClient.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return os.WriteFile(args[2], bytes, os.FileMode(0644))
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:7000", "The endpoint to use.")
}
