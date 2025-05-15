-- Drop the trigger first
DROP TRIGGER IF EXISTS game_created_trigger ON games;

-- Then drop the function
DROP FUNCTION IF EXISTS notify_game_created(); 