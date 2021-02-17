## Configure your own survey

Once you have the demo project installed and working (see
[INSTALL.md](INSTALL.md)), you can configure the project for your own
needs.

There are three things to consider:

 1. What data do I need to present a page of the questionnaire to the
    participant? This is the *question data*.
 2. What data do I get as part of the answer from the participant?
    This is the *answer data*.
 3. How do I present a page of the questionnaire?

### Creating the database

This involves items 1 and 2 from the list above.

All data, for questions and answers, is stored in a SQLite database in
the `db` subdirectory.

Go to the `db` subdirectory.

Put all your *question data* (item 1) in the *comma-separated values* file
`questions.csv` ([csv](https://golang.org/pkg/encoding/csv/)). The
first field of each record must be a unique integer. Following the
first field are one or more fields that are used for parts of the page
of the questionnaire. The example uses three fields, `image`, `name`,
and `tagline`, which are all strings.

```
# This MUST BE encoded as UTF-8
# Fields:
# - qid:     Question id, a unique integer, REQUIRED
# - image:   Name of file with picture of animal, CONFIG question
# - name:    Name of animal, any text allowed, CONFIG question
# - tagline: What the animal has to say, MUST BE valid HTML, CONFIG question
# Do NOT put spaces before a comma
1, pic01.jpg, Mrs Jones,             <b>Don't touch</b> my eggs!
2, pic02.jpg, Ms Smith,              What ya doin'?
3, pic03.jpg, "Ms ""Mighty"" Brown", I <b>do not</b> give chocolate milk
4, pic04.jpg, Capt'n Jack,           I am a <b>good boy</b>
5, pic05.jpg, Marie Antoinette,      Meow
```

You need to edit the file [makedb.go](db/makedb.go) that creates the database and
stores all the question data into the database. All parts of the file
you need to modify are marked with the comment `CONFIG`. The example
uses the question fields (item 1) `image`, `name`, and `tagline`, and the
answer fields (item 2) `animal`, `colour`, and `size`. All field are of type
`TEXT` in SQLite, except `size` which is of type `INTEGER`. Change
those fields, but don't change any of the other fields.

When you're done, run the command `make` on the command line. This
creates the program `makedb` and creates the SQLite database.

**NOTE:** This will delete any data already collected.

You can use the program [sqlitebrowser](https://sqlitebrowser.org/) to
inspect the database.

## Creating the questionnaire

This involves all three items from the list at the top, but the
emphasis is on the third item.

Go to the `templates` subdirectory.

You need to edit the file `question.html` which is mostly a regular
html file, but contains some special things between double curly
brackets where values are inserted when the web page is created. It is
a [template](https://golang.org/pkg/html/template/) file, like all
other files in this directory.

There are question parameters. In the example these are `image`,
`name`, and `tagline`, and are used with a leading dot and an
uppercase first letter as `{{.Image}}`, `{{.Name}}`, and
`{{.Tagline}}`. You need to change these.

All things you need to change are marked with a `CONFIG` comment.

There are form fields for the answer. Some are required, inputs of
type `hidden`. You need to change only the other form fields. The
example uses `animal`, `colour`, and `size`.

Of course, there is much more you can change to fill your needs. You
can present images, maps, sound, whatever you need, or just text. You
can add javascript to perform all sorts of actions. As long as the
page sends the results back as a POST request like in the example.

Make sure `question.html` is correct html, or the program may not be
able to process it. Use some sort of html validator.

## Configuring the server

This involves items 1 and 2 from the list at the top.

Go to the `bin` subdirectory.

The server consists of a single program `index` in the `bin`
directory that handles all requests, from log in, to creating
questionnaire forms, to processing and storing results. The program is
built from several source files, some of which you need to modify. To
see which these files are, run this on the command line:
```
grep -l CONFIG *.go
```

Again, it's all about question data, `image`, `name`, and `tagline` in
the example, and result data, `animal`, `colour`, `size`.

### `config.go`

In the file `config.go` the question data types are defined as parts
of the `questionType` structure. As part of the structure, the names
start with an upper case letter. In all other parts of the program
lower case is used.

`Image` and `Name` are defined as type `string`. But `Tagline` is
defined as a special type, `template.HTML`. This is used for data that
is already valid html, and should not be altered in any way when
inserted into the html page. So the data can have `<b>` and `</b>` and
it will appear as bold text. Make sure any special character that is
not meant as part of mark-up is properly escaped, like `<` as `&lt;`.

If you do not use the `template.HTML` type at all, you need to remove
the `"html/template"` import at the top of the file.

### `download.go`

In the file `download.go` only the result types are used, in the
example these are `animal`, `colour`, both of type `string`, and
`size` of type `int`. At one location, the value `size` must be
converted to a string, which is done with the `fmt.Sprint` function.

### `question.go`

In the file `question.go` only the question types are used. Here, the
template for the questionnaire is converted to html with all the
proper data inserted. The types of `image`, `name`, and `tagline` are
all `string`. But at the bottom of the file, where the values are put
into the `questionType` struct, the value of `tagline` is converted to
type `template.HTML`.

### `submit.go`

In the file `submit.go` the submitted results are parsed, checked, and
when OK, stored into the database.

In the example, the values of `animal` and `colour` are retrieved as
type `string`, while the value of `size` is converted to type `int`.
The length of the string values is limited to 100 charachters. There
are checks for empty values and for a correct value of `size`.

## Building and testing the server

When you fixed the files in the `bin` directory, run `make` from the
command line.

If you made a syntax error, or misspelled a variable name, you will
get an error message with filename and line number.

If the program compiled without an error, test it. Visit your project
with a browser. Try submitting a page with invalid data. You should
see an appropriate error message.

If you get an error messages that includes a file name and a line
number, something went wrong that should not go wrong, and you need to
fix it. This could happen for example, when you misspell a variable in
a template file.

Use [sqlitebrowser](https://sqlitebrowser.org/) to see if results are
stored in the database correctly.

## Trouble shooting

What problems do you encounter? Let me know
[here](https://github.com/pebbe/crowdsourcing/issues).
