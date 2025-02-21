package database

import (
	"database/sql"
	"errors"

	"forum/database/querries"
	"forum/structs"
)

// Either userName or userId
func GetUserProfile[T string | int](db *sql.DB, userSpecific T) (structs.Profile, error) {
	var userProfile structs.Profile
	// Query the user and their posts
	var rows *sql.Row
	var err error

	switch v := any(userSpecific).(type) {
	case string:
		rows = db.QueryRow(querries.GetUserProfileByUname, userSpecific)
	case int:
		rows = db.QueryRow(querries.GetUserProfileByID, v)
	}

	if rows.Err() != nil {
		return userProfile, errors.New("GetUserProfile 1 " + rows.Err().Error())
	}

	err = rows.Scan(&userProfile.UserID, &userProfile.UserName, &userProfile.CreatedAt, &userProfile.ArticleCount, &userProfile.CommentCount)
	if err != nil {
		return userProfile, errors.New("GetUserProfile 2 " + err.Error())
	}

	return userProfile, nil
}
