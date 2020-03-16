package main
import "github.com/satori/go.uuid"
import "encoding/json"
import "fmt"
import "database/sql"
import _ "github.com/lib/pq"
import "net/http"

var db *sql.DB

type Product struct {
	Id int
	Name string
	Description string
	Quantity int
	Composition string
	Price int
}

type Products struct {
	Products []Product
}

type User struct {
	Id int
	Name string
	Email string
	Password string
	Role int
	Status bool
}

type Users struct {
	Users []User
}

func main() {
	var err error
	db, err = sql.Open("postgres", "host=127.0.0.1 user=api password=12345678 dbname=api sslmode=disable")
	if err != nil {
		panic(err)
	}

	fmt.Println("# Starting server. To exit press CTRL+C")
	http.HandleFunc("/v1/products", handlerAllProducts)
	http.HandleFunc("/v1/products/add", handlerAddProduct)
	http.HandleFunc("/v1/products/delete", handlerDelProduct)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/v1/user/get_token", handlerGetToken)
	var address = "0.0.0.0:8000"
	fmt.Printf("server started at %s\n", address)
	error := http.ListenAndServe(address, nil)
	if error != nil {
		fmt.Println(err.Error())
	}

	db.Close ()
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
        var message = "Hello world! Go"
        w.Write([]byte(message))
}

func handlerAllProducts(w http.ResponseWriter, r *http.Request) {
	w_array := Products{}

	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", 405)
	} else {
		fmt.Println("# Querying all products")
		rows, err := db.Query("SELECT id,name,description,quantity,composition,price from products")
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			w_product := Product{}
			err := rows.Scan(&w_product.Id,&w_product.Name,&w_product.Description,&w_product.Quantity,&w_product.Composition,&w_product.Price)
			if err != nil {
				panic(err)
			}
		w_array.Products = append(w_array.Products, w_product)
		defer rows.Close()
		}
	json.NewEncoder(w).Encode(w_array)
	}
}

func handlerAddProduct(w http.ResponseWriter, r *http.Request) {

        if r.Method != "POST" {
                http.Error(w, "Method not Allowed", 405)
        } else {
          x_token := r.Header.Get("X-API-Token")
	  query := fmt.Sprintf("select u.role from sessions s left join users u on u.id = s.user_id where s.token = '%s' and ((s.added + interval '1h')> now()) and u.role = 1", x_token)
          row := db.QueryRow(query)
          var id int
          err := row.Scan(&id)
          if err != nil {
            http.Error(w, "Invalid  token", 403)
          } else {
                fmt.Println("# Querying add products")
		decoder := json.NewDecoder(r.Body)
		var g_product Product

                err := decoder.Decode(&g_product)
                if err != nil {
                        panic(err)
                }
		query := fmt.Sprintf("INSERT INTO products(name, description, quantity, composition, price) VALUES('%s', '%s', '%d', '%s', '%d') RETURNING id", g_product.Name, g_product.Description, g_product.Quantity, g_product.Composition, g_product.Price)
		fmt.Println ("# Inster QUERY: %s", query)
		rows, err := db.Query(query)
		if err != nil {
		  panic(err)
	        }
		for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
		  panic(err)
		}
	        defer rows.Close()
	        fmt.Fprintf (w, "{\"id\":%d}", id)
                }
	  }
	}
}

func handlerGetToken(w http.ResponseWriter, r *http.Request) {

        if r.Method != "POST" {
                http.Error(w, "Method not Allowed", 405)
        } else {
                fmt.Println("# Querying get Token")
                decoder := json.NewDecoder(r.Body)
                var g_user User
                err := decoder.Decode(&g_user)
                if err != nil {
                        panic(err)
                }
        	query := fmt.Sprintf("select id from users where email='%s' and password='%s' and status", g_user.Email,g_user.Password)
        	fmt.Println ("# Get Token: %s", query)
        	row := db.QueryRow(query)
		var id int
        	err = row.Scan(&id)
       		if err != nil {
       		 	http.Error(w, "User not Valid", 403) 
       		} else {
  		  query := fmt.Sprintf("select token from sessions where user_id =%d and ((added + interval '1h')> now())", id)
          	  row := db.QueryRow(query)
	  	  var token string
          	  err = row.Scan(&token)
          	  if err != nil {
	    	    token := uuid.Must(uuid.NewV4())
	    	    fmt.Println ("# Got Token: %s, id: %d", token, id)
            	    query := fmt.Sprintf("INSERT INTO sessions (user_id, token) VALUES (%d, '%s')", id, token)
	    	    _, err := db.Exec(query)
            	    if err != nil {
                      http.Error(w, "Internal Error", 500)
	            } else {
        	      fmt.Fprintf (w, "{\"token\":%s}", token)
	            }
                  } else {
                  fmt.Fprintf (w, "{\"token\":%s}", token)
	          }
 	        }
	}
}

func handlerDelProduct(w http.ResponseWriter, r *http.Request) {

        if r.Method != "DELETE" {
                http.Error(w, "Method not Allowed", 405)
        } else {
	  x_token := r.Header.Get("X-API-Token")
	  query := fmt.Sprintf("select u.role from sessions s left join users u on u.id = s.user_id where s.token = '%s' and ((s.added + interval '1h')> now()) and u.role = 1", x_token)
    	  row := db.QueryRow(query)
	  var id int
	  err := row.Scan(&id)
          if err != nil {
            http.Error(w, "Invalid  token", 403)
	  } else {
                decoder := json.NewDecoder(r.Body)
                var g_product Product
                err := decoder.Decode(&g_product)
                if err != nil {
                        panic(err)
                }
		query := fmt.Sprintf("DELETE FROM products where id = %d", g_product.Id)
		_, err = db.Exec(query)
		if err != nil {
		    http.Error(w, "Internal Error", 500)
	    	} else {
		  fmt.Fprintf (w, "Product del success, id: %d}", g_product.Id)
		}
	  }
	}
}
