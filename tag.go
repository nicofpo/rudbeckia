package rudbeckia

type Tag struct {
	Name     string `json:"name"`
	Category bool   `json:"category"`
	Lock     bool   `json:"lock"`
}
