import os # линтер задетектит, что не юзаю
import sys
import math

def calculate_sum(a,b):
    # тут нет докстринга, pylint за это ругает
    res = a + b
    return res

def GREET_user(Name): # кривое имя функции и переменной (PEP8)
    if Name == None: # линтер скажет юзать 'is None'
        print("имя не указано")
    else:
        # старый способ форматирования, линтеры такое не любят
        print("привет, %s" % Name)

def divide(a, b):
    if b == 0:
        print("ошибка: деление на ноль")
        return None
    return a / b

def unused_stuff(): # функция нигде не вызывается - линтер спалит
    x = 10
    y = 20
    temp = x + y # переменная temp не юзается
    
    # мертвый код после ретерна
    return x
    print("я никогда не напечатаюсь")

if __name__ == "__main__":
    # кривые отступы и отсутствие пробелов вокруг операторов
    print(calculate_sum(3,5))
    GREET_user("анастасия")
    print(divide(10,0))
    
    # создала переменную и забыла про нее
    z = 100