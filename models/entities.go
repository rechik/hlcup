package models

import "errors"

//go:generate easyjson -all $GOFILE

// User is user profile
type User struct {

	// адрес электронной почты пользователя. Тип - unicode-строка
	// длиной до 100 символов. Гарантируется уникальность.
	Email string `json:"email"`

	// имя и фамилия соответственно. Тип - unicode-строки длиной до 50 символов.
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	// unicode-строка "m" означает мужчской пол, а "f" - женский.
	Gender string `json:"gender"`

	// дата рождения, записанная как число секунд от начала UNIX-эпохи по UTC
	// (другими словами - это timestamp). Ограничено снизу 01.01.1930 и сверху
	// 01.01.1999-ым.
	BirthDate int64 `json:"birth_date"`

	// уникальный внешний идентификатор пользователя. Устанавливается
	// тестирующей системой и используется затем, для проверки ответов сервера.
	// 32-разрядное целое число.
	ID uint32 `json:"id"`
}

func (v *User) GetID() uint32 {
	return v.ID
}

func (v *User) Validate() error {
	switch {
	case v.ID == 0:
		return errors.New("id should be non-zero")
	case len(v.Email) > 100:
		return errors.New("email is too long")
	case len(v.FirstName) > 50:
		return errors.New("first_name is too long")
	case len(v.LastName) > 50:
		return errors.New("last_name is too long")
	case v.Gender != "m" && v.Gender != "f":
		return errors.New("invalid gender")
		//case v.BirthDate < -2147483648 || v.BirthDate > 2147483647:
		//	return fmt.Errorf("invalid birth_date: %d", v.BirthDate)
	}
	return nil
}

func (v User) IsValid() bool {
	return v.ID != 0
}

// Location is Достопримечательность
type Location struct {

	// уникальный внешний id достопримечательности. Устанавливается тестирующей
	// системой. 32-разрядное целое число.
	ID uint32 `json:"id"`

	// расстояние от города по прямой в километрах. 32-разрядное целое число.
	Distance uint32 `json:"distance"`

	// описание достопримечательности. Текстовое поле неограниченной длины.
	Place string `json:"place"`

	// название страны расположения. unicode-строка длиной до 50 символов.
	Country string `json:"country"`

	// название города расположения. unicode-строка длиной до 50 символов.
	City string `json:"city"`
}

func (v *Location) GetID() uint32 {
	return v.ID
}

func (v *Location) Validate() error {
	switch {
	case v.ID == 0:
		return errors.New("id should be non-zero")
	case len(v.Country) > 50:
		return errors.New("country is too long")
	case len(v.City) > 50:
		return errors.New("city is too long")
	}
	return nil
}

func (v Location) IsValid() bool {
	return v.ID != 0
}

// Visit is Посещение
type Visit struct {

	// дата посещения, timestamp с ограничениями: снизу 01.01.2000, а сверху 01.01.2015.
	VisitedAt int `json:"visited_at"`

	// уникальный внешний id посещения. Устанавливается тестирующей системой.
	// 32-разрядное целое число.
	ID uint32 `json:"id"`

	// id достопримечательности. 32-разрядное целое число.
	Location uint32 `json:"location"`

	// id путешественника. 32-разрядное целое число.
	User uint32 `json:"user"`

	// оценка посещения от 0 до 5 включительно. Целое число.
	Mark uint8 `json:"mark"`
}

func (v *Visit) GetID() uint32 {
	return v.ID
}

func (v *Visit) Validate() error {
	switch {
	case v.ID == 0:
		return errors.New("id should be non-zero")
	case v.Mark < 0 || v.Mark > 5:
		return errors.New("invalid mark")
	}
	return nil
}

func (v Visit) IsValid() bool {
	return v.ID != 0
}
