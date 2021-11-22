CREATE TABLE `users`
(
    id   bigint auto_increment,
    iin varchar(255) NOT NULL UNIQUE,
    username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`iin`, `username`, `password`)
VALUES ('0','admin', 'password');