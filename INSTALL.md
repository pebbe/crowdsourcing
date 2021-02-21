## Install the crowdsourcing software with the demo survey

This requires a web server with CGI enabled. It has been tested with
Apache2.

### Before you begin

You need the [Go](https://golang.org/) compiler, and the standard
build tools `gcc` and `make`.

This software needs compiling. You can't just compile it on any
machine (your local machine) and have it work on any other machine
(the machine running the web server).

If you can't compile the software on the machine the web server is
running, try to use a machine with a matching architecture, like Linux
on amd64. This should work even if the machines have different version
numbers for standard libraries, if you compile the software as static
programs (see below: Creating the server).

If the above is not an option, for instance you have a Windows
computer and the web server is running on Linux, then using
[Docker](https://www.docker.com/) could be a possibility. Create an image
in Docker which emulates the same environment as the web server, and
build the software inside that image.

### Installing the software

Everywhere you see code like `{{this}}` you need to replace it with
your own text.

Choose a directory name for your project, and download the software to
that directory of your website:

```sh
cd {{/path/to/my/website}}
git clone --depth 1 https://github.com/pebbe/crowdsourcing {{project}}
```

### Creating the server

Go to the `bin` subdirectory:

```sh
cd {{project}}/bin
```

Copy `Makefile.cfg.example` to `Makefile.cfg`, and modify it to your
needs. See the comments in the file. If you are running on Linux and
want to build a static program, you probably don't need to change
anything.

Copy `config.go.example` to `config.go`, and modify it to your needs.
See the comments in the file. At the moment, you only need to fix the
values following the line that starts with `const`, and leave
everything above it unchanged.

Run this command:

```sh
make
```

### Creating the database

Go to the `db` subdirectory, copy the example questions, and run `make`

```sh
cd {{project}}/db
cp questions.csv.example questions.csv
make
```

### Upload and test

If you aren't working on your web server yet, copy everything to the
web server.

Start your web browser, and see if it is working. You should get a
log-in screen, where you can enter you e-mail address to log in. Then
you get an e-mail message that you can use to continue to the
questionnaire.

### Set up your own survey

The next step is to change the demo survey into your own survey. To do
this, read the instructions in [CONFIG.md](CONFIG.md).

### Trouble shooting

What problems do you encounter? Let me know
[here](https://github.com/pebbe/crowdsourcing/issues).

----

When running `make`, you may get error messages like these:

```
... cannot find package "github.com/dchest/authcookie" ...
... cannot find package "github.com/mattn/go-sqlite3" ...
```

This happens with older versions of Go. To fix this, run these
commands once:

```sh
go get github.com/dchest/authcookie
go get github.com/mattn/go-sqlite3
```

----
