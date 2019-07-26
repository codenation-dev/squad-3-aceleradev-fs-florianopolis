import connectionToDB as db
import json
import numpy as np
import datetime


print("Inicio de acesso ao db: ", datetime.datetime.now())
rows = db.getHistorico()
print("Banco de dados acessado em: ", datetime.datetime.now())

months = []
monthsMedian = {}

for row in rows:
    if row[1].month not in months:
        print("Data: ", row[1])
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
        print(monthsMedian)
        print(months)
    
allMedian = []
for month in monthsMedian:
    allMedian.append(monthsMedian[month])

totMedian = np.median(allMedian)

monthPerc = {}

for month in monthsMedian:
    dif = monthsMedian[month] - totMedian
    
    perc = 100*abs(dif)/totMedian

    monthPerc[month] = perc

with open("topmonths.json", 'w') as outfile:
    json.dump(monthPerc, outfile)

print("Operação concluída em: ", datetime.datetime.now())


