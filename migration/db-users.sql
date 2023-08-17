create table users(
    id serial primary key,
    nome varchar,
    cargo varchar
);

INSERT INTO users(nome, cargo) VALUES
('Cauã Reymond', 'Ator'),
('Murylo', 'Desenvolvedor');