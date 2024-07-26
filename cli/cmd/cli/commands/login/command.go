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

package login

import (
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/security/management/kratos"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func GetSession() (string, error) {
	session := viper.GetString(flags.SessionFlag)
	if session == "" {
		return "", fmt.Errorf("no session found, please login")
	}
	return session, nil
}

func RefreshSession() (string, error) {
	endpoint := viper.GetString(flags.EndpointFlag)
	username := viper.GetString(flags.UsernameFlag)
	password := viper.GetString(flags.PasswordFlag)

	client := kratos.NewClient(endpoint)
	session, err := kratos.Login(context.Background(), client, username, password)
	if err != nil {
		return "", err
	}

	viper.Set(flags.SessionFlag, session)
	err = viper.WriteConfig()
	if err != nil {
		return "", err
	}

	return session, nil
}

var Command = &cobra.Command{
	Use:   "login",
	Short: "Login to the system.",
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := cmd.Flag(flags.EndpointFlag).Value.String()
		username := cmd.Flag(flags.UsernameFlag).Value.String()
		password := cmd.Flag(flags.PasswordFlag).Value.String()

		viper.Set(flags.EndpointFlag, endpoint)
		viper.Set(flags.UsernameFlag, username)
		viper.Set(flags.PasswordFlag, password)

		_, err := RefreshSession()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	Command.Flags().String(flags.EndpointFlag, "http://localhost:4433", "The authentication endpoint.")
	Command.Flags().String(flags.UsernameFlag, "", "")
	Command.Flags().String(flags.PasswordFlag, "", "")

	viper.BindPFlag(flags.UsernameFlag, Command.PersistentFlags().Lookup(flags.UsernameFlag))
	viper.BindPFlag(flags.PasswordFlag, Command.PersistentFlags().Lookup(flags.PasswordFlag))
	viper.BindPFlag(flags.EndpointFlag, Command.PersistentFlags().Lookup(flags.EndpointFlag))
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".bosca")
	viper.SafeWriteConfig()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
