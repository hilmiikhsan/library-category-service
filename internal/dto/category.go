package dto

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description"`
}

type GetDetailCategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetListCategoryResponse struct {
	CategoryList []Category `json:"category_list"`
	Pagination   Pagination `json:"pagination"`
}

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
