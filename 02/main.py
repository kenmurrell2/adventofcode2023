import re

def partOne():
    red = re.compile('[0-9]+(?=\sred)')
    blue = re.compile('[0-9]+(?=\sblue)')
    green = re.compile('[0-9]+(?=\sgreen)')
    game = re.compile('(?<=Game\s)[0-9]+')
    with open("data.txt") as f:
        id_sum = 0
        for line in f: 
            game_id, info = line.split(": ")
            red_sum = 0
            blue_sum = 0
            green_sum = 0
            for m in red.findall(info):
                if(red_sum < int(m)):
                    red_sum = int(m)
            for m in blue.findall(info):
                if(blue_sum < int(m)):
                    blue_sum = int(m)
            for m in green.findall(info):
                if(green_sum < int(m)):
                    green_sum = int(m)
            if red_sum <= 12 and green_sum <= 13 and blue_sum <= 14:
                id_sum += int(game.findall(game_id)[0])
        print(id_sum)

def partTwo():
    red = re.compile('[0-9]+(?=\sred)')
    blue = re.compile('[0-9]+(?=\sblue)')
    green = re.compile('[0-9]+(?=\sgreen)')
    game = re.compile('(?<=Game\s)[0-9]+')
    with open("data.txt") as f:
        id_sum = 0
        for line in f: 
            game_id, info = line.split(": ")
            red_min = 0
            green_min = 0
            blue_min = 0
            for m in red.findall(info):
                if red_min < int(m):
                    red_min = int(m)
            for m in blue.findall(info):
                if blue_min < int(m):
                    blue_min = int(m)
            for m in green.findall(info):
                if green_min < int(m):
                    green_min = int(m)
            id_sum += (red_min*blue_min*green_min)
        print(id_sum)

#partOne()
partTwo()