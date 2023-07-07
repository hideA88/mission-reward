-- +goose Up
INSERT INTO monster(id, name) VALUES
                                  (1, 'pochi'),
                                  (2, 'shiro'),
                                  (3, 'tama'),
                                  (4, 'monsterA');

INSERT INTO item(id, name) VALUES
                               (1, 'itemA'),
                               (2, 'itemB'),
                               (3, 'itemC');



-- +goose Down
DELETE FROM item;
DELETE FROM monster;
