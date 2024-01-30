DROP TABLE IF EXISTS logs;
CREATE TABLE logs (
  id         INT AUTO_INCREMENT NOT NULL,
  date      VARCHAR(255) NOT NULL,
  module     VARCHAR(255) NOT NULL,
  fonction   VARCHAR(255) NOT NULL,
  arguments  VARCHAR(255) NOT NULL,
  output    VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

