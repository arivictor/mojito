![mojito](image.png)

[![Build](https://github.com/arivictor/mojito/actions/workflows/go.yml/badge.svg)](https://github.com/arivictor/mojito/actions/workflows/go.yml)

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
logger.Config.SetLevel(mojito.DEBUG)
logger.Config.SetMessageKey("status")   // message -> status
logger.Config.SetDelimiter('|')         // \n -> |
```

and also includes structured logging as loosely typed key-value pairs:

```go
logger := mojito.New()

logger.Debug("successfully did a thing"
    "name", "mojito",
    "time", time.Now()
    "number", 3,
)
```

# ✨ Contributing

Contributions to this project are welcome. If you encounter any issues or have suggestions for improvements, please open an issue on the GitHub repository.

---

This project is licensed under the [GNU v3 License](LICENSE). Feel free to use and modify the code as per the terms of the license.
