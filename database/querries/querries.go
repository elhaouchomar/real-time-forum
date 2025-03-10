package querries

const (
	GetLatestPostsL = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
	    COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		GROUP BY p.id
		ORDER BY p.created_at DESC
		LIMIT ?
		OFFSET ?;`
	GetPostsbyUserL = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE u.username = ?
		GROUP BY p.id
		ORDER BY p.created_at DESC 
		LIMIT ?
		OFFSET ?;`
	GetPostsByMostLiked = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
	    COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		GROUP BY p.id
		ORDER BY p.like_count DESC
		LIMIT ?
		OFFSET ?;`
	GetPostsbyCategoryL = `SELECT p.*, u.username, GROUP_CONCAT(c.name, "|") AS categories,
        COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE c.name = ?
		GROUP BY p.id
		ORDER BY p.created_at DESC 
		LIMIT ?
		OFFSET ?;`
	GetPostsbyUserLikeL = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ? AND is_like = 1
		WHERE pl.is_like = 1
		GROUP BY p.id
		ORDER BY p.created_at DESC 
		LIMIT ?
		OFFSET ?;`
	GetPostByID = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE p.id = ?;`
	GetCommentsByPostL = `SELECT c.*, u.username , COALESCE(cl.is_like, "null") AS is_like
		FROM comments c 
		JOIN users u ON c.user_id = u.id 
		LEFT JOIN comment_likes cl ON c.id = cl.comment_id AND cl.user_id = ?
		WHERE post_id=?
		ORDER BY c.created_at DESC LIMIT ?`
	GetCommentsByID = `SELECT c.*, u.username , COALESCE(cl.is_like, "null") AS is_like
		FROM comments c 
		JOIN users u ON c.user_id = u.id 
		LEFT JOIN comment_likes cl ON c.id = cl.comment_id AND cl.user_id = ?
		WHERE id=?
		ORDER BY c.created_at DESC LIMIT ?`
	GetUserProfileByUname = `SELECT u.id, u.username, u.created_at, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count
		FROM users u
		LEFT JOIN posts p ON u.id = p.user_id
		LEFT JOIN comments c ON u.id = c.user_id
		WHERE u.username = ?
		GROUP BY u.id, u.username, u.created_at;`
	GetUserProfileByID = `SELECT u.id, u.username, u.created_at, COUNT(DISTINCT p.id) AS post_count, COUNT(DISTINCT c.id) AS comment_count
		FROM users u
		LEFT JOIN posts p ON u.id = p.user_id
		LEFT JOIN comments c ON u.id = c.user_id
		WHERE u.id = ?
		GROUP BY u.id, u.username, u.created_at;`
	GetCategoriesWithPostCount = `SELECT c.name, COUNT(pc.post_id) as post_count
			FROM categories c
			JOIN post_categories pc ON c.id = pc.category_id
			GROUP BY c.name 
			ORDER BY post_count DESC
			LIMIT 6`
			
	// Message queries
	GetAllMessages = `
				SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
				FROM messages m 
				JOIN users u ON m.author_id = u.id 
				ORDER BY m.timestamp DESC
			`

	GetLatestMessages = `
				SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
				FROM messages m 
				JOIN users u ON m.author_id = u.id 
				ORDER BY m.timestamp DESC 
				LIMIT ? OFFSET ?
			`

	GetMessageByID = `
				SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
				FROM messages m 
				JOIN users u ON m.author_id = u.id 
				WHERE m.id = ?
			`

	CreateMessage = `
				INSERT INTO messages (content, author_id) 
				VALUES (?, ?)
			`

	GetMessagesBetweenUsers = `
				SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
				FROM messages m 
				JOIN users u ON m.author_id = u.id 
				WHERE (m.author_id = ? AND m.recipient_id = ?) 
				   OR (m.author_id = ? AND m.recipient_id = ?) 
				ORDER BY m.timestamp DESC
			`
)
