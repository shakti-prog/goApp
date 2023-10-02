package main

import (
	// "fmt"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"io"
	"log"
	"net/http"
)

type hello struct {
	app.Compo
}

type login struct {
	app.Compo
}

type loginSuccessfull struct {
	app.Compo
}

type loginUnsuccessfull struct {
	app.Compo
}

type MyComp struct {
	app.Compo
	text string
}

type MyButton struct {
	app.Compo
}

type loginForm struct {
	app.Compo
}

var (
	globalEmail    string
	globalPassword string
)

func (c *loginForm) hanldeEmailChange(ctx app.Context, e app.Event) {
	globalEmail = ctx.JSSrc().Get("value").String()

}

func (c *loginForm) hanldePasswordChange(ctx app.Context, e app.Event) {
	globalPassword = ctx.JSSrc().Get("value").String()

}

func (c *loginForm) handleSignInClick(ctx app.Context, e app.Event) {
	if globalEmail == "" || globalPassword == "" {
		fmt.Println("Email and password are required")
		return
	}
	data := map[string]interface{}{
		"email":    globalEmail,
		"password": globalPassword,
	}
	jsonData, err1 := json.Marshal(data)
	if err1 != nil {
		fmt.Println("Error")
		fmt.Println(err1)
		return
	}
	response, err := http.Post("http://127.0.0.1:9000/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	if response.StatusCode == 200 {
		ctx.Navigate("/successfullLogin")
		fmt.Println("Sign In successfull")
	} else {
		ctx.Navigate("/unsuccessfullLogin")
		responseBody, err := io.ReadAll(response.Body)
		fmt.Println(string(responseBody))
		fmt.Println(err)
		fmt.Println(response)
	}

}

func (c *hello) handleClick(ctx app.Context, e app.Event) {
	ctx.Navigate("/signIn")

}

func (c *loginForm) Render() app.UI {
	return app.Div().Body(
		app.Link().Rel("stylesheet").Href("https://cdn.jsdelivr.net/npm/tailwindcss@2.2.15/dist/tailwind.min.css"),
		app.Div().Body(
			app.Div().Body(
				app.Div().Body(
					app.Label().Text("Your Email").For("email").Class("block mb-2 text-sm font-medium text-white"),
					app.Input().Type("email").Class("bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500").OnChange(c.hanldeEmailChange),
				),
				app.Div().Body(
					app.Label().Text("Password").For("email").Class("block mb-2 text-sm font-medium text-white"),
					app.Input().Type("password").Class("bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500").OnChange(c.hanldePasswordChange),
				),
				app.Div().Body(
					app.Button().Text("Sign In").Class("w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800").OnClick(c.handleSignInClick),
				),
			).Class("p-6 space-y-4 md:space-y-6 sm:p-8"),
		).Class("w-full bg-gray-900 rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700"),
	).Class("flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0")
}

//Main App

func (h *hello) Render() app.UI {
	return app.Div().
		Body(
			app.Link().Rel("stylesheet").Href("https://cdn.jsdelivr.net/npm/tailwindcss@2.2.15/dist/tailwind.min.css"),
			app.Div().Body(
				app.H1().Text("Welcome to go-app tutorial").Class("text-4xl font-bold text-center"),
				app.Button().Text("Sign In").Class("mt-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600").OnClick(h.handleClick),
			),
		).Class("flex h-screen justify-center items-center bg-gray-100")
}

func (h *loginSuccessfull) Render() app.UI {
	return app.Div().
		Body(
			app.Link().Rel("stylesheet").Href("https://cdn.jsdelivr.net/npm/tailwindcss@2.2.15/dist/tailwind.min.css"),
			app.Div().Body(
				app.H1().Text("Signed in successfully ").Class("text-4xl font-bold text-center"),
			),
		).Class("flex h-screen justify-center items-center bg-gray-100")
}

func (h *loginUnsuccessfull) Render() app.UI {
	return app.Div().
		Body(
			app.Link().Rel("stylesheet").Href("https://cdn.jsdelivr.net/npm/tailwindcss@2.2.15/dist/tailwind.min.css"),
			app.Div().Body(
				app.H1().Text("Sign In failed").Class("text-4xl font-bold text-center"),
			),
		).Class("flex h-screen justify-center items-center bg-gray-100")
}

func main() {
	app.Route("/", &hello{})
	app.Route("/signIn", &loginForm{})
	app.Route("/successfullLogin", &loginSuccessfull{})
	app.Route("/unsuccessfullLogin", &loginUnsuccessfull{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
