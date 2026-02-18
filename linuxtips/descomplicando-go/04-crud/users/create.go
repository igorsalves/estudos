package users

func (r *GormRepository) Create(name string) uint {
	user := &User{Name: name}
	r.db.Create(user)
	return user.ID
}
