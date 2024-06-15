# Перекрёсток

Выполнил Чернецкий Иван - https://hh.ru/resume/1f691786ff081aaeb50039ed1f7364464c7048

## Общая информация о приложении

Данная работа демострирует вариант реализации приложения, которое имитирует работу перекрёска и алгоритма для оптимальной работы светофоров.

Приложение написано на Golang и его можно легко маштабировать - добавить больше светофоров и условий. 

Перекрёсток состоит из двух перпедикулярно пересекающихся дорог и четырёх светофоров. На зёлёный свет машины могут ехать вперёд или направо.

## Работа приложения

Приложение запускает 6 докер контейнеров

- Светофор 1
- Светофор 2
- Светофор 3 
- Светофор 4
- Центр управления
- Генератор нагрузки

Информация о запуске приложения будет приведена ниже.

### Светофор

- Принимает сообщения от генератора - колличество машин (или людей, если это будет пешеходный светофор)
- Хранит ```quantity``` - колличество машин/людей, которые стоят в очереди
- Хранит ```light``` - состояние светофора на данный момент
- По запросу от Центра Управления отдёт ```quantity```
- По приказу из Центра меняет значение ```light``` и вычитает из ```quantity``` колличество машин, которое успеет проехать за время, на которое было приказано включить зелёный.

### Генератор

Просто рандомно генерирует число машин и отправляет на один из (случайно выбранный) светофоров

### Центр

Раз в цикл запрашивает у светофоров ```quantity```. Расчитывает оптимальную комбинацию и отправляет светофорам команды - цвет, на который надо переключиться, и на какое время его надо включить.

## Алгоритм

Центр оперирует следующим алгоритмом. На перекрёстке может быть конечное колличество комбинаций работы светофоров.

В самом простом случае (когда есть только машины и только 4 светофора) этих комбинаций всего 2 - два противоположных светофора зелёные, остальные два красные.

Дальше легко посчитать, сколько можно пропустить машин за единицу времени в одной и в другой комбинации. Выбирается наиболее выгодный вариант.

При маштабировании задачи (добавить ещё пешеходов и пешеходые светофоры) ситуация остаётся такой же - просто возроствает колличество комбинаций. 

Например: если мы выбираем один автомобильный светофор и делаем его зелёным, то есть всего 2 варианта что будет происходить с остальными:

- либо зелёными будут два пешеходных светофор слева от движения 
- либо автомобильный светофор напротив

Перебрав эти варинаты и расчитав колличество людей/машин, которых можно пропустить, мы найдем нужную в данный момент комбинацию.

Время, на которое надо включать выбранную комбинацию, определяется максимальностью пропускной способности комбинации на момент расчёта.

То есть чтобы ни один (по возможности) из зелёных светофоров не работал в холостую. Как только время работы комбинации заканчивается - данные уточняются и происходит новый рассчёт.

## Запуск приложения

Для запуска приложения необходимо в папке ```envs``` удалить из названий файлов ```.example```. Это переменные окружения, отвечающие за связь контейнеров между собой.

Далее в основнй папке (там где ```docker-compose```) необходимо выполнить команду ```docker compose up -d```.

Запустятся 6 описаных выше контенера, которые сразу начнуть работу.

## Заключение

Я в двух слова описал то, как вижу реализацию такой задачи. Приложение написана на Golang. Этот язык отлично подходит для работы с микросервисной архитектурой, сетевыми запросам и конкурентными/параллельными задачам. Что, собственно, широкими мазками вы можете наблюдать в коде проекта. 
Можно пользоваться стандартными библиотеками или (как в случае со светоформ) готовыми фреймворками, маштабировать задачу, добавлять различные сервисы (на пример прикрутить какой-нибудь брокер сообщений, который возьмёт на себя общение сервисов сежду собой) и т.д и т.п.

Если вас заинтересовала эта небольшая работа и вам интерена дальнейшая реализация этой или иных задач - прошу ко мне в ТГ @Gekko_Moria. Рад новым знакомствам и открыт к предложениям о работе.
