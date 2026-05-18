-- ####### MAIN ENTITIES #######

CREATE TABLE admins (
    id INTEGER PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bots (
    id INTEGER PRIMARY KEY,
    token TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ####### BOT-RELATED ENTITIES #######

CREATE TABLE commands (
    id INTEGER PRIMARY KEY,
    bot_id INTEGER NOT NULL,
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE messages (
    id INTEGER PRIMARY KEY,
    command_id INTEGER NOT NULL,
    text TEXT,
    media_type TEXT,
    media_storage_key TEXT,
    FOREIGN KEY (command_id) REFERENCES commands(id)
);

CREATE TABLE keyboards (
    id INTEGER PRIMARY KEY,
    message_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (message_id) REFERENCES messages(id)
);

CREATE TABLE reply_keyboards (
    id INTEGER PRIMARY KEY,
    message_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (message_id) REFERENCES messages(id)
);

-- ####### SETTINGS #######

CREATE TABLE settings_menu (
    id INTEGER PRIMARY KEY,
    bot_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (bot_id) REFERENCES bots(id)
);

CREATE TABLE settings (
    id INTEGER PRIMARY KEY,
    menu_id INTEGER NOT NULL,
    key TEXT UNIQUE NOT NULL,
    value TEXT NOT NULL,
    type TEXT DEFAULT 'string',
    FOREIGN KEY (menu_id) REFERENCES settings_menu(id)
);

-- ####### FILE MANAGEMENT #######

CREATE TABLE categories (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE files (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    storage_key TEXT NOT NULL,
    category_id INTEGER,
    size INTEGER,
    telegram_file_id TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- ####### RELATION TABLES #######

CREATE TABLE keyboard_buttons (
    id INTEGER PRIMARY KEY,
    keyboard_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    action_type TEXT NOT NULL,
    action_value TEXT,
    row INTEGER,
    col INTEGER,
    FOREIGN KEY (keyboard_id) REFERENCES keyboards(id)
);

CREATE TABLE reply_keyboard_buttons (
    id INTEGER PRIMARY KEY,
    keyboard_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    row INTEGER,
    col INTEGER,
    FOREIGN KEY (keyboard_id) REFERENCES reply_keyboards(id)
);

CREATE TABLE message_files (
    message_id INTEGER,
    file_id INTEGER,
    PRIMARY KEY (message_id, file_id),
    FOREIGN KEY (message_id) REFERENCES messages(id),
    FOREIGN KEY (file_id) REFERENCES files(id)
);

CREATE TABLE user_commands (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    command_id INTEGER NOT NULL,
    executed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (command_id) REFERENCES commands(id)
);

-- ####### USERS #######

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    telegram_id INTEGER UNIQUE NOT NULL,
    username TEXT,
    first_name TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE logs (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,
    command_id INTEGER,
    message_id INTEGER,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (command_id) REFERENCES commands(id),
    FOREIGN KEY (message_id) REFERENCES messages(id)
);