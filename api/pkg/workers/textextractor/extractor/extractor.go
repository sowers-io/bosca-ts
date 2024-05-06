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

package extractor

import (
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/util"
	"bosca.io/pkg/workers/common"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type ExtractRequest struct {
	Metadata *content.Metadata
	Type     string
	Name     string
}

func Extract(ctx context.Context, extractRequest ExtractRequest) error {
	svc := common.GetContentService(ctx)
	signedUrl, err := svc.GetMetadataDownloadUrl(common.GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: extractRequest.Metadata.Id,
	})
	if err != nil {
		return err
	}

	downloadUrl, err := url.Parse(signedUrl.Url)
	if err != nil {
		return err
	}

	cfg := common.GetConfiguration(ctx)
	extractorUrl, err := url.Parse(cfg.ClientEndPoints.TextExtractorApiAddress)
	if err != nil {
		return err
	}

	file, err := os.CreateTemp("/tmp", extractRequest.Metadata.Id)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	client := common.GetHttpClient(ctx)
	request := &http.Request{
		URL:    downloadUrl,
		Header: util.GetSignedUrlHeaders(signedUrl),
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	request = &http.Request{
		URL:    extractorUrl,
		Method: "POST",
		Header: map[string][]string{
			"Content-Type": {extractRequest.Metadata.ContentType},
		},
		Body: file,
	}
	response, err = client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("failed to extract text")
	}

	uploadSignedUrl, err := svc.AddMetadataSupplementary(common.GetServiceAuthorizedContext(ctx), &content.AddSupplementaryRequest{
		Id:            extractRequest.Metadata.Id,
		Type:          extractRequest.Type,
		Name:          extractRequest.Name,
		ContentType:   extractRequest.Metadata.ContentType,
		ContentLength: response.ContentLength,
	})
	if err != nil {
		return err
	}

	uploadUrl, err := url.Parse(uploadSignedUrl.Url)
	if err != nil {
		return err
	}

	request = &http.Request{
		URL:    uploadUrl,
		Method: uploadSignedUrl.Method,
		Header: util.GetSignedUrlHeaders(uploadSignedUrl),
		Body:   response.Body,
	}

	response, err = client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload extracted text: %d", response.StatusCode)
	}

	return nil
}

func Cleanup(ctx context.Context, extractRequest ExtractRequest) error {
	svc := common.GetContentService(ctx)

	_, err := svc.DeleteMetadataSupplementary(common.GetServiceAuthorizedContext(ctx), &content.SupplementaryIdRequest{
		Id:   extractRequest.Metadata.Id,
		Type: extractRequest.Type,
	})

	return err
}
