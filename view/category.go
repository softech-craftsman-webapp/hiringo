package view

type CategoryView struct {
	ID          string `json:"id"`
	CreatedByID string `json:"created_by_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryEmptyView struct {
	ID string `json:"id"`
}
