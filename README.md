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

These can be edited in that directory, if you don't identify the variables, it will try to parse the flags but may do so incorrectly right now. 


```
jms@jms-desktop:~/.tidy$ cat up
{"alias":["de"],"cmd":["docker exec -it |_var3_| /bin/sh"]}
jms@jms-desktop:~/.tidy$ 
```

Call back the aliased commands by running the Tidy CLI

```
jms@jms-desktop:~/go/src/GoTidy$ ./tidy up de 8io
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
jms@jms-desktop:~/go/src/GoTidy$ 

```

Edit commands using tidy e

```
jms@jms-desktop:~/go/src/GoTidy$ tidy e
1) *{"alias":["do"],"cmd":["aws ssm get-parameter --region |_var4_| --name |_var6_| --with-decryption"]}
2) *{"alias":["gc"],"cmd":["git commit -m |_var3_|"]}
3) *{"alias":["gp"],"cmd":["git push"]}
4) *{"alias":["ssmp"],"cmd":["aws ssm put-parameter --name |_var4_| --type |_var6_| --value |_var8_|"]}
5) *{"alias":["di"],"cmd":["docker images"]}
6) *{"alias":["ecr"],"cmd":["aws ecr get-login --no-include-email"]}
7) *{"alias":["p"],"cmd":["psql -h |_var2_| -U |_var4_| |_var_|"]}
8) *{"alias":["de"],"cmd":["docker exec -it |_var3_| /bin/sh"]}
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



### Commands:
An attempt at a CLI aliaser and my first foray into GoLang

Usage:
  tidy [command]

Available Commands:
  chore       Alias a command by adding a chore to your .tidy/up
  count       The number of chores in your up.
  e           Edit rules through a CLI menu
  help        Help about any command
  init        Configures your local directories for app usage
  up          Run your aliased command by running tidy up <cmd_here>


Flags:
  -h, --help   help for tidy

Use "tidy [command] --help" for more information about a command.

### To Do:

1. List Aliases
2. Remove Alias
3. Test new commands
5. Add optional description
6. Add ability to edit commands by giving the alias.