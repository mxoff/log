package log

var tmpl_index = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Logs</title>
    <style type="text/css">
        BODY {
            font-family: 'Monospace', serif; /* Гарнитура текста */
            font-size: 80%; /* Размер шрифта в процентах */
        }
        TABLE {
            border: 1px solid #656565; /* Рамка вокруг таблицы */
            border-collapse: collapse; /* Убираем двойные линии между ячейками */
            width:100%;
        }
        TD, TH {
            padding: 3px; /* Поля вокруг содержимого ячейки */
            height:20px;
        }
        TD {
            text-align: left; /* Выравнивание по левому краю */
          /*  text-align: center; /* Выравнивание по центру */
            border-bottom: 1px solid #656565; /* Линия внизу ячейки */
        }
        TH {
            background: #71706f; /* Цвет фона */
            color: #5977ef; /* Цвет текста */
        }
        .date{
            width:10%;
            white-space: nowrap;
            color: #5a5a5a; /* Цвет текста */
            font-size: 9pt; /
        }
        .ERRO {
            color: #ffffff; /* Цвет текста */
            background: #ea1414; /* Цвет фона */
        }
        .DEBU {
            color: #ffffff; /* Цвет текста */
            background: #5a5a5a; /* Цвет фона */
        }
        .WARN {
            color: #000000; /* Цвет текста */
            background: #ffdb00; /* Цвет фона */
        }

        .atr{
             width:5%;
           white-space: nowrap;

            color: #5a5a5a; /* Цвет текста */
            font-size: 7pt; /
        }

    </style>
</head>
<body>




<table>
{{range .}}
    <tr>
        <td class="date"> {{.Date}} </td>
        <td class="{{.Types}}"> {{.Message}} </td>
        <td> <div class="atr"> {{range .Atr}} {{.}}<br> {{end}} </div> </td>
    </tr>
{{end}}
</table>



</body>
</html>


`
