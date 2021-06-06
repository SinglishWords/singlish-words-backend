USE singlishwords;

-----------------------------------------
--              CREATE TABLES
-----------------------------------------

CREATE TABLE IF NOT EXISTS `respondent` (
--   name               type            constraints
    `id`                INT             NOT NULL    AUTO_INCREMENT,
    `age`               INT             NOT NULL,
    `gender`            ENUM ("Male", "Female", "X") 
                                        NOT NULL,
    `education`         TEXT            NOT NULL,
    `email`             VARCHAR(32)     NOT NULL,
    `native_singlish`   TINYINT         NOT NULL,
    `native_language`   VARCHAR(32),

    `location`          VARCHAR(32)     NOT NULL   
    
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `question` (
--   name         type            constraints
    `id`          INT             NOT NULL      AUTO_INCREMENT,
    `word`        VARCHAR(128)    NOT NULL,
    
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `answer` (
--   name           type             constraints
    `id`            INT              NOT NULL    AUTO_INCREMENT,
    `accossiation`  VARCHAR(32)      NOT NULL,
    
    `question_id`   INT              NOT NULL,   
    `repondent_id`  INT              NOT NULL

    PRIMARY KEY (`id`),
    FOREIGN KEY (`question_id`)  REFERENCES question(`id`),
    FOREIGN KEY (`repondent_id`) REFERENCES repondent(`id`),
);