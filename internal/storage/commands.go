package storage

var Indices = []string{
	IdxCommandsBotID,
	IdxMessagesCommandID,
	IdxFilesCategoryID,
	IdxUserCommandsUserID,
}

const (
	IdxCommandsBotID      = `CREATE INDEX IF NOT EXISTS idx_commands_bot_id ON commands(bot_id);`
	IdxMessagesCommandID  = `CREATE INDEX IF NOT EXISTS idx_messages_command_id ON messages(command_id);`
	IdxFilesCategoryID    = `CREATE INDEX IF NOT EXISTS idx_files_category_id ON files(category_id);`
	IdxUserCommandsUserID = `CREATE INDEX IF NOT EXISTS idx_user_commands_user_id ON user_commands(user_id);`
)
