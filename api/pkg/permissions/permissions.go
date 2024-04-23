package permissions

import (
	"bosca.io/pkg/configuration"
	ory "github.com/ory/client-go"
)

const PermissionsNamespace = "bosca"

func NewOryClient(cfg *configuration.ServerConfiguration) *ory.APIClient {
	orycfg := ory.NewConfiguration()
	orycfg.Servers = []ory.ServerConfiguration{}
	orycfg.OperationServers = map[string]ory.ServerConfigurations{
		"PermissionAPIService.CheckPermission": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.ReadEndPoint,
			},
		},
		"PermissionAPIService.CheckPermissionOrError": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.ReadEndPoint,
			},
		},
		"RelationshipAPIService.ExpandPermissions": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.ReadEndPoint,
			},
		},
		"RelationshipAPIService.ListRelationshipNamespaces": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.ReadEndPoint,
			},
		},
		"RelationshipAPIService.GetRelationships": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.WriteEndPoint,
			},
		},
		"RelationshipAPIService.CreateRelationship": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.WriteEndPoint,
			},
		},
		"RelationshipAPIService.PatchRelationships": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.WriteEndPoint,
			},
		},
		"RelationshipAPIService.DeleteRelationships": []ory.ServerConfiguration{
			{
				URL: cfg.Permissions.WriteEndPoint,
			},
		},
	}
	return ory.NewAPIClient(orycfg)
}
