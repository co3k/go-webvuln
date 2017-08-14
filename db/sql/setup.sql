DROP TABLE IF EXISTS webvuln_user;
CREATE TABLE IF NOT EXISTS webvuln_user (
    id INTEGER PRIMARY KEY,
    username TEXT,
    password TEXT
);

INSERT INTO webvuln_user VALUES (1, 'guest', 'd83fd40de3ca1eb68a8a520e61e78a61e2c506a08452ebe86dc4d1bfb5fc6bac');
INSERT INTO webvuln_user VALUES (2, 'co3k', '579b3db696c86164a50b0432fe6af0d5ca6b6d11b2a8d11ebfdfdcfdf067a4e4');

DROP TABLE IF EXISTS webvuln_activity;
CREATE TABLE IF NOT EXISTS webvuln_activity (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,
    body TEXT,
    created_at TEXT
);

INSERT INTO webvuln_activity VALUES (1, 1, 'こんにちはー', '1988-04-23 02:00:00');
INSERT INTO webvuln_activity VALUES (2, 1, 'だれかいませんかー', '1988-04-23 02:01:00');
INSERT INTO webvuln_activity VALUES (3, 2, 'いますよー', '1988-04-23 02:02:00');
INSERT INTO webvuln_activity VALUES (4, 1, 'ぼやきませんかー', '1988-04-23 02:03:00');
INSERT INTO webvuln_activity VALUES (5, 2, 'そうですねー', '1988-04-23 02:04:00');
INSERT INTO webvuln_activity VALUES (6, 1, 'はやくぼやいてくださいよー', '1988-04-23 02:05:00');
INSERT INTO webvuln_activity VALUES (7, 2, 'そちらが先にぼやいてくださいよー', '1988-04-23 02:06:00');
INSERT INTO webvuln_activity VALUES (8, 1, 'いいからはやくぼやけよー', '1988-04-23 02:07:00');
INSERT INTO webvuln_activity VALUES (9, 2, 'いやいやあなたが先にぼやいてよー', '1988-04-23 02:08:00');
INSERT INTO webvuln_activity VALUES (10, 1, 'なんでぼやいてくれないんですかー', '1988-04-23 02:09:00');
INSERT INTO webvuln_activity VALUES (11, 2, 'それはこちらのセリフですよー', '1988-04-23 02:10:00');

DROP TABLE IF EXISTS webvuln_session;
CREATE TABLE IF NOT EXISTS webvuln_session (
    session_id TEXT UNIQUE,
    user_id INTEGER
);
