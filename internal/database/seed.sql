-- seed.sql
-- DROP DATABASE IF EXISTS gossip;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS threads CASCADE;
DROP TABLE IF EXISTS comments CASCADE;

\c gossip_jdw6;

CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR(128) NOT NULL
);

CREATE TABLE tags (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(32) NOT NULL
);

CREATE TABLE threads (
    id VARCHAR(50) PRIMARY KEY,
    author_id VARCHAR(50),
    tag_id VARCHAR(50),
    title VARCHAR(255) NOT NULL,
    content VARCHAR(1024) NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

CREATE TABLE comments (
    id VARCHAR(50) PRIMARY KEY,
    thread_id VARCHAR(50),
    author_id VARCHAR(50),
    content VARCHAR(1024) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (thread_id) REFERENCES threads(id),
    FOREIGN KEY (author_id) REFERENCES users(id)
);

-- Insert sample users
INSERT INTO users (id, username, password) VALUES
    ('u1', 'John', encode(digest('password1', 'sha256'), 'hex')),
    ('u2', 'Jane', encode(digest('password2', 'sha256'), 'hex')),
    ('u3', 'Alice', encode(digest('password3', 'sha256'), 'hex')),
    ('u4', 'Aiken', encode(digest('password4', 'sha256'), 'hex'));

-- Insert sample tags
INSERT INTO tags (id, name) VALUES
    ('t1', 'Discussion'),
    ('t2', 'Question'),
    ('t3', 'Looking For Advice'),
    ('t4', 'Meme'),
    ('t5', 'Announcement'),
    ('t6', 'Tutorial'),
    ('t7', 'Showcase'),
    ('t8', 'Project'),
    ('t9', 'Feedback'),
    ('t10', 'Event'),
    ('t11', 'Technology News');

-- Insert sample threads
INSERT INTO threads (id, author_id, tag_id, title, content) VALUES
    ('t1', 'u2', 't1', 'Best Practices in JavaScript', 'Share your tips and tricks for writing clean and efficient JavaScript code.'),
    ('t2', 'u1', 't6', 'Introduction to Machine Learning', 'Discuss the fundamentals and latest trends in machine learning.'),
    ('t3', 'u2', 't3', 'Career Change Advice', 'Seeking advice from experienced professionals about transitioning into the tech industry.'),
    ('t4', 'u1', 't4', 'Programming Memes Collection', 'Share your favorite programming-related memes and jokes.'),
    ('t5', 'u3', 't6', 'Creating a RESTful API with Flask', 'A step-by-step guide on building a RESTful API using Flask.'),
    ('t6', 'u4', 't7', 'Showcasing My Latest Project', 'Excited to share my latest coding project. Check it out and provide feedback!'),
    ('t7', 'u3', 't9', 'Forum Improvement Suggestions', 'Share your ideas on how we can enhance the forum experience.'),
    ('t8', 'u4', 't11', 'Reporting a Bug in the App', 'Encountered a bug? Report it here with details for a quick resolution.');

-- Insert sample comments
INSERT INTO comments (id, thread_id, author_id, content, timestamp) VALUES
    ('p1', 't1', 'u1', 'Best tip! Use arrow functions for concise code.', '2024-01-05 10:15:00'),
    ('p2', 't1', 'u2', 'I also recommend using destructuring in JavaScript.', '2024-01-05 10:30:00'),
    ('p3', 't2', 'u1', 'Machine learning is fascinating. Any good resources for beginners?', '2024-01-05 11:00:00'),
    ('p4', 't3', 'u2', 'Thinking about a career change into tech. Any advice or success stories?', '2024-01-05 12:45:00'),
    ('p5', 't4', 'u1', 'A classic meme: "Why do programmers prefer dark mode?"', '2024-01-05 13:15:00'),
    ('p6', 't5', 'u3', 'Thanks for the Flask API tutorial! It helped a lot.', '2024-01-05 14:30:00'),
    ('p7', 't6', 'u4', 'Check out my latest project. Open to feedback and suggestions!', '2024-01-05 15:45:00'),
    ('p8', 't7', 'u3', 'How can we make this forum even better? Share your ideas!', '2024-01-05 16:30:00'),
    ('p9', 't8', 'u4', 'Encountered a bug in the app. Any developers available to help?', '2024-01-05 17:00:00');
