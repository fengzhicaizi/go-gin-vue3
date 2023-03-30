import request from '@/utils/request';
import type { DictType, DictData } from './dict.d';

export const getDictlistApi = () => request.get(`/v1/dict/getDictList`);

export const createDictApi = (body: DictType) => request.post(`/v1/dict/createDict`, body);

export const updateDictApi = (body: DictType) => request.put(`/v1/dict/updateDict`, body);

export const deleteDictApi = (id: number) => request.delete(`/v1/dict/deleteDict/${id}`);

export const getDictDataApi = (type: string) => request.get(`/v1/dict/getDictData/${type}`);

export const addDictDataApi = (type: string, body: DictData) => request.post(`/v1/dict/addDictData/${type}`, body);

export const updateDictDataApi = (body: DictData) => request.put(`/v1/dict/updateDictData`, body);

export const deleteDictDataApi = (id: number) => request.delete(`/v1/dict/deleteDictDataApi/${id}`);
