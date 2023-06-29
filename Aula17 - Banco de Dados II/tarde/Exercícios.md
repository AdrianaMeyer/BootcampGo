### 1.Explique o conceito de normalização e por que ele é usado.

A normalização é um processo utilizado no projeto de bancos de dados relacionais para organizar e estruturar as tabelas de forma eficiente, reduzindo a redundância e a inconsistência dos dados. Seu objetivo principal é garantir a integridade dos dados.

A normalização segue um conjunto de regras, chamadas de formas normais, que definem critérios para a organização dos atributos em tabelas. Cada uma define requisitos específicos que devem ser atendidos para que a tabela esteja nesse nível de normalização, sendo elas:
- 1NF: 	Elimina dados duplicados em atributos. Cria registros separados.
- 2NF: 	Remove colunas que não dependem da chave primária.
- 3NF: 	Remove subconjuntos de dados em várias colunas de uma tabela e crie novas tabelas, com relacionamentos entre elas.
- 4NF: Todas as dependências se foram.

### 2.Adicione um filme à tabela de filmes.

INSERT INTO `movies` VALUES (DEFAULT,NULL,NULL,'Back to the future',10.0,4,'1985-07-03 00:00:00',116,5)

### 3.Adicione um gênero à tabela de gêneros.

INSERT INTO `genres` VALUES (13,'2023-06-29 11:00:00',NULL,'Romance',13,1)

### 4.Associe o filme do Ex 2. ao gênero criado no Ex 3.

UPDATE movies SET genre_id=13 WHERE title="Back to the future"

### 5.Modifique a tabela de atores para que pelo menos um ator tenha o filme adicionado no Ex.2 como favorito.
UPDATE actors SET favorite_movie_id=22 WHERE id=47

### 6.Crie uma cópia de tabela temporária da tabela de filmes.
CREATE TEMPORARY TABLE movies_temp AS SELECT * FROM movies;

### 7.Elimine dessa tabela temporária todos os filmes que ganharam menos de 5 prêmios.
DELETE FROM movies_temp WHERE awards < 5;

### 8.Obtenha a lista de todos os gêneros que possuem pelo menos um filme.

-- Somente os nomes dos gêneros
SELECT DISTINCT genres.name
FROM genres
JOIN movies ON genres.id = movies.genre_id;

-- Os nomes dos gêneros com as respectivas quantidades 
SELECT genres.name AS genero, COUNT(movies.genre_id) AS quantidade_filmes
FROM genres
JOIN movies ON genres.id = movies.genre_id
GROUP BY genres.name;

### 9.Obtenha a lista de atores cujo filme favorito ganhou mais de 3 prêmios.

SELECT actors.id AS ator_id, CONCAT(actors.first_name, ' ', actors.last_name) AS nome_completo, actors.favorite_movie_id, movies.title, movies.awards AS quantidade_awards
FROM actors
JOIN movies ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3;

### 10.Use o plano de explicação para analisar as consultas em Ex.6 e 7.
### 11.O que são índices? Para que servem?
### 12.Crie um índice no nome na tabela de filmes.
### 13.Verifique se o índice foi criado corretamente.