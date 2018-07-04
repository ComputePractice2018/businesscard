# businesscard

# Usecases

1. Как пользователь, я хочу иметь возможность просмотреть все визитки, чтобы использовать информацию, содержащуюся в них. 
2. Как пользователь, я хочу иметь возможность добавить визитку (ФИО, должность, телефон, E-mail, место работы, адрес), чтобы пополнять список визиток. 
3. Как пользователь, я хочу иметь возможность найти нужную визитку по ФИО, чтобы быстро получить нужную информацию. 
4. Как пользователь, я хочу иметь возможность редактировать визитку, чтобы исправлять ошибки при добавлении или обновлять информацию на актуальную. 
5. Как пользователь, я хочу иметь возможность удалить визитку,чтобы не хранить ненужную информацию.

## REST API 

### GET /api/businesscard/vcards 

Ответ: 200 OK
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

Ответ: 201 Created
Location: /api/businesscard/vcards/1
    [
    }]

### GET /api/businesscard/vcards?name=ФИО

Ответ: 200 OK
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

### PUT /api/businesscard/vcards 


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

Ответ: 202 Accepted
Location: /api/businesscard/vcards/1
    [
    }]

### DELETE /api/businesscard/vcards 

Ответ: 204 Not Content
Location: /api/businesscard/vcards/1
    [
    }]

