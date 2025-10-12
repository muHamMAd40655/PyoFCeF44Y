// 代码生成时间: 2025-10-13 03:16:18
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/meta"
)

// App is the Buffalo application instance.
var App *buffalo.App

// init is the application initializer function.
func init() {
    App = buffalo.New(buffalo.Options{
        PrettyErrors: true,
    })

    // Automatically generate documentation for all routes
    App.GET("/docs", meta.Generator{"api"}, nil)

    // Add more routes here...
}

// main is the main entry point for the Buffalo application.
func main() {
    if err := App.Start(); err != nil {
        // Handle error properly and exit the application
        panic(err)
    }
}
