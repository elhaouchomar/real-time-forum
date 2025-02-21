package database

import (
	"database/sql"

	"forum/structs"
)

func CreateView(db *sql.DB, view structs.View, IsPost bool) error {
	var stmt *sql.Stmt
	var err error
	if IsPost {
		stmt, err = db.Prepare("INSERT INTO post_likes(user_id, post_id, is_like) VALUES(?,?,?)")
	} else {
		stmt, err = db.Prepare("INSERT INTO comment_likes(user_id, comment_id, is_like) VALUES(?,?,?)")
	}
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(view.UserID, view.ID, view.IsLike)
	if err != nil {
		return err
	}
	return nil
}
