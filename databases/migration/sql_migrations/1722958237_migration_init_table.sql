-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NULL,
    created_by VARCHAR(100) NULL,
    modified_at TIMESTAMP NULL,
    modified_by VARCHAR(100) NULL
);

CREATE TABLE platform (
  id SERIAL PRIMARY KEY,
  name VARCHAR(256) NOT NULL,
  created_at TIMESTAMP NULL,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP NULL,
  modified_by VARCHAR(100) NULL
);

CREATE TABLE role (
  id SERIAL PRIMARY KEY,
  name VARCHAR(256) NOT NULL
);

-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(256) NOT NULL,
  email VARCHAR(255) NULL,
  password VARCHAR(32) NOT NULL,
  role_id INT NOT NULL,
  created_at TIMESTAMP(0) NULL,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP(0) NULL,
  modified_by VARCHAR(100) NULL
);

-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE game (
  id SERIAL PRIMARY KEY,
  title VARCHAR(256) NULL,
  harga INT NULL,
  category_id INT NOT NULL,
  platform_id INT NOT NULL,
  created_at TIMESTAMP(0) NULL,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP(0) NULL,
  modified_by VARCHAR(100) NULL
);

-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE transaksi (
  id SERIAL PRIMARY KEY,
  tgl_transaksi TIMESTAMP(0) NOT NULL,
  user_id INT NOT NULL,
  total_harga INT NULL,
  created_at TIMESTAMP(0) NULL DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP(0) NULL,
  modified_by VARCHAR(100) NULL
);

-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE transaksi_detail (
  id SERIAL PRIMARY KEY,
  transaksi_id INT NOT NULL,
  game_id INT NOT NULL,
  created_at TIMESTAMP(0) NULL DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP(0) NULL,
  modified_by VARCHAR(100) NULL
    );

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS platform;
DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS game;
DROP TABLE IF EXISTS transaksi;
DROP TABLE IF EXISTS transaksi_detail;

-- +migrate StatementEnd