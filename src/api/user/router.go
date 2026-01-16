package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"otus.ru/tbw/msa-25/src/api"
	"otus.ru/tbw/msa-25/src/deps"
)

func RegisterRoutes(pathPrefix string, deps *deps.Deps) *http.ServeMux {
	router := http.NewServeMux()

	// Создание нового пользователя
	api.MuxRegisterPath(router, "POST", pathPrefix, "",
		func(w http.ResponseWriter, r *http.Request) {
			user := User{}
			ctx := context.Background()
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				api.WriteJsonError(w, r, http.StatusBadRequest, "No body", err, deps)
				return
			}

			_, err := deps.DB.NewInsert().Model(&user).Exec(ctx)
			if err != nil {
				api.WriteJsonError(w, r, http.StatusInternalServerError, "Error inserting into db", err, deps)
				return
			}
			api.WriteJson(w, r, http.StatusOK, user, nil, deps)
		})

	// Получение пользователя по ID
	api.MuxRegisterPath(router, "GET", pathPrefix, "/{userId}",
		func(w http.ResponseWriter, r *http.Request) {
			userId := r.PathValue("userId")
			user := User{}
			ctx := context.Background()

			err := deps.DB.NewSelect().Where("id = ?", userId).Model(&user).Scan(ctx)

			if err != nil {
				if err == sql.ErrNoRows {
					api.WriteJsonError(w, r, http.StatusNotFound, "Record not found", err, deps)
					return
				}
				api.WriteJsonError(w, r, http.StatusInternalServerError, "Error selecting from db", err, deps)
				return
			}
			api.WriteJson(w, r, http.StatusOK, user, nil, deps)
		})

	// Удаление пользователя по ID
	api.MuxRegisterPath(router, "DELETE", pathPrefix, "/{userId}",
		func(w http.ResponseWriter, r *http.Request) {
			userId := r.PathValue("userId")
			ctx := context.Background()

			res, err := deps.DB.NewDelete().Model((*User)(nil)).Where("id = ?", userId).Exec(ctx)
			if err != nil {
				api.WriteJsonError(w, r, http.StatusInternalServerError, "Error deleting from db", err, deps)
				return
			}

			rowsAffected, err := res.RowsAffected()
			if err != nil {
				api.WriteJsonError(w, r, http.StatusInternalServerError, "Error deleting from db, rows affected problem", err, deps)
				return
			}

			if rowsAffected == 0 {
				api.WriteJsonError(w, r, http.StatusNotFound, "Record not found", err, deps)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		})

	api.MuxRegisterPath(router, "PUT", pathPrefix, "/{userId}",
		func(w http.ResponseWriter, r *http.Request) {
			userId := r.PathValue("userId")
			user := User{}
			ctx := context.Background()

			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				api.WriteJsonError(w, r, http.StatusBadRequest, "No body", err, deps)
				return
			}

			res, err := deps.DB.NewUpdate().Model(&user).OmitZero().Where("id = ?", userId).Returning("*").Exec(ctx)
			if err != nil {
				api.WriteJsonError(w, r, http.StatusInternalServerError, "Error updating in db", err, deps)
				return
			}

			rowsAffected, err := res.RowsAffected()
			if err != nil {
				api.WriteJsonError(w, r, http.StatusInternalServerError, "Error updating in db, rows affected problem", err, deps)
				return
			}

			if rowsAffected == 0 {
				api.WriteJsonError(w, r, http.StatusNotFound, "Record not found", err, deps)
				return
			}

			api.WriteJson(w, r, http.StatusOK, user, nil, deps)
		})

	return router
}
