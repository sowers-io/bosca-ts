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

package configuration

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ServerConfiguration struct {
	Port                     int
	DatabaseConnectionString string
	OathKeeperConfiguration  string
}

func NewServerConfiguration() *ServerConfiguration {
	var configuration ServerConfiguration
	err := envconfig.Process("bosca", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	if configuration.OathKeeperConfiguration == "" {
		configuration.OathKeeperConfiguration = "../conf/accounts/oathkeeper.yaml"
	}
	if configuration.Port == 0 {
		configuration.Port = 8081
	}

	return &configuration
}
