DROP SCHEMA IF EXISTS nomikuimatch;

CREATE SCHEMA nomikuimatch;

USE nomikuimatch;

DROP TABLE IF EXISTS Area,
Genre,
User,
Restaurant,
Tags,
Favorite,
Nomikui,
Entry;

CREATE TABLE
  Area (
    id CHAR(36) PRIMARY KEY,
    areaname VARCHAR(255) NOT NULL
  );

CREATE TABLE
  Genre (
    id CHAR(36) PRIMARY KEY,
    genrename VARCHAR(255) NOT NULL
  );

CREATE TABLE
  User (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    traQid VARCHAR(255)
  );

CREATE TABLE
  Restaurant (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    areaid CHAR(36) NOT NULL,
    FOREIGN KEY (areaid) REFERENCES Area (id),
    genreid CHAR(36) NOT NULL,
    FOREIGN KEY (genreid) REFERENCES Genre (id),
    registered_at DATETIME NOT NULL,
    picture TEXT,
    comment TEXT
  );

CREATE TABLE
  Tags (
    id CHAR(36) PRIMARY KEY,
    restaurantid CHAR(36) NOT NULL,
    FOREIGN KEY (restaurantid) REFERENCES Restaurant (id),
    content VARCHAR(255) NOT NULL
  );

CREATE TABLE
  Favorite (
    userid CHAR(36),
    FOREIGN KEY (userid) REFERENCES User (id),
    restaurantid CHAR(36),
    FOREIGN KEY (restaurantid) REFERENCES Restaurant (id),
    PRIMARY KEY (userid, restaurantid)
  );

CREATE TABLE
  Nomikui (
    id CHAR(36) PRIMARY KEY,
    restaurantid CHAR(36) NOT NULL,
    FOREIGN KEY (restaurantid) REFERENCES Restaurant (id),
    organizerid CHAR(36) NOT NULL,
    FOREIGN KEY (organizerid) REFERENCES User (id),
    conducted_at DATETIME NOT NULL,
    isopen BOOLEAN NOT NULL,
    picture TEXT,
    comment TEXT
  );

CREATE TABLE
  Entry (
    userid CHAR(36),
    FOREIGN KEY (userid) REFERENCES User (id),
    nomikuiid CHAR(36),
    FOREIGN KEY (nomikuiid) REFERENCES Nomikui (id),
    present BOOLEAN NOT NULL,
    entried_at DATETIME NOT NULL,
    PRIMARY KEY (userid, nomikuiid)
  );
