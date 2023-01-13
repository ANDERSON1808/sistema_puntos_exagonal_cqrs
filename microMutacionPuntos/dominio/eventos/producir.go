package eventos

import (
	"context"
	"fmt"
	"microMutationPuntos/dominio/entidades"
)

var temaPuntosUsuario = "topic_punt_users"

func EmitirActualizationPuntosPorUsuario(usuario *entidades.Usuario) {
	if err := PublishCreated(context.Background(), usuario, temaPuntosUsuario); err != nil {
		fmt.Println("failed to publish event: ", err)
	}
}
