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

package identity

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const XSubjectId = "X-Subject"

func GetSubjectId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "failed to get grpc metadata")
	}
	userID := md.Get(XSubjectId)
	if len(userID) == 0 || userID[0] == "" {
		return "", status.Error(codes.Unauthenticated, "user is missing in metadata")
	}
	return userID[0], nil
}

func GetAuthenticatedSubjectId(ctx context.Context) (string, error) {
	id, err := GetSubjectId(ctx)
	if err != nil {
		return "", err
	}
	if id == "anonymous" {
		return "", status.Error(codes.Unauthenticated, "user is anonymous")
	}
	return id, nil
}
