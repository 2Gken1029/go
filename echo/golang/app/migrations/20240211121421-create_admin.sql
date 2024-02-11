
-- +migrate Up
CREATE TABLE IF NOT EXISTS admin (id int);
-- +migrate Down
DROP TABLE IF EXISTS admin;