// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/channels": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Channels"
                ],
                "summary": "List channels of the current user",
                "operationId": "ListChannels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/IdResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Channels"
                ],
                "summary": "Create a new channel",
                "operationId": "CreateChannel",
                "parameters": [
                    {
                        "description": "CreateChannel request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateChannelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/IdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/channels/{id}": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Channels"
                ],
                "summary": "Get channel by ID",
                "operationId": "GetChannel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ChannelResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/channels/{id}/episodes": {
            "get": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EpisodesPage"
                ],
                "summary": "List episoded of the given channel",
                "operationId": "ListEpisodes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/IdResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EpisodesPage"
                ],
                "summary": "Create a new episode",
                "operationId": "CreateEpisode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Channel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CreateEpisode request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateEpisodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/IdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/files": {
            "post": {
                "security": [
                    {
                        "Authentication": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Uploads a new audio file",
                "operationId": "UploadFile",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/IdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new user",
                "operationId": "CreateUser",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/IdResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ChannelResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "x-order": "0",
                    "example": "ch_2K9BWVNuo3sG4yM322fbP3mB6ls"
                },
                "feed_url": {
                    "type": "string",
                    "x-order": "1"
                },
                "title": {
                    "type": "string",
                    "x-order": "2",
                    "example": "Bored Owls Online Radio"
                },
                "link": {
                    "type": "string",
                    "x-order": "3",
                    "example": "https://example.com"
                },
                "authors": {
                    "type": "string",
                    "x-order": "4",
                    "example": "The Owl"
                },
                "description": {
                    "type": "string",
                    "x-order": "5",
                    "example": "Bored owls talk about whatever happens to be on their minds"
                },
                "created_at": {
                    "type": "string",
                    "format": "date-time",
                    "x-order": "6",
                    "example": "2023-01-01T01:02:03.456Z"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time",
                    "x-order": "7",
                    "example": "2023-01-01T01:02:03.456Z"
                }
            }
        },
        "CreateChannelRequest": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "x-order": "0",
                    "example": "Bored Owls Online Radio"
                },
                "link": {
                    "type": "string",
                    "x-order": "1",
                    "example": "https://example.com"
                },
                "authors": {
                    "type": "string",
                    "example": "The Owl"
                },
                "description": {
                    "type": "string",
                    "example": "Bored owls talk about whatever happens to be on their minds"
                }
            }
        },
        "CreateEpisodeRequest": {
            "type": "object",
            "properties": {
                "file_id": {
                    "type": "string",
                    "x-order": "0",
                    "example": "file_2K9BWVNuo3sG4yM322fbP3mB6ls"
                },
                "title": {
                    "type": "string",
                    "x-order": "1",
                    "example": "Bored Owls Online Radio"
                },
                "link": {
                    "type": "string",
                    "x-order": "2",
                    "example": "https://example.com"
                },
                "authors": {
                    "type": "string",
                    "x-order": "3",
                    "example": "The Owl"
                },
                "description": {
                    "type": "string",
                    "x-order": "4",
                    "example": "Bored owls talk about whatever happens to be on their minds"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "IdResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authentication": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Simple EpisodesPage Server REST API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
