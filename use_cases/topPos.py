#!/bin/python3

from collections import OrderedDict
from operator import itemgetter
import connectionToDB as db
import json
import numpy as np

# Connects to DB.
rows = db.getHistorico()

# Reads the json, groups the entries by the 'cargo' field and calculates the
# median of the 'remuneracaodomes' field for each 'cargo'.
jsons = []
for row in rows:
    jsonData = json.loads(row[2])
    jsons.append(jsonData)

posList = []
posMedian = {}

for data in jsons:
    for entry in data:
        if entry['cargo'] not in posList:
            posList.append(entry['cargo'])
            posRem = []
            for data2 in jsons:
                for entry2 in data2:
                    if entry2['cargo'] == entry['cargo']:
                        posRem.append(entry2['remuneracaodomes'])
            median = np.median(posRem)
            posMedian[entry['cargo']] = median

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