SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';
-- -----------------------------------------------------
-- Schema schema
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `doodlering` ;
USE `doodlering` ;

-- -----------------------------------------------------
-- Table `doodlering`.`Games`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Games` (
  `key` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`key`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `doodlering`.`Play`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Play` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Games_key` VARCHAR(255) NOT NULL,
  `sequence` VARCHAR(255) NOT NULL,
  `sentence` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`, `Games_key`),
  INDEX `fk_Play_Games1_idx` (`Games_key` ASC),
  CONSTRAINT `fk_Play_Games1`
    FOREIGN KEY (`Games_key`)
    REFERENCES `doodlering`.`Games` (`key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `doodlering`.`Words`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Words` (
  `word` VARCHAR(255) NOT NULL,
  `Play_id` INT NOT NULL,
  `Play_Games_key` VARCHAR(255) NOT NULL,
  CONSTRAINT `fk_Words_Play1`
    FOREIGN KEY (`Play_id` , `Play_Games_key`)
    REFERENCES `doodlering`.`Play` (`id` , `Games_key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `doodlering`.`Coordinate`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Coordinate` (
  `coordinate` VARCHAR(255) NOT NULL,
  `Play_id` INT NOT NULL,
  `Play_Games_key` VARCHAR(255) NOT NULL,
  CONSTRAINT `fk_Coordinate_Play1`
    FOREIGN KEY (`Play_id` , `Play_Games_key`)
    REFERENCES `doodlering`.`Play` (`id` , `Games_key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `doodlering`.`Sentences` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `sentence` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

ALTER TABLE `doodlering`.`Play` 
DROP COLUMN `sentence`;

ALTER TABLE `doodlering`.`Coordinate` 
CHANGE COLUMN `coordinate` `x` VARCHAR(255) NOT NULL AFTER `Play_Games_key`,
ADD COLUMN `y` VARCHAR(255) NOT NULL AFTER `x`,
ADD COLUMN `dx` VARCHAR(255) NOT NULL AFTER `y`,
ADD COLUMN `dy` VARCHAR(255) NOT NULL AFTER `dx`;

ALTER TABLE `doodlering`.`Words` 
CHANGE COLUMN `word` `label` VARCHAR(255) NOT NULL ,
ADD COLUMN `confidence` VARCHAR(255) NOT NULL AFTER `Play_Games_key`;

ALTER TABLE `Doodlering`.`Words` 
CHANGE COLUMN `confidence` `confidence` FLOAT NOT NULL ;

ALTER TABLE `Doodlering`.`Coordinate` 
CHANGE COLUMN `x` `x` FLOAT NOT NULL ,
CHANGE COLUMN `y` `y` FLOAT NOT NULL ,
CHANGE COLUMN `dx` `dx` FLOAT NOT NULL ,
CHANGE COLUMN `dy` `dy` FLOAT NOT NULL ;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 _____에서 온 사람을 보았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('정말 마음씨가 죽었나, 다시 _____을 보고 싶었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('비명을 지르며 무릎을 꿇어 인사했더니 _____이 사람들을 먹고 말았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 어떤 그림에게 ____를 주고 싶었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 그러다 고개를 들어 _____을 다시 집에 들여왔습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('죽어가는 _____를 살리고 말았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 _____을 내려다보다가 산골로 돌아가려 했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('_____에는 더 이상 구경꾼들이 보이지 않았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('전처럼 이제 _____에서 영감을 받지 못한다고 생각했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('누군가 당신의 못된 손에 _____를 주었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('한 임금님이 _____를 만들어 보았습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어떤 사람들은 _____가 되었습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그 때마다 주편의 풍경이 _____처럼 바뀌는 것이었어요.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그러다가 _____들이 돌아왔습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('옛날로 찾아가서 _____를 꾸었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사람들이 이상한 _____ 둘러싸고 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('한 사람은 _____를 정말 좋아했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('한 _____가 집에 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그런데 이제 _____를 갖고 있다는 생각이 들었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이리하여 이 _____은 자신의 구멍으로 돌아갔습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('_____을 만드는 방법을 생각해 보았습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그러다 어느 마을에서 사람들이 _____가 된 것을 보았습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어쩌면 _____이 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 마음속으로 _____에게 말했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('_____가 어린아이들을 건너가게 하고 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사실 _____가 된다면, 마음속으로부터 그것을 원했기 때문입니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('젊은이의 장미를 가져올 _____를 알게 되었답니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사람들이 있었습니다. 그 사람들은 불쌍한 _____를 보고 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('_____가 되지 않으면 이상하다는 생각이 들었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어느 날 밥을 받은 _____가 나섰습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어떤 소원을 지녔는지 모르다 보니 이 _____ 앞을 지나칠 수가 없었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사람들은 이 _____를 보고 모두 소스라치게 놀랐습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이 광경에서는 꼬마 _____가 나무에 대한 이야기를 하기 위해서 말문을 열었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('_____는 두 번 다시 잡아먹히지 않았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이윽고 사람들은 _____이 되었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('_____의 욕심이 그렇게 되었답니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('도둑이 _____를 자리에 앉히는 광경이었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이것이 _____와 한 일이란 사실에 무릎을 쳤습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이 때 이 광경을 지켜보는 _____이 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('_____를 더 이상 기억하지 않았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그러다 자리에 앉았는데 _____를 바라볼 시간이 있었습니다.');