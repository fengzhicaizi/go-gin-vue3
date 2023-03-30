import request from '@/utils/request';
import type { QueryMenusParamsType, AddMenuBodyType, UpdateMenuBodyType } from './menu.d';

export const getMenuApi = (params?: QueryMenusParamsType) => request.get('/v1/menu/getMenu', { params });

export const createMenuApi = (body: AddMenuBodyType) => request.post('/v1/menu/createMenu', body);

export const updateMenuApi = (body: UpdateMenuBodyType) => request.put(`/v1/menu/updateMenu`, body);

export const deleteMenuApi = (id: string) => request.delete(`/v1/menu/deleteMenu/${id}`);
