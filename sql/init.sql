CREATE DATABASE IF NOT EXISTS `weather_db`;

USE weather_db;

CREATE TABLE IF NOT EXISTS `registries` (
		`id` VARCHAR(255) NOT NULL,
		`cityname` VARCHAR(255) NOT NULL,
		`temperature` VARCHAR(255) NOT NULL,
		`statecode` INT(10),
		`description` VARCHAR(255) NOT NULL
);

