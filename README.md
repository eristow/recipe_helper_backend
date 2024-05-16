### Recipe Helper Backend

This app is a recipe helper that allows users to add, view, and delete recipes. It is a Go application that uses the `net/http` package to create a web server. The app currently uses an in-memory "database" to store recipes.

## TODO:

- [ ] Convert to REST API

  - [ ] Add logging
  - [ ] Create DELETE endpoint for deleting recipes
  - [ ] Create PUT endpoint for updating recipes
  - [ ] Use query strings for filtering by name instead of path parameters
  - [ ] Tests
    - [ ] Write tests for rest.go
    - [ ] Write tests for database.go
    - [ ] Write tests for recipes.go
  - [ ] Create swagger docs

- [ ] Switch from storing recipes in in-memory to storing them in a database.
- [ ] Extract HTML into React front-end
- [ ] Spruce up the page templates by making them valid HTML and adding some CSS rules.
