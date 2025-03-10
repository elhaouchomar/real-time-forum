package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"forum/database/querries"
	"forum/structs"
)

func QueryPosts(db *sql.DB, query string, args ...interface{}) ([]structs.Post, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, errors.New("QueryPosts " + err.Error())
	}
	defer rows.Close()

	var res []structs.Post
	for rows.Next() {
		var post structs.Post
		var categories sql.NullString
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content,
			&post.LikeCount, &post.DislikeCount, &post.CommentCount, &post.CreatedAt,
			&post.UserName, &categories, &post.Liked)
		if err != nil {
			return res, errors.New("QueryPosts failed to scan row: " + err.Error())
		}
		if categories.Valid {
			post.Categories = strings.Split(categories.String, "|")
		}
		res = append(res, post)
	}
	if err = rows.Err(); err != nil {
		return res, errors.New("QueryPosts " + err.Error())
	}
	return res, nil
}

func QuerryLatestPostsByUserLikes(db *sql.DB, user_id, ammount, offset int) ([]structs.Post, error) {
	return QueryPosts(db, querries.GetPostsbyUserLikeL, user_id, ammount, offset)
}

func QuerryMostLikedPosts(db *sql.DB, user_id, ammount, offset int) ([]structs.Post, error) {
	return QueryPosts(db, querries.GetPostsByMostLiked, user_id, ammount, offset)
}

func QuerryLatestPosts(db *sql.DB, user_id, ammount, offset int) ([]structs.Post, error) {
	return QueryPosts(db, querries.GetLatestPostsL, user_id, ammount, offset)
}

func QuerryPostsbyUser(db *sql.DB, username string, user_id, ammount, offset int) ([]structs.Post, error) {
	return QueryPosts(db, querries.GetPostsbyUserL, user_id, username, ammount, offset)
}

func CreatePost(db *sql.DB, UserID int, title, content string, categories []string) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO posts(user_id, title, content) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(UserID, title, content)
	if err != nil {
		return 0, err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	stmt_1, err := tx.Prepare(`INSERT INTO post_categories(category_id, post_id) VALUES(?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt_1.Close()

	for _, category := range categories {
		CategoryId, err := strconv.Atoi(category)
		if err != nil {
			return 0, err
		}
		_, err = stmt_1.Exec(CategoryId, postID)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(postID), nil
}

func GetPostByID(db *sql.DB, Postid, UserID int) (structs.Post, error) {
	var post structs.Post
	var categories sql.NullString
	err := db.QueryRow(querries.GetPostByID, UserID, Postid).Scan(
		&post.ID, &post.UserID, &post.Title, &post.Content,
		&post.LikeCount, &post.DislikeCount, &post.CommentCount,
		&post.CreatedAt, &post.UserName, &categories, &post.Liked)
	if err == sql.ErrNoRows {
		return post, fmt.Errorf("database GetPostById 1:%v", "post not found")
	} else if err != nil {
		return post, fmt.Errorf("database GetPostById 2:%v", err)
	}

	if categories.Valid {
		post.Categories = strings.Split(categories.String, "|")
	}
	return post, nil
}
