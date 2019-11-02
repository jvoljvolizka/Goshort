# Goshort
A simple url shortener written with golang

## How ?

* Send a post request to /URL with id and link
just like that:

```
curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"ID":"dalyarak","URL":"http://dalcode.jvoljvolizka.xyz"}' \
    http://127.0.0.1:3300/URL
 ```


* You can access list all the links by sending a GET request to  /URLs

***
## Nasıl ?

* Aşağıdaki gibi /URL altına ID ve link ile bir POST isteği göndermek yeterli:

```
curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"ID":"dalyarak","URL":"http://dalcode.jvoljvolizka.xyz"}' \
    http://127.0.0.1:3300/URL
 ```


* Tum linklere /URLs altına GET isteği atarak ulaşabilirsin

***

Copyright © 2019 Jvol Jvolizka <jvoljvolizka@protonmail.com>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.
