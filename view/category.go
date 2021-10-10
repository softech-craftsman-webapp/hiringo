package view

type CategoryView struct {
	ID          string `json:"id"`
	CreatedBy   string `json:"created_by"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryEmptyView struct {
	ID string `json:"id"`
}
