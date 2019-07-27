#!/bin/python3

from collections import OrderedDict
from operator import itemgetter
import connectionToDB as db
import json
import numpy as np


# Connects to DB
rows = db.getHistorico()


# Read the json, groups the entries by the 'orgao' field and calculates the
# median of 'remuneracaodomes' for the each 'orgao'.
jsons = []
for row in rows:
    jsonData = json.loads(row[2])
    jsons.append(jsonData)

orgList = []
orgMedian = {}

for data in jsons:
    for entry in data:
        if entry['orgao'] not in orgList:
            orgList.append(entry['orgao'])
            orgRem = []
            for data2 in jsons:
                for entry2 in data2:
                    if entry2['orgao'] == entry['orgao']:
                        orgRem.append(entry2['remuneracaodomes'])
            median = np.median(orgRem)
            orgMedian[entry['orgao']] = median

# Sorts list from the highest 'remuneracaodomes' to the lowest.
sortedList = OrderedDict(sorted(orgMedian.items(), key=itemgetter(1),
    reverse=True))  

# Selects the top 6 entries of the sorted list.
finalList = {}
while True:
    i = 0
    for org in sortedList:
        if i > 5:
            break
        finalList[org] = sortedList[org] 
        i = i+1
    break

# Saves to JSON.
with open('toporgs.json', 'w') as outfile:
    json.dump(finalList, outfile)
