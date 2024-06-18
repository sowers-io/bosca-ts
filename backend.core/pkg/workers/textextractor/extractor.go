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

package textextractor

import (
	protobuf "bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workers/common"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

func extractText(ctx context.Context, contentType string, metadataFile *os.File) (*http.Response, error) {
	cfg := common.GetConfiguration(ctx)
	extractorUrl, err := url.Parse(cfg.ClientEndPoints.TextExtractorApiAddress)
	if err != nil {
		return nil, err
	}
	client := common.GetHttpClient(ctx)
	file, _ := metadataFile.Stat()
	request := &http.Request{
		URL:    extractorUrl,
		Method: "POST",
		Header: map[string][]string{
			"Content-Type":   {contentType},
			"Content-Length": {fmt.Sprintf("%d", file.Size())},
		},
		ContentLength: file.Size(),
		Body:          metadataFile,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to extract text: %d", response.StatusCode)
	}
	return response, nil
}

func ExtractToSupplementaryText(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	metadataDownloadUrl, err := common.GetContentService(ctx).GetMetadataDownloadUrl(common.GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: executionContext.Metadata.Id,
	})
	if err != nil {
		return err
	}
	if metadataDownloadUrl == nil {
		slog.WarnContext(ctx, "workflow download url is nil, nothing to do", slog.String("metadata_id", executionContext.Metadata.Id))
		return nil
	}
	metadataFile, err := common.DownloadTemporaryFile(ctx, metadataDownloadUrl)
	if err != nil {
		return err
	}
	defer os.Remove(metadataFile.Name())
	response, err := extractText(ctx, executionContext.Metadata.ContentType, metadataFile)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	file, err := os.CreateTemp("/tmp", "extracted-text")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	_, err = io.Copy(file, response.Body)
	_, _ = file.Seek(0, 0)
	supplementaryId := executionContext.Activity.Outputs["supplementaryId"]
	return common.SetSupplementaryContentFile(ctx, executionContext, supplementaryId, "text/plain", file)
}
