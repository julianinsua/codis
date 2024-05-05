-- +goose Up
CREATE VIEW posts_view AS 
WITH nt AS (SELECT id, name, user_id AS "userId" FROM tags)
SELECT p.id, p.title, p.description, p.status, p.user_id, JSON_AGG(nt.*) AS post_tags 
FROM posts p 
JOIN post_tags pt on pt.post_id = p.id 
JOIN nt on nt.id = pt.tag_id
GROUP BY p.id, p.title, p.description, p.status, p.user_id;

-- +goose Down
DROP VIEW IF EXISTS posts_view;
