### tccutil2

Apple has an utility called [tccutil](http://ss64.com/osx/tccutil.html) that only allows you to reset all decisions for the specified service. 
However, tccutil doesn't offer much functionality, for that reason **tccutil2** is created

**tccutil2** is an utility made in Go that allows you modify OS X's Accessibility Database from the command line. You will be able to list
all the current service registered in the Accessibility Database, as well as disable/enable and add/remove them.

### Installation

Download from the [releases](https://github.com/JesusTinoco/go-tccutil/releases) page the latest **tccutil2** binary. As it is an
out-of-the-box utility it doesn't require of any extra configuration.

Once the binary is download you can just executed it by the command line:
> $ sudo tccutil2 --help

### Usage

**tccutil2** requires of super-user priveleges to execute the operations.

```
$ ./tccutil2 --help
tccutil2 allow you modify OS X's Accessibility Database from the command line (https://github.com/JesusTinoco/go-tccutil).

Usage:
  tccutil2 [flags]
  tccutil2 [command]

Available Commands:
  add         Add new clients to the database
  disable     Disable the given clients
  enable      Enable the given clients
  list        List all the current clients registered.
  remove      Remove clients from the database
  version     Print the version number of tccutil2

Use "tccutil2 [command] --help" for more information about a command.
```

### Examples

List all the services/clients registered
```
→ sudo ./tccutil2 list
com.divisiblebyzero.Spectacle
com.microsoft.VSCode
```

Enable services/clients
```
→ sudo ./tccutil2 enable com.divisiblebyzero.Spectacle com.microsoft.VSCode
```

Disable services/clients
```
→ sudo ./tccutil2 disable com.divisiblebyzero.Spectacle com.microsoft.VSCode
```

Remove services/clients from the database
```
→ sudo ./tccutil2 remove com.divisiblebyzero.Spectacle com.microsoft.VSCode
```

Add services/clients to the database
```
→ sudo ./tccutil2 add com.divisiblebyzero.Spectacle com.microsoft.VSCode
```

## Contributing

Bug reports and pull requests are welcome.

## License

[MIT License](LICENSE)