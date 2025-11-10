
# Стек проекта
Backend: Golang/gin/jwt
Frontend: html/css
Database: PostgreSQL
Контейнеризация: Docker,Docker-compose
---------------
# Структура проекта
Пользователь взаимодействует с web интерфейсом на html/css, с помощью http протокола отправляет запросы на сервер
На сервере backend на go обрабаывает запрос и в зависимости от запроса взаимодействует с БД на PostgreSQL

--------------
# Описание проекта
1) Маркетплейс в котором пользователь может добавлять товары в список желаемого или в корзину
В корзине пользователь может совершать покупку товара и получать id заказа со списком продуктов(симуляция оформления заказа без транзакции)
Пополнение баланса (симуляция)
Добавить возможность отмены заказа если он не принят?
Видеть статус заказа (оформлен/принят)
2) Админ может пополнять список товаров по категориям и добавлять описание,фото
Главный админ может добавлять новых администраторов или удалять их
Админ видит списки заказов с позициями каждого заказа и может изменить статус заказа с оформлен на принят
Создавать новые категории
3) Авторизация по роли admin/user
--------------
# Сущности
1) User(user_id,user_login,user_password,user_role,user_balance)
2) Wishlist(wishlist_user_id wishlist_product_id)
3) Bucket(bucket_user_id,bucket_product_id,bucket_quantity,bucket_amount)
3) Product(product_id,product_price,product_category_id,product_description,product_name)
4) Category(category_id,category_name,category_description)
5) Orders(order_id,order_user_id,order_amount,order_status,order_created,order_updated)
6) order_items(order_items_id,order_items_product_id,order_items_price,order_items_quantity)

----------------
# Методы и функции
----------------
## User
1) CreateUser создаёт нового пользователя из данных запроса. Возвращает ошибку, если не удалось сгенерировать хэш пароля.
2) Withdraw — снимает деньги с баланса пользователя. Возвращает ошибку, если на счёте недостаточно средств.
3) Deposit — пополняет баланс пользователя на указанную сумму. Возвращает ошибку, если сумма некорректна.
4) UserExist-проверка существует ли пользователь в базе данных.
5) CheckUserData-проверка запроса создания пользователя на корректность данных.
6) SaveUser-сохранение пользователя в базу данных.
7) NewUser-обработчик POST запроса на создание пользователя по эндпоинту /user
8) AdminInit-создаем роль суперадмина и сохраняем в базу данных.
----------------
## Utils
1) HashPassword генерирует хэш из пароля. Возвращает хэш и ошибку, если генерация не удалась.
2) VerifyPassword проверяет, соответствует ли пароль хэшу.
----------------
# API Эндпоинты
**1. POST /user/signup - регистрация нового пользователя**

**Запрос:**
```json
{
    "user_login":"login12",
    "user_password":"lsdDS3221ds"
}
```

**Ответ httpStatus 201 Created:**
```json
{
    "success": {
        "user_id": "f8caf404-aaf7-458e-8f77-3a7d55af2f36",
        "user_login": "login12"
    }
}
```

Ошибки:

| Код | Сообщение | Описание |
|-----|-----------|----------|
| 400 |`incorrect json body`| Некорректный формат JSON|
| 500 |`internal server error`| Ошибка на стороне сервера|
| 409 |`user already exist`|Пользователь уже существует|
----------------
**2. POST /user/login - авторизация существующего пользователя**

**Запрос:**
```json
{
    "user_login":"login12",
    "user_password":"lsdDS3221ds"
}
```

**Ответ httpStatus 200 StatusOk:**
```json
{
    "user_id": "d2c8d036-6ab8-4663-b7c9-fc83ba1a7c0f",
    "role": "user",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjExNzc4NDgsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6ImQyYzhkMDM2LTZhYjgtNDY2My1iN2M5LWZjODNiYTFhN2MwZiJ9.sbM8DNiw1E8xmKGyvHniqQiRbEGil1cmbq9x-xEOTHI"
}
```
Ошибки:

| Код | Сообщение | Описание |
|-----|-----------|----------|
| 400 | `incorrect json body` | Некорректный формат JSON |
| 500 | `internal server error` | Ошибка на стороне сервера |
| 401 | `invalid password` | Неверный пароль пользователя |
----------------
 