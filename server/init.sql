DROP SCHEMA IF EXISTS nomikuimatch;

CREATE SCHEMA nomikuimatch;

USE nomikuimatch;

DROP TABLE IF EXISTS Area,
Genre,
User,
Restaurant;

CREATE TABLE Area (id CHAR(36) PRIMARY KEY, areaname VARCHAR(255));

CREATE TABLE Genre (id CHAR(36) PRIMARY KEY, genrename VARCHAR(255));

CREATE TABLE User (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255),
  traQid VARCHAR(255)
);

CREATE TABLE Restaurant (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255),
  areaid CHAR(36),
  FOREIGN KEY (areaid) REFERENCES Area (id)
);
