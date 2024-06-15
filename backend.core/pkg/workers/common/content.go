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
	protobuf "bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/util"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetWorkflow(ctx context.Context, id string) (*content.Workflow, error) {
	contentService := GetContentService(ctx)
	return contentService.GetWorkflow(GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: id,
	})
}

func GetMetadata(ctx context.Context, id string) (*content.Metadata, error) {
	contentService := GetContentService(ctx)
	return contentService.GetMetadata(GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: id,
	})
}

func DeleteMetadata(ctx context.Context, id string) error {
	contentService := GetContentService(ctx)
	_, err := contentService.DeleteMetadata(GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: id,
	})
	return err
}

func DownloadTemporaryMetadataFile(ctx context.Context, id string) (*os.File, error) {
	client := GetContentService(ctx)
	signedUrl, err := client.GetMetadataDownloadUrl(GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return DownloadTemporaryFile(ctx, signedUrl)
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

func SetTextContent(ctx context.Context, id string, text string) error {
	contentService := GetContentService(ctx)
	signedUrl, err := contentService.GetMetadataUploadUrl(GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: id,
	})
	if err != nil {
		return err
	}
	uploadUrl, err := url.Parse(signedUrl.Url)
	if err != nil {
		return err
	}
	client := GetHttpClient(ctx)
	request := &http.Request{
		Method:        signedUrl.Method,
		URL:           uploadUrl,
		Header:        util.GetSignedUrlHeaders(signedUrl),
		Body:          io.NopCloser(strings.NewReader(text)),
		ContentLength: int64(len(text)),
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated && response.StatusCode != http.StatusAccepted && response.StatusCode != http.StatusNoContent {
		return errors.New("failed to upload")
	}
	return nil
}
