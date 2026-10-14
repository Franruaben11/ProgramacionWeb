package handlers

import (
	"encoding/json"
	"net/http"
	db "main.go/db/sqlc"
	//"strings"
	//"strconv"
)

type Enfermero struct {
	Nombre  string  `json:"nombre"`
	Contrasena string `json:"contrasena"`
}

type EnfermeroHandler struct {
	repo *db.Queries
}

func NewEnfermeroHandler(repo *db.Queries) *EnfermeroHandler {
	return &EnfermeroHandler{
		repo: repo,
	}
}


// Manejador para /enfermero
func (h *EnfermeroHandler) EnfermerosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	
	/*case http.MethodGet:
		h.getAllEnfermeros(w, r)
	*/
	case http.MethodPost:
		h.createEnfermero(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


// Manejador para /enfermero/{id}
func (h *EnfermeroHandler) EnfermeroHandler(w http.ResponseWriter, r *http.Request) {
	/*
	// Extraer ID del path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid enfermero ID", http.StatusBadRequest)
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		h.getEnfermero(w, r, id)
	case http.MethodPut:
		h.updateEnfermero(w, r, id)
	case http.MethodDelete:
		h.deleteEnfermero(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
		*/
}

func (h *EnfermeroHandler) createEnfermero(w http.ResponseWriter, r *http.Request) {
	var newEnfermero Enfermero

	err := json.NewDecoder(r.Body).Decode(&newEnfermero)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	enfermeroCreado, err := h.repo.CreateEnfermero(r.Context(), db.CreateEnfermeroParams{
        Nombre:     newEnfermero.Nombre,
        Contrasena: newEnfermero.Contrasena,
    })
    
    if err != nil {
        http.Error(w, "Error al guardar en la base de datos: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // 3. Responder con el objeto creado devuelto por la base de datos
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(enfermeroCreado)
}

/*
// POST /enfermero - Crear nuevo producto
func (h *EnfermeroHandler) createEnfermero(w http.ResponseWriter, r *http.Request) {
	var newEnfermero Enfermero

	err := json.NewDecoder(r.Body).Decode(&newEnfermero)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//hacer get DB

	//hardcodea DB
	newEnfermero.ID = len(products) + 1
	products = append(products, newEnfermero)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEnfermero)
}
*/


/*
																	CORREGIR DESP
// GET /enfermero - Listar todos los enfermeros
func (h *EnfermeroHandler) getAllEnfermeros(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}


// GET /enfermero/{id} - Obtener enfermero específico
func (h *EnfermeroHandler) getEnfermero(w http.ResponseWriter, r *http.Request, id int) {

	product, err := findProductByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// PUT /enfermero/{id} - Actualizar enfermero
func (h *EnfermeroHandler) updateEnfermero(w http.ResponseWriter, r *http.Request, id int) {
	var updatedEnfermero Enfermero

	err := json.NewDecoder(r.Body).Decode(&updatedEnfermero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	index, err := findProductIndexByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Mantener el ID original
	updatedEnfermero.ID = id
	products[index] = updatedEnfermero
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEnfermero)
}



// DELETE /enfermero/{id} - Eliminar enfermero
func (h *EnfermeroHandler) deleteEnfermero(w http.ResponseWriter, r *http.Request, id int) {
	index, err := findProductIndexByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	products = append(products[:index], products[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func (h *EnfermeroHandler) findProductByID(id int) (*Enfermero, error) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], nil
		}
	}
	return nil, errors.New("product not found")
}

// Funciones auxiliares
func findProductIndexByID(id int) (int, error) {
	for i, p := range products {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("product not found")
}

*/





















