package entities

//go:generate ffjson $GOFILE

// User is user profile
type User struct {

	// уникальный внешний идентификатор пользователя. Устанавливается
	// тестирующей системой и используется затем, для проверки ответов сервера.
	// 32-разрядное целое число.
	ID int32 `json:"id"`

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
	BirthDate int `json:"birth_date"`
}

// Location is Достопримечательность
type Location struct {

	// уникальный внешний id достопримечательности. Устанавливается тестирующей
	// системой. 32-разрядное целое число.
	ID int32 `json:"id"`

	// описание достопримечательности. Текстовое поле неограниченной длины.
	Place string `json:"place"`

	// название страны расположения. unicode-строка длиной до 50 символов.
	Country string `json:"country"`

	// название города расположения. unicode-строка длиной до 50 символов.
	City string `json:"city"`

	// расстояние от города по прямой в километрах. 32-разрядное целое число.
	Distance int32 `json:"distance"`
}

// Visit is Посещение
type Visit struct {

	// уникальный внешний id посещения. Устанавливается тестирующей системой.
	// 32-разрядное целое число.
	ID int32 `json:"id"`

	// id достопримечательности. 32-разрядное целое число.
	Location int32 `json:"location"`

	// id путешественника. 32-разрядное целое число.
	User int32 `json:"user"`

	// дата посещения, timestamp с ограничениями: снизу 01.01.2000, а сверху 01.01.2015.
	VisitedAt int `json:"visited_at"`

	// оценка посещения от 0 до 5 включительно. Целое число.
	Mark uint8
}
