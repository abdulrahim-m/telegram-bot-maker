CREATE INDEX idx_commands_bot_id ON commands(bot_id);
CREATE INDEX idx_messages_command_id ON messages(command_id);
CREATE INDEX idx_files_category_id ON files(category_id);
CREATE INDEX idx_user_commands_user_id ON user_commands(user_id);