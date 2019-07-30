#!/bin/python3

import connectionToDB as db
import json
import numpy as np
import math

# Connects to DB
rows = db.getFuncPublicoBy("id")

# Counts number of rows and gets the maximun and minimal values.
count = 0
minval = rows[0][1]
print(minval)
maxval = 0
for row in rows:
    count = count + 1
    if row[1] > 20000 and row[1] < minval:
        minval = row[1]
    if row[1] > maxval:
        maxval = row[1]

print(count)
print(minval)
print(maxval)

# Creating number of classes using Sturges Rule.
k = 1 + math.log(count)*3.3
h = (maxval - minval)/k

# Counts number of classes
classes = []

classe = 0
while classe < k:
    classes.append(minval+classe*h)
    classe = classe+1

# Count number of values per class.
classeCount = {}
for entry in classes:
    beginning = minval+classes.index(entry)*h
    if entry == classes[len(classes)-1]:
        end = maxval
    else:
        end = minval+(classes.index(entry)+1)*h
        counter = 0
    for row in rows:
        if row[1] > beginning and row[1] <= end:
            counter = counter+1
    classeCount[str(beginning)+" to "+\
            str(end)] = counter

with open("mostCommon.json", "w") as outfile:
    json.dump(classeCount, outfile)

