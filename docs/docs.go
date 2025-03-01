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
        "/items": {
            "get": {
                "description": "Show all items",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "List of items",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create new items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "New Item",
                "parameters": [
                    {
                        "description": "List of items",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ItemValues"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/items/{id}": {
            "get": {
                "description": "Search item by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Item by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "delete": {
                "description": "Delete item by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "security": [
                    {
                        "cookieAuth": []
                    }
                ],
                "description": "List of orders,need to be authnorized (sign-in to get cookie)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "List of your orders",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "cookieAuth": []
                    }
                ],
                "description": "Create order, need to be authnorized (sign-in to get cookie)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create order",
                "parameters": [
                    {
                        "description": "Order details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.InputOrder"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "security": [
                    {
                        "cookieAuth": []
                    }
                ],
                "description": "Order by id,need to be authnorized (sign-in to get cookie)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Show specific order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of your order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "description": "SignIn",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User sign-in",
                "parameters": [
                    {
                        "description": "Sign-in details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "Creation new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User sign-up",
                "parameters": [
                    {
                        "description": "Sign-up details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.OrderItem": {
            "type": "object",
            "properties": {
                "itemId": {
                    "type": "integer",
                    "example": 4
                },
                "quantity": {
                    "type": "integer",
                    "example": 7
                }
            }
        },
        "service.InputItem": {
            "type": "object",
            "required": [
                "desc",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "desc": {
                    "type": "string",
                    "example": "Compact mouse"
                },
                "name": {
                    "type": "string",
                    "example": "Wireless mouse"
                },
                "price": {
                    "type": "number",
                    "example": 19.99
                },
                "stock": {
                    "type": "integer",
                    "example": 150
                }
            }
        },
        "service.InputOrder": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.OrderItem"
                    }
                }
            }
        },
        "service.ItemValues": {
            "type": "object",
            "required": [
                "items"
            ],
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.InputItem"
                    }
                }
            }
        },
        "service.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "testmail@gmail.com"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password123"
                }
            }
        },
        "service.SignUpInput": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "testmail@gmail.com"
                },
                "firstName": {
                    "type": "string",
                    "example": "Alex"
                },
                "lastName": {
                    "type": "string",
                    "example": "Johnson"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password123"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "REST API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
