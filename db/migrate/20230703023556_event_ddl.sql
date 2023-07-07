-- +goose Up
CREATE TABLE user_login_event (
                            id BIGINT AUTO_INCREMENT,
                            user_id BIGINT NOT NULL,
                            event_at DATETIME NOT NULL,
                            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            PRIMARY KEY(id),
                            FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE TABLE kill_monster_event (
                            id BIGINT AUTO_INCREMENT,
                            user_id BIGINT NOT NULL,
                            kill_monster_id BIGINT NOT NULL,
                            event_at DATETIME NOT NULL,
                            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            PRIMARY KEY(id),
                            FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE,
                            FOREIGN KEY fx_kill_monster_id(kill_monster_id) REFERENCES monster(id) ON DELETE CASCADE
);

CREATE TABLE level_up_event (
                                    id BIGINT AUTO_INCREMENT,
                                    user_id BIGINT NOT NULL,
                                    user_monster_id BIGINT NOT NULL,
                                    level_up_size INT NOT NULL,
                                    event_at DATETIME NOT NULL,
                                    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    PRIMARY KEY(id),
                                    FOREIGN KEY fx_user_id(user_id) REFERENCES user(id) ON DELETE CASCADE,
                                    FOREIGN KEY fx_user_monster_id(user_monster_id) REFERENCES user_monster(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_login_event;
DROP TABLE kill_monster_event;
DROP TABLE level_up_event;
