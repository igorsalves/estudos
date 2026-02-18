package users

func (r *GormRepository) List() []User {
	var users []User
	r.db.Find(&users)
	return users
}
