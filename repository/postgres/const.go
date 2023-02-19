package postgres

var (
	SortQuery = map[string]string{
		"1": " ORDER BY created_at DESC ",
		"2": " ORDER BY price ASC ",
		"3": " ORDER BY price DESC ",
		"4": " ORDER BY name ASC ",
		"5": " ORDER BY name DESC ",
	}
)
