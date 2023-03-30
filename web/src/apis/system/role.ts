import request from '@/utils/request';
import type { AddRoleBodyType, UpdateRoleBodyType } from './role.d';

export const getRoleListApi = () => request.get('/v1/role/getRoleList');

export const createRoleApi = (body: AddRoleBodyType) => request.post('/v1/role/createRole', body);

export const updateRoleApi = (body: UpdateRoleBodyType) => request.put(`/v1/role/updateRole`, body);

export const deleteRoleApi = (id: string) => request.delete(`/v1/role/deleteRole/${id}`);
