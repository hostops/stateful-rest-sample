package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "context"
    "os"
    "github.com/jackc/pgx/v4/pgxpool"
)
// User ...
type User struct {
    Name string `json:"Name"`
    Lastname string `json:"Lastname"`
}
// Users ...
func handleUsers(w http.ResponseWriter, r *http.Request) {
    var users []User = []User{}
    switch r.Method {
      case "GET":
        dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	      if err != nil {
          w.WriteHeader(500);
          fmt.Fprintf(w, "{\"error\":\"Unable to connect to database: %v\"}\n", err)
          return
	      }
	      defer dbpool.Close()
        rows, err := dbpool.Query(context.Background(), "select name, lastname from users")
          if err != nil {
            w.WriteHeader(500);
            fmt.Fprintf(w, "{\"error\":\"Cannot execute statement: %v\"}\n", err)
            return
          }
        for rows.Next() {
          var name string
          var lastname string
          err := rows.Scan(&name, &lastname)
          if err != nil {
            w.WriteHeader(500);
            fmt.Fprintf(w, "{\"error\":\"Error while reading users: %v\"}\n", err)
            return
          }
          users = append(users, User{Name: name, Lastname: lastname})
	      }
         w.WriteHeader(200);
        json.NewEncoder(w).Encode(users)
        break
      case "POST":
        var user User

        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            w.WriteHeader(400);
            fmt.Fprintf(w, "{\"error\":\"Error while reading users: %v\"}\n", err)
            return
        }
        dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	      if err != nil {
          w.WriteHeader(500);
          fmt.Fprintf(w, "{\"error\":\"Unable to connect to database: %v\"}\n", err)
          return
	      }
	      defer dbpool.Close()
        _, err = dbpool.Exec(context.Background(), "INSERT INTO users(name, lastname) VALUES($1,$2)", user.Name, user.Lastname)
          if err != nil {
            w.WriteHeader(500);
            fmt.Fprintf(w, "{\"error\":\"Cannot execute statement: %v\"}\n", err)
            return
          }
          w.WriteHeader(200);
          fmt.Fprintf(w, "{\"status\":\"OK\"}\n")
        break
      default:
        fmt.Fprintf(w, "{\"error\":\"Sorry, only GET and POST methods are supported.\"}\n")
    }
}
func main() {
  fmt.Printf("Stateful rest sample started listening on port %s\n", os.Getenv("PORT"))
  http.HandleFunc("/users", handleUsers)
  log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
