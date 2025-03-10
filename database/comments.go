package database

import (
	"database/sql"
	"fmt"

	"forum/database/querries"
	"forum/structs"
)

func CreateComment(db *sql.DB, UserId, PostId int, content string) (error, int) {
	stmt, err := db.Prepare("INSERT INTO comments(user_id, post_id, content) VALUES(?,?,?)")
	if err != nil {
		return err, 0
	}
	defer stmt.Close()

	res, err := stmt.Exec(UserId, PostId, content)
	if err != nil {
		return err, 0
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, int(id)
}

func GetCommentsByPost(db *sql.DB, userId, postId int) ([]structs.Comment, error) {
	res := make([]structs.Comment, 0)
	rows, err := db.Query(querries.GetCommentsByPostL, userId, postId, -1)
	if err != nil {
		fmt.Println("GetCommentsByPost=====>", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment structs.Comment
		err := rows.Scan(
			&comment.ID,           // c.id
			&comment.PostID,       // c.post_id
			&comment.UserID,       // c.user_id
			&comment.Content,      // c.content
			&comment.LikeCount,    // c.like_count
			&comment.DislikeCount, // c.dislike_count
			&comment.CreatedAt,    // c.created_at
			&comment.UserName,     // u.username
			&comment.Liked,        // COALESCE(cl.is_like, 'null')
		)
		if err != nil {
			return res, err
		}
		res = append(res, comment)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetCommentById(db *sql.DB, UserId, commentId int) (structs.Comment, error) {
	var comment structs.Comment
	err := db.QueryRow(querries.GetCommentsByID, UserId, commentId, structs.Limit).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content,
		&comment.LikeCount, &comment.DislikeCount, &comment.CreatedAt, &comment.UserName)
	if err == sql.ErrNoRows {
		return comment, fmt.Errorf("comment with ID %d not found", commentId)
	}
	return comment, err
}
