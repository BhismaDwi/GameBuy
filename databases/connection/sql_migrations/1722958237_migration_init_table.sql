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

CREATE TABLE user (
  id SERIAL PRIMARY KEY,
  username VARCHAR(256) NOT NULL,
  email VARCHAR(255) NULL,
  password VARCHAR(32) NOT NULL,
  role_id INT NOT NULL,
  created_at TIMESTAMP NULL,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP NULL,
  modified_by VARCHAR(100) NULL,
  INDEX `fk_user_role1_idx` (`role_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_role1`
    FOREIGN KEY (`role_id`)
    REFERENCES `mydb`.`role` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

CREATE TABLE game (
  id SERIAL PRIMARY KEY,
  tittle VARCHAR(256) NULL,
  harga INT NULL,
  category_id INT NOT NULL,
  platform_id INT NOT NULL,
  created_at TIMESTAMP NULL,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP NULL,
  modified_by VARCHAR(100) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_game_category_idx` (`category_id` ASC) VISIBLE,
  INDEX `fk_game_platform1_idx` (`platform_id` ASC) VISIBLE,
  CONSTRAINT `fk_game_category`
    FOREIGN KEY (`category_id`)
    REFERENCES `mydb`.`category` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_game_platform1`
    FOREIGN KEY (`platform_id`)
    REFERENCES `mydb`.`platform` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

CREATE TABLE transaksi (
  id SERIAL PRIMARY KEY,
  tgl_transaksi TIMESTAMP NOT NULL,
  user_id INT NOT NULL,
  total_harga INT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP NULL,
  modified_by VARCHAR(100) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_transaksi_user1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_transaksi_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `mydb`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);

CREATE TABLE transaksi_detail (
  id SERIAL PRIMARY KEY,
  transaksi_id INT NOT NULL,
  game_id INT NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(100) NULL,
  modified_at TIMESTAMP NULL,
  modified_by VARCHAR(100) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_transaksi_game1_idx` (`game_id` ASC) VISIBLE,
  INDEX `fk_transaksi_copy1_transaksi1_idx` (`transaksi_id` ASC) VISIBLE,
  CONSTRAINT `fk_transaksi_game10`
    FOREIGN KEY (`game_id`)
    REFERENCES `mydb`.`game` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_transaksi_copy1_transaksi1`
    FOREIGN KEY (`transaksi_id`)
    REFERENCES `mydb`.`transaksi` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    );

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS platform;
DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS game;
DROP TABLE IF EXISTS transaksi;
DROP TABLE IF EXISTS transaksi_detail;

-- +migrate StatementEnd