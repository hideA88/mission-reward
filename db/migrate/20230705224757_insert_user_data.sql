-- +goose Up
INSERT INTO user (id, name) VALUES
                                (1, 'taro'),
                                (2, 'jiro');

INSERT INTO user_monster (id, user_id, monster_id) VALUES
                                (1, 1, 1),
                                (2, 1, 4),
                                (3, 2, 2);


INSERT INTO user_achieve_mission (id, user_id, mission_id, achieved_at) VALUES
                                (1, 1, 6, '2023-06-30 10:00:00'),
                                (2, 1, 8, '2023-07-01 10:00:00'), -- 初期所有コイン
                                (3, 2, 6, '2023-07-01 12:00:00'),
                                (4, 2, 8, '2023-07-01 12:00:00'), -- 初期所有コイン
                                (5, 2, 4, '2023-07-02 12:00:00'); -- open mission

INSERT INTO user_open_mission (id, user_id, mission_id, user_achieve_mission_id) VALUES
                                (1, 2, 7, 5);


INSERT INTO user_coin (id, user_id, user_achieve_mission_id, size) VALUES
    (1, 1, 2, 1800), -- 初期所有コイン
    (2, 2, 4, 1900); -- 初期所有コイン

-- +goose Down
DELETE FROM user_achieve_mission;
DELETE FROM user_coin;
DELETE FROM user;