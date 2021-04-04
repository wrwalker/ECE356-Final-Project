# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Double press shift to search everywhere for classes, files, tool windows, actions, and settings.
import mysql.connector

def main():
    mydb = mysql.connector.connect(
        host="localhost",
        user="root",
        password="2q8v8g5N!",
        database="Loyal"
    )

    mycursor = mydb.cursor()

    mycursor.execute("SHOW TABLES")

    for x in mycursor:
        print(x)


if __name__ == '__main__':
    main()