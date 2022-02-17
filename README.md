# fn-todo

This is a simple backend for a toy-todo app. It contains for main functions:

- get (`/get?item=<id>`) for getting indiviudal todo items
- set (`/set?item=<todo string>`) to add items to the the list
- complete (`/complete?item=<id>`) to complete an item
- list (`/list`) to list all of the items (with the optional param query `?includeCompleted=true` for completed items)

This app is separated into two repositories, the core logic is over at [jonfriesen/todo](https://github.com/jonfriesen/todo).

## Design
This repository contains the IO interface for our "buiness" logic. It establishes the endpoints (in this case functions) and exposes them, as well as does some primitivate validation. Business oriented validation should be done in the other module. To keep things simple, this repository has been scoped to just handle the input/output interface.

## Why have two repositories?
In Go, development of multiple modules within the same repository is rather difficult. With this the business logic ([jonfriesen/todo](https://github.com/jonfriesen/todo)) was written completely independently of the function interface. Once the logic was complete, validated, and ready to go, the functions "API" in this repository was designed. 