#!/bin/python3

import connectionToDB as db
import json
import numpy as np
import datetime

# Connects to DB
rows = db.getHistorico()


# Reads the json values and categorize them by the respective month.
months = []
monthsMedian = {}

for row in rows:
    if row[1].month not in months:
        months.append(row[1].month)
        
        jsons = []
        for row2 in rows:
            if row[1].month == row2[1].month:
                jsonData = json.loads(row[2])
                jsons.append(jsonData)
        
        medians = []
        for entry in jsons:
            totliq = []
            for func in entry:
                totliq.append(func["totalliquido"])
        medians.append(np.median(totliq))

        monthsMedian[row[1].month] = np.median(medians)
    

# Gets the total median of the entire period.
allMedian = []
for month in monthsMedian:
    allMedian.append(monthsMedian[month])

totMedian = np.median(allMedian)

# Calculates the percent diference between each month and the total.
monthPerc = {}

for month in monthsMedian:
    dif = monthsMedian[month] - totMedian
    
    perc = 100*dif/totMedian

    monthPerc[month] = perc

# Saves table to JSON.
with open("topmonths.json", 'w') as outfile:
    json.dump(monthPerc, outfile)



