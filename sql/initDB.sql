CREATE DATABASE IF NOT EXISTS singlishwords;
USE singlishwords;

DROP TABLE IF EXISTS `answer`;
DROP TABLE IF EXISTS `respondent`;
DROP TABLE IF EXISTS `question`;
DROP TABLE IF EXISTS `email`;
DROP TABLE IF EXISTS `association`;
DROP TABLE IF EXISTS `community_map`;

-- ---------------------------------------
--              CREATE TABLES
-- ---------------------------------------

CREATE TABLE IF NOT EXISTS `respondent` (
--   name                   type            constraints
    `id`                        INT             NOT NULL    AUTO_INCREMENT,
    `age`                       VARCHAR(256)    NOT NULL,
    `gender`                    VARCHAR(10)     NOT NULL,
    `education`                 VARCHAR(256)    NOT NULL,
    `duration_of_sgp_residence` VARCHAR(256)    NOT NULL,
    `country_of_birth`          VARCHAR(256),
    `country_of_residence`      VARCHAR(256),
    `ethnicity`                 VARCHAR(256),
    `uuid`                      VARCHAR(64),
    `is_native`                 VARCHAR(5),
    `language_spoken`           TEXT,
    `start_time`                DATETIME,
    `end_time`                  DATETIME,

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
    `id`                    INT             NOT NULL    AUTO_INCREMENT,

    `association1`          TEXT            NOT NULL,
    `association2`          TEXT            NOT NULL,
    `association3`          TEXT            NOT NULL,
    `is_recognised_word`    INT             NOT NULL,
    `time_spend`            INT             NOT NULL,
    `question_id`           INT             NOT NULL,
    `respondent_id`         INT             NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`question_id`)  REFERENCES question(`id`),
    FOREIGN KEY (`respondent_id`) REFERENCES respondent(`id`)
) CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `association` (
--   name                   type            constraints
    `id`                    INT             NOT NULL    AUTO_INCREMENT,
    `source`                VARCHAR(256)    NOT NULL,
    `target`                VARCHAR(256)    NOT NULL,
    `count`                 INT             NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE INDEX forward_association (`source`, `target`),
    UNIQUE INDEX backward_association (`target`, `source`)
) CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `community_map` (
--   name                   type            constraints
    `id`                    INT             NOT NULL    AUTO_INCREMENT,
    `word`                  VARCHAR(256)    NOT NULL,
    `community`             INT             NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`word`)
) CHARSET=utf8;
