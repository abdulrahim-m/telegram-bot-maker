package storage

var Tables = []string{
	TableAdmins,
	TableBots,
	TableCommands,
	TableMessages,
	TableKeyboards,
	TableReplayKeyboards,
	TableSettingsMenu,
	TableSettings,
	TableCategories,
	TableFiles,
	TableKeyboardButtons,
	TableReplyKeyboardButtons,
	TableMessageFiles,
	TableUserCommands,
	TableUsers,
	TableLogs,
}

const (
	// ####### MAIN ENTITIES #######
	TableAdmins = `
		CREATE TABLE IF NOT EXISTS admins (
			id INTEGER PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	TableBots = `
			CREATE TABLE IF NOT EXISTS bots (
			id INTEGER PRIMARY KEY,
			token TEXT UNIQUE NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	// ####### BOT-RELATED ENTITIES #######
	TableCommands = `
		CREATE TABLE IF NOT EXISTS commands (
			id INTEGER PRIMARY KEY,
			bot_id INTEGER NOT NULL,
			name TEXT UNIQUE NOT NULL,
			description TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	TableMessages = `
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY,
			command_id INTEGER NOT NULL,
			text TEXT,
			media_type TEXT,
			media_storage_key TEXT,
			FOREIGN KEY (command_id) REFERENCES commands(id)
		);
	`
	TableKeyboards = `
		CREATE TABLE IF NOT EXISTS keyboards (
			id INTEGER PRIMARY KEY,
			message_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			FOREIGN KEY (message_id) REFERENCES messages(id)
		);
	`
	TableReplayKeyboards = `
		CREATE TABLE IF NOT EXISTS reply_keyboards (
			id INTEGER PRIMARY KEY,
			message_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			FOREIGN KEY (message_id) REFERENCES messages(id)
		);
	`
	// ####### SETTINGS #######
	TableSettingsMenu = `
		CREATE TABLE IF NOT EXISTS settings_menu (
			id INTEGER PRIMARY KEY,
			bot_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			FOREIGN KEY (bot_id) REFERENCES bots(id)
		);
	`
	TableSettings = `
		CREATE TABLE IF NOT EXISTS settings (
			id INTEGER PRIMARY KEY,
			menu_id INTEGER NOT NULL,
			key TEXT UNIQUE NOT NULL,
			value TEXT NOT NULL,
			type TEXT DEFAULT 'string',
			FOREIGN KEY (menu_id) REFERENCES settings_menu(id)
		);
	`
	// ####### FILE MANAGEMENT #######
	TableCategories = `
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY,
			name TEXT UNIQUE NOT NULL
		);
	`
	TableFiles = `
		CREATE TABLE IF NOT EXISTS files (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			storage_key TEXT NOT NULL,
			category_id INTEGER,
			size INTEGER,
			telegram_file_id TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);
	`
	// ####### RELATION TABLES #######
	TableKeyboardButtons = `
		CREATE TABLE IF NOT EXISTS keyboard_buttons (
			id INTEGER PRIMARY KEY,
			keyboard_id INTEGER NOT NULL,
			text TEXT NOT NULL,
			action_type TEXT NOT NULL,
			action_value TEXT,
			row INTEGER,
			col INTEGER,
			FOREIGN KEY (keyboard_id) REFERENCES keyboards(id)
		);
	`
	TableReplyKeyboardButtons = `
		CREATE TABLE IF NOT EXISTS reply_keyboard_buttons (
			id INTEGER PRIMARY KEY,
			keyboard_id INTEGER NOT NULL,
			text TEXT NOT NULL,
			row INTEGER,
			col INTEGER,
			FOREIGN KEY (keyboard_id) REFERENCES reply_keyboards(id)
		);
	`
	TableMessageFiles = `
		CREATE TABLE IF NOT EXISTS message_files (
			message_id INTEGER,
			file_id INTEGER,
			PRIMARY KEY (message_id, file_id),
			FOREIGN KEY (message_id) REFERENCES messages(id),
			FOREIGN KEY (file_id) REFERENCES files(id)
		);
	`
	TableUserCommands = `
		CREATE TABLE IF NOT EXISTS user_commands (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			command_id INTEGER NOT NULL,
			executed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (command_id) REFERENCES commands(id)
		);
	`
	// ####### USERS #######
	TableUsers = `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			telegram_id INTEGER UNIQUE NOT NULL,
			username TEXT,
			first_name TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	TableLogs = `
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY,
			user_id INTEGER,
			command_id INTEGER,
			message_id INTEGER,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (command_id) REFERENCES commands(id),
			FOREIGN KEY (message_id) REFERENCES messages(id)
		);
	`
)
