package bible

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/bible/usx"
	"bosca.io/pkg/workers/common"
	"bytes"
	"context"
	"os"
)

func GenerateChapters(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	metadataFile, err := common.DownloadTemporaryMetadataFile(ctx, executionContext.Metadata.Id)
	if err != nil {
		return err
	}
	defer metadataFile.Close()
	defer os.Remove(metadataFile.Name())

	bundle, err := usx.OpenBundle(metadataFile)
	if err != nil {
		return err
	}

	svc := common.GetContentService(ctx)
	collection, err := svc.AddCollection(ctx, &content.AddCollectionRequest{
		Collection: &content.Collection{
			Name: executionContext.Metadata.Name,
			Type: content.CollectionType_standard,
		},
	})
	if err != nil {
		return err
	}

	sourceId := "workflow"
	request := &content.AddMetadataRequest{
		Metadata: &content.Metadata{
			SourceId:    &sourceId,
			TraitIds:    []string{"bible.usx.chapter"},
			ContentType: "text/plain",
		},
	}

	for _, book := range bundle.Books() {
		for _, chapter := range book.Chapters {
			request.Metadata.Name = chapter.GetUsfm()
			text := &bytes.Buffer{}
			for _, verse := range chapter.FindVerses() {
				text.WriteString(verse.GetText())
			}
			request.Metadata.ContentLength = int64(text.Len())

			response, err := svc.AddMetadata(ctx, request)
			if err != nil {
				return err
			}
			err = common.SetTextContent(ctx, response.Id, text.String())
			if err != nil {
				return err
			}
			_, err = svc.AddMetadataRelationship(ctx, &content.AddMetadataRelationshipRequest{
				MetadataId1:  executionContext.Metadata.Id,
				MetadataId2:  response.Id,
				Relationship: "usx-chapter",
			})
			if err != nil {
				return nil
			}
			_, err = svc.AddCollectionItem(ctx, &content.AddCollectionItemRequest{
				CollectionId: collection.Id,
				ItemId: &content.AddCollectionItemRequest_ChildMetadataId{
					ChildMetadataId: response.Id,
				},
			})
			if err != nil {
				return nil
			}
		}
	}
}
