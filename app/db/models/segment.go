package models

type Segment struct {
	ID   int    `json:"segmentId" uri:"id"`
	Name string `json:"name" uri:"slug" binding:"required"`
}

type Update struct {
	Add    []string `json:"add"`
	Delete []string `json:"delete"`
}

type UpdateStats struct {
	Processed []string `json:"processed"`
	Skipped   []string `json:"skipped"`
	Failed    []string `json:"failed"`
}

type Response struct {
	Added   *UpdateStats `json:"added"`
	Deleted *UpdateStats `json:"deleted"`
}
