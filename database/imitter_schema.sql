CREATE DATABASE IF NOT EXISTS sm_database;
USE sm_database;

CREATE TABLE IF NOT EXISTS user (
	id INT NOT NULL AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL,
    name VARCHAR(15) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    follows_count INT DEFAULT 0,
    PRIMARY KEY(id),
    CONSTRAINT UC_uuid UNIQUE(uuid),
    CONSTRAINT UC_email UNIQUE(email),
    CONSTRAINT UC_name UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS post (
	id INT NOT NULL AUTO_INCREMENT,
    content VARCHAR(500) NOT NULL,
    date DATETIME NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    likes_count INT DEFAULT 0,
    PRIMARY KEY(id),
    CONSTRAINT FK_post_user FOREIGN KEY(user_id)
        REFERENCES user(uuid)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS likes (
    post_id INT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    PRIMARY KEY(post_id, user_id),
    CONSTRAINT FK_likes_post FOREIGN KEY(post_id)
		REFERENCES post(id)
		ON DELETE CASCADE,
	CONSTRAINT FK_likes_user FOREIGN KEY(user_id)
		REFERENCES user(uuid)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS follows (
    user_id VARCHAR(36) NOT NULL,
    follower_id VARCHAR(36) NOT NULL,
    PRIMARY KEY(user_id, follower_id),
    CONSTRAINT FK_follows_user_id FOREIGN KEY(user_id)
        REFERENCES user(uuid)
        ON DELETE CASCADE,
    CONSTRAINT FK_follows_follower_id FOREIGN KEY(follower_id)
        REFERENCES user(uuid)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS post_comment (
	id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    comment VARCHAR(500),
    PRIMARY KEY(id),
    CONSTRAINT FK_post_comment_user FOREIGN KEY(user_id)
		REFERENCES user(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS token_blacklist (
    token_uuid VARCHAR(36) NOT NULL,
    PRIMARY KEY(token_uuid),
    CONSTRAINT UC_token_blacklist_token_uuid UNIQUE(token_uuid)
);
