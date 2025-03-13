package roleerrors

import "errors"

var (
	// Errores de validación
	ErrRoleNameEmpty        = errors.New("el nombre del rol no puede estar vacío")
	ErrRoleDescriptionEmpty = errors.New("la descripción del rol no puede estar vacía")

	// Errores de operaciones
	ErrRoleNotFound       = errors.New("rol no encontrado")
	ErrRoleAlreadyExists  = errors.New("ya existe un rol con ese nombre")
	ErrRoleCreationFailed = errors.New("error al crear el rol")
	ErrRoleUpdateFailed   = errors.New("error al actualizar el rol")
	ErrRoleDeleteFailed   = errors.New("error al eliminar el rol")

	// Errores de relaciones
	ErrRoleHasUsers         = errors.New("no se puede eliminar el rol porque tiene usuarios asignados")
	ErrPermissionNotFound   = errors.New("permiso no encontrado")
	ErrRolePermissionFailed = errors.New("error al gestionar permisos del rol")
)
