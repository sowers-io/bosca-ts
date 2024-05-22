package upload

import (
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/cmd/cli/commands/login"
	"github.com/eventials/go-tus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var Command = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file",
	Args:  cobra.ExactArgs(1),
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

		upload.Metadata["name"] = f.Name()
		upload.Metadata["filetype"] = http.DetectContentType(buf)
		uploader, err := client.CreateUpload(upload)
		if err != nil {
			return err
		}
		return uploader.Upload()
	},
}

func init() {
	Command.PersistentFlags().String(flags.EndpointFlag, "http://localhost:8099/uploads", "The endpoint to use.")
}
