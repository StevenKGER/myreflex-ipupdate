# myreflex-ipupdate
## Functionality
This tool provides an automatic whitelist update, if the server's IP address is changing (e.g. in case, where a server is hosted at home).

## How to use
Compile it yourself with ``go build`` or use the pre-compiled binaries available in the Releases tab.

Program usage:

``
./myreflex-ipupdate -username yourSpigotMCUsername -password yourMyReflexAPIPassword
``

The command line arguments can also be provided as environment variables:
- ``MYREFLEX_USERNAME``: SpigotMC username
- ``MYREFLEX_USERID``: SpigotMC userID (not required, if username set and vice versa)
- ``MYREFLEX_PASSWORD``: MyReflex-API password

Help output:
````
Usage of myreflex-ipupdate:
  -password string
        password for authenticating to MyReflex API services; can also be set as env MYREFLEX_PASSWORD
  -userID int
        userID for authenticating to MyReflex API services; can also be set as env MYREFLEX_USERID
  -username string
        username for authenticating to MyReflex API services; can also be set as env MYREFLEX_USERNAME
````

## Technical information
This tool is written in go 1.14.3 and does not use external libraries.

## Licensing
This software is licensed under the Apache License 2.0. For more information see the LICENSE file.

## Contributing
Feel free to open issues and/or pull requests if you think it's required. If you have any questions just contact us right away.
For general questions just join our [Discord](https://g.reflex.rip/discord).