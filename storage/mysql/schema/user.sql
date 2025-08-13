CREATE TABLE users
(
    id           INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    phone_number VARCHAR(11) NOT NULL UNIQUE,
    username     VARCHAR(25) NOT NULL UNIQUE,
    password     VARCHAR(50),
    total_score  INT       DEFAULT 0,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)