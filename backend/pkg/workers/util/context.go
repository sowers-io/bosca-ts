package util

import (
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search/factory"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/temporal"
	"bosca.io/pkg/util"
	"bosca.io/pkg/workers/common"
	rootContext "context"
	"github.com/qdrant/go-client/qdrant"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAITemporalClient() (client.Client, error) {
	ctx := rootContext.Background()

	cfg := configuration.NewWorkerConfiguration()
	connection, err := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	if err != nil {
		return nil, err
	}
	contentService := content.NewContentServiceClient(connection)

	searchClient, err := factory.NewSearch(cfg.Search)
	if err != nil {
		return nil, err
	}

	qdrantClient, err := qdrant.NewQdrantClient(cfg.ClientEndPoints.QdrantApiAddress)
	if err != nil {
		return nil, err
	}

	_, err = qdrantClient.CollectionsClient.Get(ctx, &go_client.GetCollectionInfoRequest{
		CollectionName: qdrant.MetadataIndex,
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if s.Code() == codes.NotFound {
				collection := &go_client.CreateCollection{
					CollectionName: qdrant.MetadataIndex,
					VectorsConfig: &go_client.VectorsConfig{
						Config: &go_client.VectorsConfig_Params{
							Params: &go_client.VectorParams{
								Size:     4096,
								Distance: go_client.Distance_Cosine,
							},
						},
					},
				}
				result, err := qdrantClient.CollectionsClient.Create(ctx, collection)
				if err != nil {
					return nil, err
				}
				if !result.Result {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	}

	httpClient := util.NewDefaultHttpClient()
	propagator := common.NewContextPropagator(cfg, httpClient, contentService, searchClient, qdrantClient)
	return temporal.NewClientWithPropagator(ctx, cfg.ClientEndPoints, propagator)
}
