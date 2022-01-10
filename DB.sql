-- MySQL Script generated by MySQL Workbench
-- Mon Jan 10 11:48:03 2022
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mydb` DEFAULT CHARACTER SET utf8 ;
USE `mydb` ;

-- -----------------------------------------------------
-- Table `mydb`.`Template`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`Template` (
)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`Macro`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`Macro` (
)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`User`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`User` (
  `ID` INT NOT NULL,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`Usergroup`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`Usergroup` (
)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`Attachement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`Attachement` (
  `ID` INT NOT NULL,
  `BelongsToTemplate` VARCHAR(45) NULL,
  `Name` VARCHAR(45) NULL,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mydb`.`SavedAttachement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mydb`.`SavedAttachement` (
  `ID` INT NOT NULL,
  `Path` VARCHAR(100) NULL,
  `BelongsToAttachement` VARCHAR(45) NULL,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
