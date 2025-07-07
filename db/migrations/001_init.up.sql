-- Table: users
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(50),
    full_name  VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: rooms
CREATE TABLE rooms
(
    id         BIGSERIAL PRIMARY KEY,
    room_name  VARCHAR(200),
    user_ids   BIGINT[],
    is_group   BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: room_history
CREATE TABLE room_history
(
    id         BIGSERIAL PRIMARY KEY,
    room_id    BIGINT,
    user_id    BIGINT,
    join_at    TIMESTAMP,
    leave_at   TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: messages
CREATE TABLE messages
(
    id          BIGSERIAL PRIMARY KEY,
    sender_id   BIGINT  NOT NULL,
    receiver_id BIGINT  NOT NULL,
    room_id     BIGINT  NOT NULL,
    image_url   VARCHAR(100),
    tree_path   VARCHAR(50),      -- path level vd: 1,3,5
    level       INTEGER NOT NULL, -- level of messages
    parent_id   BIGINT,           -- id parent of messages
    content     TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: message_status
CREATE TABLE message_status
(
    id         BIGSERIAL PRIMARY KEY,
    message_id BIGINT,
    user_id    BIGINT,
    is_read    BOOLEAN   DEFAULT FALSE,
    read_at    TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_ids_gin ON rooms USING GIN (user_ids);