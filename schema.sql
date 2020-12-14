SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';
-- -----------------------------------------------------
-- Schema doodlering
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `doodlering` DEFAULT CHARACTER SET utf8 ;
USE `doodlering` ;

-- -----------------------------------------------------
-- Table `doodlering`.`Games`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Games` (
  `key` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`key`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `doodlering`.`Play`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Play` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `Games_key` VARCHAR(255) NOT NULL,
  `sequence` VARCHAR(255) NOT NULL,
  `gen_date` VARCHAR(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`, `Games_key`),
  INDEX `fk_Play_Games1_idx` (`Games_key` ASC),
  CONSTRAINT `fk_Play_Games1`
    FOREIGN KEY (`Games_key`)
    REFERENCES `doodlering`.`Games` (`key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 10
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `doodlering`.`Coordinate`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Coordinate` (
  `Play_id` INT(11) NOT NULL,
  `Play_Games_key` VARCHAR(255) NOT NULL,
  `x` FLOAT NOT NULL,
  `y` FLOAT NOT NULL,
  `dx` FLOAT NOT NULL,
  `dy` FLOAT NOT NULL,
  INDEX `fk_Coordinate_Play1` (`Play_id` ASC, `Play_Games_key` ASC),
  CONSTRAINT `fk_Coordinate_Play1`
    FOREIGN KEY (`Play_id` , `Play_Games_key`)
    REFERENCES `doodlering`.`Play` (`id` , `Games_key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `doodlering`.`Sentences`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Sentences` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `sentence` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 81
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `doodlering`.`Play_has_Sentences`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Play_has_Sentences` (
  `Play_id` INT(11) NOT NULL,
  `Play_Games_key` VARCHAR(255) NOT NULL,
  `Sentences_id` INT(11) NOT NULL,
  PRIMARY KEY (`Play_id`, `Play_Games_key`, `Sentences_id`),
  INDEX `fk_Play_has_Sentences_Sentences1_idx` (`Sentences_id` ASC),
  INDEX `fk_Play_has_Sentences_Play1_idx` (`Play_id` ASC, `Play_Games_key` ASC),
  CONSTRAINT `fk_Play_has_Sentences_Play1`
    FOREIGN KEY (`Play_id` , `Play_Games_key`)
    REFERENCES `doodlering`.`Play` (`id` , `Games_key`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Play_has_Sentences_Sentences1`
    FOREIGN KEY (`Sentences_id`)
    REFERENCES `doodlering`.`Sentences` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

-- -----------------------------------------------------
-- Table `doodlering`.`Words`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `doodlering`.`Words` (
  `label` VARCHAR(255) NOT NULL,
  `Play_id` INT(11) NOT NULL,
  `Play_Games_key` VARCHAR(255) NOT NULL,
  `confidence` FLOAT NOT NULL)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 <span id="word">_____</span>에서 온 사람을 보았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('정말 마음씨가 죽었나, 다시 <span id="word">_____</span>을(를) 보고 싶었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('비명을 지르며 무릎을 꿇어 인사했더니 <span id="word">_____</span>(이)가 사람들을 먹고 말았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 어떤 그림에게 <span id="word">_____</span>(을)를 주고 싶었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 그러다 고개를 들어 <span id="word">_____</span>을(를) 다시 집에 들여왔습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('죽어가는 <span id="word">_____</span>(을)를 살리고 말았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 <span id="word">_____</span>(을)를 내려다보다가 산골로 돌아가려 했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('<span id="word">_____</span>에는 더 이상 구경꾼들이 보이지 않았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('전처럼 이제 <span id="word">_____</span>에서 영감을 받지 못한다고 생각했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('누군가 당신의 못된 손에 <span id="word">_____</span>(을)를 주었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('한 임금님이 <span id="word">_____</span>(을)를 만들어 보았습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어떤 사람들은 <span id="word">_____</span>(이)가 되었습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그 때마다 주편의 풍경이 <span id="word">_____</span>처럼 바뀌는 것이었어요.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그러다가 <span id="word">_____</span>들이 돌아왔습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('옛날로 찾아가서 <span id="word">_____</span>(을)를 꾸었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사람들이 이상한 <span id="word">_____</span>(을)를 둘러싸고 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('한 사람은 <span id="word">_____</span>(을)를 정말 좋아했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('한 <span id="word">_____</span>(이)가 집에 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그런데 이제 <span id="word">_____</span>(을)를 갖고 있다는 생각이 들었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이리하여 이 <span id="word">_____</span>(은)는 자신의 구멍으로 돌아갔습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('<span id="word">_____</span>(을)를 만드는 방법을 생각해 보았습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그러다 어느 마을에서 사람들이 <span id="word">_____</span>(이)가 된 것을 보았습니다. ');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어쩌면 <span id="word">_____</span>(이)가 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('당신은 마음속으로 <span id="word">_____</span>에게 말했습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('<span id="word">_____</span>(이)가 어린아이들을 건너가게 하고 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사실 <span id="word">_____</span>(이)가 된다면, 마음속으로부터 그것을 원했기 때문입니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('젊은이의 장미를 가져올 <span id="word">_____</span>(을)를 알게 되었답니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사람들이 있었습니다. 그 사람들은 불쌍한 <span id="word">_____</span>(을)를 보고 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('<span id="word">_____</span>(이)가 되지 않으면 이상하다는 생각이 들었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어느 날 밥을 받은 <span id="word">_____</span>(이)가 나섰습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('어떤 소원을 지녔는지 모르다 보니 이 <span id="word">_____</span> 앞을 지나칠 수가 없었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('사람들은 이 <span id="word">_____</span>(을)를 보고 모두 소스라치게 놀랐습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이 광경에서는 꼬마 <span id="word">_____</span>(이)가 나무에 대한 이야기를 하기 위해서 말문을 열었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('<span id="word">_____</span>(은)는 두 번 다시 잡아먹히지 않았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이윽고 사람들은 <span id="word">_____</span>(이)가 되었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('<span id="word">_____</span>의 욕심이 그렇게 되었답니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('도둑이 <span id="word">_____</span>(을)를 자리에 앉히는 광경이었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이것이 <span id="word">_____</span>와 한 일이란 사실에 무릎을 쳤습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('이 때 이 광경을 지켜보는 <span id="word">_____</span>(이)가 있었습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('<span id="word">_____</span>(을)를 더 이상 기억하지 않았습니다.');
INSERT INTO `doodlering`.`Sentences` (`sentence`) VALUES ('그러다 자리에 앉았는데 <span id="word">_____</span>(을)를 바라볼 시간이 있었습니다.');