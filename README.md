# go-env-template

Write env var values to text files with Go templates

## Usage

This program will write to stdout or a file (if `-f` is passed). Here is an
example using a template from our **testfixtures** directory.

```
$ export ABCDEF=caterpillar
$ ./go-env-template -t testfixtures/template.properties

-foo.baz.hello=caterpillar

```

## Why?

This is a inspired by the Ruby program Tiller, which is useful for 
taking config values passed to Docker containers as _environment variables_ and
then writing them into _config files_.


