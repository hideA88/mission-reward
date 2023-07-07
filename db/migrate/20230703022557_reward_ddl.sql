-- +goose Up
CREATE TABLE coin_reward (
                      id BIGINT AUTO_INCREMENT,
                      mission_id BIGINT NOT NULL,
                      size INT NOT NULL,
                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      PRIMARY KEY(id),
                      FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE
);

CREATE TABLE item_reward (
                             id BIGINT AUTO_INCREMENT,
                             mission_id BIGINT NOT NULL,
                             item_id BIGINT NOT NULL,
                             created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY(id),
                             FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE,
                             FOREIGN KEY fx_item_id(item_id) REFERENCES item(id) ON DELETE CASCADE
);

CREATE TABLE user_achieve_mission (
                                      id BIGINT AUTO_INCREMENT,
                                      user_id BIGINT NOT NULL ,
                                      mission_id BIGINT NOT NULL ,
                                      achieved_at DATETIME NOT NULL,
                                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY(id),
                                      FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE,
                                      FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE
);

CREATE TABLE user_coin (
                           id BIGINT AUTO_INCREMENT,
                           user_id BIGINT NOT NULL,
                           user_achieve_mission_id BIGINT NOT NULL,
                           size int NOT NULL,
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY(id),
                           FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE,
                           FOREIGN KEY fx_user_achieve_mission_id(user_achieve_mission_id) REFERENCES user_achieve_mission(id) ON DELETE CASCADE
);

CREATE TABLE user_item (
                           id BIGINT AUTO_INCREMENT,
                           user_id BIGINT NOT NULL,
                           item_id BIGINT NOT NULL,
                           user_achieve_mission_id BIGINT NOT NULL,
                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY(id),
                           FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE,
                           FOREIGN KEY fx_item_id(item_id) REFERENCES item(id) ON DELETE CASCADE,
                           FOREIGN KEY fx_user_achieve_mission_id(user_achieve_mission_id) REFERENCES user_achieve_mission(id) ON DELETE CASCADE
);

CREATE TABLE user_open_mission (
                                   id BIGINT AUTO_INCREMENT,
                                   user_id BIGINT NOT NULL,
                                   mission_id BIGINT NOT NULL,
                                   user_achieve_mission_id BIGINT NOT NULL,
                                   created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                   PRIMARY KEY(id),
                                   FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE,
                                   FOREIGN KEY fx_mission_id(mission_id) REFERENCES mission(id) ON DELETE CASCADE,
                                   FOREIGN KEY fx_user_achieve_mission_id(user_achieve_mission_id) REFERENCES user_achieve_mission(id) ON DELETE CASCADE
);

CREATE VIEW mission_reward AS select m.id as mission_id,
                                     m.name as mission_name,
                                     m.mission_type as mission_type,
                                     m.reward_type as reward_type,
                                     IFNULL(cr.size, 0) as reward_coin_size,
                                     IFNULL(ir.item_id, 0) as reward_item_id,
                                     IFNULL(m.reset_time, '') as reset_time,
                                     IFNULL(m.open_mission_id, 0) as open_mission_id,
                                     m.default_mission_status as default_mission_status
                              from mission as m
                                       left join coin_reward cr on m.id = cr.mission_id
                                       left join item_reward ir on m.id = ir.mission_id;

-- +goose Down
DROP VIEW  mission_reward;
DROP TABLE user_open_mission;
DROP TABLE user_item;
DROP TABLE user_coin;
DROP TABLE user_achieve_mission;
DROP TABLE item_reward;
DROP TABLE coin_reward;

