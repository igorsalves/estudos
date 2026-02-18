package main

import (
	"fmt"
	"log"

	"github.com/crud-example/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	log.Println("Conectando ao banco de dados...")
	db, err := gorm.Open(sqlite.Open("meubanco.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	log.Println("Conexão estabelecida com sucesso.")

	log.Println("Executando migração do schema...")
	db.AutoMigrate(&users.User{})
	log.Println("Migração concluída.")

	repo := users.NewRepository(db)

	log.Println("[CREATE] Criando usuário 'Jhon Doe'...")
	id := repo.Create("Jhon Doe")
	log.Printf("[CREATE] Usuário criado com ID: %d\n", id)

	log.Printf("[UPDATE] Atualizando usuário ID %d para 'Jane Doe'...\n", id)
	repo.Update(id, "Jane Doe")
	log.Printf("[UPDATE] Usuário ID %d atualizado.\n", id)

	log.Println("[LIST] Listando todos os usuários:")
	for _, u := range repo.List() {
		fmt.Printf("  -> ID: %d | Nome: %s\n", u.ID, u.Name)
	}

	log.Printf("[DELETE] Deletando usuário ID %d...\n", id)
	repo.Delete(id)
	log.Printf("[DELETE] Usuário ID %d deletado.\n", id)

	log.Println("[LIST] Listando usuários após deleção:")
	remaining := repo.List()
	if len(remaining) == 0 {
		log.Println("  -> Nenhum usuário encontrado.")
	}
	for _, u := range remaining {
		fmt.Printf("  -> ID: %d | Nome: %s\n", u.ID, u.Name)
	}
}
