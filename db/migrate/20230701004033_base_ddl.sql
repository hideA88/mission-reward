-- +goose Up
CREATE TABLE monster (
                        id BIGINT AUTO_INCREMENT,
                        name VARCHAR(255) NOT NULL,
                        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY(id)
);


CREATE TABLE item (
                     id BIGINT AUTO_INCREMENT,
                     name VARCHAR(255),
                     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                     updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                     PRIMARY KEY(id)
);

CREATE TABLE user (
                      id BIGINT AUTO_INCREMENT,
                      name VARCHAR(255),
                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                      PRIMARY KEY(id)
);



CREATE TABLE user_monster (
                             id BIGINT AUTO_INCREMENT,
                             user_id BIGINT NOT NULL,
                             monster_id BIGINT NOT NULL,
                             created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY(id),
                             FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE,
                             FOREIGN KEY fx_monster_id(monster_id) REFERENCES monster(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_monster;
DROP TABLE user;
DROP TABLE item;
DROP TABLE monster;
