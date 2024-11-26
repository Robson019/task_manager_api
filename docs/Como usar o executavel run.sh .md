# run.sh
Executável que facilita a execução dos comandos da API

### Uso geral

* ``source run.sh -backend`` ou ``./run.sh -backend``: Roda o projeto e sobe o servidor
  - Carrega todas as variáveis de ambiente do projeto. 
  - Inicializa o banco da aplicação 
  - Baixa todas as bilbiotecas do projeto 
  - Gera ou atualiza a documentação da API 
  - Aguarda o banco de dados iniciar e carrega as migrations
  - Por fim, o servidor é inicado

### SQLc

* ``source run.sh -sqlc``: Gera as funções em go a partir do código sql das queries

_**OBS.:** Você pode consultar ambos os arquivos pelo caminho: `tools/executables/scripts`_.
