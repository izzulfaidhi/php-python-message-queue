# php-python-message-queue #

## Setup ##

Have [composer](http://getcomposer.org), [python37](https://www.python.org/downloads), [docker](https://www.docker.com/products/docker-desktop) installed, then run the following command:

```bash
$ php composer.phar install
```
```bash
$ cd sender
$ pip install -r requirements.txt
```

## Usage ##
Execute commands for rabbitmq:
```bash
$ cd rabbitmq
$ docker-compose up
```

Execute commands for receiver:
```bash
$ cd receiver
$ php receiver.php
```

Execute commands for sender:

```bash
$ cd sender
$ python main.py
```