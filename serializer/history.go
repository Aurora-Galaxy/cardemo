package serializer

import "car/model"

// History_Record 关于历史记录的返回信息
type History_Record struct {
	Date   string `json:"date"`
	Time   string `json:"time"`
	PileId uint   `json:"pile_id"`
}

func BuildHistory(history model.HistoryRecord) History_Record {
	return History_Record{
		Date:   history.Date,
		Time:   history.Time,
		PileId: history.PileId,
	}
}

func BuildHistorys(items []model.HistoryRecord) (historys []History_Record) {
	for _, item := range items {
		history := BuildHistory(item)
		historys = append(historys, history)
	}
	return historys
}
