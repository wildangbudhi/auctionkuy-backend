package transaction

import "encoding/json"

type TransactionStatus struct {
	ID            *int    `json:"id"`
	SellerCommand *string `json:"seller_command"`
	BuyerCommand  *string `json:"buyer_command"`
	SellerStep    *int    `json:"seller_step"`
	BuyerStep     *int    `json:"buyer_step"`
	SellerStepMax *int    `json:"seller_step_max"`
	BuyerStepMax  *int    `json:"buyer_step_max"`
}

func (obj *TransactionStatus) MarshalJSON() ([]byte, error) {

	if (TransactionStatus{ID: obj.ID} == *obj) {
		return json.Marshal(obj.ID)
	}

	return json.Marshal(obj)

}
