
REPEATS = 3

reportFile = open("./report.csv", "w")


reportFile.write(
    """
        Два запроса:
        1) Простой GET запрос, отдающий Hello world
        2) POST запрос, сериализующий и десериализующий объект + парсинг параметра пути 
        
        4 Фреймворка:
        A -- Echo
        B -- Gin
        C -- Fiber
        D -- Http - дефолт
        
        ;A1;B1;C1;D1;A2;B2;C2;D2
        
    """
)

finish = [0,0,0,0,0,0,0,0]
for i in range(REPEATS):
    idx = i + 1
    reportFile.write("Бенчмарк "+str(idx) + ";")


    files = [
        "./out-ab/echo-" + str(idx) +".csv",
        "./out-ab/gin-" + str(idx) +".csv",
        "./out-ab/fiber-" + str(idx) +".csv",
        "./out-ab/http-" + str(idx) +".csv",
        "./out-ab/echo-p-" + str(idx) +".csv",
        "./out-ab/gin-p-" + str(idx) +".csv",
        "./out-ab/fiber-p-" + str(idx) +".csv",
        "./out-ab/http-p-" + str(idx) +".csv",
        ]

    for k, filename in enumerate(files):
        file = open(filename)
        lines = file.readlines()

        test_time = 0
        for line in lines[1:]:
            vals = line.split('\t')
            ctime = int(vals[3])
            test_time += ctime
        file.close()
        reportFile.write(str(ctime) + ";")
        finish[k] += ctime

    reportFile.write("\n")


reportFile.write("Итого;")
for x in finish:
    reportFile.write(str(x) + ";")
reportFile.write("\n")

reportFile.write("Итого по сумме тестов;")
for i in range (0, int(len(finish)/2)):
    reportFile.write(str(finish[i] + finish[i+4]) + ";")
reportFile.write("\n")

