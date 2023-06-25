// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "shreve"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/ecshreve/jepp/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/category": {
            "get": {
                "description": "Returns a list of categories.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Returns a list of categories.",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "If exists, returns up to ` + "`" + `limit` + "`" + ` random records.",
                        "name": "random",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "If exists, returns the record with the given id.",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Paging offset",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit the number of records returned",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    }
                }
            }
        },
        "/clue": {
            "get": {
                "description": "Returns a list of clues",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Returns a list of clues",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Random Clue",
                        "name": "random",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Clue ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "game",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Clue"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/game": {
            "get": {
                "description": "Returns a list of games",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Returns a list of games",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "name": "random",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Game"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Category": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer",
                    "example": 765
                },
                "name": {
                    "type": "string",
                    "example": "State Capitals"
                }
            }
        },
        "models.Clue": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "This is the answer."
                },
                "categoryId": {
                    "type": "integer",
                    "example": 804092001
                },
                "clueId": {
                    "type": "integer",
                    "example": 804002032
                },
                "gameId": {
                    "type": "integer",
                    "example": 8040
                },
                "question": {
                    "type": "string",
                    "example": "This is the question."
                }
            }
        },
        "models.Game": {
            "type": "object",
            "properties": {
                "gameDate": {
                    "type": "string",
                    "example": "2019-01-01"
                },
                "gameId": {
                    "type": "integer",
                    "example": 8040
                },
                "seasonId": {
                    "type": "integer",
                    "example": 38
                },
                "showNum": {
                    "type": "integer",
                    "example": 4532
                },
                "tapedDate": {
                    "type": "string",
                    "example": "2019-01-01"
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Jepp API Documentation",
	Description:      "This is a simple api to access jeopardy data.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
