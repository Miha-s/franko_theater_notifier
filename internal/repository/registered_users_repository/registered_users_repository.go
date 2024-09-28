package registeredusersrepository

import (
	"encoding/json"
	"os"
)

type RegisteredUsersRepository struct {
	storagePath       string
	RegisteredChatIds map[int64]bool
}

func NewRegisteredUsersRepository(storage_path string) *RegisteredUsersRepository {
	r := &RegisteredUsersRepository{
		storagePath:       storage_path,
		RegisteredChatIds: make(map[int64]bool),
	}
	r.load()
	return r
}

func (r *RegisteredUsersRepository) AddChatId(id int64) {
	r.RegisteredChatIds[id] = true
	r.save()
}

func (r *RegisteredUsersRepository) RemoveChatId(id int64) {
	delete(r.RegisteredChatIds, id)
	r.save()
}

func (r *RegisteredUsersRepository) save() error {
	file, err := os.Create(r.storagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(r.RegisteredChatIds)
}

func (r *RegisteredUsersRepository) load() error {
	file, err := os.Open(r.storagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&r.RegisteredChatIds)
}
