CREATE DATABASE sm_database;

USE sm_database;

CREATE TABLE IF NOT EXISTS user (
	id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(15) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT UC_email UNIQUE(email),
    CONSTRAINT UC_name UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS post (
	id INT NOT NULL AUTO_INCREMENT,
    content VARCHAR(500) NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT FK_post_user FOREIGN KEY(user_id)
        REFERENCES user(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS likes (
	id INT NOT NULL AUTO_INCREMENT,
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT FK_likes_post FOREIGN KEY(post_id)
		REFERENCES post(id)
		ON DELETE CASCADE,
	CONSTRAINT FK_likes_user FOREIGN KEY(user_id)
		REFERENCES user(id)
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
