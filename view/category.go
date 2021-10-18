package view

import "hiringo/model"

type CategoryView struct {
	ID          string `json:"id"`
	CreatedByID string `json:"created_by_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryEmptyView struct {
	ID string `json:"id"`
}

type PublicCategoryView struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CategoryModelToView(category model.Category) PublicCategoryView {
	return PublicCategoryView{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
