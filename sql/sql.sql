CREATE DATABASE IF NOT EXISTS Todoist;
USE Todoist;

DROP TABLE IF EXISTS Tarefas;
DROP TABLE IF EXISTS Token;
DROP TABLE IF EXISTS usuarios;


CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(30) not null,
    nick varchar(30) not null unique,
    senha varchar(150) not null,
    criadaEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE Tarefas (
    id int auto_increment primary key,
    titulo varchar(30) not null,
    descricao varchar(150) not null,
    statu tinyint(0),

    tarefa_usuario int not null,
    FOREIGN KEY (tarefa_usuario)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    criadaEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE Token (
    id int auto_increment primary key,
    id_usuario int,
    token varchar(150),

    FOREIGN KEY (id_usuario)
    REFERENCES usuarios(id)
    ON DELETE CASCADE
) ENGINE=INNODB;