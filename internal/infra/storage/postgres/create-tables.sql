DROP TABLE IF EXISTS ads;
CREATE TABLE ads (
    id uuid NOT NULL,
    title varchar(128) NOT NULL,
    description varchar(128) NOT NULL,
    Price float8 NOT NULL,
    CreatedDate date NOT NULL,
    PRIMARY KEY (id)
);
