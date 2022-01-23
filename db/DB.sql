-- MySQL Script generated by MySQL Workbench
-- Sun 23 Jan 2022 01:53:45 AM CET
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema schablone
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema schablone
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `schablone` ;
USE `schablone` ;

-- -----------------------------------------------------
-- Table `schablone`.`User`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`User` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Firstname` VARCHAR(50) NOT NULL,
  `Lastname` VARCHAR(50) NOT NULL,
  `Username` VARCHAR(45) NOT NULL,
  `Password` VARCHAR(45) NOT NULL,
  `Hash` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE INDEX `Username_UNIQUE` (`Username` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`Template`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`Template` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(45) NOT NULL,
  `Subject` VARCHAR(100) NOT NULL,
  `Content` BLOB NOT NULL,
  `IsBeingEditedBy` INT NULL,
  PRIMARY KEY (`Id`),
  INDEX `isBeingEdited_idx` (`IsBeingEditedBy` ASC) VISIBLE,
  CONSTRAINT `isBeingEdited`
    FOREIGN KEY (`IsBeingEditedBy`)
    REFERENCES `schablone`.`User` (`ID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`Macro`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`Macro` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(45) NOT NULL,
  `Content` BLOB NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE INDEX `Name_UNIQUE` (`Name` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`TemplateGroup`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`TemplateGroup` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(50) NOT NULL,
  `ParentTemplateGroup` INT NULL,
  PRIMARY KEY (`Id`),
  INDEX `parentTemplateGroup_idx` (`ParentTemplateGroup` ASC) VISIBLE,
  CONSTRAINT `parentTemplateGroup`
    FOREIGN KEY (`ParentTemplateGroup`)
    REFERENCES `schablone`.`TemplateGroup` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`Attachement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`Attachement` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `BelongsToTemplate` INT NOT NULL,
  `Name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`Id`),
  INDEX `BelongsToTemplate_idx` (`BelongsToTemplate` ASC) VISIBLE,
  CONSTRAINT `BelongsToTemplate`
    FOREIGN KEY (`BelongsToTemplate`)
    REFERENCES `schablone`.`Template` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`SavedAttachement`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`SavedAttachement` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Path` VARCHAR(100) NOT NULL,
  `BelongsToAttachement` INT NOT NULL,
  PRIMARY KEY (`ID`),
  INDEX `BelongsToAttachement_idx` (`BelongsToAttachement` ASC) VISIBLE,
  CONSTRAINT `BelongsToAttachement`
    FOREIGN KEY (`BelongsToAttachement`)
    REFERENCES `schablone`.`Attachement` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`User_TemplateGroup`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`User_TemplateGroup` (
  `BelongsToUser` INT NOT NULL,
  `TemplateGroup` INT NOT NULL,
  `UserHasWriteAccess` TINYINT NULL,
  `UserHasUserModifyAccess` TINYINT NULL,
  PRIMARY KEY (`BelongsToUser`, `TemplateGroup`),
  INDEX `User_idx` (`BelongsToUser` ASC) VISIBLE,
  INDEX `TemplateGroup_idx` (`TemplateGroup` ASC) VISIBLE,
  CONSTRAINT `User`
    FOREIGN KEY (`BelongsToUser`)
    REFERENCES `schablone`.`User` (`ID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `TemplateGroup`
    FOREIGN KEY (`TemplateGroup`)
    REFERENCES `schablone`.`TemplateGroup` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`Template_TemplateGroup`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`Template_TemplateGroup` (
  `BelongsToTemplate` INT NOT NULL,
  `TemplateGroup` INT NOT NULL,
  PRIMARY KEY (`BelongsToTemplate`, `TemplateGroup`),
  INDEX `Template_idx` (`BelongsToTemplate` ASC) VISIBLE,
  INDEX `TemplateGroup_idx` (`TemplateGroup` ASC) VISIBLE,
  CONSTRAINT `Template`
    FOREIGN KEY (`BelongsToTemplate`)
    REFERENCES `schablone`.`Template` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `TemplateGroup`
    FOREIGN KEY (`TemplateGroup`)
    REFERENCES `schablone`.`TemplateGroup` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`Macro_TemplateGroup`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`Macro_TemplateGroup` (
  `BelongsToMacro` INT NOT NULL,
  `TemplateGroup` INT NOT NULL,
  PRIMARY KEY (`BelongsToMacro`, `TemplateGroup`),
  INDEX `Macro_idx` (`BelongsToMacro` ASC) VISIBLE,
  INDEX `TemplateGroup_idx` (`TemplateGroup` ASC) VISIBLE,
  CONSTRAINT `Macro`
    FOREIGN KEY (`BelongsToMacro`)
    REFERENCES `schablone`.`Macro` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `TemplateGroup`
    FOREIGN KEY (`TemplateGroup`)
    REFERENCES `schablone`.`TemplateGroup` (`Id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `schablone`.`ActiveAPIKey`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `schablone`.`ActiveAPIKey` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `BelongsToUser` INT NOT NULL,
  `CreationDate` INT NULL,
  PRIMARY KEY (`Id`, `CreationDate`),
  INDEX `BelongsToUser_idx` (`BelongsToUser` ASC) VISIBLE,
  CONSTRAINT `BelongsToUser`
    FOREIGN KEY (`BelongsToUser`)
    REFERENCES `schablone`.`User` (`ID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
