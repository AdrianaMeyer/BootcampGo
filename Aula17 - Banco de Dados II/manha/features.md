# Exercícios
**Carregue o script o movies-db.sql no mysql workbench**

### Casos de uso

### 1 - Retorne a média de avaliação(rating) por gênero do filme

SELECT genres.name as "Gênero", 
AVG(movies.rating) as "Nota média"
FROM genres
JOIN movies
ON genres.id = movies.genre_id
GROUP BY movies.genre_id

### 2 - Retorne o nome dos filmes que um ator participou, além do nome do ator
SELECT movies.title as "Filme",
CONCAT(actors.first_name, ' ', actors.last_name) AS "Nome completo"
FROM movies
JOIN actors

### 3 - Retorne a quantidade de atores que cada filme teve

SELECT movies.title AS "Filme", 
COUNT(actors.id) AS "Qnt. Atores"
FROM movies
JOIN actor_movie ON movies.id = actor_movie.movie_id
JOIN actors ON actor_movie.actor_id = actors.id
GROUP BY movies.title;

### 4 - Retorne a quantidade de atores que cada filme teve, além do nome do filme, dos filmes que tem um rating maior ou igual a "3.0"

SELECT movies.title AS "Filme", 
COUNT(actors.id) AS "Qnt. Atores"
FROM movies
JOIN actor_movie ON movies.id = actor_movie.movie_id
JOIN actors ON actor_movie.actor_id = actors.id
WHERE movies.rating >= 3.0
GROUP BY movies.title;

