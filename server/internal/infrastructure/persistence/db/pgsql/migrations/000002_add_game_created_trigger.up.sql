-- Create a function that will be called by the trigger
CREATE OR REPLACE FUNCTION notify_game_created()
    RETURNS trigger AS
$$
BEGIN
    -- Construct the payload with game information
    PERFORM pg_notify(
            'game_created',
            json_build_object(
                    'id', NEW.id,
                    'word_id', NEW.word_id,
                    'word_length', NEW.word_length,
                    'max_attempts', NEW.max_attempts,
                    'created_at', NEW.created_at
                )::text
        );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER game_created_trigger
    AFTER INSERT
    ON games
    FOR EACH ROW
EXECUTE FUNCTION notify_game_created(); 