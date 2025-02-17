# Task Manager API

## Sumário

- [Introdução](#introdução)
- [Funcionamento do sistema](#funcionamento-do-sistema)
  - [Divisão dos módulos](#Divisão-dos-módulos)
- [Como executar o projeto?](#como-executar-o-projeto)
  - [Quais são as ferramentas necessárias?](#quais-são-as-ferramentas-necessárias)
  - [Como clonar o projeto?](#como-clonar-o-projeto)
  - [Como rodar o projeto?](#como-rodar-o-projeto)
  - [Como acessar a documentação das rotas?](#como-acessar-a-documentação-das-rotas)
- [Ferramentas de auxilio ao desenvolvimento](#ferramentas-de-desenvolvimento)

## Introdução

Este projeto foi desenvolvido para obtenção de nota da disciplina de **Devops**, ministrada pelo professor [Ítalo Carlo](https://github.com/italocarlo06).

## Funcionamento do sistema

O sistema se trata de uma API de gerenciamento de tarefas simplificada, visto que o objetivo propriamente dito da aplicação está no desenvolvimento dos containers docker com o auxílio de variáveis de ambiente para gerenciar configurações sensíveis, um usuário customizado para o banco de dados, utilizando também volumes para persistência dos dados com uso de uma network customizada para garantir a configuração de um ambiente multi-container funcional com Docker Compose. Além disso, o arquivo conta com um Health Ckeck para garantir a integridade do container.

### Divisão dos módulos

- **Conta do Usuário:** É o módulo responsável por apresentar informações básicas do perfil do usuário logado no sistema.
- **Autenticação:** O projeto possui um usuário base para que as rotas não sejam acessadas pelo usuário anônimo. Este módulo conta com login, logout e refresh do token de autenticação. 
- **Tarefas:** Módulo com o CRUD básico do projeto, onde o usuário logado consegue manter lembretes de tarefas simples.

## Como executar o projeto?

### Quais são as ferramentas necessárias?

- **Git:** É o sistema de controle de versões utilizado no projeto. Você pode fazer o download dele [aqui](https://git-scm.com/);
- **Docker:** Foco atual da disciplina, é utilizado para criar, gerenciar e remover containers. Com ele você poderá rodar o banco de dados e/ou a API do sistema de uma forma muito simples. Você pode fazer o download dele [aqui](https://docs.docker.com/desktop/);
- **Docker Compose:** Em versões mais antigas, é um comando separado, mas nas mais recentes já vem embutido no comando `docker`. O comando `docker-compose` ou `docker compose` facilita a orquestração de containers. Você pode fazer o download dele [aqui](https://docs.docker.com/compose/install/);
- **swag:** É um executável que nos permite gerar a documentação das rotas da API de uma forma simples. Você pode fazer o download dele [aqui](https://github.com/swaggo/swag);
- **migrate:** É um executável que lê as migrations (arquivos .sql) que foram criadas para realizar operações no banco e as aplica na ordem correta no banco de dados. Você pode fazer o download dele [aqui](https://github.com/golang-migrate/migrate);
- **sqlc:** É um executável gerador de código Go com segurança de tipo a partir do SQL. Será utilizado para gerar métodos e DTOs Go a partir de queries SQL (como SELECTs, INSERTs, UPDATEs etc.). Você pode fazer o download dele [aqui](https://github.com/kyleconroy/sqlc);
- **Go:** É a linguagem de programação utilizada no desenvolvimento deste backend. Você pode fazer o download dela [aqui](https://go.dev/dl/);
- **Banco de dados:** Os bancos utilizados neste backend serão o `postgreSQL` e o `redis`, que serão construídos no container docker, portanto não haverá necessidade configurá-lo previamente.

### Como clonar o projeto?

1. Acesse o repositório do projeto: https://github.com/Robson019/task_manager_api.git;
2. Na linha de comando, execute: `git clone https://github.com/Robson019/task_manager_api.git`.

### Como rodar o projeto?

1. Depois de baixar e configurar corretamente todas as dependências do tópico [Quais são as ferramentas necessárias?](#quais-são-as-ferramentas-necessárias)
2. A partir da pasta raiz do projeto, acesse: `src/api/app`;
3. Copie o arquivo `.env.example` no mesmo diretório e renomeie sua cópia como `.env` (este arquivo contém as informações de exemplo para uma configuração segura da API);
4. Retorne até a pasta raiz do projeto `cd ../../..`;
5. No terminal, execute o comando `./run.sh -backend` ou `source run.sh -backend` para executar todos os comandos de configuração do backend do projeto (Lembre-se de configurar a quebra de linha desses arquivos como LF).
6. Aguarde a finalização dos comandos e a execução do servidor.

_**OBS.:** O tópico [Ferramentas de auxilio ao desenvolvimento](#ferramentas-de-auxilio-ao-desenvolvimento) explica detalhadamente sobre o funcionamento do script run.sh_

### Como acessar a documentação das rotas?

1. Abra um navegador com o servidor em execução;
2. Acesse a URL http://localhost:8000/api/docs/index.html;
3. Acesse a rota de **login** do módulo de **autenticação**;
4. Logue com o usuário fictício `robson@gmail.com` para ter acesso as rotas não disponíveis ao usuário anônimo;
5. Com o token de acesso gerado, insira-o no campo de autorização, com o seguinte prefixo: `bearer [token gerado]`
6. Teste o CRUD de **Tarefas** ou qualquer outra rota de sua escolha.

<div id="ferramentas-de-desenvolvimento"></div>

## Ferramentas de auxilio ao desenvolvimento

### run.sh

O executável, que se encontra na raiz do projeto, auxilia na execução dos comandos extensos da ferramenta, ``sqlc`` e também **executa a aplicação**. O script auxilia na performance do desenvolvimento, e também, previne bugs relacionados às outras ferramentas. [Clique para saber como usar](./docs/Como%20usar%20o%20executavel%20run.sh%20.md)
