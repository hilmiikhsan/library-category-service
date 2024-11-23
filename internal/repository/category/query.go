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

	queryFindCategoryByID = `
		SELECT
			id,
			name,
			description
		FROM categories
		WHERE id = ?
	`

	queryFindAllCategory = `
		SELECT
			id,
			name,
			description
		FROM categories
		ORDER BY updated_at DESC
		LIMIT ?
		OFFSET ?
	`

	queryUpdateNewCategory = `
		UPDATE categories
		SET
			name = ?,
			description = ?,
			updated_at = NOW()
		WHERE id = ?
	`
)
