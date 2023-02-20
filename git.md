- set up git user email and name

```git config --global user.email="jitenp@outlook.com"```

```git config --global user.name="jiten"```

- create a directory and initilise it as a git directory

```git init```

- from the above command git directory has been created in your loical machine but not in the remote repository(in gitlab or github etc..)

- after adding/updating required files

1- add or moddify files --> ```git status```
2- add those files to stage --> ```git add helloWorld/```
3- check git status
4- commit those files --> ```git commit -m "add helloworld go app"```
5- to check ```git log```

6- to check branch ```git branch -v```
7- to rename a branch ```git branch -m master```

8- to add remote repo ```git remote add origin https://gitlab.stackroute.in/JitenP/gitdemo.git```

9- to push all commits to remote branch ```git push origin master```

10- if somebody push new changes to remote branch, you can pull those details using ```git pull origin master```