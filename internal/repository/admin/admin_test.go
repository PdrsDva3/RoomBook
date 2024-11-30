package admin

import (
	"RoomBook/internal/models"
	"RoomBook/pkg/cerr"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"golang.org/x/net/context"
	"testing"
)

func TestRepo_Create(t *testing.T) {
	// Создаем mock-объект для базы данных
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Инициализируем репозиторий
	sqlxDB := sqlx.NewDb(db, "postgres")
	repo := InitAdminRepository(sqlxDB)

	// Пример данных для создания
	adminCreate := models.AdminCreate{
		AdminBase: models.AdminBase{
			Name:  "Test",
			Email: "test@example.com",
			Phone: "1234567890",
			Photo: "photo_url",
		},
		PWD: "hashed_password",
	}

	t.Run("success", func(t *testing.T) {
		// Мокаем успешный запрос
		mock.ExpectBegin()
		mock.ExpectQuery("SELECT 1")
		mock.ExpectCommit()

		// Вызываем метод
		id, err := repo.Create(context.Background(), adminCreate)

		// Проверяем результаты
		require.NoError(t, err)
		assert.Equal(t, 0, id)

		// Проверяем все ожидания
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("transaction begin error", func(t *testing.T) {
		// Мокаем ошибку при начале транзакции
		mock.ExpectBegin().WillReturnError(err)

		// Вызываем метод
		id, err := repo.Create(context.Background(), adminCreate)

		// Проверяем ошибки
		assert.Equal(t, 0, id)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), cerr.Transaction)

		// Проверяем все ожидания
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query execution error", func(t *testing.T) {
		// Мокаем успешное начало транзакции
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO admins (name, email, phone, hashed_password, photo) VALUES ($1, $2, $3, $4, $5) returning id;`).
			WithArgs(adminCreate.Name, adminCreate.Email, adminCreate.Phone, adminCreate.PWD, adminCreate.Photo).
			WillReturnError(err) // Ошибка выполнения запроса

		// Вызываем метод
		id, err := repo.Create(context.Background(), adminCreate)

		// Проверяем ошибки
		assert.Equal(t, 0, id)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), cerr.Scan)

		// Проверяем все ожидания
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("transaction rollback error", func(t *testing.T) {
		// Мокаем ошибку при откате транзакции
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO admins (name, email, phone, hashed_password, photo) VALUES ($1, $2, $3, $4, $5) returning id;`).
			WithArgs(adminCreate.Name, adminCreate.Email, adminCreate.Phone, adminCreate.PWD, adminCreate.Photo).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(0)) // Вернем пустой id
		mock.ExpectRollback().WillReturnError(err)

		// Вызываем метод
		id, err := repo.Create(context.Background(), adminCreate)

		// Проверяем ошибки
		assert.Equal(t, 0, id)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), cerr.Rollback)

		// Проверяем все ожидания
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
