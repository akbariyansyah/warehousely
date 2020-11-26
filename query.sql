CREATE TABLE m_user(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(256) NOT NULL,
    email VARCHAR(50) NOT NULL,
    status BOOLEAN DEFAULT TRUE,
    updated DATE NOT NULL DEFAULT CURRENT_DATE,
    created DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE m_category(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created DATE NOT NULL DEFAULT CURRENT_DATE,
    updated DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE m_goods(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    stock INT NOT NULL,
    category_id BIGINT REFERENCES m_category (id),
    user_id BIGINT REFERENCES m_user (id),
    status BOOLEAN DEFAULT TRUE,
    image_path VARCHAR(100) NOT NULL,
    created DATE NOT NULL DEFAULT CURRENT_DATE,
    updated DATE NOT NULL DEFAULT CURRENT_DATE
);


CREATE TABLE m_position(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    is_delete INT NOT NULL DEFAULT 0
)
CREATE TABLE m_employee(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    position_id BIGINT NOT NULL REFERENCES m_position (id),
    id_number INT NOT NULL,
    gender INT NOT NULL,
    is_delete INT NOT NULL DEFAULT 0
)

INSERT INTO m_position (code,name,is_delete) VALUES ("SA","System Analysist",0);
INSERT INTO m_position (code,name,is_delete) VALUES ('BPS','BPS',0);
INSERT INTO m_position (code,name,is_delete) VALUES ('PRG','Programmer',0);
INSERT INTO m_position (code,name,is_delete) VALUES ('TEST','Tester',0);
INSERT INTO m_position (code,name,is_delete) VALUES ('HELP','Helpdesk',0);

