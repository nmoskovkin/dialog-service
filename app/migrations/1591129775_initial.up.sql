CREATE TABLE messages
(
    id varchar(36) NOT NULL,
    `from` varchar(36) NOT NULL,
    `to` varchar(36) NOT NULL,
    message  TEXT  NOT NULL,
    PRIMARY KEY(id)
);
