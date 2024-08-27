### Recipe Helper Backend

This app is a recipe helper that allows users to add, view, and delete recipes. It is a Go application that uses the `net/http` package to create a web server. The app currently uses an in-memory "database" to store recipes.

## TODO:

- [ ] Tests
  - [ ] Write tests for rest.go
  - [ ] Write tests for database.go
  - [ ] Write tests for recipes.go

- [ ] Create swagger docs

- [ ] Switch from storing recipes in in-memory to storing them in a database.


## Done:
- [x] Delete by ID instead of by name
- [x] Get by ID instead of by name
- [x] ID should be assigned and not 0000...
- [x] Return all recipes sorted by name
- [x] Create DELETE endpoint for deleting recipes
- [x] Add logging
- [x] Extract HTML into React front-end
  - [x] Create PUT endpoint for updating recipes
