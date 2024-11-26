-- name: InsertIntoTask :one
INSERT INTO task(title, description, status)
VALUES(@title, @description, @status) RETURNING id;

-- name: SelectTasks :many
SELECT t.id task_id,
       t.title task_title,
       t.description task_description,
       t.status task_status,
       t.created_at task_created_at,
       t.updated_at task_updated_at
FROM task t
WHERE deleted_at IS NULL
ORDER BY t.title ASC;

-- name: SelectTaskByID :one
SELECT t.id task_id,
       t.title task_title,
       t.description task_description,
       t.status task_status,
       t.created_at task_created_at,
       t.updated_at task_updated_at
FROM task t
WHERE t.id = @task_id AND deleted_at IS NULL
ORDER BY t.title ASC;

-- name: UpdateTask :exec
UPDATE task SET title = @task_title, description = @task_description,
                status = @task_status, updated_at = CURRENT_TIMESTAMP
WHERE id = @task_id AND deleted_at IS NULL;

-- name: DeleteTask :exec
UPDATE task SET deleted_at = CURRENT_TIMESTAMP
WHERE id = @task_id AND deleted_at IS NULL;