## Install the crowdsourcing software with the demo survey

This requires a web server with CGI enabled. It has been tested with
Apache2.

### Before you begin

You need the [Go](https://golang.org/) compiler, and the standard
build tools `gcc` and `make`.

The software consists of two programs. One is for initialising the
database, and can be run anywhere. You can then upload the database to
your web server.

The other program is to be run on the web server itself, and if you
can't compile it on the web server, you need to compile it on a
compatible platform. If the web server runs on Linux, and your own
machine runs a newer version of Linux, it may be that the web server
doesn't have the right C library. Trying to run the program on the
server may result in an error like this:

```
./index: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.28' not found (required by ./index)
```

One solution, if you can't compile the software on the web server, is
to use [Docker](https://www.docker.com/). Create an image in Docker
which emulates the same environment as the web server, and build the
software inside that image.

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
needs. See the comments in the file.

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
this, read the instructions in [CONFIG.md](CONFIG.md]).

### Trouble shooting

What problems do you encounter? Let me know
[here](https://github.com/pebbe/crowdsourcing/issues).

----

When running `make`, you may get an error message that starts like this:

```
headers.go:4:2: cannot find package "github.com/dchest/authcookie" in any of:
```

This happens with older versions of Go. To fix this, run these
commands once:

```sh
go get github.com/dchest/authcookie
go get github.com/mattn/go-sqlite3
```

----
