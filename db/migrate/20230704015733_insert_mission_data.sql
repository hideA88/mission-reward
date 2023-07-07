-- +goose Up
INSERT INTO mission(id, name, mission_type, reward_type, reset_time, open_mission_id, default_mission_status) VALUES
        (1, 'kill_monster_monsterA', 'killMonster', 'coin', null, null, 'open'),
        (2, 'get_total_coin_2000', 'totalCoin', 'item', null, null, 'open'),
        (3, 'level_5_monsterA', 'levelUp', 'coin', null, null, 'open'),
        (5, 'kill_monster_10', 'killMonster', 'coin', '0 10 * * 1', null, 'open'), -- 毎週月曜日10時リセット
        (6, 'daily_login_reward', 'login', 'coin', '0 4 * * *', null, 'open'), -- 毎日4時リセット
        (7, 'get_itemA', 'getItem', 'coin', null, null, 'blocked'),
        (4, 'level_5_monster_2', 'levelUp', 'coin', null, 7, 'open'), -- 外部キー制約のため順序調整
        (8, 'init', 'login', 'coin', null, null, 'open'); -- 初期コイン投入用のミッション

INSERT INTO coin_reward(id, mission_id, size) VALUES
                                                    (1, 1, 100),
                                                    (2, 3, 100),
                                                    (3, 4, 100),
                                                    (4, 5, 100),
                                                    (5, 6, 100),
                                                    (6, 7, 100),
                                                    (7, 8, 0);

INSERT INTO item_reward(id, mission_id, item_id) VALUES (1, 2, 1); -- アイテムAを付与


INSERT INTO kill_monster_mission(id, mission_id, target_type, target_monster_id, size) VALUES
            (1, 1, 'specific', 4, 1),
            (2, 5, 'all', null, 10);

INSERT INTO total_coin_mission(id, mission_id, size) VALUES (1, 2, 2000);

INSERT INTO level_up_mission(id, mission_id, level, target_type, target_monster_id, monster_size) VALUES
            (1, 3, 5, 'specific', 4, 1),
            (2, 4, 5, 'all', null, 2);

INSERT INTO get_item_mission(id, mission_id, item_id) VALUES (1, 7, 1);


-- +goose Down
DELETE FROM get_item_mission;
DELETE FROM total_coin_mission;
DELETE FROM kill_monster_mission;
DELETE FROM item_reward;
DELETE FROM coin_reward;

DELETE FROM mission;
