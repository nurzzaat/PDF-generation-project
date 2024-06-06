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
        "/forgot-password": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Forgot Password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "LogOut",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "parameters": [
                    {
                        "description": "login",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Syllabus"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "description": "syllabus",
                        "name": "syllabus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SyllabusInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/generate/{id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/literature/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "syllabus",
                        "name": "syllabus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Syllabus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/main/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "syllabus",
                        "name": "syllabus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Syllabus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/others": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "subject",
                        "name": "subject",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Syllabus"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/preface/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "syllabus",
                        "name": "syllabus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Syllabus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/question/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "syllabus",
                        "name": "syllabus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Syllabus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/text/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "syllabus",
                        "name": "syllabus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Syllabus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/topic/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "syllabus",
                        "name": "syllabus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Syllabus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/syllabus/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Syllabus"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Syllabus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/reset-password": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "Change password",
                        "name": "reset",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Password"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Confirmer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "consultation": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "faculty": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "specialist": {
                    "type": "string"
                }
            }
        },
        "models.ErrorDetail": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/models.Properties"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ErrorDetail"
                    }
                }
            }
        },
        "models.Header": {
            "type": "object",
            "properties": {
                "specialityName": {
                    "type": "string"
                },
                "subjectName": {
                    "type": "string"
                }
            }
        },
        "models.Litrature": {
            "type": "object",
            "properties": {
                "additionalLiterature": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "internetSource": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "mainLiterature": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Modules": {
            "type": "object",
            "properties": {
                "moduleName": {
                    "type": "string"
                },
                "topic": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Topic"
                    }
                }
            }
        },
        "models.Password": {
            "type": "object",
            "properties": {
                "confirmPassword": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.PrefaceInfo": {
            "type": "object",
            "properties": {
                "confirmedBy": {
                    "$ref": "#/definitions/models.Confirmer"
                },
                "discussedBy1": {
                    "$ref": "#/definitions/models.Confirmer"
                },
                "discussedBy2": {
                    "$ref": "#/definitions/models.Confirmer"
                },
                "discussion1": {
                    "type": "string"
                },
                "discussion2": {
                    "type": "string"
                },
                "insertedIn": {
                    "type": "string"
                },
                "madeBy": {
                    "$ref": "#/definitions/models.Confirmer"
                }
            }
        },
        "models.Properties": {
            "type": "object",
            "properties": {
                "additionalProp1": {
                    "type": "string"
                },
                "additionalProp2": {
                    "type": "string"
                },
                "additionalProp3": {
                    "type": "string"
                }
            }
        },
        "models.Question": {
            "type": "object",
            "properties": {
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.SuccessResponse": {
            "type": "object",
            "properties": {
                "metadata": {
                    "$ref": "#/definitions/models.Properties"
                },
                "result": {}
            }
        },
        "models.Syllabus": {
            "type": "object",
            "properties": {
                "literature": {
                    "$ref": "#/definitions/models.Litrature"
                },
                "mainInfo": {
                    "$ref": "#/definitions/models.SyllabusInfo"
                },
                "preface": {
                    "$ref": "#/definitions/models.PrefaceInfo"
                },
                "question1": {
                    "$ref": "#/definitions/models.Question"
                },
                "question2": {
                    "$ref": "#/definitions/models.Question"
                },
                "syllabusID": {
                    "type": "integer"
                },
                "text": {
                    "$ref": "#/definitions/models.Text"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Modules"
                    }
                }
            }
        },
        "models.SyllabusInfo": {
            "type": "object",
            "properties": {
                "allHours": {
                    "type": "integer"
                },
                "courseNumber": {
                    "type": "integer"
                },
                "creditNumber": {
                    "type": "integer"
                },
                "facultyName": {
                    "type": "string"
                },
                "kafedraName": {
                    "type": "string"
                },
                "lectureHours": {
                    "type": "integer"
                },
                "practiceLessons": {
                    "type": "integer"
                },
                "sro": {
                    "type": "string"
                },
                "srop": {
                    "type": "string"
                },
                "subjectInfo": {
                    "$ref": "#/definitions/models.Header"
                }
            }
        },
        "models.Text": {
            "type": "object",
            "properties": {
                "text2": {
                    "type": "string"
                },
                "text3": {
                    "type": "string"
                },
                "text4": {
                    "type": "string"
                },
                "text5": {
                    "type": "string"
                },
                "text6": {
                    "type": "string"
                },
                "text7": {
                    "type": "string"
                },
                "text8": {
                    "type": "string"
                }
            }
        },
        "models.Topic": {
            "type": "object",
            "properties": {
                "LK": {
                    "type": "integer"
                },
                "SPZ": {
                    "type": "integer"
                },
                "SRO": {
                    "type": "integer"
                },
                "literature": {
                    "type": "string"
                },
                "topicName": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "185.146.1.71",
	BasePath:         "/pdf",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
