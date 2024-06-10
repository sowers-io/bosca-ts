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

package ory

import (
	"bosca.io/pkg/security"
	"encoding/json"
	"net/http"
	"time"
)

type Session struct {
	ID                          string    `json:"id"`
	Active                      bool      `json:"active"`
	ExpiresAt                   time.Time `json:"expires_at"`
	AuthenticatedAt             time.Time `json:"authenticated_at"`
	AuthenticatorAssuranceLevel string    `json:"authenticator_assurance_level"`
	AuthenticationMethods       []struct {
		Method      string    `json:"method"`
		Aal         string    `json:"aal"`
		CompletedAt time.Time `json:"completed_at"`
	} `json:"authentication_methods"`
	IssuedAt time.Time `json:"issued_at"`
	Identity struct {
		ID             string    `json:"id"`
		SchemaID       string    `json:"schema_id"`
		State          string    `json:"state"`
		StateChangedAt time.Time `json:"state_changed_at"`
		Traits         struct {
			Consent struct {
				Newsletter bool      `json:"newsletter"`
				Tos        time.Time `json:"tos"`
			} `json:"consent"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"traits"`
		VerifiableAddresses []struct {
			ID        string    `json:"id"`
			Value     string    `json:"value"`
			Verified  bool      `json:"verified"`
			Via       string    `json:"via"`
			Status    string    `json:"status"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"verifiable_addresses"`
		RecoveryAddresses []struct {
			ID        string    `json:"id"`
			Value     string    `json:"value"`
			Via       string    `json:"via"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"recovery_addresses"`
		MetadataPublic any       `json:"metadata_public"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	} `json:"identity"`
	Devices []struct {
		ID        string `json:"id"`
		IPAddress string `json:"ip_address"`
		UserAgent string `json:"user_agent"`
		Location  string `json:"location"`
	} `json:"devices"`
}

type sessionInterceptor struct {
}

func NewSessionInterceptor() security.SessionInterceptor {
	return &sessionInterceptor{}
}

func (s *sessionInterceptor) newSession(response *http.Response) (*Session, error) {
	session := &Session{}
	err := json.NewDecoder(response.Body).Decode(&session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionInterceptor) GetSubjectId(response *http.Response) (string, error) {
	session, err := s.newSession(response)
	if err != nil {
		return "", err
	}
	return session.Identity.ID, nil
}
