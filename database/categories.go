package database

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/database/querries"
	"forum/structs"
	"strings"
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
	res := make([]structs.Post, 0)
	// TODO This Query Retern category name as USERNAME
	rows, err := db.Query(querries.GetPostsbyCategoryL, user_id, c_name, structs.Limit, offset)
	if err != nil {
		return nil, errors.New("QuerryLatestPostsByCategory " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		var categories sql.NullString
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content,
			&post.LikeCount, &post.DislikeCount, &post.CommentCount, &post.CreatedAt,
			&post.UserName, &categories, &post.Liked)
		if categories.Valid {
			post.Categories = strings.Split(categories.String, "|")
		}
		if err != nil {
			return res, errors.New("QuerryLatestPostsByCategory failed to scan row: " + err.Error())
		}
		res = append(res, post)
		fmt.Printf("PID: %d, UID: %d, CONTENT: %12s, like:%d:%d , TIME:%15s, UName: %5s, categories %v %v\n", post.ID, post.UserID, post.Content, post.LikeCount, post.LikeCount, post.CreatedAt, post.UserName, post.Categories, post.Liked)
	}

	err = rows.Err()
	if err != nil {
		return res, errors.New("QuerryLatestPostsByCategory " + err.Error())
	}
	return res, nil
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
