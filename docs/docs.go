// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/begin-task-for-user": {
            "post": {
                "description": "Begin task for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Time Tracking"
                ],
                "summary": "Begin task for user",
                "parameters": [
                    {
                        "description": "Passport number",
                        "name": "pasportNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/calculate-cost-by-user": {
            "get": {
                "description": "Возвращает затраты времени на задачи по идентификатору пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Time Tracking"
                ],
                "summary": "Затраты времени на задачи",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Серия паспорта",
                        "name": "pasportSeries",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Номер паспорта",
                        "name": "pasportNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Начало периода (в формате ISO 8601)",
                        "name": "periodFrom",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Окончание периода (в формате ISO 8601)",
                        "name": "periodTo",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список счетов (costs)",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/end-task-for-user": {
            "post": {
                "description": "Finish tracking time for a specific task",
                "produces": [
                    "application/json"
                ],
                "summary": "Finish task time tracking",
                "parameters": [
                    {
                        "description": "Passport number",
                        "name": "pasportNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/info": {
            "get": {
                "description": "Get user data by passport series and number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Passport series",
                        "name": "pasportSeries",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Passport number",
                        "name": "pasportNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/timetracking.User"
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get users data by filter and pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get users data by filter and pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/timetracking.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update user data by passport series and number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Passport series",
                        "name": "pasportSeries",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Passport number",
                        "name": "pasportNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/timetracking.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user with passport data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/timetracking.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "int32"
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user by passport series and number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Passport series",
                        "name": "pasportSeries",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Passport number",
                        "name": "pasportNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "timetracking.User": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "адрес",
                    "type": "string"
                },
                "name": {
                    "description": "имя",
                    "type": "string"
                },
                "patronymic": {
                    "description": "отчество",
                    "type": "string"
                },
                "surname": {
                    "description": "фамилия",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}