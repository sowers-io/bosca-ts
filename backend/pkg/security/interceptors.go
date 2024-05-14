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
	"context"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type interceptors struct {
	subjectFinder SubjectFinder
}

type Interceptors interface {
	UnaryInterceptor() grpc.UnaryServerInterceptor
	StreamInterceptor() grpc.StreamServerInterceptor
}

func NewSecurityInterceptors(subjectFinder SubjectFinder) Interceptors {
	return &interceptors{
		subjectFinder: subjectFinder,
	}
}

func (m *interceptors) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return m.unaryInterceptor
}

func (m *interceptors) injectSubjectId(ctx context.Context) metadata.MD {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		authorization := md.Get("Authorization")
		if authorization != nil && len(authorization) > 0 {
			subjectId, subjectType, err := m.subjectFinder.FindSubjectId(ctx, false, authorization[0])
			if err != nil {
				return nil
			}
			md.Set(identity.XSubjectId, subjectId)
			md.Set(identity.XSubjectType, subjectType)
			return md
		}
		authorization = md.Get("X-Service-Authorization")
		if authorization != nil && len(authorization) > 0 {
			subjectId, subjectType, err := m.subjectFinder.FindSubjectId(ctx, false, authorization[0])
			if err != nil {
				return nil
			}
			md.Set(identity.XSubjectId, subjectId)
			md.Set(identity.XSubjectType, subjectType)
			return md
		}
		cookies := md.Get("Cookie")
		if cookies != nil && len(cookies) > 0 {
			subjectId, subjectType, err := m.subjectFinder.FindSubjectId(ctx, true, cookies[0])
			if err != nil {
				return nil
			}
			md.Set(identity.XSubjectId, subjectId)
			md.Set(identity.XSubjectType, subjectType)
			return md
		}
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
