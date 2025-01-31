package employees

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	employeesService "github.com/almarino_meli/grupo-5-wave-15/internal/service/employees"
	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/employees"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type EmployeeHandler struct {
	service employeesService.EmployeeService
}

func NewEmployeeHandler(service employeesService.EmployeeService) EmployeeHandler {
	return EmployeeHandler{service: service}
}

func (eh *EmployeeHandler) CreateEmployee(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e employees.EmployeeDTO
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			response.Error(w, http.StatusBadRequest, "Error creating employee: Invalid input data")
			return
		}

		employeeToCreate := employees.ToEmployeeModel(e)

		if err := employeeToCreate.Validate(validate); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Error creating employee: Invalid input data")
			return
		}

		createdEmployee, err := eh.service.Create(&employeeToCreate)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.DuplicateError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error creating employee: Invalid input data",
				})
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error creating employee: Invalid input data",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error creating employee: Server ran into an error",
				})
			}
			return
		}

		createdEmployeeDTO := employees.ToEmployeeDTO(createdEmployee)

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Employee created successfully",
			"data":    createdEmployeeDTO,
		})

	}
}

func (eh *EmployeeHandler) GetAllEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e, err := eh.service.GetAll()

		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "No employees found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error getting employees: Server ran into an error",
				})
			}
			return
		}

		data := make([]employees.EmployeeDTO, 0, len(e))
		for _, employee := range e {
			data = append(data, employees.ToEmployeeDTO(employee))
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Employees retrieved successfully",
			"data":    data,
		})
	}
}

func (eh *EmployeeHandler) GetEmployeeByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Error fetching employee: Invalid input")
			return
		}

		e, err := eh.service.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error fetching employee: Invalid input",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error fetching employee: Server ran into an error",
				})
			}
			return
		}

		eDTO := employees.ToEmployeeDTO(e)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Employee retrieved successfully",
			"data":    eDTO,
		})
	}
}

func (eh *EmployeeHandler) UpdateEmployee(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Error updating employee: Invalid input")
			return
		}

		var e employees.EmployeePatch
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			response.Error(w, http.StatusBadRequest, "Error updating employee: Invalid input data")
			return
		}

		e.Id = &id

		if err := e.Validate(validate); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Error creating employee: Invalid input data")
			return
		}

		updatedEmployee, err := eh.service.Update(e)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.ValidationError{}):
				response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
					"message": "Error updating employee: Invalid input data",
				})
			case errors.Is(err, &customErrors.DuplicateError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error updating employee: Invalid input data",
				})
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error updating employee: Not found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error updating employee: Server ran into an error",
				})
			}
			return
		}

		updatedEmployeeDTO := employees.ToEmployeeDTO(updatedEmployee)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Employee updated successfully",
			"data":    updatedEmployeeDTO,
		})

	}
}

func (eh *EmployeeHandler) DeleteEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Error deleting employee: Invalid input")
			return
		}

		err = eh.service.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error deleting employee: Employee not found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error deleting employee: Server ran into an error",
				})
			}
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{})
	}
}
