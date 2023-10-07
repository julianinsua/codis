-- name: CreatePost :one
INSERT INTO posts (
	title,
	description,
	category_id,
	path,
	status
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts WHERE status=$1;

-- name: GetPost :one
SELECT * FROM posts WHERE id=$1 LIMIT 1;

-- name: GetCategoryPosts :many
SELECT * FROM posts WHERE category_id=$1;

-- name: GetTagPosts :many
SELECT * FROM posts P
WHERE EXISTS (
	SELECT 1 FROM post_tags PT
	WHERE P.id = PT.post_id 
		AND PT.tag_id = $1
); -- need to add ordering
