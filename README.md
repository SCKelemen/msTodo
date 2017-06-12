# msTodo
msTodo is a microservice plugin for the framework which exposes methods for Todo Items and lists. 
msTodo follows the CRAP schema, and it's RESTful HTTP methods are paired below.

| CRAP Verb | HTTP Verb |
|-----------|-----------|
| Create    | POST      |
| Retrieve  | GET       |
| Alter     | PATCH     |
| Purge     | DELETE    |

# Routes

/l	lists
/t	todos
/u	users
/b	board 
/z	label	

# LISTS

| CRAP Verb | HTTP Verb | Route          | Example         | Description                                         |
|-----------|-----------|----------------|-----------------|-----------------------------------------------------|
| Create    | POST      | "/l"           | POST "/l"       | Creates a List                                      |
| Retrieve  | GET       | "/l", "/l/:id" | GET "/l/123"    | Retrieves List with ID, otherwise returns all lists |
| Alter     | PATCH     | "/l/:id"       | PATCH "/l/123"  | Alters List with ID                                 |
| Purge     | DELETE    | "/l/:id"       | DELETE "/l/123" | Purges List with ID                                 |

# TODOS

| CRAP Verb | HTTP Verb | Route    | Example         | Description                 |
|-----------|----------:|----------|-----------------|-----------------------------|
| Create    | POST      | "/t"     | POST "/t"       | Creates a Todo item         |
| Retrieve  | GET       | "/t/:id" | GET "/t/123"    | Retrieves Todo item with ID |
| Alter     | PATCH     | "/t/:id" | PATCH "/t/123"  | Alters Todo item with ID    |
| Purge     | DELETE    | "/t/:id" | DELETE "/t/123" | Purges Todo item with ID    |

# LABELS

| CRAP Verb | HTTP Verb | Route    | Example         | Description             |
|-----------|----------:|----------|-----------------|-------------------------|
| Create    | POST      | "/z"     | POST "/z"       | Creates a Label         |
| Retrieve  | GET       | "/z/:id" | GET "/z/123"    | Retrieves Label with ID |
| Alter     | PATCH     | "/z/:id" | PATCH "/z/123"  | Alters Label with ID    |
| Purge     | DELETE    | "/z/:id" | DELETE "/z/123" | Purges Label with ID    |


