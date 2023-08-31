package models

type Segment struct {
	ID   int    `json:"segmentId" uri:"id"`
	Name string `json:"name" uri:"slug" binding:"required"`
}
