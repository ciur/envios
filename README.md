# Enward

Your application loads its configuration via environment variables?
Do you work locally (on you dev machine) with different groups of environment variables? 
Something like dev, test_pg, test_sqlite?
Do want an utility to load only "dev" environment varibles ?

If your answer is yes, yes yes to all above questions - then meet `enward`.

Enward is command line tool for "loading" a group of environment variable.
To be exact it prints `export X=Y` for each environment variable in specific group.

You define `.enwardrc` file, like the one below, put there are couple of
of environment variables groups:

```
[common]
PAPERMERGE__SECURITY__SECRET_KEY = 123

[dev:common]
PAPERMERGE__DATABASE__URL = postgres://coco:jumbo@db/pmgdb

[test:common]
PAPERMERGE__DATABASE__URL = sqlite3:////sqlite3.db
```
And then you can ask `enward` to export one specific group.
Group of environment varibles is same as "one profile".

To export, or to be exact, to print, environment variables of the profile "test" use:

```
    $ enward -n test
    
    export PAPERMERGE__DATABASE__URL=sqlite3:////sqlite3.db
    export PAPERMERGE__SECURITY__SECRET_KEY=123
```

Both `test` and `dev` profiles inherit environment variables from `common`. 

By default envward will look at `.enwardrc` file from the current working directory.
You can also provide path to specific config file via `-c` flag:

```
    $ enward -n test1 -c parser/test_data/config3.ini
    
    export x=1
    export y=2
```

Finally, you can print all profiles in specific configuration with `-l` flag:

```
    $ enward -l

    1. common
    2. dev inherits from common
    3. test inherits from common
```


## Configuration File Syntax

[SomeParent]
commonVar1 = a
commonVar2 = b
commonVar3 = c

[profileName1:SomeParent:defaultSwitch]
X1 = val1
X2 = val2

[profileName2]
X1 = anothervalue1
X2 = anothervalue2
Y1 = coco

"defaultSwitch" part is not used. Initial idea was to have of one the profiles defined as 'default'; but
in the end there is no practical use for it - thus it has no effect. Just don't use it.

In example above "ProfileName1" has its own variables (X1 and X2) and also inherits another
set of variables from "SomeParent": commonVar1, commonVar2, commonVar3.

"profileName2" has no parent.

Hierarchy is only one level: a profile can have a single parent (say P), and its parent (P) cannot
have other parents i.e. parent/child hirarchy is only one level deep.
