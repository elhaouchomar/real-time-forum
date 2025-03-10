package database

import (
	"database/sql"
	"forum/database/querries"
	"forum/structs"
)

func GetCategoriesWithPostCount(db *sql.DB) (map[string]int, error) {
	categories := make(map[string]int)
	rows, err := db.Query(`
		SELECT c.name, COUNT(pc.post_id) as post_count
		FROM categories c
		LEFT JOIN post_categories pc ON c.id = pc.category_id
		GROUP BY c.name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var postCount int
		if err := rows.Scan(&name, &postCount); err != nil {
			return nil, err
		}
		categories[name] = postCount
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func QuerryLatestPostsByCategory(db *sql.DB, user_id int, c_name string, offset int) ([]structs.Post, error) {
	return QueryPosts(db, querries.GetPostsbyCategoryL, user_id, c_name, structs.Limit, offset)
}

func IsCategoryValid(category string) bool {
	categories := []string{"General", "Entertainment", "Health", "Business", "Sports", "Technology"}
	for _, c := range categories {
		if c == category {
			return true
		}
	}
	return false
}
