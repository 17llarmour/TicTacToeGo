import time
import requests
import json
from random import *

print("Running Client")

while True:
    time.sleep(1)
    r = requests.get('http://localhost/state')
    print(r.text)
    decoded = json.loads(r.text)

    for x in range(3):
        println = ""
        for y in range(3):
            println += decoded[x][y]
        print(println)

    box = randint(0,8)

    url = "http://localhost/play" + "?place=" + str(box)
    r = requests.post(url)