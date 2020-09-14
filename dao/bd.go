package dao

import (
	"database/sql"
	"errors"
	"log"

	"github.com/THEcanon001/AdminPw/encriptacion"

	"github.com/THEcanon001/AdminPw/model"
	_ "github.com/lib/pq"
)

type connection struct {
	driver string
	user   string
	pwd    string
	server string
	port   string
	dbname string
}

const encripted string = "1d65cfb3e5f8519f32724f58ba3813faef8a10c374707d2422948257b9336b31d73f036b11016c5c4e578985cc6e41198739944ebb507f647ba9f31d57277f70945a2d626c98531377c59e09a3e9bfe9c96ccb2690b17a56a73b6ff61d2f2aede0254e046842464724eaeb333d3eb829ec33d6ee54"

func getConnection(key string) *sql.DB {
	ds := encriptacion.Decrypt(encripted, key)
	db, err := sql.Open("postgres", ds) //abro conexion
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//ObtenerUsuario devuelve el hash y la salt del usuario asociado
func ObtenerUsuario(us string, key string) (model.Usuario, error) {
	sql := `SELECT id, usuario, password, salt FROM
				usuarios
				WHERE usuario = $1`

	db := getConnection(key)
	defer db.Close()
	rows, err := db.Query(sql, us)
	if err != nil {
		return model.Usuario{}, err
	}
	defer rows.Close()
	var u model.Usuario
	if rows.Next() {
		err = rows.Scan(&u.ID, &u.Usuario, &u.Password, &u.Salt)
		if err != nil {
			return model.Usuario{}, err
		}
	}
	return u, nil
}

//InsertarUsuario inserta un nuevo usuario
func InsertarUsuario(u model.Usuario, key string) error {
	sql := `INSERT INTO
				usuarios (usuario, password, salt)
				VALUES ($1, $2, $3)`

	db := getConnection(key)
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(u.Usuario, u.Password, u.Salt)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}
	return nil
}

//AgregarDato agrega datos codificados al usuario activo
func AgregarDato(ID int, d model.Data, key string) error {
	sql := `INSERT INTO
	data (nombre, password, idusuario)
	VALUES ($1, $2, $3)`

	db := getConnection(key)
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(d.Nombre, d.Password, ID)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}
	return nil
}

//VerDatos lista todos los datos sin decodificar de la base de datos
func VerDatos(ID int, key string) (datos map[int]model.Data, err error) {
	sql := `SELECT id, nombre, password, idusuario FROM
				data
				WHERE idusuario = $1`
	datos = make(map[int]model.Data)
	db := getConnection(key)
	defer db.Close()
	rows, err := db.Query(sql, ID)
	if err != nil {
		return
	}
	defer rows.Close()
	var contador int = 1
	for rows.Next() {
		d := model.Data{}
		err = rows.Scan(&d.ID, &d.Nombre, &d.Password, &d.IDusuario)
		if err != nil {
			return
		}
		datos[contador] = d
		contador++
	}
	return datos, nil
}
