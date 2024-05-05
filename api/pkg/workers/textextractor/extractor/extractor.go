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
	"io"
	"net/http"
	"net/url"
	"os"
)

func Extract(ctx context.Context, metadata *content.Metadata) (string, error) {
	svc := common.GetContentService(ctx)
	signedUrl, err := svc.GetMetadataDownloadUrl(ctx, &protobuf.IdRequest{
		Id: metadata.Id,
	})
	if err != nil {
		return "", err
	}

	downloadUrl, err := url.Parse(signedUrl.Url)
	if err != nil {
		return "", err
	}

	cfg := common.GetConfiguration(ctx)
	extractorUrl, err := url.Parse(cfg.ClientEndPoints.TextExtractorApiAddress)
	if err != nil {
		return "", err
	}

	file, err := os.CreateTemp("/tmp", metadata.Id)
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())

	client := common.GetHttpClient(ctx)
	request := &http.Request{
		URL:    downloadUrl,
		Header: util.GetSignedUrlHeaders(signedUrl),
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	request = &http.Request{
		URL:    extractorUrl,
		Method: "POST",
		Body:   file,
	}
	response, err = client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("text extractor API returned non-OK status code")
	}

	return "", nil
}

func Cleanup(ctx context.Context, metadata *content.Metadata) {

}
