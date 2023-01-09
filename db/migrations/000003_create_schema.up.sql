ALTER TABLE `avatar`
 ADD `avatar_text` varchar(45) NOT NULL,
 ADD `image_url` varchar(255),
 ADD `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;