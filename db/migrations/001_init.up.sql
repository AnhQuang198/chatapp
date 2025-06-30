-- Table: users
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(50),
    full_name  VARCHAR(200),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Table: rooms
CREATE TABLE rooms
(
    id         BIGSERIAL PRIMARY KEY,
    room_name  VARCHAR(200),
    user_ids   VARCHAR(200),
    is_group   BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Table: room_history
CREATE TABLE room_history
(
    id         BIGSERIAL PRIMARY KEY,
    room_id    INTEGER,
    user_id    INTEGER,
    join_at    TIMESTAMP,
    leave_at   TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Table: messages
CREATE TABLE messages
(
    id         BIGSERIAL PRIMARY KEY,
    sender_id  INTEGER NOT NULL,
    room_id    INTEGER NOT NULL,
    image_url  VARCHAR(100),
    tree_path  VARCHAR(50), -- path level vd: 1,3,5
    level      INTEGER,     -- level of messages
    parent_id  INTEGER,     -- id parent of messages
    content    TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Table: message_status
CREATE TABLE message_status
(
    id         BIGSERIAL PRIMARY KEY,
    message_id INTEGER,
    user_id    INTEGER,
    is_read    BOOLEAN DEFAULT FALSE,
    read_at    TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);