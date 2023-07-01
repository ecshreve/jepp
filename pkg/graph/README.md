# GraphQL

## Sample Queries

```{graphql}
{
  game(gameID: 3135) {
    id
    season {
      id
    }
    show
    airDate
    tapeDate
    cluesConnection(first: 20) {
      edges {
        node {
          id
          question
          answer
          category {
            name
          }
        }
      }
      pageInfo {
        startCursor
        endCursor
        hasNextPage
      }
    }
  }
}
```