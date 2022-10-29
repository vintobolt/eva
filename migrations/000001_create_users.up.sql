CREATE TABLE users 
(
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE DEFAULT NULL,
    username TEXT UNIQUE NOT NULL,
    fullname TEXT NOT NULL,
    rolename TEXT DEFAULT 'nobody',
    passwd TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    token TEXT
);

INSERT INTO users (email, username, fullname, rolename, passwd, active, created_at, updated_at)
VALUES 
('vintobolt@protonmail.com', 'vintobolt', 'Oleksandr Zatserklianyi', 'administrator', '13wia0aw', TRUE, NOW(), NOW()),
('finegripper@gmail.com', 'finegripper', 'Vitaliy Kovalenko', 'doctor', '29wia9aw', TRUE, NOW(), NOW()),
('aniyoole@protonmail.com', 'aniyoole', 'Aniyo Oluwadamilare', 'assistant', '39wia8aw', TRUE, NOW(), NOW());