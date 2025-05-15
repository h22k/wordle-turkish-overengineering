CREATE TABLE word_pool
(
    id   SERIAL PRIMARY KEY,
    word VARCHAR(7) NOT NULL UNIQUE
);

CREATE TABLE games
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    word_id      INT                      NOT NULL REFERENCES word_pool (id),
    word_length  INT                      NOT NULL,
    max_attempts INT                      NOT NULL,
    is_active    BOOLEAN                  NOT NULL DEFAULT TRUE,

    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE guesses
(
    id             UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    game_id        UUID                     NOT NULL REFERENCES games (id) ON DELETE CASCADE,
    word           VARCHAR(7)               NOT NULL,
    attempt_number INT                      NOT NULL,
    session_id     VARCHAR(255)             NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE (game_id, attempt_number, session_id),
    UNIQUE (game_id, session_id, word)
);

CREATE INDEX idx_guesses_game_session ON guesses(game_id, session_id);

CREATE INDEX idx_word_pool_word ON word_pool(word);

CREATE INDEX idx_games_word_id ON games(word_id);

CREATE INDEX idx_games_is_active ON games(is_active) WHERE is_active = true;