<?php

require_once __DIR__ . '/vendor/autoload.php';
use PhpAmqpLib\Connection\AMQPStreamConnection;
$dotenv = Dotenv\Dotenv::createImmutable(__DIR__);
$dotenv->load();

$HOST = $_ENV['HOST'];
$PORT = $_ENV['PORT'];
$USERNAME = $_ENV['USERNAME'];
$PASSWORD = $_ENV['PASSWORD'];

$connection = new AMQPStreamConnection($HOST, $PORT, $USERNAME, $PASSWORD, "/");
$channel = $connection->channel();

$channel->queue_declare('test', false, false, false, false);

echo " [*] Waiting for messages. To exit press CTRL+C\n";

$callback = function ($msg) {
    echo ' [x] Received ', $msg->body, "\n";
};

$channel->basic_consume('test', '', false, true, false, false, $callback);

while ($channel->is_consuming()) {
    $channel->wait();
}

$channel->close();
$connection->close();
?>