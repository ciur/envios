# What is Enward?

Enward is command line tool for working "loading" a group of environment variable.

You define `.enwardrc` file, like for example below there are three groups
of environment variables (three profiles):

```
[common]
PAPERMERGE__SECURITY__SECRET_KEY = 123

[dev:common]
PAPERMERGE__DATABASE__URL = postgres://coco:jumbo@db/pmgdb

[test:common]
PAPERMERGE__DATABASE__URL = sqlite3:////sqlite3.db
```

And then you ask enward to "export" one for those profiles.
To export environment variables from profile "test" use command:

```
    $ enward -n test
    
    export PAPERMERGE__DATABASE__URL=sqlite3:////sqlite3.db
    export PAPERMERGE__SECURITY__SECRET_KEY=123
```

Both `test` and `dev` profiles inherit environment variables from `common`. 

By default envward will look at `.enwardrc` file from the current working directory.
You can also provide path to specific config file via `-c` flag:

```
    $enward -n test1 -c parser/test_data/config3.ini
    
    export x=1
    export y=2
```
