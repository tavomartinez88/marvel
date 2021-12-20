package dao

type CollaboratorsDao struct {
	ColID string `json:"colid"`
	LastSync string `json:"last_sync"`
	Editors []string `json:"editors,omitempty"`
	Writers []string `json:"writers,omitempty"`
	Colorists []string `json:"colorists,omitempty"`
}

func (c CollaboratorsDao) ID() (jsonField string, value interface{}) {
	value=c.ColID
	jsonField="colid"
	return
}
