-- users
CREATE TABLE IF NOT EXISTS users (
                                     id            BIGSERIAL PRIMARY KEY,
                                     name          VARCHAR(255) NOT NULL,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
    );

-- todo_lists
CREATE TABLE IF NOT EXISTS todo_lists (
                                          id          BIGSERIAL PRIMARY KEY,
                                          title       VARCHAR(255) NOT NULL,
    description TEXT
    );

-- связь пользователь-список
CREATE TABLE IF NOT EXISTS users_lists (
                                           id       BIGSERIAL PRIMARY KEY,
                                           user_id  BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list_id  BIGINT NOT NULL REFERENCES todo_lists(id) ON DELETE CASCADE,
    UNIQUE (user_id, list_id)
    );

-- элементы
CREATE TABLE IF NOT EXISTS todo_items (
                                          id          BIGSERIAL PRIMARY KEY,
                                          title       VARCHAR(255) NOT NULL,
    description TEXT,
    done        BOOLEAN NOT NULL DEFAULT FALSE
    );

-- связь список-элемент
CREATE TABLE IF NOT EXISTS list_items (
                                          id       BIGSERIAL PRIMARY KEY,
                                          item_id  BIGINT NOT NULL REFERENCES todo_items(id) ON DELETE CASCADE,
    list_id  BIGINT NOT NULL REFERENCES todo_lists(id) ON DELETE CASCADE,
    UNIQUE (list_id, item_id)
    );

-- Индексы на FK (ускоряют JOIN/DELETE)
CREATE INDEX IF NOT EXISTS idx_users_lists_user_id ON users_lists(user_id);
CREATE INDEX IF NOT EXISTS idx_users_lists_list_id ON users_lists(list_id);
CREATE INDEX IF NOT EXISTS idx_list_items_list_id ON list_items(list_id);
CREATE INDEX IF NOT EXISTS idx_list_items_item_id ON list_items(item_id);