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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

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

func DownloadTemporarySupplementaryFile(ctx context.Context, id string, supplementaryId *content.WorkflowActivityParameterValue) (*os.File, error) {
	client := GetContentService(ctx)
	signedUrl, err := client.GetMetadataSupplementaryDownloadUrl(GetServiceAuthorizedContext(ctx), &content.SupplementaryIdRequest{
		Id:   id,
		Type: supplementaryId.GetSingleValue(),
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

func SetContent(ctx context.Context, metadataId string, data []byte) error {
	contentService := GetContentService(ctx)
	signedUrl, err := contentService.GetMetadataUploadUrl(GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: metadataId,
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
		Body:          io.NopCloser(bytes.NewBuffer(data)),
		ContentLength: int64(len(data)),
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated && response.StatusCode != http.StatusAccepted && response.StatusCode != http.StatusNoContent {
		b, _ := io.ReadAll(response.Body)
		return errors.New("failed to upload: " + string(b))
	}
	return nil
}

func SetSupplementaryContent(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext, supplementaryId *content.WorkflowActivityParameterValue, contentType string, data []byte) error {
	svc := GetContentService(ctx)
	uploadSignedUrl, err := svc.AddMetadataSupplementary(GetServiceAuthorizedContext(ctx), &content.AddSupplementaryRequest{
		Id:            executionContext.Metadata.Id,
		Type:          supplementaryId.GetSingleValue(),
		Name:          "",
		ContentType:   contentType,
		ContentLength: int64(len(data)),
	})
	if err != nil {
		return err
	}
	uploadUrl, err := url.Parse(uploadSignedUrl.Url)
	if err != nil {
		return err
	}
	request := &http.Request{
		URL:           uploadUrl,
		Method:        uploadSignedUrl.Method,
		Header:        util.GetSignedUrlHeaders(uploadSignedUrl),
		ContentLength: int64(len(data)),
		Body:          io.NopCloser(bytes.NewBuffer(data)),
	}
	response, err := GetHttpClient(ctx).Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload extracted text: %d", response.StatusCode)
	}
	return nil
}

func SetSupplementaryContentFile(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext, supplementaryId *content.WorkflowActivityParameterValue, contentType string, file *os.File) error {
	info, _ := file.Stat()
	svc := GetContentService(ctx)
	uploadSignedUrl, err := svc.AddMetadataSupplementary(GetServiceAuthorizedContext(ctx), &content.AddSupplementaryRequest{
		Id:            executionContext.Metadata.Id,
		Type:          supplementaryId.GetSingleValue(),
		Name:          "",
		ContentType:   contentType,
		ContentLength: info.Size(),
	})
	if err != nil {
		return err
	}
	uploadUrl, err := url.Parse(uploadSignedUrl.Url)
	if err != nil {
		return err
	}
	request := &http.Request{
		URL:           uploadUrl,
		Method:        uploadSignedUrl.Method,
		Header:        util.GetSignedUrlHeaders(uploadSignedUrl),
		ContentLength: info.Size(),
		Body:          file,
	}
	response, err := GetHttpClient(ctx).Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload extracted text: %d", response.StatusCode)
	}
	return nil
}
