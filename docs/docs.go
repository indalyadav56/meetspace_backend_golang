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
        "/v1/auth/forgot-password": {
            "post": {
                "description": "Forgot password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "forgot-password",
                "parameters": [
                    {
                        "description": "forgot password request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "description": "Login user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "login-user",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login user successfully"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/logout": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User logout User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "user-logout",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LogoutRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/refresh-token": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "refresh jwt token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "refresh token",
                "parameters": [
                    {
                        "description": "refresh token request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/register": {
            "post": {
                "description": "Register User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "register-user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Register user successfully"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/send-email": {
            "post": {
                "description": "Send email to user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "send-email",
                "parameters": [
                    {
                        "description": "send email request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.SendEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/auth/verify-email": {
            "post": {
                "description": "Verify email otp.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "verify-email",
                "parameters": [
                    {
                        "description": "verify email request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.VerifyEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/chat/groups": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Add Chat group",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Group"
                ],
                "summary": "add-chat-group",
                "parameters": [
                    {
                        "description": "Add chat group details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AddChatGroup"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/chat/messages/{chat_room_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get chat messages by room id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Message"
                ],
                "summary": "get chat messages by room id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chat Room ID",
                        "name": "chat_room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/chat/room/contact": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "UserLogin User account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "UserLogin User account",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/chat/rooms": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "GetChatRooms",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "GetChatRooms",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "CreateChatRoom",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "CreateChatRoom",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "DeleteChatRoom",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat-Room"
                ],
                "summary": "DeleteChatRoom",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/user/check-email": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User create",
                "parameters": [
                    {
                        "description": "User create details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/user/profile": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user profile",
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get all users",
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user-update",
                "parameters": [
                    {
                        "description": "User create details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User create",
                "parameters": [
                    {
                        "description": "User create details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "types.AddChatGroup": {
            "type": "object",
            "required": [
                "title",
                "user_ids"
            ],
            "properties": {
                "title": {
                    "type": "string"
                },
                "user_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "types.CreateUserData": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "types.ForgotPasswordRequest": {
            "type": "object",
            "required": [
                "email",
                "new_password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "types.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "types.LogoutRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "types.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "types.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "types.SendEmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "types.VerifyEmailRequest": {
            "type": "object",
            "required": [
                "email",
                "otp"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                }
            }
        },
        "utils.Metadata": {
            "type": "object",
            "properties": {
                "api_version": {
                    "type": "string",
                    "default": "v1"
                },
                "server_info": {
                    "$ref": "#/definitions/utils.ServerInfo"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/utils.Metadata"
                },
                "request_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "time_stamp": {
                    "type": "string"
                }
            }
        },
        "utils.ServerInfo": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
