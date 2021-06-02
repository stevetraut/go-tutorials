---
title: Develop a RESTful API with Go and Gin
description: Using the Gin web framework to develop a RESTful API.
author: straut
tags: Go
date_published: 
---

# Develop a RESTful API with Go and Gin

This tutorial introduces the basics of writing a RESTful web service API with Go and the [Gin Web Framework](https://gin-gonic.com/docs/).

You'll get the most out of this tutorial if you have a basic familiarity with Go and its tooling. If this is your first exposure to Go, please see [Tutorial: Get started with Go](https://golang.org/doc/tutorial/getting-started) for a quick introduction.

Written in Go, the Gin framework simplifies many coding tasks associated with building web applications, including web services.

Through this tutorial you will be building a server with two REST endpoints. Your example project is a repository of vintage jazz records. You'll be learning Go best practices.

In this tutorial, you will:

1. Create a project for your code.
2. List API endpoints.
3. Design and seed the data.
4. Write a handler to return all items.
5. Write a handler to add a new item.
6. Write a handle to return a specific item.

## Prerequisites

*   **A Google account.** You'll need an account with Google. A Gmail account will do, for example. You can [create a free account](https://accounts.google.com/signup/v2/webcreateaccount?flowName=GlifWebSignIn&flowEntry=SignUp).

    This tutorial uses the Google Cloud Shell Editor, an IDE with all of the basic environment dependencies you'll need. 

## Create a project for your code

To begin, create a project for the code you'll write.

1. Ensure that the **cloudshell_open** folder is selected.

1. Click the <walkthrough-editor-spotlight spotlightId="menu-file">File Menu</walkthrough-editor-spotlight>, then click **New Folder**.

1. In the **New Folder** dialog, enter `web-service-gin` for the folder name, then click **OK**.

1. Click the <walkthrough-editor-spotlight spotlightId="menu-file">File Menu</walkthrough-editor-spotlight>, then click **Open Workspace**.

1. In the **Open Workspace** dialog, select the cloudshell_open/web-service-gin folder you just created, then click **Open**.

1. If Cloud Shell isn't open, click the <walkthrough-cloud-shell-icon></walkthrough-cloud-shell-icon> button to open it.

1. In Cloud Shell, change to the web-service-gin directory you created.

    ```bash
    cd ~/cloudshell_open/web-service-gin
    ```

1. Create a module in which you can manage dependencies.

    Run the `go mod init` command, giving it the path of the module your code will be in.

    ```bash
    go mod init example.com/web-service-gin
    ```

    This command creates a go.mod file in which dependencies you might add will be listed for tracking. For more, be sure to see [Managing dependencies](https://golang.org/doc/modules/managing-dependencies).

Next, you'll capture the design for the API you're writing.

## List API endpoints

You'll be building the start of an API that provides access to a clearinghouse of vintage recordings on vinyl. So you'll need to provide endpoints through which a client can get and add albums on behalf of users.

When developing an API, you typically begin with planning the endpoints your API will have. Your API will be more useful to client developers if the endpoints are intuitive.

Here are the endpoints you'll support in this tutorial.

`/albums`

- `GET` – Get a list of all albums, returned as JSON.
- `POST` – Add a new album from request data sent as JSON.

`/albums/:id`

- `GET` – Get an album by its ID, returning the album data as JSON.

Next, you'll design data structures for handling data.

## Design and seed the data

To keep things simple for the tutorial, you'll store the data accessed by the API in memory. A more typical API would have a database behind it, with API code interacting with database access code.

Note that storing album data in memory means that every time you stop and start the server, the set of albums stored in memory will reset.

1. Click the <walkthrough-editor-spotlight spotlightId="menu-file">File Menu</walkthrough-editor-spotlight>, then click **New File**.
1. In the **New File** dialog, name the file **web-service-gin.go**, then click **OK**.
1. Into web-service-gin.go, at the top of the file, paste the package declaration below.

    ```golang
    package main
    ```

1. Beneath the package declaration, paste the following declaration of an `album` struct. You'll use this to store album data in memory.

    Struct tags such as ``json:"artist"`` specify what a field's name should be when the struct's contents are serialized into JSON. Without them, the generated JSON would use the struct's capitalized field names – a style not typical in JSON.

    ```golang
    // album represents data about a record album.
    type album struct {
        ID     string  `json:"id"`
        Title  string  `json:"title"`
        Artist string  `json:"artist"`
        Price  float32 `json:"price"`
    }
    ```

1. Beneath the struct declaration you just added, paste the following slice of pointers to `album` structs containing data you'll use to start.

    ```golang
    // albums to seed record album data.
    var albums = []*album{
        {ID: "48590", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "48583", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "48581", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
    }

    ```

Next, you'll write code to implement your first endpoint.

## Write a handler to return all items

When the client makes a request at `GET /albums`, you want to return all of the albums as JSON. 

To do this, you'll write the following:

*   Logic to prepare a response
*   Code to map the request path to your logic

Note that this is the reverse of how they'll be executed at run time, but you're adding dependencies first, then the code that depends on them.

1. Beneath the code you added in the preceding section, paste the following code to get the album list.

    This `getAllAlbums` function creates JSON from the slice of `album` structs, writing the JSON into the response.

    ```
    // getAllAlbums returns the list of all albums as JSON.
    func getAllAlbums(c *gin.Context) {
        c.JSON(http.StatusOK, albums)
    }
    ```

    You'll have errors until you import required packages. You can fix those by hovering the mouse pointer over the errors, then using the Quick Fix feature to add imports, or you can add the imports in a later step below.

    In this code, you:

    *   Write a `getAllAlbums` function that takes a <code>[gin.Context](https://pkg.go.dev/github.com/gin-gonic/gin#Context)</code> parameter. <code>Context</code> is the most important part of Gin, carrying request details, validating and serializing JSON, and more. Note that you could have given this function any name – neither Gin nor Go require a particular function name format.
    *   Call the <code>[Context.JSON function](https://pkg.go.dev/github.com/gin-gonic/gin#Context.JSON)</code> to serialize the struct into JSON. The function's first argument is the HTTP status code you want to send to the client. Here, you're using a passing a <code>[StatusOK constant](https://pkg.go.dev/net/http#StatusOK)</code> from the <code>net/http</code> package to indicate <code>200 OK</code>.

    > Note that you can replace <code>Context.JSON</code> with a call to <code>[Context.IndentedJSON function](https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON)</code> to display JSON that's a bit easier to read while debugging.

1. Near the top of the web-service-gin.go file, just beneath the `albums` slice declaration, paste the code below to map the handler function to an endpoint path.

    This sets up an association in which requests to the `/albums` endpoint path are handled by the `getAllAlbums` function you wrote.

    ```
    func main() {
        router := gin.Default()
        router.GET("/albums", getAllAlbums)

        router.Run(":8080")
    }
    ```

    In this code, you:

    *   Initialize a Gin router using the <code>[Default function](https://pkg.go.dev/github.com/gin-gonic/gin#Default)</code>.
    *   Use the [<code>GET function</code>](https://pkg.go.dev/github.com/gin-gonic/gin#RouterGroup.GET) to associate the <code>GET</code> HTTP method and <code>/albums</code> path with a handler function.
    *   Use the <code>[Run function](https://pkg.go.dev/github.com/gin-gonic/gin#Engine.Run)</code> to attach the router to an <code>http.Server</code> and start the server.

1. If you haven't added required imports yet, beneath the package declaration, paste the following `import` statement:

    ```golang
    import (
        "net/http"

        "github.com/gin-gonic/gin"
    )
    ```

### Run the code

1. Begin tracking the Gin module as a dependency.

    Use the <code>[go mod tidy command](https://golang.org/ref/mod#go-mod-tidy)</code> to add the github.com/gin-gonic/gin module as a dependency for your module.

    ```bash
    go mod tidy
    ```
1. Open the <walkthrough-editor-spotlight spotlightId="debug-configuration">Debug view</walkthrough-editor-spotlight>, then click the **Start Debugging** button.

    The first time you start debugging, you'll be prompted to configure launch setting. Just close the launch.json file and click **Start Debugging** again. 

    Once the code is running, you have a running HTTP server to which you can send requests.

1. In Cloud Shell, use the following command to test the web service endpoint.

    ```bash
    curl -H "Content-Type: application/json" http://localhost:8080/albums
    ```

    The command should display the data you seeded the service with. The following output is indented for easier reading.

    ```
    [
            {
                    "id": "48590",
                    "title": "Blue Train",
                    "artist": "John Coltrane",
                    "price": 199.99
            },
            {
                    "id": "48583",
                    "title": "Jeru",
                    "artist": "Gerry Mulligan",
                    "price": 17.99
            },
            {
                    "id": "48581",
                    "title": "Sarah Vaughan and Clifford Brown",
                    "artist": "Sarah Vaughan",
                    "price": 39.99
            }
    ]

    ```

In the next section, you'll add code to handle a `POST` request to add an item.

## Write a handler to add a new item

When the client makes a `POST ` request at `/albums`, you want to add the album described in the request body to the data you've already got.

To do this, you'll write the following:

*   Logic to add the new album to the existing list, then return the updated list as a response.
*   A bit of code to route the `POST` request to your logic.

1. Beneath the `getAllAlbums` function you added in the preceding section, paste the following code to add an album.

    The `addAlbum` function will add the data you receive to the list of albums.

    ```golang
    // addAlbum adds an album from JSON received in the request body.
    func addAlbum(c *gin.Context) {
        var a album

        // Call ShouldBindJSON to confirm that the
        // request body JSON is valid for the struct.
        if err := c.ShouldBindJSON(&a); err != nil {
            c.JSON(http.StatusInternalServerError,
                gin.H{"error": err.Error()})
            return
        }

        // Add the new album to the slice.
        albums = append(albums, &a)
        // Return the slice as JSON.
        c.JSON(http.StatusOK, albums)
    }
    ```

    In this code, you:

    *   Use the <code>[Context.ShouldBindJSON function](https://pkg.go.dev/github.com/gin-gonic/gin#Context.ShouldBindJSON)</code> to retrieve the request body and validate it. If it's invalid, your code returns an HTTP 500 error and returns.
    *   Append the <code>album</code> struct initialized from the JSON to the <code>albums</code> slice.
    *   Use the <code>[Context.JSON function](https://pkg.go.dev/github.com/gin-gonic/gin#Context.JSON)</code> to serialize the <code>albums</code> slice to JSON for a response.

1. Replace your <code>main</code> function with the following code.

    ```
    func main() {
        router := gin.Default()
        router.GET("/albums", getAllAlbums)
        router.POST("/albums", addAlbum)

        router.Run(":8080")
    }
    ```

    In this code, you:

    *   Associate the `POST` method at the `/albums` path with the `addAlbum` function.

### Run the code

1. If the server is still running from the last section, stop it.
1. Open the <walkthrough-editor-spotlight spotlightId="debug-configuration">Debug view</walkthrough-editor-spotlight>, then click the **Start Debugging** button.
1. In Cloud Shell, use the following command to test the web service endpoint.

    ```bash
    curl -H "Content-Type: application/json" -d '{"id": "48590","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}' http://localhost:8080/albums
    ```

    The command should display the data you already had, but with the data you're adding appended. Here, the example output is indented for easier reading.

    ```
    [
            {
                    "id": "48590",
                    "title": "Blue Train",
                    "artist": "John Coltrane",
                    "price": 199.99
            },
            {
                    "id": "48583",
                    "title": "Jeru",
                    "artist": "Gerry Mulligan",
                    "price": 17.99
            },
            {
                    "id": "48581",
                    "title": "Sarah Vaughan and Clifford Brown",
                    "artist": "Sarah Vaughan",
                    "price": 39.99
            },
            {
                    "id": "48590",
                    "title": "The Modern Sound of Betty Carter",
                    "artist": "Betty Carter",
                    "price": 49.99
            }
    ]

    ```

In the next section, you'll add code to handle a `GET` for a specific item.

## Write a handler to return a specific item 

When the client makes a request to `GET /albums/:id`, you want to return the album whose ID matches the `id` path parameter.

To do this, you will:

*   Add logic to retrieve the requested album.
*   Map the path to the logic.
1. Beneath the `addAlbum` function you added in the preceding section, paste the following code to retrieve a specific album.

    This `getAlbumByID` function will extract the ID in the request path, then locate an album that matches.

    ```
    // getAlbumByID locates the album whose ID value matches the id
    // parameter sent by the client, then returns that album as a response.
    func getAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Loop through the list of albums, looking for
        // an album whose ID value matches the parameter.
        for _, a := range albums {
            if a.ID == id {
                c.JSON(http.StatusOK, a)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
    }
    ```

    In this code, you:

    *   Use the `Context.Param` function to retrieve the `id` path parameter from the URL.
    *   Loop through the `album` structs in the slice, looking for one whose `ID` field value matches the `id` parameter value. If it's found, you serialize that `album` struct to JSON and return it as a response with a `200 OK` HTTP code.
    *   Return an HTTP 404 error with [`http.StatusNotFound`](https://pkg.go.dev/net/http#StatusNotFound) if the album isn't found.

1. Finally, replace your <code>main</code> function with the following code.

    ```
    func main() {
        router := gin.Default()
        router.GET("/albums", getAllAlbums)
        router.GET("/albums/:id", getAlbumByID)
        router.POST("/albums", addAlbum)

        router.Run(":8080")
    }
    ```

    In this code, you:

    *   Associate the `/albums/:id` path with the `getAlbumByID` function.

### Run the code

1. If the server is still running from the last section, stop it.
1. Open the <walkthrough-editor-spotlight spotlightId="debug-configuration">Debug view</walkthrough-editor-spotlight>, then click the **Start Debugging** button.
1. In Cloud Shell, use the following command to test the web service endpoint.

    ```bash
    curl localhost:8080/albums/48583
    ```

    The command should display JSON for the album whose ID you used. If the album wasn't found, you'll get JSON with an error message.

    ```
    {
            "id": "48583",
            "title": "Jeru",
            "artist": "Gerry Mulligan",
            "price": 17.99
    }

    ```

## Conclusion

Congratulations! You've just used Go and Gin to write a simple RESTful web service.

Suggested next topics:

*   If you're new to Go, you'll find useful best practices described in [Effective Go](https://golang.org/doc/effective_go) and [How to write Go code](https://golang.org/doc/code).
*   The [Go Tour](https://tour.golang.org/welcome/1) is a great step-by-step introduction to Go fundamentals.
*   For more about Gin, see the [Gin Web Framework package documentation](https://pkg.go.dev/github.com/gin-gonic/gin) or the [Gin Web Framework docs](https://gin-gonic.com/docs/).

<walkthrough-inline-feedback></walkthrough-inline-feedback>