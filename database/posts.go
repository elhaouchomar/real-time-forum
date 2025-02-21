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

func QuerryLatestPostsByUserLikes(db *sql.DB, user_id, ammount, offset int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query(querries.GetPostsbyUserLikeL, user_id, ammount, offset)
	if err != nil {
		return nil, errors.New("QuerryLatestPosts " + err.Error())
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
			return res, errors.New("QuerryLatestPosts failed to scan row: " + err.Error())
		}
		res = append(res, post)
		// fmt.Printf("PID: %d, UID: %d, CONTENT: %12s, like:%d:%d , TIME:%15s, UName: %5s, categories %v %v\n", post.ID, post.UserID, post.Content, post.LikeCount, post.LikeCount, post.CreatedAt, post.UserName, post.Categories, post.Liked)
	}

	err = rows.Err()
	if err != nil {
		return res, errors.New("QuerryLatestPosts " + err.Error())
	}
	return res, nil
}


func QuerryMostLikedPosts(db *sql.DB, user_id, ammount, offset int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query(querries.GetPostsByMostLiked, user_id, ammount, offset)
	if err != nil {
		return nil, errors.New("QuerryLatestPosts " + err.Error())
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
			return res, errors.New("QuerryLatestPosts failed to scan row: " + err.Error())
		}
		res = append(res, post)
		// fmt.Printf("PID: %d, UID: %d, CONTENT: %12s, like:%d:%d , TIME:%15s, UName: %5s, categories %v %v\n", post.ID, post.UserID, post.Content, post.LikeCount, post.LikeCount, post.CreatedAt, post.UserName, post.Categories, post.Liked)
	}

	err = rows.Err()
	if err != nil {
		return res, errors.New("QuerryLatestPosts " + err.Error())
	}
	return res, nil
}

func QuerryLatestPosts(db *sql.DB, user_id, ammount, offset int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query(querries.GetLatestPostsL, user_id, ammount, offset)
	if err != nil {
		return nil, errors.New("QuerryLatestPosts " + err.Error())
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
			return res, errors.New("QuerryLatestPosts failed to scan row: " + err.Error())
		}
		res = append(res, post)
		fmt.Printf("PID: %d, UID: %d, CONTENT: %12s, like:%d:%d , TIME:%15s, UName: %5s, categories %v %v\n", post.ID, post.UserID, post.Content, post.LikeCount, post.LikeCount, post.CreatedAt, post.UserName, post.Categories, post.Liked)
	}

	err = rows.Err()
	if err != nil {
		return res, errors.New("QuerryLatestPosts " + err.Error())
	}
	return res, nil
}

func QuerryPostsbyUser(db *sql.DB, username string, user_id, ammount, offset int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query(querries.GetPostsbyUserL, user_id, username, ammount, offset)
	if err != nil {
		return nil, errors.New("QuerryPostsbyUser " + err.Error())
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
			return res, errors.New("QuerryPostsbyUser " + err.Error())
		}
		res = append(res, post)
		fmt.Printf("Post ID: %d, User ID: %d, Title: %15s, Content: %15s, Created At: %s\n", post.ID, post.UserID, post.Title, post.Content, post.CreatedAt)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.New("QuerryPostsbyUser " + err.Error())
	}
	return res, nil
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

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Println("CreatePost 7", err)
		return 0, err
	}
	PostId := int(postID)
	return PostId, nil
}

func GetPostByID(db *sql.DB, Postid, UserID int) (structs.Post, error) {
	var post structs.Post
	var categories sql.NullString
	err := db.QueryRow(querries.GetPostByID, UserID, Postid).Scan( // TODO Is Liked Return NULL
		&post.ID, &post.UserID, &post.Title, &post.Content, // instead of real value
		&post.LikeCount, &post.DislikeCount, &post.CommentCount, // related to POST.html page
		&post.CreatedAt, &post.UserName, &categories,
		&post.Liked)
	if err == sql.ErrNoRows {
		return post, fmt.Errorf("databse GetPostById 1:%v", "post not found")
	} else if err != nil {
		return post, fmt.Errorf("databse GetPostById 2:%v", err)
	}

	if categories.Valid {
		post.Categories = strings.Split(categories.String, "|")
	}
	return post, nil
}
