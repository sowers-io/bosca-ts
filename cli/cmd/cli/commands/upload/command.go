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
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/cmd/cli/commands/login"
	"github.com/eventials/go-tus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strings"
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

		cfg := tus.DefaultConfig()

		session, err := login.GetSession()
		if err != nil {
			return err
		}

		cfg.Header.Set("Authorization", "Bearer "+session)

		client, err := tus.NewClient(endpoint, cfg)
		if err != nil {
			return err
		}
		upload, err := tus.NewUploadFromFile(f)
		if err != nil {
			return err
		}
		buf := make([]byte, 0, 8096)
		f.Read(buf)
		f.Seek(0, 0)

		parts := strings.Split(f.Name(), string(os.PathSeparator))
		if len(args) == 2 {
			upload.Metadata["id"] = args[1]
		}
		upload.Metadata["name"] = parts[len(parts)-1]
		upload.Metadata["filetype"] = http.DetectContentType(buf)
		upload.Metadata["traits"] = cmd.Flag(flags.TraitFlag).Value.String()
		uploader, err := client.CreateUpload(upload)
		if err != nil {
			return err
		}
		return uploader.Upload()
	},
}

func init() {
	Command.PersistentFlags().String(flags.TraitFlag, "", "Trait ID")
	Command.PersistentFlags().String(flags.EndpointFlag, "http://localhost:7001/files", "The endpoint to use.")
}
