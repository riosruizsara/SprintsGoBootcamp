package server

import (
	"net/http"

	sectionsHd "github.com/almarino_meli/grupo-5-wave-15/internal/handler/sections"
	sectionsLd "github.com/almarino_meli/grupo-5-wave-15/internal/loader/sections"
	sectionsRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/sections"
	sectionSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/sections"

	buyersHd "github.com/almarino_meli/grupo-5-wave-15/internal/handler/buyers"
	buyersLd "github.com/almarino_meli/grupo-5-wave-15/internal/loader/buyers"
	buyersRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/buyers"
	buyersSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/buyers"

	productsHd "github.com/almarino_meli/grupo-5-wave-15/internal/handler/products"
	productsLd "github.com/almarino_meli/grupo-5-wave-15/internal/loader/products"
	productsRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/products"
	productsSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/products"

	employeesHd "github.com/almarino_meli/grupo-5-wave-15/internal/handler/employees"
	employeesLd "github.com/almarino_meli/grupo-5-wave-15/internal/loader/employees"
	employeesRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/employees"
	employeesSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/employees"

	warehousesHd "github.com/almarino_meli/grupo-5-wave-15/internal/handler/warehouses"
	warehousesLd "github.com/almarino_meli/grupo-5-wave-15/internal/loader/warehouses"
	warehousesRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/warehouses"
	warehousesSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/warehouses"

	sellersHd "github.com/almarino_meli/grupo-5-wave-15/internal/handler/sellers"
	sellersLd "github.com/almarino_meli/grupo-5-wave-15/internal/loader/sellers"
	sellersRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/sellers"
	sellersSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/sellers"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ConfigServerChi struct {
	ServerAddress            string
	ProductsLoaderFilePath   string
	SellersLoaderFilePath    string
	BuyersLoaderFilePath     string
	EmployeesLoaderFilePath  string
	SectionsLoaderFilePath   string
	WarehousesLoaderFilePath string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.ProductsLoaderFilePath != "" {
			defaultConfig.ProductsLoaderFilePath = cfg.ProductsLoaderFilePath
		}
		if cfg.SellersLoaderFilePath != "" {
			defaultConfig.SellersLoaderFilePath = cfg.SellersLoaderFilePath
		}
		if cfg.BuyersLoaderFilePath != "" {
			defaultConfig.BuyersLoaderFilePath = cfg.BuyersLoaderFilePath
		}
		if cfg.EmployeesLoaderFilePath != "" {
			defaultConfig.EmployeesLoaderFilePath = cfg.EmployeesLoaderFilePath
		}
		if cfg.SectionsLoaderFilePath != "" {
			defaultConfig.SectionsLoaderFilePath = cfg.SectionsLoaderFilePath
		}
		if cfg.WarehousesLoaderFilePath != "" {
			defaultConfig.WarehousesLoaderFilePath = cfg.WarehousesLoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:            defaultConfig.ServerAddress,
		productsLoaderFilePath:   defaultConfig.ProductsLoaderFilePath,
		sellersLoaderFilePath:    defaultConfig.SellersLoaderFilePath,
		buyersLoaderFilePath:     defaultConfig.BuyersLoaderFilePath,
		employeesLoaderFilePath:  defaultConfig.EmployeesLoaderFilePath,
		sectionsLoaderFilePath:   defaultConfig.SectionsLoaderFilePath,
		warehousesLoaderFilePath: defaultConfig.WarehousesLoaderFilePath,
	}
}

type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// productsLoaderFilePath is the path to the file that contains the products
	productsLoaderFilePath string
	// sellersLoaderFilePath is the path to the file that contains the sellers
	sellersLoaderFilePath string
	// buyersLoaderFilePath is the path to the file that contains the buyers
	buyersLoaderFilePath string
	// employeesLoaderFilePath is the path to the file that contains the employees
	employeesLoaderFilePath string
	// sectionsLoaderFilePath is the path to the file that contains the sections
	sectionsLoaderFilePath string
	// warehousesLoaderFilePath is the path to the file that contains the warehouses
	warehousesLoaderFilePath string
}

func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	warehousesLoader := warehousesLd.NewWarehouseJSONFile(a.warehousesLoaderFilePath)
	warehousesDB, err := warehousesLoader.Load()
	if err != nil {
		return
	}

	sectionsLoader := sectionsLd.NewSectionJSONFile(a.sectionsLoaderFilePath)
	sectionsDB, err := sectionsLoader.Load()
	if err != nil {
		return
	}

	buyersLoader := buyersLd.NewBuyerJSONFile(a.buyersLoaderFilePath)
	buyersDB, err := buyersLoader.Load()
	if err != nil {
		return
	}

	productsLoader := productsLd.NewProductJSONFile(a.productsLoaderFilePath)
	productsDB, err := productsLoader.Load()
	if err != nil {
		return
	}

	employeesLoader := employeesLd.NewEmployeeJSON(a.employeesLoaderFilePath)
	employeesDB, err := employeesLoader.Load()
	if err != nil {
		return
	}

	sellersLoader := sellersLd.NewSellerJSONFile(a.sellersLoaderFilePath)
	sellersDB, err := sellersLoader.Load()
	if err != nil {
		return
	}

	// - repositories
	productsRepository := productsRp.NewProductMap(productsDB)
	sellersRepository := sellersRp.NewSellerMap(sellersDB)
	employeesRepository := employeesRp.NewEmployeeRepository(employeesDB)
	warehousesRepository := warehousesRp.NewWarehouseMap(warehousesDB)
	buyersRepository := buyersRp.NewBuyerMap(buyersDB)
	sectionsRepository := sectionsRp.NewSectionMap(sectionsDB)

	// - services
	productsService := productsSv.NewProductServiceDefault(productsRepository, sellersRepository)
	sellersService := sellersSv.NewSellerDefault(sellersRepository)
	employeesService := employeesSv.NewEmployeeService(employeesRepository, warehousesRepository)
	warehousesService := warehousesSv.NewWarehouseServiceDefault(warehousesRepository)
	buyersService := buyersSv.NewBuyerDefault(buyersRepository)
	sectionService := sectionSv.NewSectionServiceDefault(sectionsRepository)

	// - handleres/controllers
	productsHandler := productsHd.NewProductController(productsService)
	sellersHandler := sellersHd.NewSellerDefault(sellersService)
	employeesHandler := employeesHd.NewEmployeeHandler(employeesService)
	warehousesHandler := warehousesHd.NewWarehouseController(warehousesService)
	buyersHandler := buyersHd.NewBuyerDefault(buyersService)
	sectionsHandler := sectionsHd.NewSectionController(sectionService)

	// router
	rt := chi.NewRouter()

	// validator
	validate := validator.New()

	// - endpoints
	rt.Route("/api/v1/buyers", func(rt chi.Router) {
		rt.Get("/", buyersHandler.GetAllBuyers())
		rt.Get("/{id}", buyersHandler.GetBuyer())
		rt.Post("/", buyersHandler.PostBuyer(validate))
		rt.Patch("/{id}", buyersHandler.PatchBuyer(validate))
		rt.Delete("/{id}", buyersHandler.DeleteBuyer())
	})

	rt.Route("/api/v1/products", func(rt chi.Router) {
		rt.Get("/", productsHandler.GetAll())
		rt.Get("/{id}", productsHandler.GetByID())
		rt.Post("/", productsHandler.Create(validate))
		rt.Patch("/{id}", productsHandler.Update(validate))
		rt.Delete("/{id}", productsHandler.Delete())
	})

	rt.Route("/api/v1/sellers", func(rt chi.Router) {
		rt.Post("/", sellersHandler.CreateSeller(validate))
		rt.Get("/", sellersHandler.GetAllSellers())
		rt.Get("/{id}", sellersHandler.GetSellerById())
		rt.Patch("/{id}", sellersHandler.UpdateSeller(validate))
		rt.Delete("/{id}", sellersHandler.DeleteSeller())
	})

	rt.Route("/api/v1/sections", func(rt chi.Router) {
		rt.Get("/", sectionsHandler.GetAllSections())
		rt.Get("/{id}", sectionsHandler.GetSectionsByID())
		rt.Post("/", sectionsHandler.CreateSections())
		rt.Delete("/{id}", sectionsHandler.DeleteSection())
		rt.Patch("/{id}", sectionsHandler.UpdateSection())
	})

	rt.Route("/api/v1/employees", func(rt chi.Router) {
		rt.Post("/", employeesHandler.CreateEmployee(validate))
		rt.Get("/", employeesHandler.GetAllEmployees())
		rt.Get("/{id}", employeesHandler.GetEmployeeByID())
		rt.Patch("/{id}", employeesHandler.UpdateEmployee(validate))
		rt.Delete("/{id}", employeesHandler.DeleteEmployee())
	})

	rt.Route("/api/v1/warehouses", func(rt chi.Router) {
		rt.Get("/", warehousesHandler.GetAll())
		rt.Get("/{id}", warehousesHandler.GetByID())
		rt.Post("/", warehousesHandler.Create())
		rt.Patch("/{id}", warehousesHandler.Update())
		rt.Delete("/{id}", warehousesHandler.Delete())
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
