package admin

import (
	"RoomBook/internal/models"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_Get(t *testing.T) {
	// Инициализируем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Не удалось создать mock DB: %v", err)
	}
	defer db.Close()

	// Инициализируем объект репозитория с мока
	repo := &Repo{db: sqlx.NewDb(db, "postgres")}

	// Ожидаемый результат
	expectedAdmin := &models.Admin{
		ID: 1,
		AdminBase: models.AdminBase{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Phone: "+123456789",
			Photo: "0000",
		},
	}

	// Настроим мок: при выполнении запроса с id=1 возвращаем результат
	rows := sqlmock.NewRows([]string{"name", "email", "phone", "photo"}).
		AddRow(expectedAdmin.Name, expectedAdmin.Email, expectedAdmin.Phone, expectedAdmin.Photo)

	mock.ExpectQuery(`SELECT name, email, phone, photo from admins WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	// Выполняем тестируемую функцию
	admin, err := repo.Get(context.Background(), 1)
	t.Logf("expected: %v\n  \t \t\t\tgive  : %v", expectedAdmin, admin)

	// Проверяем что ошибок нет
	assert.NoError(t, err)

	// Проверяем, что возвращенный результат совпадает с ожидаемым
	assert.NotNil(t, admin)
	assert.Equal(t, expectedAdmin.ID, admin.ID)
	assert.Equal(t, expectedAdmin.Name, admin.Name)
	assert.Equal(t, expectedAdmin.Email, admin.Email)
	assert.Equal(t, expectedAdmin.Phone, admin.Phone)
	assert.Equal(t, expectedAdmin.Photo, admin.Photo)

	// Проверяем, что все ожидаемые запросы были выполнены
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepo_Get_NoAdminFound(t *testing.T) {
	// Инициализируем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Не удалось создать mock DB: %v", err)
	}
	defer db.Close()

	// Инициализируем объект репозитория с мока
	repo := &Repo{db: sqlx.NewDb(db, "postgres")}

	// Настроим мок: при запросе с id=1 возвращаем пустой результат (нет администратора)
	mock.ExpectQuery(`SELECT name, email, phone, photo from admins WHERE id = \$1`).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	// Выполняем тестируемую функцию
	admin, err := repo.Get(context.Background(), 1)
	t.Logf("expected: %v\n  \t \t\t\tgive  : %v", sql.ErrNoRows, err)

	// Проверяем, что ошибка вернулась (ожидаем ошибку поиска)
	assert.Error(t, err)
	assert.Nil(t, admin)

	// Проверяем, что все ожидаемые запросы были выполнены
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepo_Create_Success(t *testing.T) {
	// Инициализация мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Не удалось создать mock DB: %v", err)
	}
	defer db.Close()

	// Инициализация репозитория с мокированным db
	repo := &Repo{db: sqlx.NewDb(db, "postgres")}

	// Ожидаемые данные для администратора
	adminCreate := models.AdminCreate{
		AdminBase: models.AdminBase{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Phone: "+123456789",
			Photo: "0000",
		},
		PWD: "0000",
	}

	// Ожидаемый id, который будет возвращен при вставке
	expectedID := 1

	// Настроим мок для транзакции
	mock.ExpectBegin() // Начинаем транзакцию
	mock.ExpectQuery(`INSERT into admins \(name, email, phone, hashed_password, photo\) VALUES \(\$1, \$2, \$3, \$4, \$5\) returning id`).
		WithArgs(adminCreate.Name, adminCreate.Email, adminCreate.Phone, adminCreate.PWD, adminCreate.Photo).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID)) // Возвращаем id
	mock.ExpectCommit() // Совершаем commit транзакции

	// Выполняем метод
	id, err := repo.Create(context.Background(), adminCreate)
	t.Logf("expected: %v\n  \t \t\t\tgive  : %v", expectedID, id)

	// Проверяем, что ошибок нет
	assert.NoError(t, err)

	// Проверяем, что id совпадает с ожидаемым
	assert.Equal(t, expectedID, id)

	// Проверяем, что все запросы были выполнены
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
