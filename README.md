# Nil SQL

Add Serialization Support to Null SQL Types.

## Usage

You can use the types in this package as a direct replacement for the `sql.NullX` types.

```go
package awesome

import "github.com/go-nm/nilsql"

type Thing struct {
  Description   nilsql.String
  MyIntNum      nilsql.Int64
  IsActive      nilsql.Bool
  MyFloatingNum nilsql.Float64
}
```
