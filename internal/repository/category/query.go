package category

const (
	queryInsertNewCategory = `
		INSERT INTO categories
		(
			name,
			description
		) VALUES (?, ?)
	`

	queryFindCategoryByName = `
		SELECT
			id,
			name,
			description
		FROM categories
		WHERE name = ?
	`
)
