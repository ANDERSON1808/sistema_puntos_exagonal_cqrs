package eventos

import (
	"context"
	"fmt"
	"microMutationPuntos/dominio/entidades"
	"microMutationPuntos/infraestructura/Modelos"
)

var temaPuntosUsuario = "topic_punt_users"
var temaRedimirPuntos = "topic_redimir_punto_usuario"

func EmitirActualizationPuntosPorUsuario(usuario *entidades.Usuario) {
	if err := PublishCreated(context.Background(), usuario, temaPuntosUsuario); err != nil {
		fmt.Println("failed to publish event: ", err)
	}
}

func EmitirRedimirPuntos(m Modelos.RedimirPuntos) error {
	if err := PublishRedimir(context.Background(), m, temaRedimirPuntos); err != nil {
		fmt.Println("failed to publish event: ", err)
		return err
	}
	return nil
}
