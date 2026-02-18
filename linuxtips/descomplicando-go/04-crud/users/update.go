package users

func (r *GormRepository) Update(id uint, name string) {
	var user User
	r.db.First(&user, id)
	r.db.Model(&user).Update("name", name)
}
