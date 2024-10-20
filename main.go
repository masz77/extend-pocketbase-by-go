package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/labstack/echo/v5"
)

// type User struct {
// 	Id       string
// 	Email    string
// 	// id       string
// 	// username string
// 	// email    string
// 	// name     string
// 	// created  string
// 	// updated  string
// }

func main() {
	app := pocketbase.New()
	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// e.Router.Use(apis.CORS())
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		e.Router.GET("/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "hello world")
		})

		e.Router.GET("/api/custom/try-db", func(c echo.Context) error {
			_users := []struct {
				Id    string `db:"id" json:"_idx"`
				Email string `db:"email" json:"email"`
			}{}
			// _dao := app.Dao()

			// collection, err2 := _dao.FindCollectionByNameOrId("users")
			// if err2 != nil {
			// 	return nil
			// }
			// print(collection)
			// return c.JSON(http.StatusOK, collection)

			// _users := []User{}
			// err1 := _dao.DB().NewQuery("SELECT id,username,email, name, created, updated FROM users").All(&_users)
			// _users, err1 := _dao.FindRecordsByIds("users", []string{"5gu1o2vtq01y6jz", "acd1f76iu1e5512"})
			err1 := app.Dao().DB().
							Select("id", "email").
							From("users").
							Limit(100).
							OrderBy("created ASC").
							All(&_users)
			if err1 != nil {
				return nil
			}
			return c.JSON(http.StatusOK, _users)
		})

		e.Router.GET("/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")

			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		} /* optional middlewares */)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
