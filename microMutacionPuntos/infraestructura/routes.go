package infraestructura

func (s *Server) registerRoutes() {
	puntos := s.engine.Group("api/v1/puntos")
	puntos.Post("/redimir", s.RedimirPuntoHandler)
	puntos.Post("/acumulacion", s.AcumularPuntoHandler)
}
