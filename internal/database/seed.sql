-- seed.sql
DROP DATABASE IF EXISTS gossip;

CREATE DATABASE IF NOT EXISTS gossip;

USE gossip;


-- schema
CREATE TABLE IF NOT EXISTS users (
                                     id INT PRIMARY KEY AUTO_INCREMENT,
                                     name VARCHAR(128) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS tags (
                                    id INT PRIMARY KEY AUTO_INCREMENT,
                                    name VARCHAR(16) NOT NULL
);

CREATE TABLE IF NOT EXISTS threads (
                                       id INT PRIMARY KEY AUTO_INCREMENT,
                                       author_id INT,
                                       tag_id INT,
                                       title VARCHAR(255) NOT NULL,
                                       content VARCHAR(1024) NOT NULL,
                                       FOREIGN KEY (author_id) REFERENCES users(id),
                                       FOREIGN KEY (tag_id) REFERENCES tags(id)
);

CREATE TABLE IF NOT EXISTS posts (
                                     id INT PRIMARY KEY AUTO_INCREMENT,
                                     thread_id INT,
                                     author_id INT,
                                     content VARCHAR(1024) NOT NULL,
                                     timestamp DATETIME NOT NULL,
                                     FOREIGN KEY (thread_id) REFERENCES threads(id),
                                     FOREIGN KEY (author_id) REFERENCES users(id)
);

-- Insert sample users
INSERT INTO users (id, name) VALUES
                                     (1, 'CVWO'),
                                     (2, 'Kenneth');

-- Insert sample tags
INSERT INTO tags (id, name) VALUES
                                (1, 'Opinion'),
                                (2, 'Satirical');

-- Insert sample threads
INSERT INTO threads (id, author_id, tag_id, title, content) VALUES
                                                                (1, 2, 2, 'Sample Thread 1', 'Sample thread content 1'),
                                                                (2, 1, 2, 'Sample Thread 2', 'Sample thread content 2');

-- Insert sample posts
INSERT INTO posts (id, thread_id, author_id, content, timestamp) VALUES
                                                                     (1, 1, 1, 'Sample post content 1', NOW()),
                                                                     (2, 1, 2, 'Sample post content 2', NOW());

