package serializer

import "car/model"

// Pile 关于充电桩的返回信息
type Pile struct {
	Id               uint   `json:"id"`
	Status           int    `json:"status"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	ReserveStartTime string `json:"reserve_start_time"`
	ReserveEndTime   string `json:"reserve_end_time"`
}

// BuildPile 序列化充电桩
func BuildPile(pile model.ChargingPile) Pile {
	return Pile{
		Id:               pile.ID,
		Status:           pile.Status,
		StartTime:        pile.StartTime,
		EndTime:          pile.EndTime,
		ReserveStartTime: pile.ReserveStartTime,
		ReserveEndTime:   pile.ReserveEndTime,
	}
}

func BuildPiles(items []model.ChargingPile) (piles []Pile) {
	for _, item := range items {
		pile := BuildPile(item)
		piles = append(piles, pile)
	}
	return piles
}
