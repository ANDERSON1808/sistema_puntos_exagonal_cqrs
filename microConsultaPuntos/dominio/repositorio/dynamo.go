package repositorio

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"microConsultaPuntos/dominio/entidades"
	"strconv"
)

type UsuarioRepositorio struct {
	db *dynamodb.DynamoDB
}

var NombreTablaUsuario = "usuarios"

func NewUsuarioRepositorio(db *dynamodb.DynamoDB) *UsuarioRepositorio {
	return &UsuarioRepositorio{db: db}
}

func (r *UsuarioRepositorio) Close() {
	//
}

func (r UsuarioRepositorio) CrearUsuario(usuario *entidades.Usuario) (err error) {
	_, err = r.db.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(usuario.UsuarioId)),
			},
			"NombreUsuario": {
				S: aws.String(usuario.NombreUsuario),
			},
			"TotalPuntos": {
				S: aws.String(usuario.TotalPuntos),
			},
		},
		TableName: &NombreTablaUsuario,
	})
	return
}

func (r UsuarioRepositorio) BuscarUsuarioId(usuarioId int) (respuesta entidades.Usuario, err error) {
	result, err := r.db.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(usuarioId)),
			},
		},
		TableName: &NombreTablaUsuario,
	})

	if err != nil {
		fmt.Println("Error BuscarUsuarioId", err)
		return
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &respuesta)
	if err != nil {
		fmt.Println("Error convertir modelo usuarioId", err)
		return
	}
	return
}

func (r UsuarioRepositorio) ActualizarUsuario(usuario *entidades.Usuario) (err error) {
	_, err = r.db.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("NombreUsuario"),
			"#W": aws.String("TotalPuntos"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":NombreUsuario": {
				S: aws.String(usuario.NombreUsuario),
			},
			":TotalPuntos": {
				S: aws.String(usuario.TotalPuntos),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(usuario.UsuarioId)),
			},
		},
		TableName:        &NombreTablaUsuario,
		UpdateExpression: aws.String("SET #N = :NombreUsuario, #W = :TotalPuntos"),
	})
	if err != nil {
		fmt.Println("Error al actualizar usuario", err)
		return
	}
	return
}
