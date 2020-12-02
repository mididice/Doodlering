SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';
-- -----------------------------------------------------
-- Schema schema
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `schema` ;
USE `schema` ;

-- -----------------------------------------------------
-- Table `schema`.`Games`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schema`.`Games` (
  `key` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`key`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schema`.`Play`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schema`.`Play` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Games_key` VARCHAR(255) NOT NULL,
  `sequence` VARCHAR(255) NOT NULL,
  `sentence` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`, `Games_key`),
  INDEX `fk_Play_Games1_idx` (`Games_key` ASC),
  CONSTRAINT `fk_Play_Games1`
    FOREIGN KEY (`Games_key`)
    REFERENCES `schema`.`Games` (`key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schema`.`Words`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schema`.`Words` (
  `word` VARCHAR(255) NOT NULL,
  `Play_id` INT NOT NULL,
  `Play_Games_key` VARCHAR(255) NOT NULL,
  CONSTRAINT `fk_Words_Play1`
    FOREIGN KEY (`Play_id` , `Play_Games_key`)
    REFERENCES `schema`.`Play` (`id` , `Games_key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schema`.`Coordinate`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schema`.`Coordinate` (
  `coordinate` VARCHAR(255) NOT NULL,
  `Play_id` INT NOT NULL,
  `Play_Games_key` VARCHAR(255) NOT NULL,
  CONSTRAINT `fk_Coordinate_Play1`
    FOREIGN KEY (`Play_id` , `Play_Games_key`)
    REFERENCES `schema`.`Play` (`id` , `Games_key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `schema`.`Sentences` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `sentence` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

ALTER TABLE `schema`.`Play` 
DROP COLUMN `sentence`;

ALTER TABLE `schema`.`Coordinate` 
CHANGE COLUMN `coordinate` `x` VARCHAR(255) NOT NULL AFTER `Play_Games_key`,
ADD COLUMN `y` VARCHAR(255) NOT NULL AFTER `x`,
ADD COLUMN `dx` VARCHAR(255) NOT NULL AFTER `y`,
ADD COLUMN `dy` VARCHAR(255) NOT NULL AFTER `dx`;

ALTER TABLE `schema`.`Words` 
CHANGE COLUMN `word` `label` VARCHAR(255) NOT NULL ,
ADD COLUMN `confidence` VARCHAR(255) NOT NULL AFTER `Play_Games_key`;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
