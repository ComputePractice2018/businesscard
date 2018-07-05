# businesscard

## Usecases

1. Как пользователь, я хочу иметь возможность просмотреть все визитные карты, чтобы использовать эту информацию.
1. Как пользователь, я хочу иметь возможность добавить визитную карту (ФИО, Должность, Телефон, E-mail, Место работы, Адрес), чтобы пополнить список визитных карт.
1. Как пользователь, я хочу иметь возможность изменять содержимое моей визитки, чтобы исправлять ошибки при добавлении.
1. Как пользователь, я хочу иметь возможность найти визитную карту по указанному в ней полю "ФИО".
1. Как пользователь, я хочу иметь возможность удалить мою визитную карту, чтобы не хранить неактуальную информацию.

## REST API

### GET /api/businesscard/vcards

Ответ 200 OK
```json
[{
    "name": "ФИО",
    "position": "Должность",
    "phone": "+7-999-999-99-99",
    "email": "user@domain.ru",
    "workplace": "Место работы",
    "address": "Адрес"
}]
```

### POST /api/businesscard/vcards

Тело запроса:

```json
    {
        "name": "ФИО",
        "position": "Должность",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "workplace": "Место работы",
        "address": "Адрес"
    }
```

Ответ 201 Created
Location: /api/businesscard/vcards/1

### PUT /api/businesscard/vcards/

Тело запроса:

```json
    {
        "name": "ФИО",
        "position": "Должность",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "workplace": "Место работы",
        "address": "Адрес"
    }
```

Ответ 202 Accepted
Location: /api/businesscard/vcards/1

### GET /api/businesscard/vcards?name=ФИО

Тело запроса:

Ответ 200 OK
```json
    [{
        "name": "ФИО",
        "position": "Должность",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "workplace": "Место работы",
        "address": "Адрес"
    }]
```

### DELETE /api/businesscard/vcards/

Ответ 204 No Content
