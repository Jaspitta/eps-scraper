package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
}

// The struct with the .Count property we refere to in the index.html
type Count struct {
	Count int
}

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

// just a struct for the page content
type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		// pre seeded contacts
		Contacts: []Contact{
			newContact("Giovanni", "gio.rossi@mail.com"),
			newContact("Mario", "mar.rossi@mail.com"),
		},
	}
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	data := newData()
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		data.Contacts = append(data.Contacts, newContact(name, email))
		return c.Render(200, "display", data)
	})

	e.Logger.Fatal(e.Start(":42069"))

}

// func getNextEPS(symbol string, client *http.Client) string {
// 	// preparing request
// 	pathElems := []string{"https://www.earningswhispers.com/api/getstocksdata/", symbol}
// 	path := strings.Join(pathElems, "")
// 	req, _ := http.NewRequest("GET", path, nil)
// 	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
// 	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
// 	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
// 	req.Header.Add("X-Requested-With", "XMLHttpRequest")
// 	req.Header.Add("Referer", "https://www.earningswhispers.com/stocks/STEM")
//
// 	// performing http req
// 	resp, _ := client.Do(req)
//
// 	defer resp.Body.Close()
//
// 	// reading body as bytes
// 	body, _ := io.ReadAll(resp.Body)
//
// 	return formatResp(body, symbol)
// }
//
// func formatResp(body []byte, symbol string) (result string) {
// 	var j interface{}
// 	json.Unmarshal(body, &j)
//
// 	// converting json to map
// 	m := j.(map[string]interface{})
//
// 	// printing needed value from map
// 	// only up to char 10 to remove the timestamp
// 	mess := []string{symbol, "\t --> \t", m["nextEPSDate"].(string)[0:10], "\n"}
// 	result = strings.Join(mess, "")
// 	return
// }
