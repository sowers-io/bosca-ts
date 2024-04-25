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

package common

import "context"

type Authorization struct {
	HeaderValue string
}

func (t Authorization) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.HeaderValue,
	}, nil
}

func (Authorization) RequireTransportSecurity() bool {
	// KJB: TODO: I'm leaning towards letting SSL between services within Kubernetes being managed at a different layer
	//            So, for now I'll leave this as false.  But, it's still open for consideration.
	return false
}
