-- name: IsUserExists :one
SELECT
    EXISTS (
        SELECT username FROM t_user WHERE username = @username
    );