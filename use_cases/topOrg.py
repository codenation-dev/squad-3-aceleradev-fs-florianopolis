#!/bin/python3

from collections import OrderedDict
from operator import itemgetter
import connectionToDB as db
import json
import numpy as np


# Connects to DB
rows = db.getFuncPublicoBy("orgao")


# Read the DB, groups the entries by the 'orgao' field and calculates the
# median of 'remuneracaodomes' for the each 'orgao'.
orgList = []
orgMedian = {}
for row in rows:
    if row[0] not in orgList:
        orgList.append(row[0])
        orgRem = []
        for row2 in rows:
            if row2[0] == row[0]:
                orgRem.append(row2[1])
        median = np.median(orgRem)
        orgMedian[row[0]] = median

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
