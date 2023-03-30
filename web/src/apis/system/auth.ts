import request from '@/utils/request';
import type { LoginBodyType, QueryParamsType, AuthBodyType, UpdateAuthBodyType } from './auth.d';

// 登录
export const loginApi = (body: LoginBodyType) => request.post('/v1/auth/login', body);

//　获取所有的用户
export const getAuthListApi = (params: QueryParamsType) => request.get('/v1/auth/getAuthList', { params });

export const createAuthApi = (body: AuthBodyType) => request.post('/v1/auth/createAuth', body);

// 修改人员信息
export const updateAuthApi = (body: UpdateAuthBodyType) => request.put(`/v1/auth/updateAuth`, body);

// 删除人员
export const deleteAuthApi = (id: number) => request.delete(`/v1/auth/deleteAuth/${id}`);

// 重置密码
export const resetPasswordApi = (id: number) => request.patch(`/v1/auth/resetPassword/${id}`);

// 根据token获取用户信息
export const getAuthByTokenApi = () => request.get(`/v1/auth/getAuthByToken`);
