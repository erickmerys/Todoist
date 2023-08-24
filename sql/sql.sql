CREATE DATABASE IF NOT EXISTS Todoist;
USE Todoist;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS Tarefas;

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