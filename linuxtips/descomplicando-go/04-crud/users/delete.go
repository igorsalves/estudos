package users

func (r *GormRepository) Delete(id uint) {
	var user User
	r.db.First(&user, id)
	r.db.Delete(&user)
}
