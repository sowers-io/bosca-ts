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

package security

import (
	"bosca.io/pkg/security/identity"
	"bosca.io/pkg/util"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type SubjectFinder interface {
	FindSubjectId(ctx context.Context, cookie bool, authorization string) (string, string, error)
}

type subjectFinder struct {
	endpoint                  *url.URL
	client                    *http.Client
	serviceAccountId          string
	serviceAccountTokenHeader string
	interceptor               SessionInterceptor
}

func NewSubjectFinder(endpoint string, serviceAccountId string, serviceAccountToken string, interceptor SessionInterceptor) SubjectFinder {
	endpointUrl, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("failed to parse endpoint %s: %v", endpoint, err)
	}
	return &subjectFinder{
		endpoint:                  endpointUrl,
		interceptor:               interceptor,
		serviceAccountId:          serviceAccountId,
		serviceAccountTokenHeader: "Token " + serviceAccountToken,
		client:                    util.NewDefaultHttpClient(),
	}
}

func (m *subjectFinder) FindSubjectId(ctx context.Context, cookie bool, authorization string) (string, string, error) {
	request := &http.Request{
		Method: "GET",
		Header: map[string][]string{},
		URL:    m.endpoint,
	}

	if authorization == m.serviceAccountTokenHeader {
		return m.serviceAccountId, identity.SubjectTypeServiceAccount, nil
	} else if cookie {
		request.Header["Cookie"] = []string{authorization}
	} else {
		request.Header["Authorization"] = []string{authorization}
	}

	r, err := m.client.Do(request)
	if err != nil {
		log.Printf("failed to get subject: %v", err)
		return "", "", err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Printf("failed to get session: %v", r.Status)
		return "", "", fmt.Errorf("failed to get session: %d", r.StatusCode)
	}

	subjectId, err := m.interceptor.GetSubjectId(r)
	if err != nil {
		log.Printf("failed to get subject: %v", err)
		return "", "", err
	}
	return subjectId, identity.SubjectTypeUser, nil
}
