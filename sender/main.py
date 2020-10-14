import pika
import psutil

from env import PORT, HOSTNAME, USERNAME, PASSWORD

credentials = pika.PlainCredentials(USERNAME, PASSWORD)
connection = pika.BlockingConnection(pika.ConnectionParameters(HOSTNAME, PORT, USERNAME, credentials))
channel = connection.channel()

channel.queue_declare(queue='test')

hdd = psutil.disk_usage('.')

message = 'Available disk space: {} GB'.format(hdd.free / (2**30))
channel.basic_publish(exchange='',
                      routing_key='test',
                      body=message)

print("message sent: {}".format(message))

connection.close()
