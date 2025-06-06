package order

import (
	"context"
	"encoding/json"

	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/domain/order"
	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/domain/protocol/httpclient"
)

type Service struct {
	HTTPClient httpclient.HttpInterface
}

func New(HTTPClient httpclient.HttpInterface) *Service {
	return &Service{
		HTTPClient: HTTPClient,
	}
}

func (r *Service) RetrieveOrders(ctx context.Context, page int, partnerKey string) ([]order.Order, error) {
	params := httpclient.RequestParams{
		Method:      "GET",
		URL:         "/orders",
		QueryParams: map[string]string{"page": string(page), "partner_key": partnerKey},
	}

	response, err := r.HTTPClient.MakeRequest(ctx, params)
	if err != nil {
		return nil, err
	}

	var ordersResponse order.OrdersResponse
	if err := json.Unmarshal(response.Body, &ordersResponse); err != nil {
		return nil, err
	}

	return ordersResponse.Data, nil
}
