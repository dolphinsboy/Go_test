## mplayer
### 测试

```
bogon:mplayer_go guosong$ go test mlib
ok      mlib    0.006s
```

### 功能测试

构建二进制文件:

```
go build
```

测试
```
bogon:mplayer_go guosong$ ./mplayer_go

        lib list --View the existing music lib
        lib add <name><artist><genre><source><type> -- Add a music to the music lib
        lib remove <name> --Remove the specified music from the lib
        play <name> -- Play the specified music

Enter command->
list
Unreconginzed command: list
Enter command->
lib list
Enter command->
lib add a b c d e
Enter command->
lib list
1 : a b d e
Enter command->
play a
Unsupported music type e
Enter command->
```

## Game

```
Command> h   

                Commands:
                        login <username><level><exp>
                        logout <username>
                        send <message>
                        listplayer
                        quit(q)
                        help(h)
                
Command> login guosong12 3 4
Command> listplayer
1 : &{guosong 1 2 0 <nil>}
2 : &{guosong12 3 4 0 <nil>}
Command> send hello
guosong received message hello
guosong12 received message hello
Command> 

```
