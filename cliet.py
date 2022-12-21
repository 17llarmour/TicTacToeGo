import time
import requests
import json
from random import *
from pygame import *

def printBoard():
    for x in range(3):
        println = ""
        for y in range(3):
            println += decoded[x][y]
        print(println)


def drawBoard():
    pass


print("Running Client")
init()
width = 900
height = 900
screen = display.set_mode((width,height))
end_program = False
box = -1






while not end_program:
    for e in event.get():
        if e.type == QUIT:
            end_program = True
        if e.type == KEYDOWN:

            if e.key == K_0:
                box = 0
            if e.key == K_1:
                box = 1
            if e.key == K_2:
                box = 2
            if e.key == K_3:
                box = 3
            if e.key == K_4:
                box = 4
            if e.key == K_5:
                box = 5
            if e.key == K_6:
                box = 6
            if e.key == K_7:
                box = 7
            if e.key == K_8:
                box = 8

    time.sleep(1)
    r = requests.get('http://localhost/state')
    print(r.text)
    decoded = json.loads(r.text)
    printBoard()
    url = "http://localhost/play" + "?place=" + str(box)
    r = requests.post(url)
