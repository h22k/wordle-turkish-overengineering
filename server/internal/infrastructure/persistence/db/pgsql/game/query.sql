-- name: CreateGame :one
INSERT INTO games (word_id, max_attempts, word_length)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetActiveGame :one
SELECT g.*, w.word as secret_word
FROM games g
         JOIN word_pool w ON w.id = g.word_id
WHERE g.is_active = true
ORDER BY g.created_at DESC
LIMIT 1;

-- name: FindGameById :one
SELECT g.*, w.word as secret_word
FROM games g
         JOIN word_pool w ON w.id = g.word_id
WHERE g.id = $1;

-- name: CreateGuess :one
INSERT INTO guesses (game_id, word, attempt_number, session_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetGameGuesses :many
SELECT *
FROM guesses
WHERE game_id = $1
  AND session_id = $2
ORDER BY attempt_number;

-- name: GetGameGuessesCount :one
SELECT COUNT(id)
FROM guesses
WHERE game_id = $1
  AND session_id = $2;

-- name: MakeGameInactive :one
UPDATE games
SET is_active  = false,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetRandomWord :one
SELECT id, word
FROM word_pool
ORDER BY RANDOM()
LIMIT 1;

-- name: IsValidGuess :one
SELECT EXISTS(SELECT 1
              FROM word_pool
              WHERE word = $1);

-- name: AddWordToPool :one
INSERT INTO word_pool (word)
VALUES ($1)
RETURNING *;

-- name: GetAllWords :many
SELECT word
FROM word_pool;

-- name: FindWord :one
SELECT *
FROM word_pool
WHERE word = $1;

-- name: IsWordExists :one
SELECT EXISTS(SELECT 1
              FROM word_pool
              WHERE word = $1);