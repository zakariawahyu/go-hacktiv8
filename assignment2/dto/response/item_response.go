package response

import "github.com/zakariawahyu/go-hacktiv8/assignment2/entity"

type ItemsResponse struct {
	ItemID      int64  `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func NewItemsResponseArray(items []entity.Items) []ItemsResponse {
	itemsResp := []ItemsResponse{}
	for _, value := range items {
		item := ItemsResponse{
			ItemID:      value.ItemID,
			ItemCode:    value.ItemCode,
			Description: value.Description,
			Quantity:    value.Quantity,
		}
		itemsResp = append(itemsResp, item)
	}
	return itemsResp
}
