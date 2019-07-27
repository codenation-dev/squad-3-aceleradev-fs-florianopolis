import mysql.connector
import os

dbInfo = {
    'Host': os.getenv('MYSQL_HOST'),
    'Port': os.getenv('MYSQL_PORT'),
    'Database': os.getenv('MYSQL_DATABASE'),
    'User': os.getenv('MYSQL_USER'),
    'Password': os.getenv('MYSQL_PASSWORD'),
}

# function GETHISTORICO--------------------
# --Purpose: takes the last 48 entries of
#           the DB table HISTORICO and
#           returns them.
#--Params: None.
#--Return: rows.
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

#function GETFUNCPUBLICOBY------------------
#--Purpose: selects from the DB table
#           FUNCPUBLICO the values from a
#           given column and the 
#           correspondent salary.
#--Params: column (string)
#--Return: rows
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




