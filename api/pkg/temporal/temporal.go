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

package temporal

import (
	"bosca.io/pkg/configuration"
	"context"
	"go.temporal.io/sdk/workflow"
)
import "go.temporal.io/sdk/client"

func NewClient(ctx context.Context, cfg *configuration.ClientEndpoints) (client.Client, error) {
	c, err := client.DialContext(ctx, client.Options{
		HostPort: cfg.TemporalApiAddress,
	})
	return c, err
}

func NewClientWithPropagator(ctx context.Context, cfg *configuration.ClientEndpoints, propagator workflow.ContextPropagator) (client.Client, error) {
	c, err := client.DialContext(ctx, client.Options{
		HostPort:           cfg.TemporalApiAddress,
		ContextPropagators: []workflow.ContextPropagator{propagator},
	})
	return c, err
}
