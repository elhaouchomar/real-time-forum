package database

import (
	"database/sql"
	"fmt"
)

// Crate HasUserLikedPost in database/likes.go Based on the Satment Query Above
func HasUserLikedPost(db *sql.DB, userId, postId int, post bool) (bool, error) {
	var count int
	//	query := `SELECT COUNT(*) FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 1`
	query := ""
	if post {
		query = `SELECT COUNT(*) FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 1`
	} else {
		query = `SELECT COUNT(*) FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 1`
	}
	err := db.QueryRow(query, userId, postId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Crate LikePost in database/likes.go
func LikePost(db *sql.DB, userId, postId int, post bool) error {
	query := ""
	if post {
		query = `INSERT INTO post_likes (user_id, post_id, is_like) VALUES (?, ?, 1)`
	} else {
		query = `INSERT INTO comment_likes (user_id, comment_id, is_like) VALUES (?, ?, 1)`
	}
	_, err := db.Exec(query, userId, postId)
	fmt.Println(err)
	return err
}

// Crate UnlikePost in database/likes.go
func UnlikePost(db *sql.DB, userId, postId int, post bool) error {
	query := ""
	if post {
		query = `DELETE FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 1`
	} else {
		query = `DELETE FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 1`
	}
	_, err := db.Exec(query, userId, postId)
	return err
}

// Crate GetPostLikeCount in database/likes.go
func GetPostLikeCount(db *sql.DB, postId int, post bool) (int, error) {
	var count int
	query := ""
	if post {
		query = `SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 1`
	} else {
		query = `SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 1`
	}
	err := db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Crate HasUserDislikedPost in database/likes.go
func HasUserDislikedPost(db *sql.DB, userId, postId int, post bool) (bool, error) {
	var count int
	query := ""
	if post {
		query = `SELECT COUNT(*) FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 0`
	} else {
		query = `SELECT COUNT(*) FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 0`
	}
	err := db.QueryRow(query, userId, postId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Crate DislikePost
func DislikePost(db *sql.DB, userId, postId int, post bool) error {
	query := ""
	if post {
		query = `INSERT INTO post_likes (user_id, post_id, is_like) VALUES (?, ?, 0)`
	} else {
		query = `INSERT INTO comment_likes (user_id, comment_id, is_like) VALUES (?, ?, 0)`
	}
	_, err := db.Exec(query, userId, postId)
	return err
}

// Crate UndislikePost in
func UndislikePost(db *sql.DB, userId, postId int, post bool) error {
	query := ""
	if post {
		query = `DELETE FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 0`
	} else {
		query = `DELETE FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 0`
	}
	_, err := db.Exec(query, userId, postId)
	return err
}

// Crate GetPostDislikeCount in
func GetPostDislikeCount(db *sql.DB, postId int, post bool) (int, error) {
	var count int
	query := ""
	if post {
		query = `SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 0`
	} else {
		query = `SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 0`
	}
	err := db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// / in the query above, we need after each like to update the like_count in the post table
// Crate UpdatePostLikeCount in database/likes.go
func UpdatePostLikeCount(db *sql.DB, postId int, post bool) error {
	query := ""
	if post {
		query = `UPDATE posts SET like_count = (SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 1) WHERE id = ?`
	} else {
		query = `UPDATE comments SET like_count = (SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 1) WHERE id = ?`
	}
	_, err := db.Exec(query, postId, postId)
	return err
}

// the same for dislike
// Crate UpdatePostDislikeCount in database/likes.go
func UpdatePostDislikeCount(db *sql.DB, postId int, post bool) error {
	query := ""
	if post {
		query = `UPDATE posts SET dislike_count = (SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 0) WHERE id = ?`
	} else {
		query = `UPDATE comments SET dislike_count = (SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 0) WHERE id = ?`
	}
	_, err := db.Exec(query, postId, postId)
	return err
}

/* we Will Create Function To Like And Dislike Comments >>
We have 2 tables in database :
	"comments": `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY NOT NULL,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		like_count INTEGER DEFAULT 0,
		dislike_count INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		CONSTRAINT no_duplicates UNIQUE (post_id, user_id, content)
		);`,


			"comment_likes": `CREATE TABLE IF NOT EXISTS comment_likes (
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		is_like BOOLEAN NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, comment_id)
		);`,

*/
/////////// This section Bellong to Reaction in the Comments Like And Dislike /////////////
// // Crate HasUserLikedComment in database/likes.go

// func HasUserLikedComment(db *sql.DB, userId, commentId int) (bool, error) {
// 	var count int
// 	query := `SELECT COUNT(*) FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 1`
// 	err := db.QueryRow(query, userId, commentId).Scan(&count)
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }

// // Crate LikeComment in database/likes.go
// func LikeComment(db *sql.DB, userId, commentId int) error {
// 	query := `INSERT INTO comment_likes (user_id, comment_id, is_like) VALUES (?, ?, 1)`
// 	_, err := db.Exec(query, userId, commentId)
// 	return err
// }

// // Crate UnlikeComment in database/likes.go
// func UnlikeComment(db *sql.DB, userId, commentId int) error {
// 	query := `DELETE FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 1`
// 	_, err := db.Exec(query, userId, commentId)
// 	return err
// }

// // Crate GetCommentLikeCount in database/likes.go
// func GetCommentLikeCount(db *sql.DB, commentId int) (int, error) {
// 	var count int
// 	query := `SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 1`
// 	err := db.QueryRow(query, commentId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

// // Crate UpdateCommentLikeCount in database/likes.go
// func UpdateCommentLikeCount(db *sql.DB, commentId int) error {
// 	query := `UPDATE comments SET like_count = (SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 1) WHERE id = ?`
// 	_, err := db.Exec(query, commentId, commentId)
// 	return err
// }

// // Crate HasUserDislikedComment in database/likes.go
// func HasUserDislikedComment(db *sql.DB, userId, commentId int) (bool, error) {
// 	var count int
// 	query := `SELECT COUNT(*) FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 0`
// 	err := db.QueryRow(query, userId, commentId).Scan(&count)
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }

// // Crate DislikeComment in database/likes.go
// func DislikeComment(db *sql.DB, userId, commentId int) error {
// 	query := `INSERT INTO comment_likes (user_id, comment_id, is_like) VALUES (?, ?, 0)`
// 	_, err := db.Exec(query, userId, commentId)
// 	return err
// }

// // Crate UndislikeComment in database/likes.go
// func UndislikeComment(db *sql.DB, userId, commentId int) error {
// 	query := `DELETE FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 0`
// 	_, err := db.Exec(query, userId, commentId)
// 	return err
// }

// // Crate GetCommentDislikeCount in database/likes.go
// func GetCommentDislikeCount(db *sql.DB, commentId int) (int, error) {
// 	var count int
// 	query := `SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 0`
// 	err := db.QueryRow(query, commentId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

// // Crate UpdateCommentDislikeCount in database/likes.go
// func UpdateCommentDislikeCount(db *sql.DB, commentId int) error {
// 	query := `UPDATE comments SET dislike_count = (SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 0) WHERE id = ?`
// 	_, err := db.Exec(query, commentId, commentId)
// 	return err
// }
