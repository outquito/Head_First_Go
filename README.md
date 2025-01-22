# Head_First_Go
Вивчаємо програмування по книзі Head First Go (headfirstgo.com)

Всім читачам "ПРИВІТ" починаю свій шлях вивчення мови програмування Go! (20.01.2025) 
Всі прикладм, нотатки та свої думки буду викладати тут, по худу читання книги.

# 1. Знайомства з Go
Перша програма на Go (helloGo.go) була написана на веб-сайті https://paly.golang.org

    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, Go!")
    }

    package main - Це строка повідомляє, що весь код цього файлу відноситься до капету "main"

    import "fmt" - Означає, що ми будемо використовувати код форматованого тексту з пакета "fmt"

    func main() - Функция "main" відіграє важливу роль - саме вона виконується при запуску програми

    fmt.Println - Для цього вона викликає функцію "Println" з пакету "fmt"

    ("Hello, Go!") - Ця строка виводить повідомлення "Hello, Go!" в терміналі



ВИКОРИСТАННЯ ФУНКЦІЙ З ДВОМА ПАКЕТАМИ

У файлі twoPackages.go описано як викликати два пакета, завдяки ().

    package main

    import (
        "math"
        "strings"
    )

    func main() {
        math.Floor(2.75)
        strings.Title("head first go")
    }

    math.Floor(2.75) - викликаємо функцію Floor з пакету "math"
    strings.Title("head first go") - викликаємо функцію Title з пакету "strings"

Після того як пакети math та strings будут імпортовані, ми зможемо викликати функції Floor та Title. На даний час данний файл не дає нам ніяких результатів.


ПОВЕРНЕННЯ ЗНАЧЕНЬ ФУНКЦІЙ

returnFunctionValues.go

    package main

    import (
        "fmt"
        "math"
        "strings"
    )

    func main() {
        fmt.Println(math.Floor(2.75))
        fmt.Println(strings.Title("head first go"))
    }

byThePool.go "У басейна" це файл завдання з книги де потрібно заповнити прорущені слова коду щоб від дав результат "Cannonball!!!", використовуючи пропоновані фрагмети коду тільки один раз.

    package ____

    import (
        ____
    )

    ____ main() {
        fmt.Println(____)
    }

main, Println, "Connonball!!!", "math", "fmt", func.