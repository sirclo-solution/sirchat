# Sirchat is Sirclo Chat APP

Sirchat is a Golang Framework or plugin to build Sirclo Chat Apps in a flash with the latest platform features. Sirchat can customize the order and appearance of information and guide users through your apps capabilities by initial your blocks, updating, composing, and stacking blocks.

Read [the documentation](https://pkg.go.dev/github.com/sirclo-solution/sirchat) to explore the basic and advanced concepts of Sirchat for Golang.

## Setup

```bash
go install github.com/sirclo-solution/sirchat
```

or

```bash
go get -u github.com/sirclo-solution/sirchat
```

## Initialization

Create an app by calling the constructor, which is a top-level export.

```go
app := apps.NewApps(apps.AppConfig{
    AppSecret: SECRET_KEY,
})

// Example Command
app.Command("/exampleOne", commandHandler func(c context.Context) (interface{}, error){
    /* Process logic here */
})

// start service
app.Start(apps.AppServerConfig{
    Port:    "8080",
})
```

## Example

See a more complete example of using the block kit [here](https://github.com/sirclo-solution/sirchat/tree/main/examples)