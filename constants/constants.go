package constants

var (
	INSERT_POST             = `INSERT INTO posts (title, description) VALUES ($1, $2)`
	GET_POST_BY_ID          = `SELECT id, title, description FROM posts WHERE id = $1`
	GET_COMMENTS_BY_POST_ID = `SELECT id, author, message FROM comments WHERE post_id = $1 ORDER BY created_at DESC`
	INSERT_COMMENTS         = `INSERT INTO comments (post_id, author, message) VALUES ($1, $2, $3) RETURNING id`
)

var (
	HTML_COMMENT = `<!DOCTYPE html>
    <html>
    <head>
        <title>Comments</title>
        <style>
            body { font-family: Arial, sans-serif; margin: 20px; }
            .comment { border: 1px solid #ddd; padding: 10px; margin-bottom: 10px; }
            .author { font-weight: bold; margin-bottom: 5px; }
        </style>
    </head>
    <body>
        <h1>Comments for Post ID:  %s  </h1>`
)
