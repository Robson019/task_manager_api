// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Robson Gominho",
            "email": "rag2@aluno.ifal.edu.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/profile": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Rota que retorna todas as informações do usuário logado.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Conta do usuário"
                ],
                "summary": "Pesquisar dados do perfil do usuário logado",
                "operationId": "FindProfile",
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.Account"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Rota que permite que um usuário se autentique no sistema utilizando seu endereço de e-mail e senha.\n| E-mail              | Senha     | Função                                                            |\n|---------------------|-----------|-------------------------------------------------------------------|\n| robson@gmail.com | Test1234! | Usuário do sistema. |",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas de autenticação"
                ],
                "summary": "Fazer login no sistema",
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para que o login seja realizado.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.Authorization"
                        }
                    },
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Rota que permite que um usuário faça logout no sistema.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas de autenticação"
                ],
                "summary": "Fazer logout no sistema",
                "operationId": "Logout",
                "responses": {
                    "204": {
                        "description": "Requisição realizada com sucesso."
                    },
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas de autenticação"
                ],
                "summary": "Gerar um novo par de tokens para autenticação",
                "operationId": "Refresh",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para que o login seja realizado.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RefreshDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.Authorization"
                        }
                    },
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Rota que permite a listagem das tarefas.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tarefas"
                ],
                "summary": "Listar tarefas",
                "operationId": "FindTasks",
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Task"
                            }
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Rota que permite a criação de uma tarefa pelo usuário.\n| Status  |\n|---------|\n| pending |\n| progress |\n| done |",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tarefas"
                ],
                "summary": "Criar tarefa",
                "operationId": "CreateTask",
                "parameters": [
                    {
                        "description": "JSON com todos os dados necessários para criar uma tarefa.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.TaskDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.DefaultPostResponse"
                        }
                    },
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Rota que permite a busca de uma tarefa pelo id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tarefas"
                ],
                "summary": "Buscar tarefa",
                "operationId": "FindTaskByID",
                "parameters": [
                    {
                        "type": "string",
                        "default": "03b3aecd-1b52-4357-875c-298a4bc60132",
                        "description": "ID da tarefa.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/response.Task"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Rota que permite a atualização de uma tarefa pelo usuário.\n| Status  |\n|---------|\n| pending |\n| progress |\n| done |",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tarefas"
                ],
                "summary": "Atualizar tarefa",
                "operationId": "UpdateTask",
                "parameters": [
                    {
                        "type": "string",
                        "default": "03b3aecd-1b52-4357-875c-298a4bc60132",
                        "description": "ID da tarefa.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "JSON com todos os dados necessários para criar uma tarefa.",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.TaskDTO"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Requisição realizada com sucesso."
                    },
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Rota que permite apagar uma tarefa.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tarefas"
                ],
                "summary": "Deletar tarefa",
                "operationId": "DeleteTask",
                "parameters": [
                    {
                        "type": "string",
                        "default": "03b3aecd-1b52-4357-875c-298a4bc60132",
                        "description": "ID da tarefa.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Requisição realizada com sucesso."
                    },
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.LoginDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "robson@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "Test1234!"
                }
            }
        },
        "request.RefreshDTO": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "request.TaskDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Task Description"
                },
                "status": {
                    "type": "string",
                    "example": "pending"
                },
                "title": {
                    "type": "string",
                    "example": "Task Title"
                }
            }
        },
        "response.Account": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/response.Role"
                }
            }
        },
        "response.Authorization": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.DefaultPostResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "response.ErrorMessage": {
            "type": "object",
            "properties": {
                "duplicated_fields": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "invalid_fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.InvalidField"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "response.InvalidField": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "field_name": {
                    "type": "string"
                }
            }
        },
        "response.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.Task": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Task Manager API",
	Description:      "API de gerenciamento de tarefas para disciplina de DEVOPS",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
