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

const encripted string = "9facb884924d3760630d6f093163d17c472b061a867ab479f2f3c4108fa5a16a0e2d92de6ee2c63d88bd5b8b0a979b5038215b34f5c7a225ef579da28205b30f451d9a987c019e190121f359b2eddfee4dd0e7fc966f687a32329d0ab45c0b9dc9f84c7bf37e7dbebbda98a673047d7a966ca155b0"

func getConnection() *sql.DB {
	ds := encriptacion.Decrypt(encripted, "01f5cebdcc5ef0a2861317086ab7e24f2091a9a551fb8c1f86857f543e1d2ef0")
	db, err := sql.Open("postgres", ds) //abro conexion
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//ObtenerUsuario devuelve el hash y la salt del usuario asociado
func ObtenerUsuario(us string) (model.Usuario, error) {
	sql := `SELECT id, usuario, password, salt FROM
				usuarios
				WHERE usuario = $1`

	db := getConnection()
	defer db.Close()
	rows, err := db.Query(sql, us)
	if err != nil {
		return model.Usuario{}, err
	}
	defer rows.Close()
	var u model.Usuario
	if rows.Next() {
		err = rows.Scan(&u.ID, &u.User, &u.Password, &u.Salt)
		if err != nil {
			return model.Usuario{}, err
		}
	}
	return u, nil
}

//InsertarUsuario inserta un nuevo usuario
func InsertarUsuario(u model.Usuario) error {
	sql := `INSERT INTO
				usuarios (usuario, password, salt)
				VALUES ($1, $2, $3)`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(u.User, u.Password, u.Salt)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}
	return nil
}

//ModificarUsuario actualiza el usuario pasado como parametro
func ModificarUsuario(u model.Usuario) error {
	sql := `UPDATE usuarios
			SET usuario = $1, password = $2, salt = $3
			WHERE id = $4`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.User, u.Password, u.Salt, u.ID)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}
	return nil
}

//EliminarUsuario elimina el usuario pasado como parametro y todos sus datos
func EliminarUsuario(u model.Usuario) error {
	eliminarDatoUsuario(u.ID)
	sql := `DELETE FROM usuarios
			WHERE id = $1`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.ID)
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
func AgregarDato(ID int, d model.Data) error {
	sql := `INSERT INTO
	data (name, userdata, password, iduser)
	VALUES ($1, $2, $3, $4)`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(d.Name, d.User, d.Password, ID)
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
func VerDatos(ID int) (datos map[int]model.Data, err error) {
	sql := `SELECT id, name, userdata, password, iduser FROM
				data
				WHERE iduser = $1 
				order by id`
	datos = make(map[int]model.Data)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query(sql, ID)
	if err != nil {
		return
	}
	defer rows.Close()
	var contador int = 1
	for rows.Next() {
		d := model.Data{}
		err = rows.Scan(&d.ID, &d.Name, &d.User, &d.Password, &d.IDusuario)
		if err != nil {
			return
		}
		datos[contador] = d
		contador++
	}
	return datos, nil
}

//ModificarDato actualiza el dato pasado como parametro
func ModificarDato(ID int, d model.Data) error {
	sql := `UPDATE data
			SET name = $1, userdata = $2, password = $3
			WHERE id = $4 and iduser = $5`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(d.Name, d.User, d.Password, d.ID, ID)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}
	return nil
}

//EliminarDato elimina el dato pasado como parametro
func EliminarDato(ID int, idData int) error {
	sql := `DELETE FROM data
			WHERE id = $1 and iduser = $2`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(idData, ID)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}
	return nil
}

func eliminarDatoUsuario(ID int) error {
	sql := `DELETE FROM data
			WHERE iduser = $1`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(ID)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba una fila afectada")
	}
	return nil
}
