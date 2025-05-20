CREATE DATABASE IF NOT EXISTS stats;

CREATE TABLE IF NOT EXISTS stats.views
(
    post_id UInt32,
    user    String,
    date    Date
)
    ENGINE = MergeTree()
        ORDER BY (post_id, date);

CREATE TABLE IF NOT EXISTS stats.likes
(
    post_id UInt32,
    user    String,
    date    Date
)
    ENGINE = MergeTree()
        ORDER BY (post_id, date);

CREATE TABLE IF NOT EXISTS stats.comments
(
    post_id UInt32,
    user    String,
    date    Date
)
    ENGINE = MergeTree()
        ORDER BY (post_id, date);
