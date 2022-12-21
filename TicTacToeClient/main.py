import time as t
import requests
import json
from random import *
from pygame import *

def printBoard(decoded):
    for x in range(3):
        println = ""
        for y in range(3):
            println += decoded[x][y]
        print(println)


def drawBoard(decoded):
    screen.fill((255,255,255))
    draw.rect(screen, (31,90,255), (275,0,50,900))
    draw.rect(screen, (31,90,255), (575,0,50,900))
    draw.rect(screen, (31,90,255), (0,275,900,50))
    draw.rect(screen, (31,90,255), (0,575,900,50))
    yOffset = 125
    for x in range(3):
        xOffset = 125
        for y in range(3):
            if decoded[x][y] == "x":
                draw.rect(screen, (0, 0, 0), (xOffset-75, yOffset, 200, 50))
                draw.rect(screen, (0, 0, 0), (xOffset, yOffset-75, 50, 200))
            if decoded[x][y] == "o":
                draw.circle(screen, (0,0,0), (xOffset + 25,yOffset + 25), 100, width = 50 )
            xOffset += 300
        yOffset += 300
    display.flip()


def handleWeb(box):
    url = "http://localhost/play" + "?place=" + str(box)
    r = requests.post(url)
    r = requests.get('http://localhost/state')
    print(r.text)
    decoded = json.loads(r.text)
    printBoard(decoded)
    drawBoard(decoded)


def checkWin():
    r = requests.get("http://localhost/win")
    # print(r.text)
    decodedWin = json.loads(r.text)
    if decodedWin[2] == "9" and decodedWin[1].lower() != "true":
        print("Draw")
        print(decodedWin)
        end_program = True
    elif decodedWin[1].lower() == "true":
        print(decodedWin)
        print(decodedWin[0], "has won")
        end_program = True
    else:
        end_program = False
    return end_program

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
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_1:
                box = 1
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_2:
                box = 2
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_3:
                box = 3
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_4:
                box = 4
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_5:
                box = 5
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_6:
                box = 6
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_7:
                box = 7
                handleWeb(box)
                end_program = checkWin()
            if e.key == K_8:
                box = 8
                handleWeb(box)
                end_program = checkWin()
   # handleWeb(box)

