// Code generated by jzero. DO NOT EDIT.
// type: version

package simpleapi

import (
	"context"
	"simpleapi/simpleapi-go/model/simpleapi/types"

	"github.com/jzero-io/restc"
)

var (
	_ = context.Background()
)

type VersionGetter interface {
	Version() VersionInterface
}

type VersionInterface interface {
	// API /api/v1/version
	GetVersion(ctx context.Context, param types.GetVersionRequest) (*types.GetVersionResponse, error)

	VersionExpansion
}

type versionClient struct {
	client restc.Interface
}

func newVersionClient(c *SimpleapiClient) *versionClient {
	return &versionClient{
		client: c.RESTClient(),
	}
}

func (x *versionClient) GetVersion(ctx context.Context, param types.GetVersionRequest) (*types.GetVersionResponse, error) {
	var resp *types.GetVersionResponse
	err := x.client.Verb("GET").
		SubPath(
			"/api/v1/version",
		).
		Params().
		Body(nil).
		Do(ctx).
		Into(&resp, true)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
