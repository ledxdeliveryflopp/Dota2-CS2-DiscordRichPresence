# Dota 2 / CS:GO 2 Discord Rich Presence (RPC)
Integration of Discord rich presence for the Dota 2 and CS:GO 2

[EN](#About) | [RU](#что-это)

## About

This is a application which adds Dota 2 / CS:GO 2 as a Discord Rich Presence.

This application extracts in-game information using Valve GSI.
This application modifies in no way whatsoever any data from the game, it is completely safe and will not get you banned.

## Debugging requirements

1. Golang 1.23
2. Discord
3. Dota 2 / CS:GO 2

## Features

### Dota 2

1. Displays the name of the character you are playing on.
2. Displays the HP of the character in %.
3. Displays the character icon.
4. Displays the character level.
5. Displays your KDA for the current match.
6. Displays your current gold for the current match.

### CS:GO 2

1. Displays the name of the map you are currently playing on.
2. Displays the round number in the current match.
3. Displays the number of rounds won from CT and T.
4. Displays your team(CT/T) for the current round.
5. Displays your HP and armor for the current round.
6. Displays your KDA and mvps count for the current match.


## Images

**Dota 2**

![plot](/git/static/dota.jpg)

**CS:GO 2**

![plot](/git/static/cs.jpg)


## How to Install

1. Go to "Releases" tab and download the latest version's zip or rar archive.
2. Open "settings.json" and set you steam ID **(If you don't plan to play CS:GO, you can write any string)**:
```json
{
  "steam_id": "<you steam ID>" // string
}
```
3. Transfer or copy the "gamestate_integration_cs2-gsi.cfg" file to - <CS:GO folder>\game\core\cfg\
4. Transfer or copy the "gamestate_integration" folder to - <Dota 2 folder>\dota\cfg\
5. Launch the app after Discord.

[//]: # (## Know issue)

## Что это?

Интеграция Discord rich presence для Dota 2 и CS:GO 2.

Информация получаемая для установки статуса в Discord берется с помощью GSI интеграции от Valve

За использование **нельзя** получить бан так-как приложение никак не взаимодействует напрямую с игрой.

## Зависимости необходимые для деббагинга

1. Golang 1.23
2. Discord
3. Dota 2 / CS:GO 2

### Dota 2

1. Отображает имя персонажа на котором вы играете.
2. Отображает HP персонажа в %.
3. Отображает иконку персонажа.
4. Отображает уровень персонажа.
5. Отображает ваше KDA в текущей игре.
6. Отображает ваше золото в текущей игре.

### CS:GO 2

1. Отображает название карты на которой вы сейчас играете.
2. Отображает номер текущего раунда.
3. Отображает количество выигранных раундов CT/T.
4. Отображает за какую сторону вы играете.
5. Отображает ваше значение HP и брони в текущем раунде.
6. Отображает ваше KDA и количество MVP в текущей игре.

## Изображения

**Dota 2**

![plot](/git/static/dota.jpg)

**CS:GO 2**

![plot](/git/static/cs.jpg)

## Как использовать

1. Перейдите в вкладку "Releases" и скачайте zip или rar архив с последнего релиза.
2.  Отройте файл "settings.json" и установите ваш Steam ID **(Если не планируете играть в CS:GO, то можно установить любую строку)**:
```json
{
  "steam_id": "<you steam ID>" // string
}
```
3. Перенесите или скопируйте файл "gamestate_integration_cs2-gsi.cfg" в - <папка с CS:GP>\game\core\cfg\
4. Перенесите или скопируйте папку "gamestate_integration" в - <папка с Dota 2>\game\cfg\
5. Запустите приложение после запуска Discord.

[//]: # (## Известные ошибки)


## Что планируется добавить
1. Поддержку режимов в для статуса в CS:GO.
2. Русскую локализацию для текста в статусах для CS:GO.