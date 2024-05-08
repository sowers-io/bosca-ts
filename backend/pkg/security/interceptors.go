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
	"bosca.io/pkg/identity"
	"bosca.io/pkg/util"
	"context"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"net/url"
)

type interceptors struct {
	endpoint                  *url.URL
	client                    *http.Client
	serviceAccountId          string
	serviceAccountTokenHeader string
	interceptor               SessionInterceptor
}

type Interceptors interface {
	UnaryInterceptor() grpc.UnaryServerInterceptor
	StreamInterceptor() grpc.StreamServerInterceptor
}

type SessionInterceptor interface {
	GetSubjectId(response *http.Response) (string, error)
}

func NewSecurityInterceptors(endpoint string, serviceAccountId string, serviceAccountToken string, interceptor SessionInterceptor) Interceptors {
	endpointUrl, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("failed to parse endpoint %s: %v", endpoint, err)
	}
	return &interceptors{
		endpoint:                  endpointUrl,
		interceptor:               interceptor,
		serviceAccountId:          serviceAccountId,
		serviceAccountTokenHeader: "Token " + serviceAccountToken,
		client:                    util.NewDefaultHttpClient(),
	}
}

func (m *interceptors) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return m.unaryInterceptor
}

func (m *interceptors) injectSubjectId(ctx context.Context) metadata.MD {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		request := &http.Request{
			Header: map[string][]string{},
			URL:    m.endpoint,
		}
		authorization := md.Get("Authorization")
		if authorization != nil && len(authorization) > 0 {
			request.Header["Authorization"] = authorization
		} else {
			serviceAuthorization := md.Get("X-Service-Authorization")
			if serviceAuthorization != nil && len(serviceAuthorization) > 0 && serviceAuthorization[0] == m.serviceAccountTokenHeader {
				md.Set(identity.XSubjectId, m.serviceAccountId)
				md.Set(identity.XSubjectType, identity.SubjectTypeServiceAccount)
				return md
			}
			cookies := md.Get("Cookie")
			if cookies != nil && len(cookies) > 0 {
				request.Header["Cookie"] = cookies
			}
		}
		if len(request.Header) == 0 {
			return nil
		}
		r, err := m.client.Do(request)
		if err != nil {
			log.Printf("failed to get session: %v", err)
			return nil
		}
		defer r.Body.Close()
		subjectId, err := m.interceptor.GetSubjectId(r)
		if err != nil {
			log.Printf("failed to get subject: %v", err)
			return nil
		}
		md.Set(identity.XSubjectId, subjectId)
		md.Set(identity.XSubjectType, identity.SubjectTypeUser)
		return md
	}
	return nil
}

func (m *interceptors) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md := m.injectSubjectId(ctx)
	if md != nil {
		ctx = metadata.NewIncomingContext(ctx, md)
	}
	return handler(ctx, req)
}

func (m *interceptors) StreamInterceptor() grpc.StreamServerInterceptor {
	return m.streamInterceptor
}

func (m *interceptors) streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	wrappedStream := middleware.WrapServerStream(stream)
	md := m.injectSubjectId(wrappedStream.Context())
	if md != nil {
		wrappedStream.WrappedContext = metadata.NewOutgoingContext(wrappedStream.Context(), md)
	}
	return handler(srv, wrappedStream)
}
