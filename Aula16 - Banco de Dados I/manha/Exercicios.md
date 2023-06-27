## 1 - O que é redundância de dados?
A redundância de dados em um banco de dados ocorre quando as mesmas informações são armazenadas repetidamente em diferentes locais. Isso pode ser intencional ou não. A redundância de dados pode ser considerada um problema em sistemas de banco de dados, pois pode causar:

- Ineficiência de armazenamento: ocupa mais espaço de armazenamento do que o necessário, o que pode resultar em custos adicionais para aquisição e manutenção de infraestrutura de armazenamento.

- Inconsistência de dados: Se as informações redundantes não forem atualizadas corretamente, diferentes cópias dos dados podem se tornar inconsistentes. Por exemplo, se uma alteração for feita em uma cópia de um dado e não em outras cópias, isso resultará em divergências e inconsistências nos dados.

- Dificuldade de manutenção: A redundância aumenta a complexidade da manutenção dos dados. Atualizar várias cópias dos mesmos dados requer mais esforço e pode causar mais erros.

- Aumento da probabilidade de anomalias: Maiores chances de ocorrerem problemas durante as operações de inserção, atualização ou exclusão de dados. Por exemplo, durante a atualização de uma cópia de dados, outras cópias podem ser afetadas de maneira inconsistente, levando a resultados imprecisos ou incorretos.

## 2 - O que é inconsistência de dados?
A inconsistência de dados ocorre quando informações redundantes não são atualizadas corretamente, assim, diferentes cópias dos dados podem se tornar inconsistentes. Ou então, podemos dizer que os dados são inconsistentes quando uma mesma busca gera resultados diferentes.

## 3 - Nomeie pelo menos 2 bancos de dados entre Consistência e Tolerância de Particionamento.
MongoDB e Redis

## 4 - O que significa SQL?
Significa "Structured Query Language". É uma linguagem usada para gerenciar bancos de dados relacionais. Ela fornece uma maneira padronizada e universal de interagir com bancos de dados, permitindo realizar diversas operações, como consulta, inserção, atualização e exclusão de dados.

## 5 - Em uma tabela de dados, o que é a Primary Key ou Chave Primária?
A Primary Key ou Chave Primária é usada para identificar exclusivamente cada registro em uma tabela de banco de dados. A Primary Key garante a integridade dos dados, evita duplicatas e fornece um meio eficiente de pesquisa e relacionamento entre tabelas. As Primary Keys podem ser compostas por um único campo ou por vários campos, formando uma chave primária composta. Ao criar uma tabela, é necessário especificar qual coluna ou conjunto de colunas será usado como Primary Key.

## 6 - Em uma tabela de dados, o que é a Foreign Key ou Chave Estrangeira?
A Foreign Key ou Chave Estrangeira, é um conceito em bancos de dados relacionais que estabelece uma relação entre duas tabelas. Ela é usada para garantir a integridade referencial, permitindo que uma tabela se relacione com outra tabela por meio de uma chave primária, pois ela faz referência à Primary Key de outra tabela, criando uma ligação entre os registros.

## 7 - O que significa que um relacionamento entre duas entidades é UM para UM?
A cardinalidade do relacionamento é considerada "um para um ou 1:1" quando um único registro de uma entidade está associado a no máximo um registro da outra entidade, e vice-versa. Ou seja, para cada registro em uma entidade, há no máximo um registro correspondente na outra entidade.

## 8 - O que significa que um relacionamento entre duas entidades é UM para MUITOS?
A cardinalidade do relacionamento é considerada "um para muitos ou 1:N" quando um registro de uma entidade pode estar associado a vários registros da outra entidade, mas cada registro dessa outra entidade está associado a no máximo um registro da primeira entidade. Ou seja, para cada registro em uma entidade, pode haver vários registros correspondentes na outra entidade.

## 9 - O que significa que um relacionamento entre duas entidades é de MUITOS para MUITOS?
A cardinalidade do relacionamento é considerada "muitos para muitos ou N:N" quando múltiplos registros de uma entidade podem estar associados a vários registros da outra entidade, e vice-versa. 

## 10 - Defina o que é um banco de dados.
Um banco de dados é uma coleção organizada de informações estruturadas, armazenadas de forma persistente em um sistema de armazenamento. Ele é projetado para permitir a criação, armazenamento, recuperação, modificação e exclusão eficiente de dados.

Os bancos de dados são amplamente utilizados em diversas áreas e desempenham um papel crucial no armazenamento e gerenciamento de grandes volumes de informações de maneira estruturada e acessível. Eles fornecem uma estrutura robusta para gerenciar dados, garantir a integridade dos dados e suportar a tomada de decisões baseada em informações.

## 11 - Defina o conceito de atributo.
Em bancos de dados, os atributos são características ou propriedades que descrevem os dados armazenados em uma tabela. Também são conhecidos como campos ou colunas. Cada tabela em um banco de dados é composta por um conjunto de atributos que definem os tipos de dados e as informações a serem armazenadas. Cada atributo possui um nome exclusivo e um tipo de dado que determina o formato dos valores que podem ser armazenados nele, como texto, número, data, entre outros.
