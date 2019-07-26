import mysql.connector

dbInfo = {
    'Host': '93.188.162.171',
    'Port': '3306',
    'Database': 'PROJETOUATI',
    'User': 'root',
    'Password': 'mbdsqd3fln',
}

def getHistorico():
    cnx = mysql.connector.connect(
            host=dbInfo['Host'],
            port=dbInfo['Port'],
            database=dbInfo['Database'],
            user=dbInfo['User'],
            password=dbInfo['Password'],
            )
    cursor = cnx.cursor()

    selectRows = ('SELECT * FROM HISTORICO')

    cursor.execute(selectRows)

    rows = cursor.fetchall()

    return rows

    def close():
        cursor.close()
        cnx.close()

def getFuncPublicoBy(column):
    cnx = mysql.connector.connect(
            host=dbInfo['Host'],
            port=dbInfo['Port'],
            database=dbInfo['Database'],
            user=dbInfo['User'],
            password=dbInfo['Password'],
            )
    cursor = cnx.cursor()

    selection = ('SELECT %s, remuneracaodomes FROM FUNCPUBLICO' % column)

    cursor.execute(selection)

    rows = cursor.fetchall()

    return rows

    def close():
        cursor.close()
        cnx.close()



