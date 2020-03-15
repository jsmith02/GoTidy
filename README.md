# GoTidy

A golang CLI aliasing tool

--------------------------------------------------------------------
### What it does
Tidy sets up a simple file structure in your $HOME directory called .tidy. 


In that directory it stores maps of commands you provide to it via Tidy's CLI

Add a chore to tidy:

```
jms@jms-desktop:~/$ tidy chore
Enter Command Alias (You'll use this to call the function later e.g. tidy dc for docker compose ): 
de
Enter command with flags e.g. (aws ssm get-parameter --name |_var_| --with-decryption): 
docker exec -it |_var_| /bin/sh
Alias configured for de
```

See all your configs by running list.

```
jms@jms-desktop:~/go/src/GoTidy$ tidy chore list
gp : git push
ssmp : aws ssm put-parameter --name |_var_| --type |_var_| --value |_var_|
di : docker images
ecr : aws ecr get-login --no-include-email
p : psql -h |_var2_| -U |_var_| |_var_|
de : docker exec -it |_var_| /bin/sh
dont : aws ssm get-paramter --name |_var_| --with-decryption

```

These can be edited in that directory, if you don't identify the variables, it will try to parse the flags but may do so incorrectly right now. 


```
jms@jms-desktop:~/.tidy$ cat up
{"alias":["de"],"cmd":["docker exec -it |_var_| /bin/sh"]}
jms@jms-desktop:~/.tidy$ 
```

Call back the aliased commands by running the Tidy CLI

```
jms@jms-desktop:~/go/src/GoTidy$ ./tidy up de 8io
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
jms@jms-desktop:~/go/src/GoTidy$ 

```

Tidy will indicate when you've not provided what it needs to execute a command:

```
Incorrect number of arguments passed for alias, "ssmp".
Arguments needed: 3
Arguments provided: 1
```

Edit commands using tidy e

```
jms@jms-desktop:~/go/src/GoTidy$ tidy e
1) *{"alias":["do"],"cmd":["aws ssm get-parameter --region |_var_| --name |_var_| --with-decryption"]}
2) *{"alias":["gc"],"cmd":["git commit -m |_var_|"]}
3) *{"alias":["gp"],"cmd":["git push"]}
4) *{"alias":["ssmp"],"cmd":["aws ssm put-parameter --name |_var_| --type |_var_| --value |_var8_|"]}
5) *{"alias":["di"],"cmd":["docker images"]}
6) *{"alias":["ecr"],"cmd":["aws ecr get-login --no-include-email"]}
7) *{"alias":["p"],"cmd":["psql -h |_var_| -U |_var_| |_var_|"]}
8) *{"alias":["de"],"cmd":["docker exec -it |_var_| /bin/sh"]}
Which config do you want to edit?

```

No duplicate aliases are allowed:

```
jms@jms-desktop:~/go/src/GoTidy$ tidy chore
Enter Command Alias (You'll use this to call the function later e.g. tidy dc for docker compose ): 
de
Woops, that alias is already used!
jms@jms-desktop:~/go/src/GoTidy$ 

```

Remove commands from your up with delete or del. 


```
jms@jms-desktop:~/go/src/GoTidy$ tidy del di
Removed di from up
```

### Commands:
Usage:
  tidy [command]

Available Commands:
  chore       Alias a command by adding a chore to your .tidy/up
  count       The number of chores in your up.
  delete      Delete value from up by specifying alias, can also be called with del
  edit        Edit rules through a CLI menu, specify alias to edit specific command
  help        Help about any command
  init        Configures your local directories for app usage
  up          Run your aliased command by running tidy up <cmd_here>

Flags:
  -h, --help   help for tidy
Use "tidy [command] --help" for more information about a command.


### To Do:

1. Add optional description
2. Add export
3. Add import
