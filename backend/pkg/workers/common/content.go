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

import (
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/util"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)

func GetMetadata(ctx context.Context, id string) (*content.Metadata, error) {
	contentService := GetContentService(ctx)
	return contentService.GetMetadata(GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: id,
	})
}

func DownloadTemporaryFile(ctx context.Context, signedUrl *content.SignedUrl) (*os.File, error) {
	downloadUrl, err := url.Parse(signedUrl.Url)
	if err != nil {
		return nil, err
	}
	file, err := os.CreateTemp("/tmp", signedUrl.Id)
	if err != nil {
		return nil, err
	}
	client := GetHttpClient(ctx)
	request := &http.Request{
		URL:    downloadUrl,
		Header: util.GetSignedUrlHeaders(signedUrl),
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(0, 0)
	return file, err
}
