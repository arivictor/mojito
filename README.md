# mojito

Light-weight customisable structured logger for your Go applications.

## 🚀 Installation

```shell
go get -u github.com/arivictor/mojito
```

## ⚡️ Quick Start
Mojito is ready to-go out of the box:

```go
logger := mojito.New()

m.Debug("hello")    // won't log
m.Error("world")    // will log
```

but comes with a lot of customizability! Here's an example:

```go
logger := mojito.New()
logger.Config.SetLevel(mojito.Error)
logger.Config.SetMessageKey("status")   // message -> status
logger.Config.SetDelimiter('|')         // \n -> |
```

# 👨‍💻 Contributing

Please open a pull request with listed changes. Any issues please feel free to contact me.

---

Released under the [GNU GENERAL PUBLIC LICENSE](LICENSE).