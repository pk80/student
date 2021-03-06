// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A sample API to demonstrate how to write OpenAPI Specification.",
    "title": "Sample API",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/sample_api",
  "paths": {
    "/persons": {
      "get": {
        "description": "Returns a list containing persons with pagging support.",
        "summary": "Gets some persons",
        "parameters": [
          {
            "type": "integer",
            "description": "Number of persons returned",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Page Number",
            "name": "pageNumber",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of Person",
            "schema": {
              "$ref": "#/definitions/Persons"
            }
          }
        }
      },
      "post": {
        "description": "Adds a new person to the person list.",
        "summary": "Creates a person",
        "parameters": [
          {
            "description": "The person to create",
            "name": "person",
            "in": "body",
            "schema": {
              "$ref": "./person.yml#/Person"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Person succesfully created."
          },
          "400": {
            "description": "Personc couldn't have been created."
          }
        }
      }
    },
    "/persons/{username}": {
      "get": {
        "description": "Returns a single person for its username.",
        "summary": "Gets a person",
        "parameters": [
          {
            "type": "string",
            "description": "The person's username.",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A person.",
            "schema": {
              "$ref": "./person.yml#/Person"
            }
          },
          "404": {
            "description": "The person doesn't exist."
          }
        }
      }
    }
  },
  "definitions": {
    "Persons": {
      "type": "array",
      "items": {
        "$ref": "./person.yml#/Person"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A sample API to demonstrate how to write OpenAPI Specification.",
    "title": "Sample API",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/sample_api",
  "paths": {
    "/persons": {
      "get": {
        "description": "Returns a list containing persons with pagging support.",
        "summary": "Gets some persons",
        "parameters": [
          {
            "type": "integer",
            "description": "Number of persons returned",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Page Number",
            "name": "pageNumber",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of Person",
            "schema": {
              "$ref": "#/definitions/Persons"
            }
          }
        }
      },
      "post": {
        "description": "Adds a new person to the person list.",
        "summary": "Creates a person",
        "parameters": [
          {
            "description": "The person to create",
            "name": "person",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/person"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Person succesfully created."
          },
          "400": {
            "description": "Personc couldn't have been created."
          }
        }
      }
    },
    "/persons/{username}": {
      "get": {
        "description": "Returns a single person for its username.",
        "summary": "Gets a person",
        "parameters": [
          {
            "type": "string",
            "description": "The person's username.",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A person.",
            "schema": {
              "$ref": "#/definitions/person"
            }
          },
          "404": {
            "description": "The person doesn't exist."
          }
        }
      }
    }
  },
  "definitions": {
    "Persons": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/person"
      }
    },
    "person": {
      "required": [
        "username"
      ],
      "properties": {
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  }
}`))
}
