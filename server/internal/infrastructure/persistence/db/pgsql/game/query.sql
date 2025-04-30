-- name: CreateGame :one
INSERT INTO games (secret_word, max_attempts, word_length)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetActiveGame :one
SELECT *
FROM games
WHERE is_active = true
ORDER BY created_at DESC
LIMIT 1;

-- name: FindGameById :one
SELECT *
FROM games
WHERE id = $1;

-- name: CreateGuess :one
INSERT INTO guesses (game_id, word, attempt_number, session_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetGameGuesses :many
SELECT *
FROM guesses
WHERE game_id = $1
  AND session_id = $2
ORDER BY attempt_number ASC;

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

-- name: GetRandomSecretWord :one
SELECT word
FROM word_pool
WHERE is_answer = true
ORDER BY RANDOM()
LIMIT 1;

-- name: IsValidGuess :one
SELECT EXISTS(SELECT 1
              FROM word_pool
              WHERE word = $1
                AND is_valid = true);

-- name: AddWordToPool :one
INSERT INTO word_pool (word, is_answer, is_valid)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAllAnswerWords :many
SELECT word
FROM word_pool
WHERE is_answer = true;

-- name: GetAllValidWords :many
SELECT word
FROM word_pool
WHERE is_valid = true;

-- name: FindWord :one
SELECT *
FROM word_pool
WHERE is_valid = true
  AND word = $1;