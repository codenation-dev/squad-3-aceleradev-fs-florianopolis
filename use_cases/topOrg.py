#!/bin/python3

from collections import OrderedDict
from operator import itemgetter
import connectionToDB as db
import json
import numpy as np

rows = db.getHistorico()

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

sortedList = OrderedDict(sorted(orgMedian.items(), key=itemgetter(1),
    reverse=True))  

finalList = {}
while True:
    i = 0
    for org in sortedList:
        if i > 5:
            break
        finalList[org] = sortedList[org] 
        i = i+1
    break

with open('toporgs.json', 'w') as outfile:
    json.dump(finalList, outfile)
