INSERT INTO usuarios (nome)
VALUES (?)
RETURNING id, nome;
