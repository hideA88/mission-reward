-- +goose Up
CREATE TABLE mission (
                         id BIGINT AUTO_INCREMENT,
                         name VARCHAR(255) NOT NULL,
                         mission_type ENUM(
                             'killMonster',
                             'totalCoin',
                             'levelUp',
                             'login',
                             'getItem'
                             ) NOT NULL,
                         reward_type ENUM(
                             'coin',
                             'item'
                             ) NOT NULL,
                         reset_time VARCHAR(255),
                         open_mission_id BIGINT,
                         default_mission_status ENUM(
                             'open',
                             'close',
                             'blocked'
                         ),
                         created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY(id),
                         FOREIGN KEY fx_open_mission_id(open_mission_id) REFERENCES mission(id) ON DELETE CASCADE
);

create index mission_type_index on mission(mission_type);

CREATE TABLE kill_monster_mission (
                               id BIGINT AUTO_INCREMENT,
                               mission_id BIGINT NOT NULL,
                               target_type ENUM(
                                   'specific',
                                   'all'
                                   ) NOT NULL,
                               target_monster_id BIGINT,
                               size INT NOT NULL,
                               created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY(id),
                               FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE,
                               FOREIGN KEY fx_target_monster_id(target_monster_id) REFERENCES monster(id) ON DELETE CASCADE
);

CREATE TABLE total_coin_mission (
                                      id BIGINT AUTO_INCREMENT,
                                      mission_id BIGINT NOT NULL,
                                      size INT NOT NULL,
                                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY(id),
                                      FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE
);

CREATE TABLE level_up_mission (
                                    id BIGINT AUTO_INCREMENT,
                                    mission_id BIGINT NOT NULL,
                                    level INT NOT NULL,
                                    target_type ENUM(
                                        'specific',
                                        'all'
                                        ) NOT NULL,
                                    target_monster_id BIGINT,
                                    monster_size INT NOT NULL,
                                    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    PRIMARY KEY(id),
                                    FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE,
                                    FOREIGN KEY fx_target_monster_id(target_monster_id) REFERENCES monster(id) ON DELETE CASCADE
);

CREATE TABLE get_item_mission (
                                  id BIGINT AUTO_INCREMENT,
                                  mission_id BIGINT NOT NULL,
                                  item_id BIGINT NOT NULL,
                                  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY(id),
                                  FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE,
                                  FOREIGN KEY fx_item_id(item_id) REFERENCES item(id) ON DELETE CASCADE
);



-- +goose Down
DROP TABLE get_item_mission;
DROP TABLE level_up_mission;
DROP TABLE total_coin_mission;
DROP TABLE kill_monster_mission;

DROP TABLE mission;
