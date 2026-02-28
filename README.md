# API-estudantes
<br>

<h2> Este foi o projeto que aprendi no Curso Golang do Zero </h2>
API para gerenciar os estudantes no Curso Golang do Zero.

<h2> Introdução </h2>
O objetivo da aplicação é criar um sistema para controlar os estudantes cadastrados e ativos na plataforma do curso. 

## Tecnologias utilizadas
* VS Code
* Golang (Go)
* Insomnia
* Echo v4
* GORM
* SQLite (extensão do VS Code)
* SQLite Viewer (extensão do VS Code)
* C++ (extensão do VS Code)
* Zerolog


## Rotas

| Método | URL             | Descrição                                        |
| ------ | --------------  | -------------------------------------------------|
| GET    | /estudantes     | Listar todos os alunos                           |
| POST   | /estudantes     | Cadastrar um estudante                           |
| GET    | /estudantes/:id | Pegar a informação de um estudante específico    |
| PUT    | /estudantes/:id | Atualizar informações de um estudante específico |
| DELETE | /estudantes/:id | Deletar um estudante específico                  |

## Estrutura do estudante:
- Nome
- CPF
- E-mail
- Idade
- Ativo