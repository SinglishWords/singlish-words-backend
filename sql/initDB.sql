CREATE DATABASE IF NOT EXISTS singlishwords;
USE singlishwords;

DROP TABLE IF EXISTS `answer`;
DROP TABLE IF EXISTS `respondent`;
DROP TABLE IF EXISTS `question`;
DROP TABLE IF EXISTS `email`;

-- ---------------------------------------
--              CREATE TABLES
-- ---------------------------------------

CREATE TABLE IF NOT EXISTS `respondent` (
--   name                   type            constraints
    `id`                    INT             NOT NULL    AUTO_INCREMENT,
    `age`                   INT             NOT NULL,
    `gender`                VARCHAR(10)     NOT NULL,
    `education`             VARCHAR(256)    NOT NULL,

    `country_of_birth`      VARCHAR(256) ,
    `country_of_residence`  VARCHAR(256) ,
    `ethnicity`             VARCHAR(256) ,
    `uuid`                  VARCHAR(64) ,
    `is_native`             VARCHAR(5)     ,
    `language_spoken`       TEXT,
    `start_time`            DATETIME        ,
    `end_time`              DATETIME        ,

    PRIMARY KEY (`id`)
) CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `email` (
--   name                   type            constraints
   `email`                 VARCHAR(320),

   `want_lucky_draw`       VARCHAR(5)     DEFAULT 'no',
   `want_update`           VARCHAR(5)     DEFAULT 'no',
   `time_on_pages`         TEXT

) CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `question` (
--   name         type            constraints
  `id`          INT             NOT NULL      AUTO_INCREMENT,
  `word`        VARCHAR(128)    NOT NULL,
  `enable`      INT             DEFAULT 1,
  `count`       INT             DEFAULT 1,

  UNIQUE (`word`),
  PRIMARY KEY (`id`)
) CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `answer` (
--   name           type             constraints
    `id`            INT             NOT NULL    AUTO_INCREMENT,

    `association1`  TEXT            NOT NULL,
    `association2`  TEXT            NOT NULL,
    `association3`  TEXT            NOT NULL,
    `time_spend`    INT             NOT NULL,
    `question_id`   INT             NOT NULL,
    `respondent_id` INT             NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`question_id`)  REFERENCES question(`id`),
    FOREIGN KEY (`respondent_id`) REFERENCES respondent(`id`)
) CHARSET=utf8;