#!/bin/python3

from collections import OrderedDict
from operator import itemgetter
import connectionToDB as db
import json
import numpy as np

# Connects to DB.
rows = db.getFuncPublicoBy("cargo")

# Reads the DB, groups the entries by the 'cargo' field and calculates the
# median of the 'remuneracaodomes' field for each 'cargo'.
posList = []
posMedian = {}
for row in rows:
    if row[0] not in posList:
        posList.append(row[0])
        posRem = []
        for row2 in rows:
            if row2[0] == row[0]:
                posRem.append(row2[1])
        median = np.median(posRem)
        posMedian[row[0]] = median

# Sorts list by from the highest 'remuneracaodomes' to the lowest.
sortedList = OrderedDict(sorted(posMedian.items(), key=itemgetter(1),
    reverse=True))  

# Gets the top 6 entries from the list.
finalList = {}
while True:
    i = 0
    for org in sortedList:
        if i > 5:
            break
        finalList[org] = sortedList[org] 
        i = i+1
    break

# Saves to JSON file.
with open('toppos.json', 'w') as outfile:
    json.dump(finalList, outfile)
