# php-python-message-queue #

## Setup ##

Ensure you have [composer](http://getcomposer.org) and [python37](https://www.python.org/downloads)  installed, then run the following command:

```bash
$ php composer.phar install
```
```bash
$ cd sender
$ pip install -r requirements.txt
```

## Usage ##

Execute commands:

```bash
$ cd receiver
$ php receiver.php
```

Then on the other Terminal do:

```bash
$ cd sender
$ python main.py
```

You should see the message arriving to the process on the other Terminal