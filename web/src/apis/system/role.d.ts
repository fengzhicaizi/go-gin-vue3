export type AddRoleBodyType = {
	name?: string;
	menuIds?: string;
};

export type UpdateRoleBodyType = AddRoleBodyType & {
	id: number;
};

export type BindMenusBodyType = {
	roleId: number;
	menuIds: string;
};
