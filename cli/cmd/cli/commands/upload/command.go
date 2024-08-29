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

package upload

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "upload [file]",
	Short: "Upload a file",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		endpoint := cmd.Flag(flags.EndpointFlag).Value.String()
		cmd.Printf("Uploading to %v", endpoint)
		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}
		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}
		s, _ := f.Stat()
		size := s.Size()
		metadata, err := client.AddMetadata(ctx, &content.AddMetadataRequest{
			Metadata: &content.Metadata{
				Name:          args[0],
				ContentType:   "application/octet-stream",
				ContentLength: &size,
				LanguageTag:   "en",
			},
		})
		if err != nil {
			return err
		}
		signedUrl, err := client.GetMetadataUploadUrl(ctx, &bosca.IdRequest{
			Id: metadata.Id,
		})
		if err != nil {
			return err
		}

		cmd.Printf("\r\n\r\n%v\r\n\r\n", signedUrl)

		// req, err := http.NewRequest(signedUrl.Method, strings.Replace(signedUrl.Url, "bosca.minio", "minio", 1), f)
		req, err := http.NewRequest(signedUrl.Method, signedUrl.Url, f)
		for _, h := range signedUrl.Headers {
			req.Header.Add(h.Name, h.Value)
		}
		for k, v := range signedUrl.Attributes {
			req.Header.Add(k, v)
		}
		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			s, _ := ioutil.ReadAll(res.Body)
			return errors.New(res.Status + ": " + string(s))
		}
		_, err = client.SetMetadataReady(ctx, &content.MetadataReadyRequest{Id: metadata.Id})
		return err
	},
}

func init() {
	Command.PersistentFlags().String(flags.TraitFlag, "", "Trait ID")
	Command.PersistentFlags().String(flags.EndpointFlag, "localhost:7000", "The endpoint to use.")
}
