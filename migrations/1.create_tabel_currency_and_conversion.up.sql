CREATE TABLE IF NOT EXISTS `currency` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE IF NOT EXISTS `conversion` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `currencyIDFrom` int,
  `currencyIDTo` int,
  `rate` int
);

ALTER TABLE `conversion` ADD FOREIGN KEY (`currencyIDFrom`) REFERENCES `currency` (`id`);

ALTER TABLE `conversion` ADD FOREIGN KEY (`currencyIDTo`) REFERENCES `currency` (`id`);
