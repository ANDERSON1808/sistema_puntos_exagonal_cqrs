package repositorio

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"microMutationPuntos/dominio/entidades"
	"strconv"
)

type PuntosRepositorio struct {
	db *dynamodb.DynamoDB
}

var TableName = "gestion_puntos"

func NewPuntosRepositorio(db *dynamodb.DynamoDB) *PuntosRepositorio {
	return &PuntosRepositorio{db: db}
}

func (r PuntosRepositorio) AcumularPuntos(puntos *entidades.Puntos) (err error) {
	_, err = r.db.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(puntos.PuntoId)),
			},
			"UsuarioId": {
				S: aws.String(puntos.UsuarioId),
			},
			"Punto": {
				S: aws.String(puntos.Punto),
			},
			"DetalleMovimiento": {
				S: aws.String(puntos.DetalleMovimiento),
			},
		},
		TableName: &TableName,
	})
	return err
}

func (r PuntosRepositorio) BuscarPuntoId(punto int) (respuesta entidades.Puntos, err error) {
	result, err := r.db.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(punto)),
			},
		},
		TableName: &TableName,
	})

	if err != nil {
		fmt.Println("Error BuscarPuntoId", err)
		return
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &respuesta)
	if err != nil {
		fmt.Println("Error convirtiendo modelo puntos id", err)
		return
	}
	return
}
