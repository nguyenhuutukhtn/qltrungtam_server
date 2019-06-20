package model

type Course struct {
	Id            int    `json:"Id,omitempty"`
	Name          string `json:"Name,omitempty"`
	Type          string `json:"Type"`
	StartDay      string `json:"StartDay"`
	Teacher       string `json:"Teacher"`
	FinishDay     string `json:"FinishDay"`
	StudentNumber int    `json:"StudentNumber"`
	Fee           int    `json:"Fee"`
	DiscountInfo  string `json:"DiscountInfo"`
}
