CREATE TABLE habit_categories
(
    id            SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP    NOT NULL,
    updated_at    TIMESTAMP    NOT NULL
);

CREATE TABLE habits
(
    id          SERIAL PRIMARY KEY,
    category_id INT          NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP    NOT NULL,

    FOREIGN KEY (category_id) REFERENCES habit_categories (id)
);

CREATE TABLE habit_records
(
    id          SERIAL PRIMARY KEY,
    habit_id    INT         NOT NULL,
    record_date TIMESTAMP   NOT NULL,
    result      VARCHAR(50) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,

    FOREIGN KEY (habit_id) REFERENCES habits (id)
);

CREATE TABLE tags
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP    NOT NULL
);

CREATE TABLE habit_tags
(
    habit_id INT NOT NULL,
    tag_id   INT NOT NULL,

    PRIMARY KEY (habit_id, tag_id),
    FOREIGN KEY (habit_id) REFERENCES habits (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id)
);

CREATE TABLE goals
(
    id          SERIAL PRIMARY KEY,
    description TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);

CREATE TABLE habit_goals
(
    habit_id INT NOT NULL,
    goal_id  INT NOT NULL,

    PRIMARY KEY (habit_id, goal_id),
    FOREIGN KEY (habit_id) REFERENCES habits (id),
    FOREIGN KEY (goal_id) REFERENCES goals (id)
);

CREATE TABLE events
(
    id         SERIAL PRIMARY KEY,
    habit_id   INT       NOT NULL,
    subject    TEXT      NOT NULL,
    start_at   TIMESTAMP NOT NULL,
    end_at     TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    FOREIGN KEY (habit_id) REFERENCES habits (id)
);